package command

import (
	"testing"
)

func TestCommand(t *testing.T) {
	t.Run("command", func(t *testing.T) {
		tv := &tv{}

		onCommand := &onCommand{
			device: tv,
		}

		offCommand := &offCommand{
			device: tv,
		}

		onButton := &button{
			command: onCommand,
		}
		onButton.press()

		offButton := &button{
			command: offCommand,
		}
		offButton.press()
	})
}
