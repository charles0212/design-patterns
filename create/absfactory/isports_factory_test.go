package absfactory

import (
	"fmt"
	"testing"
)

func TestAbstractFactory(t *testing.T) {
	t.Run("AbstractFactory", func(t *testing.T) {
		adidasFactory, _ := getSportsFactory("adidas")
		nikeFactory, _ := getSportsFactory("nike")

		nikeShoe := nikeFactory.makeShoe()
		nikeShirt := nikeFactory.makeShirt()

		adidasShoe := adidasFactory.makeShoe()
		adidasShirt := adidasFactory.makeShirt()

		printShoeDetails(nikeShoe)
		printShirtDetails(nikeShirt)

		printShoeDetails(adidasShoe)
		printShirtDetails(adidasShirt)
	})
}

func printShoeDetails(s iShoe) {
	fmt.Printf("Logo: %s", s.getLogo())
	fmt.Println()
	fmt.Printf("Size: %d", s.getSize())
	fmt.Println()
}

func printShirtDetails(s iShirt) {
	fmt.Printf("Logo: %s", s.getLogo())
	fmt.Println()
	fmt.Printf("Size: %d", s.getSize())
	fmt.Println()
}
