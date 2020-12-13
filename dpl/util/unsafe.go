package util

import (
	"reflect"
	"unsafe"
)

func UnsafeExtend(object, parent interface{}) unsafe.Pointer {
	v1 := reflect.ValueOf(object)
	v2 := reflect.ValueOf(parent)
	p1 := unsafe.Pointer(v1.Pointer())
	p2 := unsafe.Pointer(v2.Pointer())
	t1 := v1.Type().Elem()
	t2 := v2.Type().Elem()
	if t1.Size() < t2.Size() {
		panic("object offset base pointer out of range")
	}
	for i := uintptr(0); i < t2.Size(); i++ {
		*(*byte)(unsafe.Pointer((uintptr)(unsafe.Pointer(p1)) + i)) = *(*byte)(unsafe.Pointer((uintptr)(unsafe.Pointer(p2)) + i))
	}
	return p1
}

func UnsafeConvert(value interface{}) unsafe.Pointer {
	return unsafe.Pointer(reflect.ValueOf(value).Pointer())
}
