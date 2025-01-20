package application_test

import (
	"app/internal/application"
	"testing"

	"github.com/stretchr/testify/require"
)

var cfgAppDefault = &application.ConfigApplicationDefault{
	Addr: ":8080",
}

func Test_SetUp(t *testing.T) {
	app := application.NewApplicationDefault(cfgAppDefault)
	err := app.SetUp()
	require.Nil(t, err)
}

func Test_TearDown(t *testing.T) {
	app := application.NewApplicationDefault(cfgAppDefault)
	err := app.TearDown()
	require.Nil(t, err)
}
