package routing

import (
	"context"

	"github.com/gptlocal/netools/common/net"
	"github.com/gptlocal/netools/features"
	"github.com/gptlocal/netools/transport"
)

// Dispatcher is a feature that dispatches inbound requests to outbound handlers based on rules.
// Dispatcher is required to be registered in a Xray instance to make Xray function properly.
//
// netool:api:stable
type Dispatcher interface {
	features.Feature

	// Dispatch returns a Ray for transporting data for the given request.
	Dispatch(ctx context.Context, dest net.Destination) (*transport.Link, error)
	DispatchLink(ctx context.Context, dest net.Destination, link *transport.Link) error
}

// DispatcherType returns the type of Dispatcher interface. Can be used to implement common.HasType.
//
// netool:api:stable
func DispatcherType() interface{} {
	return (*Dispatcher)(nil)
}
