package idgenerators

import (
	"fmt"
	"github.com/google/uuid"
)

func getUUID() (uuid.UUID, error) {
	return uuid.NewUUID()
}

func Get100UUIDs() *[]uuid.UUID {
	res := make([]uuid.UUID, 100)
	for i := 0; i < 100; i++ {
		id, err := getUUID()
		if err != nil {
			fmt.Println(err.Error())
		}
		res[i] = id
	}

	return &res
}
