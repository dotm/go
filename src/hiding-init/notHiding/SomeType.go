package v0

//SomeType MUST NOT BE INITIALIZED DIRECTLY like this: `SomeType{}`.
//Use NewSomeType instead!
type SomeType struct {
	A int
	B string
}

//NewSomeType doesn't hide anything. It just warn programmer through doc string
func NewSomeType(a int, b string) *SomeType {
	////you can do any kind of configuration here
	return &SomeType{A: a, B: b}
}
