package main

import (
	"fmt"
	"strconv"
	"unsafe"
)

type (
	SuperStruct struct {
		id   int
		name string
	}
)

func main() {

	const ItemCt = 1 << 20

	fmt.Println("testing", ItemCt, "items")

	// test a std list
	stdlist := make([]*SuperStruct, ItemCt)
	for i, _ := range stdlist {
		stdlist[i] = &SuperStruct{id: i, name: "std" + strconv.Itoa(i)}
	}
	for _, item := range stdlist {
		fmt.Printf("%p: %#v\n", item, item)
	}

	// test a uintptr list pointing to stdlist items
	uintptrlist := make([]uintptr, len(stdlist))
	for i, item := range stdlist {
		uintptrlist[i] = (uintptr)(unsafe.Pointer(item))
	}
	for i, _ := range uintptrlist {
		item := (*SuperStruct)(unsafe.Pointer(uintptrlist[i]))
		fmt.Printf("%p: %#v\n", item, item)
	}

	// test byte array of pointers to stdlist items
	// traverse via pointer arithmetic
	ptrsize := unsafe.Sizeof(&SuperStruct{})
	ptrlist := make([]byte, ptrsize*len(stdlist))
	ptr := (uintptr)(unsafe.Pointer(&ptrlist[0]))
	for _, s := range stdlist {
		*(**SuperStruct)(unsafe.Pointer(ptr)) = s
		ptr += ptrsize
	}

	ptr = (uintptr)(unsafe.Pointer(&ptrlist[0]))
	for i := 0; i < ItemCt; i++ {
		item := *(**SuperStruct)(unsafe.Pointer(ptr))
		fmt.Printf("%p: %#v\n", item, item)
		ptr += ptrsize
	}

}
