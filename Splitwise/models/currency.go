package models

const (
	USD = iota
	INR
)

type Currency struct {
	Code int
}
