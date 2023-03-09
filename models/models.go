package models

import "html/template"

type Board struct {
	Title string
	Posts []Post
}

type Post struct {
	Board     string
	Content   string
	Html      template.HTML
	Timestamp string
}
