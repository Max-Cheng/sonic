// +build !noasm !appengine
// Code generated by asm2asm, DO NOT EDIT.

package avx2

//go:nosplit
//go:noescape
//goland:noinspection ALL
func __native_entry__() uintptr

var (
    _subr__f64toa      = __native_entry__() + 903
    _subr__i64toa      = __native_entry__() + 3915
    _subr__lspace      = __native_entry__() + 429
    _subr__lzero       = __native_entry__() + 13
    _subr__quote       = __native_entry__() + 5328
    _subr__skip_array  = __native_entry__() + 19163
    _subr__skip_object = __native_entry__() + 19198
    _subr__skip_one    = __native_entry__() + 16306
    _subr__u64toa      = __native_entry__() + 4008
    _subr__unquote     = __native_entry__() + 7125
    _subr__value       = __native_entry__() + 11812
    _subr__vnumber     = __native_entry__() + 14464
    _subr__vsigned     = __native_entry__() + 15778
    _subr__vstring     = __native_entry__() + 13587
    _subr__vunsigned   = __native_entry__() + 16037
)

const (
    _stack__f64toa = 120
    _stack__i64toa = 24
    _stack__lspace = 8
    _stack__lzero = 8
    _stack__quote = 80
    _stack__skip_array = 128
    _stack__skip_object = 128
    _stack__skip_one = 128
    _stack__u64toa = 8
    _stack__unquote = 72
    _stack__value = 392
    _stack__vnumber = 312
    _stack__vsigned = 16
    _stack__vstring = 112
    _stack__vunsigned = 8
)

var (
    _ = _subr__f64toa
    _ = _subr__i64toa
    _ = _subr__lspace
    _ = _subr__lzero
    _ = _subr__quote
    _ = _subr__skip_array
    _ = _subr__skip_object
    _ = _subr__skip_one
    _ = _subr__u64toa
    _ = _subr__unquote
    _ = _subr__value
    _ = _subr__vnumber
    _ = _subr__vsigned
    _ = _subr__vstring
    _ = _subr__vunsigned
)

const (
    _ = _stack__f64toa
    _ = _stack__i64toa
    _ = _stack__lspace
    _ = _stack__lzero
    _ = _stack__quote
    _ = _stack__skip_array
    _ = _stack__skip_object
    _ = _stack__skip_one
    _ = _stack__u64toa
    _ = _stack__unquote
    _ = _stack__value
    _ = _stack__vnumber
    _ = _stack__vsigned
    _ = _stack__vstring
    _ = _stack__vunsigned
)
