package greet

func functionWithMock(g IGreeter) string {
	return g.Greet("world")
}
