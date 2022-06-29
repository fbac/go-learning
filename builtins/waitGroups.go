/*
	waitGroups example:
	- sync goroutines
		- add, wait and more methods
*/
package main

import (
	"fmt"
	"runtime"
	"sync"
)

type potion struct {
	id   int
	name string
}

func craft(id int, name string) potion {
	for i := 0; i < 5; i++ {
		fmt.Printf("crafting %s potion...\n", name)
	}
	// signal that the goroutine is done
	wg.Done()
	return potion{id, name}
}

// required: create a waitGroup
var wg sync.WaitGroup

func main() {
	fmt.Println("let's craft some potions")

	// wait for n wg.Done()
	wg.Add(3)
	go craft(0, "health")
	go craft(1, "agility")
	go craft(2, "strength")
	// package runtime holds runtime info
	// such as arch, cpu, goroutines...
	fmt.Printf("crafting %v potions", runtime.NumGoroutine())
	wg.Wait()
}
