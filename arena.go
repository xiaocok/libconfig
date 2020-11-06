/*
 * The MIT License (MIT)
 *
 * Copyright (c) 2018 Aliaksandr Valialkin
 *
 * Permission is hereby granted, free of charge, to any person obtaining a copy
 * of this software and associated documentation files (the "Software"), to deal
 * in the Software without restriction, including without limitation the rights
 * to use, copy, modify, merge, publish, distribute, sublicense, and/or sell
 * copies of the Software, and to permit persons to whom the Software is
 * furnished to do so, subject to the following conditions:
 *
 * The above copyright notice and this permission notice shall be included in all
 * copies or substantial portions of the Software.
 *
 * THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
 * IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
 * FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
 * AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
 * LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
 * OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN THE
 * SOFTWARE.
 *
 * Author: Aliaksandr Valialkin <valyala@gmail.com>
 */
package libconfig

import (
	"strconv"
)

// Arena may be used for fast creation and re-use of Values.
//
// Typical Arena lifecycle:
//
//     1) Construct Values via the Arena and Value.Set* calls.
//     2) Marshal the constructed Values with Value.MarshalTo call.
//     3) Reset all the constructed Values at once by Arena.Reset call.
//     4) Go to 1 and re-use the Arena.
//
// It is unsafe calling Arena methods from concurrent goroutines.
// Use per-goroutine Arenas or ArenaPool instead.
type Arena struct {
	b []byte
	c cache
}

// Reset resets all the Values allocated by a.
//
// Values previously allocated by a cannot be used after the Reset call.
func (a *Arena) Reset() {
	a.b = a.b[:0]
	a.c.reset()
}

// NewObject returns new empty object value.
//
// New entries may be added to the returned object via Set call.
//
// The returned object is valid until Reset is called on a.
func (a *Arena) NewObject() *Value {
	v := a.c.getValue()
	v.t = TypeObject
	v.o.reset()
	return v
}

// NewArray returns new empty array value.
//
// New entries may be added to the returned array via Set* calls.
//
// The returned array is valid until Reset is called on a.
func (a *Arena) NewArray() *Value {
	v := a.c.getValue()
	v.t = TypeArray
	v.a = v.a[:0]
	return v
}

// NewString returns new string value containing s.
//
// The returned string is valid until Reset is called on a.
func (a *Arena) NewString(s string) *Value {
	v := a.c.getValue()
	v.t = typeRawString
	bLen := len(a.b)
	a.b = escapeString(a.b, s)
	v.s = b2s(a.b[bLen+1 : len(a.b)-1])
	return v
}

// NewStringBytes returns new string value containing b.
//
// The returned string is valid until Reset is called on a.
func (a *Arena) NewStringBytes(b []byte) *Value {
	v := a.c.getValue()
	v.t = typeRawString
	bLen := len(a.b)
	a.b = escapeString(a.b, b2s(b))
	v.s = b2s(a.b[bLen+1 : len(a.b)-1])
	return v
}

// NewNumberFloat64 returns new number value containing f.
//
// The returned number is valid until Reset is called on a.
func (a *Arena) NewNumberFloat64(f float64) *Value {
	v := a.c.getValue()
	v.t = TypeNumber
	bLen := len(a.b)
	a.b = strconv.AppendFloat(a.b, f, 'g', -1, 64)
	v.s = b2s(a.b[bLen:])
	return v
}

// NewNumberInt returns new number value containing n.
//
// The returned number is valid until Reset is called on a.
func (a *Arena) NewNumberInt(n int) *Value {
	v := a.c.getValue()
	v.t = TypeNumber
	bLen := len(a.b)
	a.b = strconv.AppendInt(a.b, int64(n), 10)
	v.s = b2s(a.b[bLen:])
	return v
}

// NewNumberString returns new number value containing s.
//
// The returned number is valid until Reset is called on a.
func (a *Arena) NewNumberString(s string) *Value {
	v := a.c.getValue()
	v.t = TypeNumber
	v.s = s
	return v
}

// NewNull returns null value.
func (a *Arena) NewNull() *Value {
	return valueNull
}

// NewTrue returns true value.
func (a *Arena) NewTrue() *Value {
	return valueTrue
}

// NewFalse return false value.
func (a *Arena) NewFalse() *Value {
	return valueFalse
}
