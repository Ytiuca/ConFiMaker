package main

type Widget interface {
	Name() string
	ToPython() string
	ToGetter() string
	ToArg() string
}
