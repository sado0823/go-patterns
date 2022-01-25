package proxy

import "fmt"

type (
	// Output 代理模式
	Output interface {
		Do(action string)
	}

	Terminal struct{}

	TerminalProxy struct {
		*Terminal
	}
)

func (t *TerminalProxy) Do(action string) {
	if t.Terminal == nil {
		t.Terminal = new(Terminal)
	}

	if action == "old" {
		t.Terminal.Do(action)
		return
	}

	fmt.Println("terminal proxy do: ", action)
}

func (t *Terminal) Do(action string) {
	fmt.Printf("terminal do: %s \n", action)
}
