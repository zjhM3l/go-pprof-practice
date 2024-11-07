package animal

import (
	"github.com/zjhM3l/go-pprof-practice/animal/canidae/dog"
	"github.com/zjhM3l/go-pprof-practice/animal/canidae/wolf"
	"github.com/zjhM3l/go-pprof-practice/animal/felidae/cat"
	"github.com/zjhM3l/go-pprof-practice/animal/felidae/tiger"
	"github.com/zjhM3l/go-pprof-practice/animal/muridae/mouse"
)

var (
	AllAnimals = []Animal{
		&dog.Dog{},
		&wolf.Wolf{},

		&cat.Cat{},
		&tiger.Tiger{},

		&mouse.Mouse{},
	}
)

type Animal interface {
	Name() string
	Live()

	Eat()
	Drink()
	Shit()
	Pee()
}
