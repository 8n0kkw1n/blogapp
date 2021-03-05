package models

//Posts - структура постов
type Posts struct {
	Title   string `form:"title,text:"`
	Content string `form:"content,text:"`
	Date    string `form:"-"`
}
