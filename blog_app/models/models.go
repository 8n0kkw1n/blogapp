package models

import "time"

//Posts - структура постов
type Posts struct {
	ID      int    `form:"-"`
	Title   string `form:"title,text:"`
	Content string `form:"content,text:"`
	Date    string `form:"-"`
}

//DateNow - сегодняшняя дата добавляем в блог
func DateNow() string {
	d := time.Now()
	date := d.Format("2006-01-02 15:04:05")
	return date
}
