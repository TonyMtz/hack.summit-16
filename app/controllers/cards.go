package controllers

import (
	"github.com/revel/revel"
	"github.com/TonyMtz/hack.summit-16.service.git/app/models"
)

type Cards struct {
	*revel.Controller
}

func (c Cards) List() revel.Result {
	tag1 := models.Tag{Id:1, Name:"Tag 1"}
	tag2 := models.Tag{Id:2, Name:"Tag 2"}
	tag3 := models.Tag{Id:3, Name:"Tag 3"}
	cards := []models.Card{
		models.Card{Id:1, Title:"Test 1", Desc:"Wow!", Tags:[]models.Tag{tag1, tag3}},
		models.Card{Id:2, Title:"Test 2", Desc:"Wow!!"},
		models.Card{Id:3, Title:"Test 3", Desc:"Wow!!!", Tags: []models.Tag{tag1, tag2, tag3}},
	}
	return c.RenderJson(cards)
}