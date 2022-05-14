package visitor

import (
	"fmt"
	"testing"
)

func TestVisitor(t *testing.T) {
	t.Run("visitor", func(t *testing.T) {
		square := &square{side: 2}
		circle := &circle{radius: 3}
		rectangle := &rectangle{l: 2, b: 3}

		areaCalculator := &areaCalculator{}

		square.accept(areaCalculator)
		circle.accept(areaCalculator)
		rectangle.accept(areaCalculator)

		fmt.Println()
		middleCoordinates := &middleCoordinates{}
		square.accept(middleCoordinates)
		circle.accept(middleCoordinates)
		rectangle.accept(middleCoordinates)
	})
}
