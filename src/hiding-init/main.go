package main

import (
	"fmt"
	v0 "hiding-init/notHiding"
	someType "hiding-init/someType"
	v1 "hiding-init/v1"
)

func main() {
	type0 := v0.SomeType{}
	fmt.Println(type0)
	//this approach doesn't hide anything
	//it just warn users through doc string of SomeType

	type1 := v1.NewType()
	fmt.Println(&type1)
	//this approach disallows: privateType{}
	//but you can't return a pointer from NewType: *Type

	type2A := someType.SomeType{}
	fmt.Println(type2A)
	type2B := someType.NewSomeType(3, "three")
	fmt.Println(type2B)
	//this approach still allow initialization to zero values
}
