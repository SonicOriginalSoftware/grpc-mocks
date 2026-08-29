//revive:disable:package-comments
package clientconn

import (
	"context"

	"google.golang.org/grpc"
)

// Invoker fills in the reply of a recorded unary RPC. Tests supply one so the
// mock never needs to know which service it is standing in for.
type Invoker func(ctx context.Context, method string, args, reply any) error

// Mock implements grpc.ClientConnInterface for testing.
type Mock struct {
	grpc.ClientConnInterface

	// InvokeFn fills in the reply. When nil, Invoke records the call and
	// leaves the reply untouched.
	InvokeFn Invoker
	// Err is returned by Invoke when set, in preference to calling InvokeFn.
	Err error
	// Ctx is the context received by Invoke.
	Ctx context.Context
	// Method is the full RPC method received by Invoke.
	Method string
	// Args is the request received by Invoke.
	Args any
	// Calls tracks how many unary RPCs were invoked.
	Calls int
}

// Invoke records the unary RPC, then reports the configured error or hands the
// reply to InvokeFn.
func (c *Mock) Invoke(
	ctx context.Context, method string, args, reply any, _ ...grpc.CallOption,
) error {
	c.Calls++
	c.Ctx = ctx
	c.Method = method
	c.Args = args

	if err := ctx.Err(); err != nil {
		return err
	}
	if c.Err != nil {
		return c.Err
	}
	if c.InvokeFn == nil {
		return nil
	}

	return c.InvokeFn(ctx, method, args, reply)
}
