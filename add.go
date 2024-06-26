// Package add provides utilities for addition operations
package add

import (
	"golang.org/x/exp/constraints"
)

type Number interface {
	constraints.Integer | constraints.Float
}

// Add takes two Number and sums them up. More information on the addition operation can be found [here].
// [here]: https://www.mathsisfun.com/numbers/addition.html
func Add[T Number](x, y T) T {
	return x + y
}
