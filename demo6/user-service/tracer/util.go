package tracer

import (
	"github.com/gin-gonic/gin"
	"github.com/opentracing/opentracing-go"
	"net/http"
)

const SPAN_CTX_KEY = "__SPAN_CTX__"

func ExtractSpanContext(c *gin.Context) {
	tracer := opentracing.GlobalTracer()
	spanCtx, _ := tracer.Extract(opentracing.HTTPHeaders, opentracing.HTTPHeadersCarrier(c.Request.Header))
	if spanCtx != nil {
		c.Set(SPAN_CTX_KEY, spanCtx)
	}
}

// propagate span context information when call other service
func InjectSpanContext(c *gin.Context, req *http.Request) {
	if v, existed := c.Get(SPAN_CTX_KEY); existed {
		if ctx, ok := v.(opentracing.SpanContext); ok {
			tracer := opentracing.GlobalTracer()
			tracer.Inject(ctx, opentracing.HTTPHeaders, opentracing.HTTPHeadersCarrier(req.Header))
		}
	}
}
