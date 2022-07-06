package benchmarkingtest

import (
	"fmt"
	"testing"
)

func TestSumInt(t *testing.T) {
	x := SumInt(1, 2)
	if x != 3 {
		t.Error("expected 3, got:", x)
	}
}

func TestConcatenateString(t *testing.T) {
	s := ConcatenateString("Hello", " World!")
	if s != "Hello World!" {
		t.Error("expected Hello World!, got:", s)
	}
}

func ExampleSumInt(t *testing.T) {
	x := SumInt(1, 2)
	fmt.Println(x)
	// Output:
	// 3
}

/*
	In order to benchmark we'll just simply create some bench tests
	More info in the test/bench docs
*/

// BenchmarkSumInt-8	1000000000	0.2654 ns/op
func BenchmarkSumInt(b *testing.B) {
	for i := 0; i < b.N; i++ {
		SumInt(5, 5)
	}
}

// BenchmarkConcatenateString-8	69000822	15.96 ns/op
func BenchmarkConcatenateString(b *testing.B) {
	for i := 0; i < b.N; i++ {
		ConcatenateString("Hello", " World!")
	}
}
