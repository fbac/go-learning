# new vs make

## new()

`new()` will return a pointer to a variable of the type we declared. It's the only function capable of returning a pointer to a native type variable in one line.

The following code:

```go
package main

import "fmt"

main() {
    var myInt = new(int)
    fmt.Printf("%T\n", myInt)
}
```

Would output `*int`.

And it is equivalent to:

```go
package main

func main() {
    var ptrx int
    x := &ptrx
}
```

Caveats:
- Types like `map`, `slice` and `channels` are composite types, meaning they're wrappers types that hold pointers to the in-memory fields that compose them.
Calling new to initialize one of this composite types would result in an unusable composite type.

- For a slice it will return a nil array, and no item can be added into a nil array.
- Same would happen with structs or maps.

## make()

`Zeroing` is a form of initializing all the variable fields to its default values, that is 0 for integers, nil for structs, "" for strings, etc.

`make()` zeroes the underlying type, creating a valid map, slice or struct ready to be used, since all the values are default and there are no nil fields.
