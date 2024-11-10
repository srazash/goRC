package user

import (
	"fmt"
	"math/rand"
)

type colour int

const (
	Red colour = iota
	Green
	Blue
	Indigo
	Yellow
	Orange
	Purple
	Brown
	Gold
	Silver
	White
	Black
)

func (c colour) string() string {
	switch c {
	case Red:
		return "Red"
	case Green:
		return "Green"
	case Blue:
		return "Blue"
	case Indigo:
		return "Indigo"
	case Yellow:
		return "Yellow"
	case Orange:
		return "Orange"
	case Purple:
		return "Purple"
	case Brown:
		return "Brown"
	case Gold:
		return "Gold"
	case Silver:
		return "Silver"
	case White:
		return "White"
	case Black:
		return "Black"
	default:
		return "Unknown"
	}
}

type animal int

const (
	Hedgehog animal = iota
	Badger
	Cat
	Dog
	Mouse
	Snake
	Owl
	Monkey
	Eagle
	Shark
	Octopus
	Tiger
	Bear
	Lion
	Lizard
	Fox
	Wolf
	Penguin
)

func (a animal) string() string {
	switch a {
	case Hedgehog:
		return "Hedgehog"
	case Badger:
		return "Badger"
	case Cat:
		return "Cat"
	case Dog:
		return "Dog"
	case Mouse:
		return "Mouse"
	case Snake:
		return "Snake"
	case Owl:
		return "Owl"
	case Monkey:
		return "Monkey"
	case Eagle:
		return "Eagle"
	case Shark:
		return "Shark"
	case Octopus:
		return "Octopus"
	case Tiger:
		return "Tiger"
	case Bear:
		return "Bear"
	case Lion:
		return "Lion"
	case Lizard:
		return "Lizard"
	case Fox:
		return "Fox"
	case Wolf:
		return "Wolf"
	case Penguin:
		return "Penguin"
	default:
		return "Unknown"
	}
}

func Randonym() string {
	c := colour(rand.Intn(11))
	a := animal(rand.Intn(17))
	n := rand.Intn(900) + 100

	return fmt.Sprintf("%s%s%d", c.string(), a.string(), n)
}
