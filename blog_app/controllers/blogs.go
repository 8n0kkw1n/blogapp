package controllers

import (
	"database/sql"
	"fmt"
	"hwgo/5/blog_app/models"
	"log"

	"github.com/astaxie/beego"
)

//BlogsController - all blogs controller
type BlogsController struct {
	beego.Controller
	Db *sql.DB
}

//Get -
func (c *BlogsController) Get() {
	posts, err := getAllPosts(c.Db)
	if err != nil {
		log.Println(err)
		return
	}

	c.Data["Posts"] = posts
	c.TplName = "index.tpl"
}

//getAllPosts — получение всего списка
func getAllPosts(db *sql.DB) ([]models.Posts, error) {
	rows, err := db.Query("select * from `gb_blog`.`blog_lists` order by date desc")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	res := make([]models.Posts, 0, 1)
	for rows.Next() {
		post := models.Posts{}

		if err := rows.Scan(&post.ID, &post.Title, &post.Content, &post.Date); err != nil {
			log.Println(err)
			continue
		}

		res = append(res, post)
	}

	return res, nil
}

//Open -
func (c *BlogsController) Open() {
	id := c.Ctx.Request.URL.Query().Get("id")
	if len(id) == 0 {
		log.Print("id is empty")
		return
	}

	post, err := getPost(c.Db, id)
	if err != nil {
		c.Ctx.ResponseWriter.WriteHeader(404)
		return
	}

	c.Data["Post"] = post
	c.TplName = "post.tpl"
}

//Write - открываем форму поста
func (c *BlogsController) Write() {
	c.TplName = "write.tpl"
}

//getPost — открытие поста
func getPost(db *sql.DB, id string) (models.Posts, error) {
	post := models.Posts{}

	row := db.QueryRow(fmt.Sprintf("select * from `gb_blog`.`blog_lists` where `blog_lists`.id = %v", id))
	err := row.Scan(&post.ID, &post.Title, &post.Content, &post.Date)
	if err != nil {
		return post, err
	}

	return post, nil
}

//Post -
func (c *BlogsController) Post() {
	post := models.Posts{}
	if err := c.ParseForm(&post); err != nil {
		beego.Error("Could parse form", err)
	} else {
		c.Data["Posts"] = post
	}

	if err := createPost(c.Db, post); err != nil {
		c.Ctx.ResponseWriter.WriteHeader(500)
	}

	c.Redirect("/", 302)
}

func createPost(db *sql.DB, post models.Posts) error {
	post.Date = models.DateNow()

	res, err := db.Exec("insert into `gb_blog`.`blog_lists` (title, content, date) values (?, ?, ?)",
		post.Title, post.Content, post.Date)

	affect, err := res.RowsAffected()
	if err != nil {
		return err
	}

	fmt.Println("Insert:", affect)
	return nil
}

//Edit -
func (c *BlogsController) Edit() {
	id := c.Ctx.Request.URL.Query().Get("id")
	if len(id) == 0 {
		log.Print("empty id")
		return
	}

	post, err := getPost(c.Db, id)

	if err != nil {
		c.Ctx.ResponseWriter.WriteHeader(404)
		return
	}
	c.Data["Post"] = post
	c.TplName = "edit.tpl"
}

//EditPost -
func (c *BlogsController) EditPost() {
	id := c.Ctx.Input.Param(":id")
	if len(id) == 0 {
		log.Print("empty id")
		return
	}

	post := models.Posts{}
	if err := c.ParseForm(&post); err != nil {
		beego.Error("Could parse form", err)
	} else {
		c.Data["Posts"] = post
	}

	if err := updatePost(c.Db, id, post); err != nil {
		c.Ctx.ResponseWriter.WriteHeader(500)
	}

	c.Redirect("/", 302)
}

//updatePost -
func updatePost(db *sql.DB, id string, post models.Posts) error {
	post.Date = models.DateNow()

	res, err := db.Exec("UPDATE `gb_blog`.`blog_lists` SET title = ?, content = ?, date = ? WHERE id = ?",
		post.Title, post.Content, post.Date, id)

	affect, err := res.RowsAffected()
	if err != nil {
		return err
	}

	fmt.Println("Insert:", affect)
	return err
}

//DeletePost -
func (c *BlogsController) DeletePost() {
	id := c.Ctx.Input.Param(":id")
	if len(id) == 0 {
		log.Print("empty id")
		return
	}

	err := deletePost(c.Db, id)
	if err != nil {
		return
	}

	c.Redirect("/", 302)
}

//deletePost - удаление поста
func deletePost(db *sql.DB, id string) error {
	res, err := db.Exec("DELETE FROM `gb_blog`.`blog_lists` WHERE id = ?", id)
	if err != nil {
		return err
	}

	affect, err := res.RowsAffected()
	if err != nil {
		return err
	}

	fmt.Println("Insert:", affect)
	return err
}
