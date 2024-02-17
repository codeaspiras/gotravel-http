package altcontext

import (
	"context"

	"github.com/codeaspiras/gotravel-http/internal/errs"
	"github.com/codeaspiras/gotravel-http/internal/validator"
)

type validatorCtx uint

const validatorCtxKey validatorCtx = iota

func CtxWithValidator(ctx context.Context, v validator.Validator) context.Context {
	return context.WithValue(ctx, validatorCtxKey, v)
}

func GetValidator(ctx context.Context) (validator.Validator, error) {
	v := ctx.Value(validatorCtxKey)
	if v == nil {
		return nil, errs.ErrEmptyCtx
	}

	if rv, castable := v.(validator.Validator); castable {
		return rv, nil
	}

	return nil, errs.ErrInvalidCtxValue
}
