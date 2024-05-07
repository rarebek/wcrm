package middleware

// import (
// 	"net/http"

// 	"go.opentelemetry.io/otel/attribute"

// 	"api-gateway/api/response"
// 	"api-gateway/internal/pkg/otlp"
// )

// func Tracing(next http.Handler) http.Handler {
// 	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
// 		rw := response.NewResponseWriter(w, http.StatusOK)
// 		// tracing
// 		ctx, span := otlp.Start(r.Context(), "", r.URL.Path)
// 		// add request id to header
// 		w.Header().Add(RequestIDHeader, span.SpanContext().TraceID().String())
// 		next.ServeHTTP(rw, r.WithContext(ctx))
// 		// add attributes
// 		span.SetAttributes(
// 			attribute.Key("http.method").String(r.Method),
// 			attribute.Key("http.url").String(r.URL.Path),
// 			attribute.Key("http.status_code").Int(rw.StatusCode()),
// 		)
// 		// end completes the span
// 		span.End()
// 	})
// }
