package models

import "github.com/gofrs/uuid"

func (ms *ModelSuite) Test_Task() {
	id, err := uuid.FromString("468d02bb-98ac-4496-af51-62c3c1f55530")
	ms.NoError(err)
	newTask := Task{
		ID:          id,
		Name:        "Make tests",
		Description: "Write all tests for actions",
	}

	ms.NoError(ms.DB.Create(&newTask))

}
