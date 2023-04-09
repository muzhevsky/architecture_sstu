package main

import (
	"fmt"
	"uid_generator/main/idgenerators"
)

func main() {
	ids1 := idgenerators.Get100UUIDs()
	for i := 0; i < len(*ids1); i++ {
		fmt.Println((*ids1)[i])
	}

	idgenerators.InitGenerator(1, 1)
	idgenerators.StartGenerator()
	ids2 := idgenerators.Get100SnowFlakeIDs()
	for i := 0; i < len(*ids2); i++ {
		fmt.Println((*ids2)[i])
	}
}
