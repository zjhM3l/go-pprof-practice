package felidae

import "github.com/zjhM3l/go-pprof-practice/animal"

type Felidae interface {
	animal.Animal
	Climb()
	Sneak()
}
