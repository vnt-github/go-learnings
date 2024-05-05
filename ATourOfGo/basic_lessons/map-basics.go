package main

import (
	"fmt"
)

type Vertex struct {
	Lat  float64
	Long float64
}

// Map literals are like struct literals, but the keys are required.
var mapLiteral1 = map[string]Vertex{
	"key1": Vertex{1.1, 1.2},
	"key2": Vertex{2.1, 2.2},

	// NOTE: at the end of the literal there is a comma ,
}

// If the top-level type is just a type name, you can omit it from the elements of the literal.
var mapLiteral2 = map[string]Vertex{
	"key_a": {1.1, 1.2},
	"key_b": {2.1, 2.2},
}

func main() {
	// The zero value of a map is nil. A nil map has no keys, nor can keys be added.
	// The make function returns a map of the given type, initialized and ready for use.

	// var m map[string]Vertex
	// m = make(map[string]Vertex)
	m := make(map[string]Vertex)

	m["test"] = Vertex{1.1, 1.2}
	fmt.Println(m)

	fmt.Println(mapLiteral1)
	fmt.Println(mapLiteral2)

	// inserting or updating a value is m[key] = elem

	m["new_value1"] = Vertex{2.1, 2.2}
	m["test"] = Vertex{3.1, 3.2}
	fmt.Println(m)

	// retrieve a value elem = m[key]
	var testValue float64
	testValue = m["test"].Lat
	fmt.Println(testValue)

	// delete an element delete(m, key)
	delete(m, "test")
	// If key is not in the map, then elem is the zero value for the map's element type.
	fmt.Println(m["test"])

	// Test that a key is present with a two-value assignment: elem, ok = m[key]
	// If key is in m, ok is true. If not, ok is false.
	elem, isPresent := m["test"]
	// If key is not in the map, then elem is the zero value for the map's element type.
	fmt.Println(elem, isPresent)
}
