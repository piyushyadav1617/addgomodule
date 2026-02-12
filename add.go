// Package addgomodule provides method to add to numbers
package addgomodule

import (
	"golang.org/x/exp/constraints"
)

type Number interface {
	constraints.Integer | constraints.Float
}

// Add adds two numbers.
// See [Math is Fun] for more about addition.
//
// [Math is Fun]: https://www.mathisfun.com/numbers/addition.html
func Add[T Number](a, b T) T {
	return a + b
}
