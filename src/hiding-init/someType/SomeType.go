package sometype

//SomeType demonstrate how we can disable initialization of a struct outside its package
type SomeType struct {
	a int
	b string
}

//NewSomeType hides the initialization outside of its package by making its field private
func NewSomeType(a int, b string) *SomeType {
	return &SomeType{a: a, b: b}
}
