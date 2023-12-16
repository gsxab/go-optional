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
	OrLazyOpt(func() Optional[T]) Optional[T]
	OrLazyOptE(func() (Optional[T], error)) (Optional[T], error)
	OrLazyOptC(context.Context, func(context.Context) Optional[T]) Optional[T]
	OrLazyOptCE(context.Context, func(context.Context) (Optional[T], error)) (Optional[T], error)

	OrZero() T
	OrValue(T) T
	OrLazy(func() T) T
	OrLazyE(func() (T, error)) (T, error)
	OrLazyC(context.Context, func(context.Context) T) T
	OrLazyCE(context.Context, func(context.Context) (T, error)) (T, error)

	Ptr() *T
	PtrOrLazyPtr(func() *T) *T
	PtrOrLazyPtrE(func() (*T, error)) (*T, error)
	PtrOrLazyPtrC(context.Context, func(context.Context) *T) *T
	PtrOrLazyPtrCE(context.Context, func(context.Context) (*T, error)) (*T, error)
}
