package testmodule

import "fmt"

type Mod struct {
	Name string
}

func (t *Mod) Say()  {
	fmt.Println(t.Name+"v-4.0.0")
}
