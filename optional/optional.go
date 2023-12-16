/*
 * SPDX-License-Identifier: MIT
 *
 * Copyright (c) 2023 Gsxab
 *
 * Permission is hereby granted, free of charge, to any person obtaining a copy of
 * this software and associated documentation files (the "Software"), to deal in
 * the Software without restriction, including without limitation the rights to
 * use, copy, modify, merge, publish, distribute, sublicense, and/or sell copies of
 * the Software, and to permit persons to whom the Software is furnished to do so,
 * subject to the following conditions:
 *
 * The above copyright notice and this permission notice shall be included in all
 * copies or substantial portions of the Software.
 *
 * THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
 * IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY, FITNESS
 * FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE AUTHORS OR
 * COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER LIABILITY, WHETHER
 * IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM, OUT OF OR IN
 * CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN THE SOFTWARE.
 */

package optional

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"

	"github.com/gsxab/go-optional/abstract"
)

type Optional[T any] struct {
	value *T
}

// Constructors.

func New[T any](ptr *T) Optional[T] {
	if ptr == nil {
		return Optional[T]{}
	}
	copy := *ptr
	return Optional[T]{value: &copy}
}

func NewNoCopy[T any](ptr *T) Optional[T] {
	return Optional[T]{value: ptr}
}

func Empty[T any]() Optional[T] {
	return Optional[T]{}
}

func NewFromValue[T any](v T) Optional[T] {
	return Optional[T]{value: &v}
}

func NewFromPair[T any](v T, ok bool) Optional[T] {
	if !ok {
		return Optional[T]{}
	}
	return Optional[T]{value: &v}
}

func NewFromErrorPair[T any](v T, err error) Optional[T] {
	if err != nil {
		return Optional[T]{}
	}
	return Optional[T]{value: &v}
}

// Value getters.

func (o Optional[T]) HasValue() bool {
	return o.value != nil
}

func (o Optional[T]) IsEmpty() bool {
	return o.value == nil
}

func (o Optional[T]) RequireValue() {
	if o.IsEmpty() {
		panic("optional is required")
	}
}

func (o Optional[T]) MustValue() T {
	o.RequireValue()
	return *o.value
}

func (o Optional[T]) ValueNoCheck() T {
	return *o.value
}

func (o Optional[T]) Value() (T, bool) {
	if o.HasValue() {
		return *o.value, true
	}
	var zero T
	return zero, false
}

func (o Optional[T]) ValueOk() (T, bool) {
	return o.Value()
}

func (o Optional[T]) Ptr() *T {
	return o.value
}

// Defaulted optional getters.

func (o Optional[T]) Or(v abstract.Optional[T]) abstract.Optional[T] {
	if o.HasValue() {
		return o
	}
	return v
}

func (o Optional[T]) OrLazyOpt(gen func() abstract.Optional[T]) abstract.Optional[T] {
	if o.HasValue() {
		return o
	}
	return gen()
}

func (o Optional[T]) OrLazyOptE(gen func() (abstract.Optional[T], error)) (abstract.Optional[T], error) {
	if o.HasValue() {
		return o, nil
	}
	return gen()
}

func (o Optional[T]) OrLazyOptC(ctx context.Context, gen func(context.Context) abstract.Optional[T]) abstract.Optional[T] {
	if o.HasValue() {
		return o
	}
	return gen(ctx)
}

func (o Optional[T]) OrLazyOptCE(ctx context.Context, gen func(context.Context) (abstract.Optional[T], error)) (abstract.Optional[T], error) {
	if o.HasValue() {
		return o, nil
	}
	return gen(ctx)
}

// Defaulted value getters.

func (o Optional[T]) OrZero() T {
	if o.HasValue() {
		return *o.value
	}
	var zero T
	return zero
}

func (o Optional[T]) OrValue(v T) T {
	if o.HasValue() {
		return *o.value
	}
	return v
}

