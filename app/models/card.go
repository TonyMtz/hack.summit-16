package models

type Card struct {
	Id    int
	Title string
	Desc  string
	Tags  []Tag
}