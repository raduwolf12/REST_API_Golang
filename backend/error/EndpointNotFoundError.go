package error

import (
	"errors"
	"net/http"
)

type EndpointNotFoundError struct {
	StatusCode int
	Err error
}

func (e *EndpointNotFoundError) Init() {
	e.StatusCode = http.StatusNotFound
	e.Err = errors.New("not found!")
}

func infoNf() *EndpointNotFoundError {
	e := new(EndpointNotFoundError)
	e.StatusCode = http.StatusNotFound
	e.Err = errors.New("not found!")
	return e
}



func (r *EndpointNotFoundError) Error() string {
	return r.Err.Error()
}

func (r *EndpointNotFoundError) String() string {
	return "404 nu exista asa endpoint"
}

func (r *EndpointNotFoundError) GetStatusCode() interface{} {
	return r.StatusCode
}
func (r *EndpointNotFoundError) GetErr() interface{} {
	return r.Err
}