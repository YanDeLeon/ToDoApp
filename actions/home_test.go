package actions

import "net/http"

func (as *ActionSuite) Test_HomeHandler() {
	res := as.HTML("/").Get()

	as.Equal(http.StatusOK, res.Code)
}
