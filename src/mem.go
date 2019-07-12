package main

import (
	"fmt"
	"unsafe"
)

type experiment struct {
	name      string
	awardName string
}

type dataFile struct {
	s3Key      string
	dlCount    uint64      // Ability to mark memory size and signed/unsigned
	experiment *experiment // Asterisk means pointer to variable in heap (C style)
}

func main() {
	valFile := dataFile{} // Empty initialization on stack zeros out fields and sets pointer to nil
	valFile.print()       // Note that the string is not nil because it is a value type

	valFile = dataFile{"2019/2/11/Something.bigWig", 6, nil} // Order of definition. All fields required.
	valFile.print()

	ptrFile := new(dataFile) // Memory zeroed out, located on the heap
	ptrFile.print()

	ptrFile = newFile("1999/4/15/Something.bigBed")

	var bigFloat float64
	fmt.Printf("%d bytes\n", unsafe.Sizeof(bigFloat))
	var smallPosInt uint8
	fmt.Printf("%d bytes\n", unsafe.Sizeof(smallPosInt))

	valFile.modifyCopy()
	valFile.print()

	valFile.modifyRef()
	valFile.print()
}

func (file *dataFile) print() {
	fmt.Printf("%+v\n", file)
}

func newFile(s3key string) *dataFile {
	file := &dataFile{s3Key: s3key} // Named initialization. Other values are zeroed. Converted to pointer with ampersand
	fmt.Printf("%T\n", file)  // Pointer type since we t
	return file                     // Not valid C code, returning pointer to stack object
}

func (file dataFile) modifyCopy() {
	file.experiment = &experiment{}
}

func (file *dataFile) modifyRef() {
	file.experiment = &experiment{}
}
