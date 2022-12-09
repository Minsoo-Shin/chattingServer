package test

import (
	"fmt"
	"testing"
)

func Test_DeleteMapKeyValue(t *testing.T) {
	m := make(map[string]int)
	m["mukul"] = 10
	m["mayank"] = 9
	m["deepak"] = 8

	fmt.Println(m)

	fmt.Println("Deleting the key named deepak from the map")

	delete(m, "deepak")

	fmt.Println(m)
}
