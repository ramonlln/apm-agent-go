// +build gofuzz

package elasticapm

import (
	"context"
	"crypto/tls"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"strings"
	"time"

	"github.com/pkg/errors"

	"github.com/elastic/apm-agent-go/internal/fastjson"
	"github.com/elastic/apm-agent-go/model"
	"github.com/elastic/apm-agent-go/stacktrace"
	"github.com/elastic/apm-server/processor"
	error_processor "github.com/elastic/apm-server/processor/error"
	metric_processor "github.com/elastic/apm-server/processor/metric"
	transaction_processor "github.com/elastic/apm-server/processor/transaction"
)

func Fuzz(data []byte) int {
	type Payload struct {
		Service      *model.Service       `json:"service"`
		Process      *model.Process       `json:"process,omitempty"`
		System       *model.System        `json:"system,omitempty"`
		Errors       []*model.Error       `json:"errors"`
		Transactions []*model.Transaction `json:"transactions"`
	}
	var payload Payload
	if err := json.Unmarshal(data, &payload); err != nil {
		return 0
	}

	tracer := DefaultTracer
	tracer.Transport = &gofuzzTransport{}
	tracer.SetCaptureBody(CaptureBodyAll)

	setContext := func(in *model.Context, out *Context) error {
		if in == nil {
			return nil
		}
		for _, item := range in.Custom {
			out.SetCustom(item.Key, item.Value)
		}
		for k, v := range in.Tags {
			out.SetTag(k, v)
		}
		if in.Request != nil {
			var body io.Reader
			var postForm url.Values
			if in.Request.Body != nil {
				body = strings.NewReader(in.Request.Body.Raw)
				if in.Request.Body.Form != nil {
					postForm = in.Request.Body.Form
				}
			}
			req, err := http.NewRequest(in.Request.Method, in.Request.URL.Full, body)
			if err != nil {
				return err
			}
			capturedBody := tracer.CaptureHTTPRequestBody(req)
			if in.Request.Socket != nil {
				req.RemoteAddr = in.Request.Socket.RemoteAddress
				if in.Request.Socket.Encrypted {
					req.TLS = new(tls.ConnectionState)
				}
			}
			req.PostForm = postForm
			if in.User != nil && in.User.Username != "" {
				req.SetBasicAuth(in.User.Username, "")
			}

			var major, minor int
			if n, err := fmt.Sscanf(in.Request.HTTPVersion, "%d.%d", &major, &minor); err != nil {
				return err
			} else if n != 2 {
				return errors.Errorf("invalid HTTP version %s", in.Request.HTTPVersion)
			}
			req.ProtoMajor = major
			req.ProtoMinor = minor

			if in.Request.Headers != nil {
				if in.Request.Headers.UserAgent != "" {
					req.Header.Set("User-Agent", in.Request.Headers.UserAgent)
				}
				if in.Request.Headers.ContentType != "" {
					req.Header.Set("Content-Type", in.Request.Headers.ContentType)
				}
				if in.Request.Headers.Cookie != "" {
					for _, v := range strings.Split(in.Request.Headers.Cookie, ";") {
						req.Header.Add("Cookie", v)
					}
				}
			}

			out.SetHTTPRequest(req)
			out.SetHTTPRequestBody(capturedBody)
		}
		if in.Response != nil {
			out.SetHTTPStatusCode(in.Response.StatusCode)
			if in.Response.Finished != nil {
				out.SetHTTPResponseFinished(*in.Response.Finished)
			}
			if in.Response.HeadersSent != nil {
				out.SetHTTPResponseHeadersSent(*in.Response.HeadersSent)
			}
			if in.Response.Headers != nil {
				h := make(http.Header)
				h.Set("Content-Type", in.Response.Headers.ContentType)
				out.SetHTTPResponseHeaders(h)
			}
		}
		return nil
	}

	for _, t := range payload.Transactions {
		if t == nil {
			continue
		}
		tx := tracer.StartTransaction(t.Name, t.Type)
		tx.Result = t.Result
		tx.Timestamp = time.Time(t.Timestamp)
		if setContext(t.Context, &tx.Context) != nil {
			return 0
		}
		for _, s := range t.Spans {
			span := tx.StartSpan(s.Name, s.Type, nil)
			span.Timestamp = tx.Timestamp.Add(time.Duration(s.Start * float64(time.Millisecond)))
			if s.Context != nil && s.Context.Database != nil {
				span.Context.SetDatabase(DatabaseSpanContext{
					Instance:  s.Context.Database.Instance,
					Statement: s.Context.Database.Statement,
					Type:      s.Context.Database.Type,
					User:      s.Context.Database.User,
				})
			}
			span.Duration = time.Duration(s.Duration * float64(time.Millisecond))
			span.End()
		}
		tx.Duration = time.Duration(t.Duration * float64(time.Millisecond))
		tx.End()
	}

	for _, e := range payload.Errors {
		if e == nil {
			continue
		}
		var err *Error
		if e.Log.Message != "" {
			err = tracer.NewErrorLog(ErrorLogRecord{
				Message:       e.Log.Message,
				MessageFormat: e.Log.ParamMessage,
				Level:         e.Log.Level,
				LoggerName:    e.Log.LoggerName,
			})
		} else {
			ee := exceptionError{e.Exception}
			if e.Exception.Code.String != "" {
				err = tracer.NewError(stringCodeException{ee})
			} else {
				err = tracer.NewError(float64CodeException{ee})
			}
		}
		if setContext(e.Context, &err.Context) != nil {
			return 0
		}
		err.Culprit = e.Culprit
		err.Timestamp = time.Time(e.Timestamp)
		err.Send()
	}

	return 0
}

type float64CodeException struct {
	exceptionError
}

func (e float64CodeException) Code() float64 {
	return e.x.Code.Number
}

type stringCodeException struct {
	exceptionError
}

func (e stringCodeException) Code() string {
	return e.x.Code.String
}

type exceptionError struct {
	x model.Exception
}

func (e exceptionError) Type() string {
	return e.x.Type
}

func (e exceptionError) Error() string {
	return e.x.Message
}

func (e exceptionError) StackTrace() []stacktrace.Frame {
	if len(e.x.Stacktrace) == 0 {
		return nil
	}
	frames := make([]stacktrace.Frame, len(e.x.Stacktrace))
	for i, f := range e.x.Stacktrace {
		frames[i].Function = f.Function
		frames[i].File = f.File
		frames[i].Line = f.Line
	}
	return frames
}

type gofuzzTransport struct {
	writer fastjson.Writer
}

func (t *gofuzzTransport) SendErrors(ctx context.Context, payload *model.ErrorsPayload) error {
	t.writer.Reset()
	payload.MarshalFastJSON(&t.writer)
	t.process(error_processor.NewProcessor())
	return nil
}

func (t *gofuzzTransport) SendMetrics(ctx context.Context, payload *model.MetricsPayload) error {
	t.writer.Reset()
	payload.MarshalFastJSON(&t.writer)
	t.process(metric_processor.NewProcessor())
	return nil
}

func (t *gofuzzTransport) SendTransactions(ctx context.Context, payload *model.TransactionsPayload) error {
	t.writer.Reset()
	payload.MarshalFastJSON(&t.writer)
	t.process(transaction_processor.NewProcessor())
	return nil
}

func (t *gofuzzTransport) process(p processor.Processor) {
	raw := make(map[string]interface{})
	if err := json.Unmarshal(t.writer.Bytes(), &raw); err != nil {
		panic(err)
	}
	if err := p.Validate(raw); err != nil {
		panic(err)
	}
	if _, err := p.Decode(raw); err != nil {
		panic(err)
	}
}
