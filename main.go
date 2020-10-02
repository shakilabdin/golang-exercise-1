package main

import "fmt"

func main() {
	exercise_one()
	exercise_two()
	exercise_three()
}

//exercise 1
func exercise_one() {
	x := 42
	y := "James Bond"
	z := true

	fmt.Println(x, y, z)
	fmt.Println(x)
	fmt.Println(y)
	fmt.Println(z)
}

//exercise 2
func exercise_two() {
	var x int
	var y string
	var z bool

	fmt.Println(x, y, z)
}

//exercise 3
func exercise_three() {
	type what int
	var x what
	fmt.Println(x)
	fmt.Printf("%T\n", x)
	x = 42
	fmt.Println(x)
}

//exercise 4
func exercise_four() {

}
