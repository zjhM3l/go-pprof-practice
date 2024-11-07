package muridae

import "github.com/zjhM3l/go-pprof-practice/animal"

type Muridae interface {
	animal.Animal
	Hole()
	Steal()
}
