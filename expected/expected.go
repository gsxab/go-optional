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

type Expected[T any] struct {
	value *T
	err   error
}

// Constructors.

func New[T any](v T, err error) Expected[T] {
	if err != nil {
		return Expected[T]{err: err}
	}
	return Expected[T]{value: &v}
}

func NewNoCopy[T any](ptr *T) Expected[T] {
	return Expected[T]{value: ptr}
}

func NewFromError[T any](err error) Expected[T] {
	return Expected[T]{err: err}
}

func NewValue[T any](v T) Expected[T] {
	return Expected[T]{value: &v}
}

// Value getters.

func (e Expected[T]) HasValue() bool {
	return e.err == nil
}

func (e Expected[T]) IsEmpty() bool {
	return e.err != nil
}

func (e Expected[T]) RequireValue() {
	if e.IsEmpty() {
		panic("expected is required")
	}
}

func (e Expected[T]) MustValue() T {
	e.RequireValue()
	return *e.value
}

func (e Expected[T]) ValueNoCheck() T {
	return *e.value
}

func (e Expected[T]) Value() (T, error) {
	if e.HasValue() {
		return *e.value, nil
	}
	var zero T
	return zero, e.err
}

func (e Expected[T]) ValueOk() (T, bool) {
	if e.HasValue() {
		return *e.value, true
	}
	var zero T
	return zero, false
}

func (e Expected[T]) Error() error {
	return e.err
}

func (e Expected[T]) Ptr() *T {
	return e.value
}

// Defaulted optional getters.

func (e Expected[T]) Or(v abstract.Optional[T]) abstract.Optional[T] {
	if e.HasValue() {
		return e
	}
	return v
}

func (e Expected[T]) OrLazy(gen func() abstract.Optional[T]) abstract.Optional[T] {
	if e.HasValue() {
		return e
	}
	return gen()
}

func (e Expected[T]) OrLazyWithErr(gen func() (abstract.Optional[T], error)) (abstract.Optional[T], error) {
	if e.HasValue() {
		return e, nil
	}
	return gen()
}

func (e Expected[T]) OrLazyWithCtx(ctx context.Context, gen func(context.Context) abstract.Optional[T]) abstract.Optional[T] {
	if e.HasValue() {
		return e
	}
	return gen(ctx)
}

func (e Expected[T]) OrLazyWithCtxErr(ctx context.Context, gen func(context.Context) (abstract.Optional[T], error)) (abstract.Optional[T], error) {
	if e.HasValue() {
		return e, nil
	}
	return gen(ctx)
}

// Defaulted value getters.

func (e Expected[T]) ValueOrZero() T {
	if e.HasValue() {
		return *e.value
	}
	var zero T
	return zero
}

func (e Expected[T]) ValueOrValue(v T) T {
	if e.HasValue() {
		return *e.value
	}
	return v
}

func (e Expected[T]) ValueOrLazy(gen func() T) T {
	if e.HasValue() {
		return *e.value
	}
	return gen()
}

func (e Expected[T]) ValueOrLazyWithErr(gen func() (T, error)) (T, error) {
	if e.HasValue() {
		return *e.value, nil
	}
	return gen()
}

func (e Expected[T]) ValueOrLazyWithCtx(ctx context.Context, gen func(context.Context) T) T {
	if e.HasValue() {
		return *e.value
	}
	return gen(ctx)
}

func (e Expected[T]) ValueOrLazyWithCtxErr(ctx context.Context, gen func(context.Context) (T, error)) (T, error) {
	if e.HasValue() {
		return *e.value, nil
	}
	return gen(ctx)
}

// Defaulted pointer getters.

func (e Expected[T]) PtrOrPtr(v *T) *T {
	if e.HasValue() {
		return e.value
	}
	return v
}

func (e Expected[T]) PtrOrLazyPtr(gen func() *T) *T {
	if e.HasValue() {
		return e.value
	}
	return gen()
}

func (e Expected[T]) PtrOrLazyPtrWithErr(gen func() (*T, error)) (*T, error) {
	if e.HasValue() {
		return e.value, nil
	}
	return gen()
}

func (e Expected[T]) PtrOrLazyPtrWithCtx(ctx context.Context, gen func(context.Context) *T) *T {
	if e.HasValue() {
		return e.value
	}
	return gen(ctx)
}

func (e Expected[T]) PtrOrLazyPtrWithCtxErr(ctx context.Context, gen func(context.Context) (*T, error)) (*T, error) {
	if e.HasValue() {
		return e.value, nil
	}
	return gen(ctx)
}

// If present.

func (e Expected[T]) Foreach(callback func(T)) {
	if e.HasValue() {
		callback(*e.value)
	}
}

func (e Expected[T]) ForeachWithErr(callback func(T) error) error {
	if e.HasValue() {
		return callback(*e.value)
	}
	return nil
}

func (e Expected[T]) ForeachWithCtx(ctx context.Context, callback func(context.Context, T)) {
	if e.HasValue() {
		callback(ctx, *e.value)
	}
}

func (e Expected[T]) ForeachWithCtxErr(ctx context.Context, callback func(context.Context, T) error) error {
	if e.HasValue() {
		return callback(ctx, *e.value)
	}
	return nil
}

func (e Expected[T]) ForeachPtr(callback func(*T)) {
	if e.HasValue() {
		callback(e.value)
	}
}

func (e Expected[T]) ForeachPtrWithErr(callback func(*T) error) error {
	if e.HasValue() {
		return callback(e.value)
	}
	return nil
}

func (e Expected[T]) ForeachPtrWithCtx(ctx context.Context, callback func(context.Context, *T)) {
	if e.HasValue() {
		callback(ctx, e.value)
	}
}

func (e Expected[T]) ForeachPtrWithCtxErr(ctx context.Context, callback func(context.Context, *T) error) error {
	if e.HasValue() {
		return callback(ctx, e.value)
	}
	return nil
}

// Extra methods.

// Stringer.

func (e Expected[T]) String() string {
	if e.HasValue() {
		v := any(*e.value)
		if stringer, ok := v.(fmt.Stringer); ok {
			return fmt.Sprintf("Expected[%s]", stringer)
		}
		return fmt.Sprintf("Expected[%s]", v)
	}
	return fmt.Sprintf("Expected[Unexpected, %s]", e.err.Error())
}

// JSONMarshaller

func (e Expected[T]) MarshalJSON() ([]byte, error) {
	if e.IsEmpty() {
		return []byte("null"), nil
	}

	marshal, err := json.Marshal(*e.value)
	if err != nil {
		return nil, err
	}
	return marshal, nil
}

func (e *Expected[T]) UnmarshalJSON(data []byte) error {
	if len(data) <= 0 || bytes.Equal(data, []byte("null")) {
		e.value = nil
		return nil
	}

	err := json.Unmarshal(data, e.value)
	if err != nil {
		return err
	}

	return nil
}

// Container.
