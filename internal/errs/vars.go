package errs

import "errors"

var (
	// ErrEmptyCtx happens when it's not possible to extract something from value because it's not filled
	ErrEmptyCtx = errors.New("empty context")
	// ErrInvalidCtxValue happens when the context.Context.Value returns something not castable to the what it should be
	ErrInvalidCtxValue = errors.New("invalid context")
)
