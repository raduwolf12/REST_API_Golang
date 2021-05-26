package error

import (
	"errors"
	"net/http"
)

type RequestError struct {
	StatusCode int
	Err error
}

func (r *RequestError) Error() string {
	return r.Err.Error()
}

func (e *RequestError) Init() {
	e.StatusCode = http.StatusInternalServerError
	e.Err =errors.New("StatusInternalServerError!")
}

func infoR() *RequestError {
	e := new(RequestError)
	e.StatusCode = http.StatusInternalServerError
	e.Err = errors.New("StatusInternalServerError!")
	return e
}

func (r *RequestError) String() string {
	return "404 nu exista asa endpoint"
}

func (r *RequestError) GetStatusCode() interface{} {
	return r.StatusCode
}
func (r *RequestError) GetErr() interface{} {
	return r.Err
}