func (o Optional[T]) OrLazy(gen func() T) T {
	if o.HasValue() {
		return *o.value
	}
	return gen()
}

func (o Optional[T]) OrLazyE(gen func() (T, error)) (T, error) {
	if o.HasValue() {
		return *o.value, nil
	}
	return gen()
}

func (o Optional[T]) OrLazyC(ctx context.Context, gen func(context.Context) T) T {
	if o.HasValue() {
		return *o.value
	}
	return gen(ctx)
}

func (o Optional[T]) OrLazyCE(ctx context.Context, gen func(context.Context) (T, error)) (T, error) {
	if o.HasValue() {
		return *o.value, nil
	}
	return gen(ctx)
}

// Defaulted pointer getters.

func (o Optional[T]) PtrOrPtr(v *T) *T {
	if o.HasValue() {
		return o.value
	}
	return v
}

func (o Optional[T]) PtrOrLazyPtr(gen func() *T) *T {
	if o.HasValue() {
		return o.value
	}
	return gen()
}

func (o Optional[T]) PtrOrLazyPtrE(gen func() (*T, error)) (*T, error) {
	if o.HasValue() {
		return o.value, nil
	}
	return gen()
}

func (o Optional[T]) PtrOrLazyPtrC(ctx context.Context, gen func(context.Context) *T) *T {
	if o.HasValue() {
		return o.value
	}
	return gen(ctx)
}

func (o Optional[T]) PtrOrLazyPtrCE(ctx context.Context, gen func(context.Context) (*T, error)) (*T, error) {
	if o.HasValue() {
		return o.value, nil
	}
	return gen(ctx)
}

// If present / absent.

func (o Optional[T]) Foreach(callback func(T)) {
	if o.HasValue() {
		callback(*o.value)
	}
}

func (o Optional[T]) ForeachE(callback func(T) error) error {
	if o.HasValue() {
		return callback(*o.value)
	}
	return nil
}

func (o Optional[T]) ForeachC(ctx context.Context, callback func(context.Context, T)) {
	if o.HasValue() {
		callback(ctx, *o.value)
	}
}

func (o Optional[T]) ForeachCE(ctx context.Context, callback func(context.Context, T) error) error {
	if o.HasValue() {
		return callback(ctx, *o.value)
	}
	return nil
}

func (o Optional[T]) ForeachPtr(callback func(*T)) {
	if o.HasValue() {
		callback(o.value)
	}
}

func (o Optional[T]) ForeachPtrE(callback func(*T) error) error {
	if o.HasValue() {
		return callback(o.value)
	}
	return nil
}

func (o Optional[T]) ForeachPtrC(ctx context.Context, callback func(context.Context, *T)) {
	if o.HasValue() {
		callback(ctx, o.value)
	}
}

func (o Optional[T]) ForeachPtrCE(ctx context.Context, callback func(context.Context, *T) error) error {
	if o.HasValue() {
		return callback(ctx, o.value)
	}
	return nil
}

// Extra methods.

// Stringer.

func (o Optional[T]) String() string {
	if o.HasValue() {
		v := any(*o.value)
		if stringer, ok := v.(fmt.Stringer); ok {
			return fmt.Sprintf("Optional[%s]", stringer)
		}
		return fmt.Sprintf("Optional[%s]", v)
	}
	return "Optional[]"
}

// JSONMarshaller

func (o Optional[T]) MarshalJSON() ([]byte, error) {
	if o.IsEmpty() {
		return []byte("null"), nil
	}

	marshal, err := json.Marshal(*o.value)
	if err != nil {
		return nil, err
	}
	return marshal, nil
}

func (o *Optional[T]) UnmarshalJSON(data []byte) error {
	if len(data) <= 0 || bytes.Equal(data, []byte("null")) {
		o.value = nil
		return nil
	}

	err := json.Unmarshal(data, o.value)
	if err != nil {
		return err
	}

	return nil
}

// Container.
