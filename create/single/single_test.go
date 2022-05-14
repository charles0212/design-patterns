package single

import (
	"fmt"
	"testing"
)

func TestSingle(t *testing.T) {
	t.Run("Single", func(t *testing.T) {
		for i := 0; i < 30; i++ {
			go getInstance()
		}

		// Scanln is similar to Scan, but stops scanning at a newline and
		// after the final item there must be a newline or EOF.
		fmt.Scanln()
	})
}

func TestSingleSafe(t *testing.T) {
	t.Run("SingleSafe", func(t *testing.T) {
		for i := 0; i < 30; i++ {
			go getInstanceSafe()
		}

		// Scanln is similar to Scan, but stops scanning at a newline and
		// after the final item there must be a newline or EOF.
		fmt.Scanln()
	})
}
