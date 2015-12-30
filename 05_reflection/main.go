package main

import (
	"fmt"
	"reflect"
)

type Bird struct {
	Name           string
	LifeExpectance int
}

func (Bird) fly() {
	fmt.Printf("I am flying....")
}

func main() {
	sparrow := &Bird{Name: "Sparrow", LifeExpectance: 3}
	s := reflect.ValueOf(sparrow).Elem()
	typeofT := s.Type()
	fields := s.NumField()
	for i := 0; i < fields; i++ {
		f := s.Field(i)
		fmt.Printf("%d: %s %s = %v \n", i, typeofT.Field(i).Name, f.Type(), f.Interface())
	}
}
