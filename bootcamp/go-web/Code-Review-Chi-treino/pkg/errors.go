package pkg

import (
	"fmt"
	"net/http"
)

// RegistrationExistsError representa um erro quando o registro já existe
type defaultErr struct {
	Message string
	Status  int
}

type errRegistrationAlreadyExists struct {
	defaultErr
}

var ErrRegistrationAlreadyExists error = fmt.Errorf(
	"%v", errRegistrationAlreadyExists{
		defaultErr{
			Message: "Registrarion already exists",
			Status:  http.StatusConflict,
		},
	},
)
