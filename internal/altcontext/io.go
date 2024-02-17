package altcontext

import (
	"context"

	"github.com/codeaspiras/gotravel-http/internal/errs"
	"github.com/codeaspiras/gotravel/stdio"
)

type ioCtx uint

const ioCtxKey ioCtx = iota

func CtxWithIO(ctx context.Context, io stdio.IO) context.Context {
	return context.WithValue(ctx, ioCtxKey, io)
}

func GetIO(ctx context.Context) (stdio.IO, error) {
	v := ctx.Value(ioCtxKey)
	if v == nil {
		return nil, errs.ErrEmptyCtx
	}

	if io, castable := v.(stdio.IO); castable {
		return io, nil
	}

	return nil, errs.ErrInvalidCtxValue
}
