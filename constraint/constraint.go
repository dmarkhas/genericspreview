package main

import "fmt"

type Addable interface {
	type int, float64, string
}

func sum[T Addable](a, b T) T {
	return a + b
}

func main() {
	fmt.Printf("%[1]T %[1]v\n", sum(10, 30)) // arguments must be of the same type
	fmt.Printf("%[1]T %[1]v\n", sum(1.5, 4.2))
	fmt.Printf("%[1]T %[1]v\n", sum("foo", "bar"))

	// uncomment for error
	// fmt.Println(sum("foo", 1))

}
