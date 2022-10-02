package model

type Album struct {
	Id string `json:"id"`
	Title string `json:"title"`
}

var Albums = []Album{
	{Id: "1", Title: "Gone with the wind"},
	{Id: "2", Title: "War and Peace"},
}