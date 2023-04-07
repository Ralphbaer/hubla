package http

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"net/http"
	"strconv"
	"strings"
	"time"

	"github.com/Ralphbaer/hubla/backend/common/hlog"
)

// RequestInfo is a struct design to store http access log data
type RequestInfo struct {
	Method        string
	Username      string
	URI           string
	Referer       string
	RemoteAddress string
	Status        int
	Date          time.Time
	Duration      time.Duration
	UserAgent     string
	CorrelationID string
	Protocol      string
	Size          int
	Body          string
}

// ResponseMetricsWrapper is a Wrapper responsible for collect the response data such as status code and size
// It implements built-in ResponseWriter interface.
type ResponseMetricsWrapper struct {
	http.ResponseWriter
	StatusCode int
	Size       int
	Body       string
}

// WriteHeader is built-in ResponseWriter.WriteHeader interface implementation
func (rec *ResponseMetricsWrapper) WriteHeader(code int) {
	rec.StatusCode = code
	rec.ResponseWriter.WriteHeader(code)
}

// Write is built-in ResponseWriter.Write interface implementation
func (rec *ResponseMetricsWrapper) Write(bytes []byte) (int, error) {
	if rec.StatusCode == 0 {
		rec.StatusCode = http.StatusOK
	}

	rec.Body = string(bytes)

	size, err := rec.ResponseWriter.Write(bytes)
	rec.Size += size
	return size, err
}

// NewRequestInfo creates an instance of RequestInfo
func NewRequestInfo(w http.ResponseWriter, r *http.Request) *RequestInfo {
	username, referer := "-", "-"

	if r.URL.User != nil {
		if name := r.URL.User.Username(); name != "" {
			username = name
		}
	}

	if r.Header.Get("Referer") != "" {
		referer = r.Header.Get("Referer")
	}

	body := ""

	if r.ContentLength > 0 {
		b, err := ioutil.ReadAll(r.Body)

		if err == nil {
			body = string(b)
			r.Body = ioutil.NopCloser(bytes.NewBuffer(b))
		}
	}

	return &RequestInfo{
		Method:        r.Method,
		URI:           r.URL.String(),
		Username:      username,
		Referer:       referer,
		UserAgent:     r.Header.Get(headerUserAgent),
		CorrelationID: w.Header().Get(headerCorrelationID),
		RemoteAddress: GetRemoteAddress(r),
		Protocol:      r.Proto,
		Date:          time.Now().UTC(),
		Body:          body,
	}
}

// CLFString produces a log entry format similar to Common Log Format (CLF)
// Ref: https://httpd.apache.org/docs/trunk/logs.html#common
func (r *RequestInfo) CLFString() string {
	return fmt.Sprint(strings.Join([]string{
		r.RemoteAddress,
		"-",
		r.Username,
		`"` + r.Method,
		r.URI,
		`"` + r.Protocol,
		strconv.Itoa(r.Status),
		strconv.Itoa(r.Size),
		r.Referer,
		r.UserAgent,
	}, " "))
}

// String implements fmt.Stringer interface and produces a log entry using RequestInfo.CLFExtendedString.
func (r *RequestInfo) String() string {
	return r.CLFString()
}

func (r *RequestInfo) debugRequestString() string {
	return strings.Join([]string{
		r.CLFString(),
		r.Referer,
		r.UserAgent,
		r.CorrelationID,
		r.Body,
	}, " ")
}

func (r *RequestInfo) debugResponseString(w *ResponseMetricsWrapper) string {
	return strings.Join([]string{
		r.CLFString(),
		r.Referer,
		r.UserAgent,
		r.CorrelationID,
		w.Body,
	}, " ")
}

// FinishRequestInfo calculates the duration of RequestInfo automatically using time.Now()
// It also set StatusCode and Size of RequestInfo passed by ResponseMetricsWrapper
func (r *RequestInfo) FinishRequestInfo(rw *ResponseMetricsWrapper) {
	r.Duration = time.Now().UTC().Sub(r.Date)
	r.Status = rw.StatusCode
	r.Size = rw.Size
}

type logMiddleware struct {
	Logger hlog.Logger
}

// LogMiddlewareOption represents the log middleware function as an implementation
type LogMiddlewareOption func(l *logMiddleware)

// WithLogger is a functional option for logMiddleware
func WithLogger(logger hlog.Logger) LogMiddlewareOption {
	return func(l *logMiddleware) {
		l.Logger = logger
	}
}

// buildOpts creates an instance of logMiddleware with options
func buildOpts(opts ...LogMiddlewareOption) *logMiddleware {
	mid := &logMiddleware{
		Logger: &hlog.GoLogger{},
	}

	for _, opt := range opts {
		opt(mid)
	}

	return mid
}

// WithLog is a middleware to log access to http server.
// It logs access log according to Apache Standard Logs which uses Common Log Format (CLF)
// Ref: https://httpd.apache.org/docs/trunk/logs.html#common
func WithLog(opts ...LogMiddlewareOption) func(next http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			info := NewRequestInfo(w, r)

			mid := buildOpts(opts...)

			logger := mid.Logger.WithFields(map[string]interface{}{
				headerCorrelationID: info.CorrelationID,
			})

			rw := ResponseMetricsWrapper{w, 200, 0, ""}

			logger.Debug(info.debugRequestString())

			ctx := hlog.ContextWithLogger(r.Context(), logger)

			next.ServeHTTP(w, r.WithContext(ctx))

			info.FinishRequestInfo(&rw)

			logger.Debug(info.debugResponseString(&rw))
			logger.Infoln(info)
		})
	}
}
