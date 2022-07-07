package adapter

import (
	"context"

	M "github.com/sagernet/sing/common/metadata"
)

type Inbound interface {
	Service
	Type() string
	Tag() string
}

type InboundContext struct {
	Inbound     string
	Network     string
	Source      M.Socksaddr
	Destination M.Socksaddr
	Domain      string
	Protocol    string
	Outbound    string

	// cache

	SniffEnabled             bool
	SniffOverrideDestination bool

	SourceGeoIPCode string
	GeoIPCode       string
}

type inboundContextKey struct{}

func WithContext(ctx context.Context, inboundContext *InboundContext) context.Context {
	return context.WithValue(ctx, (*inboundContextKey)(nil), inboundContext)
}

func ContextFrom(ctx context.Context) *InboundContext {
	metadata := ctx.Value((*inboundContextKey)(nil))
	if metadata == nil {
		return nil
	}
	return metadata.(*InboundContext)
}

func AppendContext(ctx context.Context) (context.Context, *InboundContext) {
	metadata := ContextFrom(ctx)
	if metadata != nil {
		return ctx, metadata
	}
	metadata = new(InboundContext)
	return WithContext(ctx, metadata), nil
}
