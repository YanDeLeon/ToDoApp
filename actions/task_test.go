package actions

import (
	"net/http"
	"net/url"
	"todoapp/models"

	"github.com/gofrs/uuid"
)

func (as *ActionSuite) Test_List_Task() {
	newTask := models.Task{
		Name:        "Make tests",
		Description: "Write all tests for actions",
	}

	as.NoError(as.DB.Create(&newTask))

	res := as.HTML("/").Get()
	as.Equal(http.StatusOK, res.Code)
	as.Contains(res.Body.String(), "Make tests")
	as.Contains(res.Body.String(), "tests for actions")
}

func (as *ActionSuite) Test_New_Task() {
	res := as.HTML("/task/new").Get()
	as.Equal(http.StatusOK, res.Code)
	as.Contains(res.Body.String(), "Create New Task")
	as.Contains(res.Body.String(), "Description")
}

func (as *ActionSuite) Test_Create_Task() {
	res := as.HTML("/task/new").Get()
	as.Equal(http.StatusOK, res.Code)

	task := url.Values{
		"Name":        []string{"Make tests"},
		"Description": []string{"Write all tests for actions"},
	}

	as.TableChange("tasks", 1, func() {
		res := as.HTML("/task/create").Put(task)
		as.Equal(http.StatusSeeOther, res.Code)
		as.Equal(res.Location(), "/")
	})

}

func (as *ActionSuite) Test_Show_Task() {
	id, err := uuid.FromString("468d02bb-98ac-4496-af51-62c3c1f55530")
	as.NoError(err)
	newTask := models.Task{
		ID:          id,
		Name:        "Make tests",
		Description: "Write all tests for actions",
	}

	as.NoError(as.DB.Create(&newTask))

	res := as.HTML("/").Get()
	as.Equal(http.StatusOK, res.Code)

	res = as.HTML("/task/468d02bb-98ac-4496-af51-62c3c1f55530/show").Get()
	as.Equal(http.StatusOK, res.Code)
	as.Contains(res.Body.String(), "Make tests")
	as.Contains(res.Body.String(), "tests for actions")
}

func (as *ActionSuite) Test_Edit_Task() {
	id, err := uuid.FromString("468d02bb-98ac-4496-af51-62c3c1f55530")
	as.NoError(err)
	newTask := models.Task{
		ID:          id,
		Name:        "Make tests",
		Description: "Write all tests for actions",
	}

	as.NoError(as.DB.Create(&newTask))

	res := as.HTML("/").Get()
	as.Equal(http.StatusOK, res.Code)

	res = as.HTML("/task/468d02bb-98ac-4496-af51-62c3c1f55530/edit").Get()
	as.Equal(http.StatusOK, res.Code)
	as.Contains(res.Body.String(), "Edit Task:")

}

func (as *ActionSuite) Test_Update_Task() {
	id, err := uuid.FromString("468d02bb-98ac-4496-af51-62c3c1f55530")
	as.NoError(err)
	newTask := models.Task{
		ID:          id,
		Name:        "Make tests",
		Description: "Write all tests for actions",
	}

	as.NoError(as.DB.Create(&newTask))

	res := as.HTML("/").Get()
	as.Equal(http.StatusOK, res.Code)

	res = as.HTML("/task/468d02bb-98ac-4496-af51-62c3c1f55530/edit").Get()
	as.Equal(http.StatusOK, res.Code)
	values := url.Values{
		"Name":        []string{"New Name"},
		"Description": []string{"New description"},
	}
	as.TableChange("tasks", 0, func() {
		res = as.HTML("/task/468d02bb-98ac-4496-af51-62c3c1f55530/update").Put(values)
		as.Equal(http.StatusSeeOther, res.Code)
	})

	res = as.HTML("/").Get()
	as.Contains(res.Body.String(), "New Name")
	as.Contains(res.Body.String(), "New description")

}

func (as *ActionSuite) Test_Delete_Task() {
	id, err := uuid.FromString("468d02bb-98ac-4496-af51-62c3c1f55530")
	as.NoError(err)
	newTask := models.Task{
		ID:          id,
		Name:        "Make tests",
		Description: "Write all tests for actions",
	}

	as.NoError(as.DB.Create(&newTask))

	as.TableChange("tasks", -1, func() {
		res := as.HTML("/task/468d02bb-98ac-4496-af51-62c3c1f55530/delete").Delete()
		as.Equal(http.StatusSeeOther, res.Code)
		as.Equal(res.Location(), "/")
	})

	res := as.HTML("/").Get()
	as.Equal(http.StatusOK, res.Code)
	as.NotContains(res.Body.String(), "Make tests")
	as.NotContains(res.Body.String(), "Write all tests for actions")
}

func (as *ActionSuite) Test_Change_Task_Status() {
	id, err := uuid.FromString("468d02bb-98ac-4496-af51-62c3c1f55530")
	as.NoError(err)
	newTask := models.Task{
		ID:          id,
		Name:        "Make tests",
		Description: "Write all tests for actions",
	}

	as.NoError(as.DB.Create(&newTask))

	res := as.HTML("/").Get()
	as.Equal(http.StatusOK, res.Code)
	as.Contains(res.Body.String(), "Pending")

	as.TableChange("tasks", 0, func() {
		res := as.HTML("/task/468d02bb-98ac-4496-af51-62c3c1f55530/update-status").Put(newTask)
		as.Equal(http.StatusSeeOther, res.Code)
		as.Equal(res.Location(), "/")
	})

	res = as.HTML("/").Get()
	as.Equal(http.StatusOK, res.Code)
	as.Contains(res.Body.String(), "Finished")

}
