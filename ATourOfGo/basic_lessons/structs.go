package main

import (
	"fmt"
	"reflect"
)

// The type keyword provides a way to name and define new types in Go, making your code more organized, modular, and reusable. Type can be used for non struct purposes as well.
// The struct keyword, used within a type declaration, provides the structure and blueprint for the data you want to group together, enabling you to create variables of that type.
// hence struct has to be with type but type can be without struct
type vertex struct {
	X int
	Y int
}

/*
A struct literal denotes a newly allocated struct value by listing the values of its fields.
You can list just a subset of fields by using the Name: syntax. (And the order of named fieds is irrelevant.)
The special prefix & returns a pointer to the struct value.
*/

var (
	v1 = vertex{1, 2} // has type Vertex
	v2 = vertex{X: 1} // Y:0 is implicit
	v3 = vertex{}     // X:0 and Y:0
	v4 = vertex{X: 1, Y: 2}
	p1 = &vertex{1, 2} // has type *Vertex
)

func main() {
	v := vertex{1, 2}
	p := &v

	// To access the field X of a struct when we have the struct pointer p we could write (*p).X.
	// However, that notation is cumbersome, so the language permits us instead to write just p.X, without the explicit dereference.
	p.X = 1e9
	fmt.Println(v, p, &v, p.X, v.X, &v.X, &p.X)
	fmt.Println(reflect.TypeOf(v), reflect.TypeOf(p), reflect.TypeOf(v.X))
	fmt.Println(v1, p1, v2, v3, v4)
}
