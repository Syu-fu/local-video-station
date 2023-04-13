package controllers_test

import (
	"testing"

	"api/controllers"
	"api/controllers/testdata"
)

var (
	aCon *controllers.VideoController
	tCon *controllers.TagController
)

func TestMain(m *testing.M) {
	vSer := testdata.NewServiceMock()
	tSer := testdata.NewTagServiceMock()
	aCon = controllers.NewVideoController(vSer)
	tCon = controllers.NewTagController(tSer)

	m.Run()
}
