package main

type Command struct {
	Name         string
	API          string
	Org          string
	Space        string
	LoginOptions string `yaml:"login-options"`
}
