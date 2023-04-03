package controllers_test

import (
	"testing"

	"api/controllers"
	"api/controllers/testdata"
)

var aCon *controllers.VideoController

func TestMain(m *testing.M) {
	ser := testdata.NewServiceMock()
	aCon = controllers.NewVideoController(ser)

	m.Run()
}
