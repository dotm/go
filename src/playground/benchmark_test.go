package main

import "testing"

func BenchmarkAnything(b *testing.B) {
	for i := 0; i < b.N; i++ {
		//put anything here
	}
}
