package helpers

import gonanoid "github.com/matoous/go-nanoid"

func GenerateID() string {
	id, _ := gonanoid.Generate("1234567890abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ", 10)
	return id
}
