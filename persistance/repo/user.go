package main

import "github.com/google/uuid"

type User struct {
	UUID     uuid.UUID
	Username string
	Email    string
}
