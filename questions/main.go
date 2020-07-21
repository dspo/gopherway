package main

import (
	"fmt"
)

type ReaderA interface {
	Read([]byte)(int, error)
}

type ReaderB = interface {
	Read([]byte)(int, error)
}

type ReaderC = interface {
	Read([]byte)(int, error)
}

type AMaker interface {
	MakeReader() ReaderA
}

type BMaker interface {
	MakeReader() ReaderB
}

type CMaker interface {
	MakeReader() ReaderC
}

type readerC struct {

}

func (r *readerC) Read([]byte) (int, error) {
	return 0, nil
}

type cMaker struct {

}

func (r *cMaker) MakeReader() ReaderC {
	return new(readerC)
}

func f1() {
	var b BMaker
	b = new(cMaker)
	_, okA := b.(AMaker)
	_, okB := b.(CMaker)
	fmt.Println(okA, okB)
}

func f2() {
	var c CMaker
	c = new(cMaker)
	_, okA := c.(AMaker)
	_, okB := c.(BMaker)
	fmt.Println(okA, okB)
}

func main() {
	f1()
	f2()
}
