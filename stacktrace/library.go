// Code generated by "go generate". DO NOT EDIT.

package stacktrace

import (
	"strings"

	"github.com/elastic/apm-agent-go/internal/radix"
)

var libraryPackages = radix.NewFromMap(map[string]interface{}{
	"archive/tar":                       true,
	"archive/zip":                       true,
	"bufio":                             true,
	"bytes":                             true,
	"compress/bzip2":                    true,
	"compress/flate":                    true,
	"compress/gzip":                     true,
	"compress/lzw":                      true,
	"compress/zlib":                     true,
	"container/heap":                    true,
	"container/list":                    true,
	"container/ring":                    true,
	"context":                           true,
	"crypto":                            true,
	"crypto/aes":                        true,
	"crypto/cipher":                     true,
	"crypto/des":                        true,
	"crypto/dsa":                        true,
	"crypto/ecdsa":                      true,
	"crypto/elliptic":                   true,
	"crypto/hmac":                       true,
	"crypto/internal/cipherhw":          true,
	"crypto/md5":                        true,
	"crypto/rand":                       true,
	"crypto/rc4":                        true,
	"crypto/rsa":                        true,
	"crypto/sha1":                       true,
	"crypto/sha256":                     true,
	"crypto/sha512":                     true,
	"crypto/subtle":                     true,
	"crypto/tls":                        true,
	"crypto/x509":                       true,
	"crypto/x509/pkix":                  true,
	"database/sql":                      true,
	"database/sql/driver":               true,
	"debug/dwarf":                       true,
	"debug/elf":                         true,
	"debug/gosym":                       true,
	"debug/macho":                       true,
	"debug/pe":                          true,
	"debug/plan9obj":                    true,
	"encoding":                          true,
	"encoding/ascii85":                  true,
	"encoding/asn1":                     true,
	"encoding/base32":                   true,
	"encoding/base64":                   true,
	"encoding/binary":                   true,
	"encoding/csv":                      true,
	"encoding/gob":                      true,
	"encoding/hex":                      true,
	"encoding/json":                     true,
	"encoding/pem":                      true,
	"encoding/xml":                      true,
	"errors":                            true,
	"expvar":                            true,
	"flag":                              true,
	"fmt":                               true,
	"go/ast":                            true,
	"go/build":                          true,
	"go/constant":                       true,
	"go/doc":                            true,
	"go/format":                         true,
	"go/importer":                       true,
	"go/internal/gccgoimporter":         true,
	"go/internal/gcimporter":            true,
	"go/internal/srcimporter":           true,
	"go/parser":                         true,
	"go/printer":                        true,
	"go/scanner":                        true,
	"go/token":                          true,
	"go/types":                          true,
	"hash":                              true,
	"hash/adler32":                      true,
	"hash/crc32":                        true,
	"hash/crc64":                        true,
	"hash/fnv":                          true,
	"html":                              true,
	"html/template":                     true,
	"image":                             true,
	"image/color":                       true,
	"image/color/palette":               true,
	"image/draw":                        true,
	"image/gif":                         true,
	"image/internal/imageutil":          true,
	"image/jpeg":                        true,
	"image/png":                         true,
	"index/suffixarray":                 true,
	"internal/cpu":                      true,
	"internal/nettrace":                 true,
	"internal/poll":                     true,
	"internal/race":                     true,
	"internal/singleflight":             true,
	"internal/syscall/unix":             true,
	"internal/syscall/windows":          true,
	"internal/syscall/windows/registry": true,
	"internal/syscall/windows/sysdll":   true,
	"internal/testenv":                  true,
	"internal/testlog":                  true,
	"internal/trace":                    true,
	"io":                                true,
	"io/ioutil":                         true,
	"log":                               true,
	"log/syslog":                        true,
	"math":                              true,
	"math/big":                          true,
	"math/bits":                         true,
	"math/cmplx":                        true,
	"math/rand":                         true,
	"mime":                              true,
	"mime/multipart":                    true,
	"mime/quotedprintable":              true,
	"net":                            true,
	"net/http":                       true,
	"net/http/cgi":                   true,
	"net/http/cookiejar":             true,
	"net/http/fcgi":                  true,
	"net/http/httptest":              true,
	"net/http/httptrace":             true,
	"net/http/httputil":              true,
	"net/http/internal":              true,
	"net/http/pprof":                 true,
	"net/internal/socktest":          true,
	"net/mail":                       true,
	"net/rpc":                        true,
	"net/rpc/jsonrpc":                true,
	"net/smtp":                       true,
	"net/textproto":                  true,
	"net/url":                        true,
	"os":                             true,
	"os/exec":                        true,
	"os/signal":                      true,
	"os/signal/internal/pty":         true,
	"os/user":                        true,
	"path":                           true,
	"path/filepath":                  true,
	"plugin":                         true,
	"reflect":                        true,
	"regexp":                         true,
	"regexp/syntax":                  true,
	"runtime":                        true,
	"runtime/cgo":                    true,
	"runtime/debug":                  true,
	"runtime/internal/atomic":        true,
	"runtime/internal/sys":           true,
	"runtime/pprof":                  true,
	"runtime/pprof/internal/profile": true,
	"runtime/race":                   true,
	"runtime/trace":                  true,
	"sort":                           true,
	"strconv":                        true,
	"strings":                        true,
	"sync":                           true,
	"sync/atomic":                    true,
	"syscall":                        true,
	"testing":                        true,
	"testing/internal/testdeps":      true,
	"testing/iotest":                 true,
	"testing/quick":                  true,
	"text/scanner":                   true,
	"text/tabwriter":                 true,
	"text/template":                  true,
	"text/template/parse":            true,
	"time":                           true,
	"unicode":                        true,
	"unicode/utf16":                  true,
	"unicode/utf8":                   true,
	"unsafe":                         true,
	"vendor/golang_org/x/crypto/chacha20poly1305":                   true,
	"vendor/golang_org/x/crypto/chacha20poly1305/internal/chacha20": true,
	"vendor/golang_org/x/crypto/cryptobyte":                         true,
	"vendor/golang_org/x/crypto/cryptobyte/asn1":                    true,
	"vendor/golang_org/x/crypto/curve25519":                         true,
	"vendor/golang_org/x/crypto/poly1305":                           true,
	"vendor/golang_org/x/net/http2/hpack":                           true,
	"vendor/golang_org/x/net/idna":                                  true,
	"vendor/golang_org/x/net/internal/nettest":                      true,
	"vendor/golang_org/x/net/lex/httplex":                           true,
	"vendor/golang_org/x/net/nettest":                               true,
	"vendor/golang_org/x/net/proxy":                                 true,
	"vendor/golang_org/x/text/secure":                               true,
	"vendor/golang_org/x/text/secure/bidirule":                      true,
	"vendor/golang_org/x/text/transform":                            true,
	"vendor/golang_org/x/text/unicode":                              true,
	"vendor/golang_org/x/text/unicode/bidi":                         true,
	"vendor/golang_org/x/text/unicode/norm":                         true,
	"github.com/elastic/apm-agent-go":                               true,
})

// RegisterLibraryPackage registers the given packages as being
// well-known library path prefixes. This must not be called
// concurrently with any other functions or methods in this
// package; it is expected to be used by init functions.
func RegisterLibraryPackage(pkg ...string) {
	for _, pkg := range pkg {
		libraryPackages.Insert(pkg, true)
	}
}

// RegisterApplicationPackage registers the given packages as being
// an application path. This must not be called concurrently with
// any other functions or methods in this package; it is expected
// to be used by init functions.
//
// It is not typically necessary to register application paths. If
// a package does not match a registered *library* package path
// prefix, then the path is considered an application path. This
// function exists for the unusual case that an application exists
// within a library (e.g. an example program).
func RegisterApplicationPackage(pkg ...string) {
	for _, pkg := range pkg {
		libraryPackages.Insert(pkg, false)
	}
}

// IsLibraryPackage reports whether or not the given package path is
// a well-known library path (stdlib or apm-agent-go).
func IsLibraryPackage(pkg string) bool {
	if strings.HasSuffix(pkg, "_test") {
		return false
	}
	prefix, v, ok := libraryPackages.LongestPrefix(pkg)
	if !ok || v == false {
		return false
	}
	return prefix == pkg || pkg[len(prefix)] == '/'
}
