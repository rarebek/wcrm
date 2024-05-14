package middleware

import (
	"bufio"
	"errors"
	"net"
	"net/http"

	"github.com/gin-gonic/gin"
	"go.opentelemetry.io/otel"
	"go.opentelemetry.io/otel/attribute"
	"go.opentelemetry.io/otel/trace"
)

// Tracing middleware function
func Tracing(c *gin.Context) {
	// Create a new context with tracing span
	ctx := c.Request.Context()
	tracer := otel.GetTracerProvider().Tracer("your-service-name")
	ctx, span := tracer.Start(ctx, "HTTP "+c.Request.Method+" "+c.FullPath(), trace.WithSpanKind(trace.SpanKindServer))
	defer span.End()

	// Add TraceID to response header
	c.Writer.Header().Add("TraceID", span.SpanContext().TraceID().String())

	// Create a custom response writer to capture status code
	rw := &responseWriter{c.Writer, http.StatusOK}

	// Serve the request with the modified response writer and context
	c.Writer = rw
	c.Request = c.Request.WithContext(ctx)

	// Call the next handler
	c.Next()

	// Add attributes to the span
	span.SetAttributes(
		attribute.String("http.method", c.Request.Method),
		attribute.String("http.url", c.FullPath()),
		attribute.Int("http.status_code", rw.statusCode),
	)
}

// Custom responseWriter to capture status code
type responseWriter struct {
	gin.ResponseWriter
	statusCode int
}

// WriteHeader method implements the http.ResponseWriter interface
func (rw *responseWriter) WriteHeader(statusCode int) {
	rw.statusCode = statusCode
	rw.ResponseWriter.WriteHeader(statusCode)
}

// Write method implements the http.ResponseWriter interface
func (rw *responseWriter) Write(data []byte) (int, error) {
	return rw.ResponseWriter.Write(data)
}

// WriteString method implements the gin.ResponseWriter interface
func (rw *responseWriter) WriteString(s string) (int, error) {
	return rw.ResponseWriter.WriteString(s)
}

// Hijack method implements the http.Hijacker interface (optional)
func (rw *responseWriter) Hijack() (net.Conn, *bufio.ReadWriter, error) {
	if hijacker, ok := rw.ResponseWriter.(http.Hijacker); ok {
		return hijacker.Hijack()
	}
	return nil, nil, errors.New("response writer does not support hijacking")
}

// CloseNotify method implements the http.CloseNotifier interface (optional)
func (rw *responseWriter) CloseNotify() <-chan bool {
	if notifier, ok := rw.ResponseWriter.(http.CloseNotifier); ok {
		return notifier.CloseNotify()
	}
	// Return a dummy channel if CloseNotify is not supported
	return make(chan bool)
}
