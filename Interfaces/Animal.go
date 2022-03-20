package Interfaces

import (
	"alias/Errors"
)

type Animal interface {
	MakeNoise() (string, *Errors.AnimalErrors)
}
