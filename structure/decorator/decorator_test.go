package decorator

import (
	"fmt"
	"testing"
)

func TestDecorator(t *testing.T) {
	t.Run("Decorator", func(t *testing.T) {
		pizza := &veggieMania{}

		//Add cheese topping
		pizzaWithCheese := &cheeseTopping{
			pizza: pizza,
		}

		//Add tomato topping
		pizzaWithCheeseAndTomato := &tomatoTopping{
			pizza: pizzaWithCheese,
		}

		fmt.Printf("Price of veggieMania with tomato and cheese topping is %d\n", pizzaWithCheeseAndTomato.getPrice())
	})
}
