// Package writes tests ref/writes edges
package writes

import "fmt"

type S struct {
	//- @F defines/binding F
	F int

	//- @Nested defines/binding Nested
	Nested *S
}

func f() {
	//- @x defines/binding X
	var x int

	//- @x ref/writes X
	//- @y defines/binding Y
	x, y := 1, 2

	//- @y ref Y
	//- @y ref/writes Y
	y++

	//- @#0i defines/binding I
	//- @#1i ref I
	//- @#2i ref/writes I
	//- @#3i ref I
	for i := 0; i < 10; i = i + 1 {
	}

	//- @z defines/binding Z
	z := S{Nested: &S{}}
	//- @z ref Z
	//- @F ref/writes F
	z.F = 42
	//- @z ref Z
	//- @Nested ref Nested
	//- @F ref/writes F
	z.Nested.F = 52

	//- @x ref X
	//- @y ref Y
	fmt.Println(x, y)
}
