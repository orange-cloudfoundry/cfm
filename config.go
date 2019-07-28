package main

type Config struct {
	CurrentGroup string
	Targets      []Target
}

type Target struct {
	Api   string
	Alias string
	Group string
}
