package models

type Card struct {
	Id       string
	Title    string
	Desc     string
	Provider string
	Tags     []Tag
}