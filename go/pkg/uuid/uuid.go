package uuid

import "github.com/google/uuid"

func Gen() uuid.UUID {
	u, err := uuid.NewUUID()
	if err != nil {
		panic(err)
	}
	return u
}
