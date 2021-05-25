package main


import (
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

)

func TestLoginPacientiOK(t *testing.T) {
	req, err := http.NewRequest("POST", "/login?email=mionesc@yahoo.com&cnp=123456789", nil)
	if err != nil {
		t.Fatal(err)
	}
	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(login)
	handler.ServeHTTP(rr, req)
	if status := rr.Code; status != http.StatusOK {
		t.Errorf("handler returned wrong status code: got %v want %v",
			status, http.StatusOK)
	}
}
func TestLoginDoctoriOK(t *testing.T) {
	req, err := http.NewRequest("POST", "/loginDoctor?email=mionesc@yahoo.com&cnp=123456789", nil)
	if err != nil {
		t.Fatal(err)
	}
	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(login)
	handler.ServeHTTP(rr, req)
	if status := rr.Code; status != http.StatusOK {
		t.Errorf("handler returned wrong status code: got %v want %v",
			status, http.StatusOK)
	}
}


func TestLoginPacientiStatusForbiden(t *testing.T) {
	req, err := http.NewRequest("POST", "/login?email=mionesc@yahoo.com&cnp=12", nil)
	if err != nil {
		t.Fatal(err)
	}
	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(login)
	handler.ServeHTTP(rr, req)
	if status := rr.Code; status != http.StatusForbidden {
		t.Errorf("handler returned wrong status code: got %v want %v",
			status, http.StatusForbidden)
	}
}
func TestLoginDoctoriStatusForbiden(t *testing.T) {
	req, err := http.NewRequest("POST", "/loginDoctor?email=mionesc@yahoo.com&cnp=231", nil)
	if err != nil {
		t.Fatal(err)
	}
	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(login)
	handler.ServeHTTP(rr, req)
	if status := rr.Code; status != http.StatusForbidden {
		t.Errorf("handler returned wrong status code: got %v want %v",
			status, http.StatusForbidden)
	}
}

func TestGetAllDoctori(t *testing.T) {
	req, err := http.NewRequest("POST", "/alldoctori", nil)
	if err != nil {
		t.Fatal(err)
	}
	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(ReturnAllDoctori)
	handler.ServeHTTP(rr, req)
	if status := rr.Code; status != http.StatusOK {
		t.Errorf("handler returned wrong status code: got \n %v want \n %v",
			status, http.StatusOK)
	}

	// Check the response body is what we expect.
	expected := "[{\"Id\":1,\"nume\":\"Mariana\",\"prenume\":\"Ionescu\",\"mail\":\"mionesc@yahoo.com\",\"specialitate\":\"Dentist\",\"cnp\":\"123456789\"},{\"Id\":2,\"nume\":\"Ion\",\"prenume\":\"Marin\",\"mail\":\"irad@yahoo.com\",\"specialitate\":\"Ortoped\",\"cnp\":\"123456789\"}]"
	if strings.TrimRight(rr.Body.String(), "\n") != expected {
		t.Errorf("handler returned unexpected body: got \n %v want \n %v",
			rr.Body.String(), expected)
	}
}

func TestGetAllPacienti(t *testing.T) {
	req, err := http.NewRequest("POST", "/allpacienti", nil)
	if err != nil {
		t.Fatal(err)
	}
	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(ReturnAllPacienti)
	handler.ServeHTTP(rr, req)
	if status := rr.Code; status != http.StatusOK {
		t.Errorf("handler returned wrong status code: got \n %v want \n %v",
			status, http.StatusOK)
	}

	// Check the response body is what we expect.
	expected := "[{\"Id\":1,\"nume\":\"Mariana\",\"prenume\":\"Petrescu\",\"mail\":\"mionesc@yahoo.com\",\"cnp\":\"123456789\"},{\"Id\":2,\"nume\":\"Ion\",\"prenume\":\"Raducan\",\"mail\":\"irad@yahoo.com\",\"cnp\":\"123456789\"},{\"Id\":3,\"nume\":\"Ion\",\"prenume\":\"Costin\",\"mail\":\"icost@yahoo.com\",\"cnp\":\"123456789\"}]"
	if strings.TrimRight(rr.Body.String(), "\n") != expected {
		t.Errorf("handler returned unexpected body: got \n %v want \n %v",
			rr.Body.String(), expected)
	}
}