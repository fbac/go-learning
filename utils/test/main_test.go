// Same package as main
package main

import (
	"fmt"
	"testing"
)

// TestNewPotion is the expected name, and the input is the expected signature
func TestNewPotion(t *testing.T) {
	p := NewPotion(0, "test")
	if p.id != 0 {
		t.Error("Expected p.id: 0, got:", p.id)
	}
}

// TestString is the expected name, and the input is the expected signature
func TestString(t *testing.T) {
	p := NewPotion(0, "test")
	pString := p.String()
	if pString != "test" {
		t.Error("Expected string: test, got:", pString)
	}
}

func ExampleString() {
	p := NewPotion(0, "test")
	fmt.Println(p.String())
	// Output:
	// test
}
