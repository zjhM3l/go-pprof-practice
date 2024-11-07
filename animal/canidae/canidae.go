package canidae

import "github.com/zjhM3l/go-pprof-practice/animal"

type Canidae interface {
	animal.Animal
	Run()
	Howl()
}
