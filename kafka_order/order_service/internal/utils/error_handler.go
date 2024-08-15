package utils

import (
	"errors"
	"fmt"
)

var (
	ErrFindByID error
	ErrFindByName error
	ErrPathVar = errors.New("error when converting string to int for path variable")
)

func NewErrFindById(id int) error {
	ErrFindByID = fmt.Errorf("no package found with ID %d", id)
	return ErrFindByID
}

func NewErrFindByName(name string) (error) {
	ErrFindByName = fmt.Errorf("no package found with name %s", name)
	return ErrFindByName
}
