// Copyright (c) 2024 Uber Technologies, Inc.
//
// Permission is hereby granted, free of charge, to any person obtaining a copy
// of this software and associated documentation files (the "Software"), to deal
// in the Software without restriction, including without limitation the rights
// to use, copy, modify, merge, publish, distribute, sublicense, and/or sell
// copies of the Software, and to permit persons to whom the Software is
// furnished to do so, subject to the following conditions:
//
// The above copyright notice and this permission notice shall be included in
// all copies or substantial portions of the Software.
//
// THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
// IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
// FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
// AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
// LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
// OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN
// THE SOFTWARE.

package interceptor

import (
	"context"
	"go.uber.org/yarpc/api/middleware"
	"go.uber.org/yarpc/api/transport"
)

type (
	// UnaryOutbound defines transport interceptor for `UnaryOutbound`s.
	//
	// UnaryOutbound interceptor MAY do zero or more of the following: change the
	// context, change the request, change the returned response, handle the
	// returned error, call the given outbound zero or more times.
	//
	// UnaryOutbound interceptor MUST always return a non-nil Response or error,
	// and they MUST be thread-safe
	//
	// UnaryOutbound interceptor is re-used across requests and MAY be called
	// multiple times on the same request.
	UnaryOutbound = middleware.UnaryOutbound

	// OnewayOutbound defines transport interceptor for `OnewayOutbound`s.
	//
	// OnewayOutbound interceptor MAY do zero or more of the following: change the
	// context, change the request, change the returned ack, handle the returned
	// error, call the given outbound zero or more times.
	//
	// OnewayOutbound interceptor MUST always return an Ack (nil or not) or an
	// error, and they MUST be thread-safe.
	//
	// OnewayOutbound interceptor is re-used across requests and MAY be called
	// multiple times on the same request.
	OnewayOutbound = middleware.OnewayOutbound

	// StreamOutbound defines transport interceptor for `StreamOutbound`s.
	//
	// StreamOutbound interceptor MAY do zero or more of the following: change the
	// context, change the requestMeta, change the returned Stream, handle the
	// returned error, call the given outbound zero or more times.
	//
	// StreamOutbound interceptor MUST always return a non-nil Stream or error,
	// and they MUST be thread-safe
	//
	// StreamOutbound interceptors is re-used across requests and MAY be called
	// multiple times on the same request.
	StreamOutbound = middleware.StreamOutbound
)

var (
	// NopUnaryOutbound is a no-operation unary outbound middleware.
	NopUnaryOutbound transport.UnaryOutbound = nopUnaryOutbound{}
)

type nopUnaryOutbound struct{}

// Call processes a unary request and returns a nil response and no error.
func (nopUnaryOutbound) Call(ctx context.Context, req *transport.Request) (*transport.Response, error) {
	return nil, nil
}

// Start starts the outbound middleware. It is a no-op.
func (nopUnaryOutbound) Start() error {
	return nil
}

// Stop stops the outbound middleware. It is a no-op.
func (nopUnaryOutbound) Stop() error {
	return nil
}

// IsRunning checks if the outbound middleware is running. Always returns false.
func (nopUnaryOutbound) IsRunning() bool {
	return false
}

// Transports returns the transports associated with this middleware. Always returns nil.
func (nopUnaryOutbound) Transports() []transport.Transport {
	return nil
}

// UnaryOutboundFunc adapts a function into a UnaryOutbound middleware.
type UnaryOutboundFunc func(ctx context.Context, req *transport.Request) (*transport.Response, error)

// Call executes the function as a UnaryOutbound call.
func (f UnaryOutboundFunc) Call(ctx context.Context, req *transport.Request) (*transport.Response, error) {
	return f(ctx, req)
}

// Start starts the UnaryOutboundFunc middleware. It is a no-op.
func (f UnaryOutboundFunc) Start() error {
	return nil
}

// Stop stops the UnaryOutboundFunc middleware. It is a no-op.
func (f UnaryOutboundFunc) Stop() error {
	return nil
}

// IsRunning checks if the UnaryOutboundFunc middleware is running. Always returns false.
func (f UnaryOutboundFunc) IsRunning() bool {
	return false
}

// Transports returns the transports associated with this middleware. Always returns nil.
func (f UnaryOutboundFunc) Transports() []transport.Transport {
	return nil
}