package controller

import (
	"github.com/nikhil-thomas/icndb-app-operator/pkg/controller/funapp"
)

func init() {
	// AddToManagerFuncs is a list of functions to create controllers and add them to a manager.
	AddToManagerFuncs = append(AddToManagerFuncs, funapp.Add)
}
