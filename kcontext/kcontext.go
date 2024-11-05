package kcontext

import (
	"context"

	"github.com/Ccheers/protoc-gen-go-kratos-http/audit"
	khttp "github.com/go-kratos/kratos/v2/transport/http"
)

type khttpContextKey struct{}

func SetKHTTPContextWithContext(ctx context.Context, kctx khttp.Context) context.Context {
	return context.WithValue(ctx, khttpContextKey{}, kctx)
}

func GetKHTTPContextWithContext(ctx context.Context) (khttp.Context, bool) {
	v, ok := ctx.Value(khttpContextKey{}).(khttp.Context)
	return v, ok
}

type khttpAuditContextKey struct{}

func SetKHTTPAuditContextWithContext(ctx context.Context, audit *audit.Audit) context.Context {
	return context.WithValue(ctx, khttpAuditContextKey{}, audit)
}

func GetKHTTPAuditContextWithContext(ctx context.Context) (*audit.Audit, bool) {
	v, ok := ctx.Value(khttpAuditContextKey{}).(*audit.Audit)
	return v, ok
}
