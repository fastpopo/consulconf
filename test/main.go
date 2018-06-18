package main

import (
	"fmt"
	"time"

	"github.com/fastpopo/gconf"
	"github.com/fastpopo/consulconf"
)

func main() {
	builder := gconf.NewConfBuilder()

	jsonSource := gconf.NewJsonFileConfSource("test.json")
	source := consulconf.NewConsulConfSource("165.213.198.242:8500")
	source.SetOnConfChangedCallback(consulChanged)


	builder.Add(source)
	builder.Add(jsonSource)
	root := builder.Build()

	pairs := root.ToArray()

	for _, p := range pairs {
		fmt.Printf("Key: %v, Value: %v\n", p.Key, p.Value)
	}

	time.Sleep(time.Minute * 3)
}

func consulChanged(confChanges gconf.ConfChanges) {
	fmt.Printf("[consulChanged] Changed: %v\n", confChanges.GetNumOfChanges())

	if confChanges.GetNumOfChanges() == 0 {
		return
	}

	changes := confChanges.GetChanges()

	for _, c := range changes {
		fmt.Printf("%v\n", c.String())
	}
}
