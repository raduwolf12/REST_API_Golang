package error
import (
	"errors"
	"net/http"
)

type LoginError struct {
	StatusCode int
	Err error
}

func (e *LoginError) Init() {
	e.StatusCode = http.StatusForbidden
	e.Err = errors.New("StatusForbidden!")
}

func infoLE() *LoginError {
	e := new(LoginError)
	e.StatusCode = http.StatusForbidden
	e.Err = errors.New("StatusForbidden!")
	return e
}



func (r *LoginError) Error() string {
	return r.Err.Error()
}

func (r *LoginError) String() string {
	return "404 nu exista asa endpoint"
}

func (r *LoginError) GetStatusCode() interface{} {
	return r.StatusCode
}
func (r *LoginError) GetErr() interface{} {
	return r.Err
}