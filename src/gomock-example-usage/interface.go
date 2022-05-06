package greet

//go:generate $GOPATH/bin/mockgen -source=interface.go -destination=./mock/mock.go -package=mock

//IGreeter is the interface
type IGreeter interface {
	Greet(string) string
}
