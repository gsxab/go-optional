package abstract

import "context"

type Optional[T any] interface {
	HasValue() bool
	IsEmpty() bool
	RequireValue()
	MustValue() T
	ValueNoCheck() T
	ValueOk() (T, bool)

	Or(v Optional[T]) Optional[T]
	OrLazy(func() Optional[T]) Optional[T]
	OrLazyWithErr(func() (Optional[T], error)) (Optional[T], error)
	OrLazyWithCtx(context.Context, func(context.Context) Optional[T]) Optional[T]
	OrLazyWithCtxErr(context.Context, func(context.Context) (Optional[T], error)) (Optional[T], error)

	ValueOrZero() T
	ValueOrValue(T) T
	ValueOrLazy(func() T) T
	ValueOrLazyWithErr(func() (T, error)) (T, error)
	ValueOrLazyWithCtx(context.Context, func(context.Context) T) T
	ValueOrLazyWithCtxErr(context.Context, func(context.Context) (T, error)) (T, error)

	Ptr() *T
	PtrOrLazyPtr(func() *T) *T
	PtrOrLazyPtrWithErr(func() (*T, error)) (*T, error)
	PtrOrLazyPtrWithCtx(context.Context, func(context.Context) *T) *T
	PtrOrLazyPtrWithCtxErr(context.Context, func(context.Context) (*T, error)) (*T, error)
}
