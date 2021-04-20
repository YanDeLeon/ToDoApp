package actions

import (
	"net/http"
	"todoapp/models"

	"github.com/gobuffalo/buffalo"
	"github.com/gobuffalo/pop/v5"
	"github.com/gofrs/uuid"
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
	vQ := tx.PaginateFromParams(c.Params())
	vQ.Order("created_at asc")

	if err := vQ.All(&tasks); err != nil {
		return err
	}

	c.Set("tasks", tasks)
	return c.Render(http.StatusOK, r.HTML("index.html"))
}

func Create(c buffalo.Context) error {
	task := models.Task{}
	tx := c.Value("tx").(*pop.Connection)

	if err := c.Bind(&task); err != nil {
		return errors.WithStack(errors.Wrap(err, "add task bind error"))
	}

	if err := tx.Create(&task); err != nil {
		return errors.WithStack(errors.Wrap(err, "create task error"))
	}

	return c.Redirect(http.StatusSeeOther, "rootPath()")
}

func Show(c buffalo.Context) error {
	id, err := uuid.FromString(c.Param("id"))

	if err != nil {
		return errors.WithStack(errors.Wrap(err, "loading id error"))
	}

	task, err := findTask(c, id)

	if err != nil {
		return errors.WithStack(errors.Wrap(err, "finding task error"))
	}

	c.Set("task", task)

	return c.Render(http.StatusOK, r.HTML("task/show.html"))
}

func Edit(c buffalo.Context) error {
	id, err := uuid.FromString(c.Param("id"))

	if err != nil {
		return errors.WithStack(errors.Wrap(err, "loading id error"))
	}

	task, err := findTask(c, id)

	if err != nil {
		return errors.WithStack(errors.Wrap(err, "finding task error"))
	}

	c.Set("task", task)

	return c.Render(http.StatusOK, r.HTML("task/edit.html"))
}

func Delete(c buffalo.Context) error {
	id, err := uuid.FromString(c.Param("id"))

	if err != nil {
		return errors.WithStack(errors.Wrap(err, "loading id error"))
	}

	tx := c.Value("tx").(*pop.Connection)

	task, err := findTask(c, id)

	if err != nil {
		return errors.WithStack(errors.Wrap(err, "finding task error"))
	}

	if err := tx.Destroy(&task); err != nil {
		return errors.WithStack(errors.Wrap(err, "destroy task error"))
	}

	return c.Redirect(http.StatusSeeOther, "rootPath()")
}

//Update updates xD
func Update(c buffalo.Context) error {
	id, err := uuid.FromString(c.Param("id"))

	if err != nil {
		return errors.WithStack(errors.Wrap(err, "loading id error"))
	}

	tx := c.Value("tx").(*pop.Connection)

	task, err := findTask(c, id)

	if err != nil {
		return errors.WithStack(errors.Wrap(err, "finding task error"))
	}

	if err := c.Bind(&task); err != nil {
		return errors.WithStack(errors.Wrap(err, "add task bind error"))
	}

	if err := tx.Update(&task); err != nil {
		return errors.WithStack(errors.Wrap(err, "create task error"))
	}

	return c.Redirect(http.StatusSeeOther, "rootPath()")
}

func ChangeStatus(c buffalo.Context) error {
	id, err := uuid.FromString(c.Param("id"))

	if err != nil {
		return errors.WithStack(errors.Wrap(err, "loading id error"))
	}

	tx := c.Value("tx").(*pop.Connection)

	task, err := findTask(c, id)

	if err != nil {
		return errors.WithStack(errors.Wrap(err, "finding task error"))
	}

	task.Finished = !task.Finished

	if err := tx.Update(&task); err != nil {
		return errors.WithStack(errors.Wrap(err, "create task error"))
	}

	return c.Redirect(http.StatusSeeOther, "rootPath()")
}

func findTask(c buffalo.Context, id uuid.UUID) (models.Task, error) {

	tx := c.Value("tx").(*pop.Connection)

	task := models.Task{}

	if err := tx.Find(&task, id); err != nil {
		return task, err
	}

	return task, nil
}
