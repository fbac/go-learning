/*
	This lacks of real use.
	Do not use, it's just an experiment.
*/
package main

import (
	"fmt"
	"reflect"
)

func main() {

	structCeption := struct {
		id      int
		name    string
		struct2 struct {
			id      int
			name    string
			struct3 struct {
				id   int
				name string
			}
		}
	}{id: 1, name: ".struct", struct2: struct {
		id      int
		name    string
		struct3 struct {
			id   int
			name string
		}
	}{id: 2, name: ".struct.struct", struct3: struct {
		id   int
		name string
	}{id: 3, name: ".struct.struct.struct"}}}

	fmt.Println(structCeption)

	// reflection
	structCeptionType := reflect.TypeOf(structCeption)
	structCeptionFields := structCeptionType.NumField()
	for i := 0; i < structCeptionFields; i++ {
		field := structCeptionType.Field(i)
		fmt.Printf("Field Type: %s: %s Kind: %s\n", field.Name, field.Type.Name(), field.Type.Kind())
	}
}
