package fschaos

import (
	"log"
	"os"
	//"syscall"
	"unsafe"
)

const (
	AlignSize = 4096 //align min size

	BlockSize = 4096
)

//todo memory
//func OpenFile(name string, flag int, perm os.FileMode) (file *os.File, err error) {
//	return os.OpenFile(name, syscall.O_DIRECT|flag, perm)
//}

func OpenFile(name string, flag int, perm os.FileMode) (file *os.File, err error) {
	return os.OpenFile(name, flag, perm)
}

func alignment(block []byte, AlignSize int) int {
	return int(uintptr(unsafe.Pointer(&block[0])) & uintptr(AlignSize-1))
}

func AlignedBlock(BlockSize int) []byte {
	block := make([]byte, BlockSize+AlignSize)
	if AlignSize == 0 {
		return block
	}
	a := alignment(block, AlignSize)
	offset := 0
	if a != 0 {
		offset = AlignSize - a
	}
	block = block[offset : offset+BlockSize]
	if BlockSize != 0 {
		a = alignment(block, AlignSize)
		if a != 0 {
			log.Fatal("Failed to align block")
		}
	}
	return block
}
