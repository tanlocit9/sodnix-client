package generator

import "github.com/google/uuid"

// GenerateUUIDv4 returns a new UUIDv4
func GenerateUUIDv4() uuid.UUID {
	return uuid.New()
}

// GenerateUUIDv7 returns a new UUIDv7 or panics (fail-fast strategy)
func GenerateUUIDv7() uuid.UUID {
	id, err := uuid.NewV7()
	if err != nil {
		panic("failed to generate UUIDv7: " + err.Error())
	}
	return id
}
