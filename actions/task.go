package actions

import (
	"net/http"
	"todoapp/models"

	"github.com/gobuffalo/buffalo"
	"github.com/gobuffalo/pop/v5"
	"github.com/pkg/errors"
)

//  New renders the pago to create a new task default implementation.
func New(c buffalo.Context) error {
	task := models.Task{}

	c.Set("task", task)
	return c.Render(http.StatusOK, r.HTML("task/new.html"))
}

func List(c buffalo.Context) error {
	tasks := models.Tasks{}
	tx := c.Value("tx").(*pop.Connection)

	if err := tx.All(&tasks); err != nil {
		return err
	}

	c.Set("tasks", tasks)
	return c.Render(http.StatusOK, r.HTML("index.html"))
}

func Create(c buffalo.Context) error {
	task := models.Task{}
	tx := c.Value("tx").(*pop.Connection)

	form := struct {
		Name        string
		Description string
	}{}

	if err := c.Bind(&form); err != nil {
		return errors.WithStack(errors.Wrap(err, "add task bind error"))
	}

	task.Name = form.Name
	task.Description = form.Description

	if err := tx.Eager().Create(&task); err != nil {
		return errors.WithStack(errors.Wrap(err, "create task error"))
	}

	return c.Redirect(http.StatusSeeOther, "rootPath()")
}
