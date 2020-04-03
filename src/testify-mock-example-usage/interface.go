package greet

//go:generate $GOPATH/bin/mockery -all -output ./mocks

//IGreeter is the interface
type IGreeter interface {
	Greet(string) string
}
