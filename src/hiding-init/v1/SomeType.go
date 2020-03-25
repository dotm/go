package v1

//Type demonstrate how we can disable initialization of a struct outside its package
type Type interface {
	GetA() int
	GetB() string
	SetA(int)
	SetB(string)
}

type privateType struct {
	a int
	b string
}

//Make privateType conforms to public Type
func (b *privateType) GetA() int       { return b.a }
func (b *privateType) GetB() string    { return b.b }
func (b *privateType) SetA(val int)    { b.a = val }
func (b *privateType) SetB(val string) { b.b = val }

//NewType is the only way to initialize the struct outside of its package
func NewType() Type {
	//you can do any kind of configuration here
	return new(privateType)
}
