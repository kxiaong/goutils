package tracing

import (
	"context"

	common "code.byted.org/bytedtrace/bytedtrace-common/go/const"
	b "code.byted.org/bytedtrace/interface-go"
	"code.byted.org/gopkg/logid"
)

// PropagateContext : 繁衍一个新的context: 从原始请求context中提取上下文信息，写入新的context
func PropagateContext(ctx context.Context, spanType, spanName string) (context.Context, b.Span) {
	span := b.GetSpanFromContext(ctx)
	newSpan, newCtx := b.StartCustomSpan(ctx, spanType, spanName, b.FollowsFrom(span.GetContext()))
	return newCtx, newSpan
}

func NewServerSpan(name string, tagKV map[string]interface{}) (b.Span, context.Context) {
	ctx := context.WithValue(context.Background(), common.LogIdKey, logid.GenLogID())
	root, ctx := b.StartServerSpan(ctx, name)
	for k, v := range tagKV {
		root.SetTag(k, v)
	}
	return root, ctx
}

func CtxAddKV(ctx context.Context, k string, v interface{}) context.Context {
	span := b.GetSpanFromContext(ctx)
	if span != nil {
		span.SetTag(k, v)
	}
	return ctx
}

func CtxAttachLogID(ctx context.Context, logID string) context.Context {
	newCtx := context.WithValue(ctx, common.LogIdKey, logID)
	return newCtx
}
