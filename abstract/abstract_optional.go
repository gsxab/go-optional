/*
 * SPDX-License-Identifier: Apache-2.0
 *
 * Copyright (c) 2023 Gsxab
 *
 * Licensed under the Apache License, Version 2.0 (the "License");
 * you may not use this file except in compliance with the License.
 * You may obtain a copy of the License at
 *
 *     http://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing, software
 * distributed under the License is distributed on an "AS IS" BASIS,
 * WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 * See the License for the specific language governing permissions and
 * limitations under the License.
 */

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
