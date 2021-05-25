package main

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	_ "github.com/lib/pq"
	"html/template"
	"initialize/class"
	"log"
	"net/http"
	"strconv"
)

//
//type Pacient struct {
//		Id string `json:"Id"`
//		Nume string `json:"nume"`
//		Prenume string `json:"prenume"`
//		Mail string `json:"mail"`
//		Cnp string `json:"cnp"`
//	}
//type Doctor struct {
//	Id string `json:"Id"`
//	Nume string `json:"nume"`
//	Prenume string `json:"prenume"`
//	Mail string `json:"mail"`
//	Specialitate string `json:"specialitate"`
//	Cnp string `json:"cnp"`
//}

var Doctori []class.Doctor
var Pacienti []class.Pacient

func ReturnAllDoctori(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Endpoint Hit: returnAllArticles")
	json.NewEncoder(w).Encode(Doctori)
}
func ReturnAllPacienti(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Endpoint Hit: returnAllArticles")
	json.NewEncoder(w).Encode(Pacienti)
}

func homePage(w http.ResponseWriter, r *http.Request) {
	//fmt.Fprintf(w, "Welcome to the HomePage!")
	//fmt.Println("Endpoint Hit: homePage")

	http.ServeFile(w, r, "html/index.html")
}

func ReturnSinglePacient(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	key := vars["id"]

	for _, pacient := range Pacienti {
		id := strconv.Itoa(pacient.GetId())
		if id == key {
			json.NewEncoder(w).Encode(pacient)
			w.WriteHeader(http.StatusOK)
			return
		}
	}
}

func ReturnSingleDoctor(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	key := vars["id"]

	for _, doctor := range Doctori {
		id := strconv.Itoa(doctor.GetId())
		if id == key {
			json.NewEncoder(w).Encode(doctor)
			w.WriteHeader(http.StatusOK)
			return
		}
	}
}

type Handler func(http.ResponseWriter, *http.Request)

func (fn Handler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Headers", "Content-Type")
	//if e := fn(w, r); e != nil { // e is *appError, not os.Error.
	//	http.Error(w, e.Message, e.Code)
	//}
}

func login(w http.ResponseWriter, r *http.Request) {
	fmt.Println("method:", r.Method) //get request method
	if r.Method == "GET" {
		t, _ := template.ParseFiles("indexOld.html")
		t.Execute(w, nil)
	} else {
		r.ParseForm()
		// logic part of log in
		fmt.Println("cnp:", r.Form["cnp"])
		fmt.Println("email:", r.Form["email"])
	}

	key := r.Form["email"]
	key1 := r.Form["cnp"]

	for _, pacient := range Pacienti {
		if pacient.Mail == key[0] && pacient.Cnp == key1[0] {
			//json.NewEncoder(w).Encode(pacient)
			http.ServeFile(w, r, "html/dashBoardPacienti.html")
			w.WriteHeader(http.StatusOK)
			return
		}
	}
	errorHandler(w, r, http.StatusForbidden)
}

func loginDoctor(w http.ResponseWriter, r *http.Request) {
	fmt.Println("method:", r.Method) //get request method
	if r.Method == "GET" {
		t, _ := template.ParseFiles("indexOld.html")
		t.Execute(w, nil)
	} else {
		r.ParseForm()
		// logic part of log in
		fmt.Println("cnp:", r.Form["cnp"])
		fmt.Println("email:", r.Form["email"])
	}

	key := r.Form["email"]
	key1 := r.Form["cnp"]

	for _, doctor := range Doctori {
		if doctor.Mail == key[0] && doctor.Cnp == key1[0] {
			//json.NewEncoder(w).Encode(doctor)
			http.ServeFile(w, r, "html/dashBoardDoctori.html")
			w.WriteHeader(http.StatusOK)
			return
		}
	}
	errorHandler(w, r, http.StatusForbidden)
}



func errorHandler(w http.ResponseWriter, r *http.Request, status int) {
	w.WriteHeader(status)
	if status == http.StatusNotFound {
		http.ServeFile(w, r, "html/ups.html")
	}
	if status == http.StatusForbidden {
		http.ServeFile(w, r, "html/ups.html")
	}
}

func handleRequests() {

	myRouter := mux.NewRouter().StrictSlash(true)

	myRouter.PathPrefix("/js/").Handler(http.StripPrefix("/js/", http.FileServer(http.Dir("html/js"))))
	myRouter.PathPrefix("/css/").Handler(http.StripPrefix("/css/", http.FileServer(http.Dir("html/css"))))
	myRouter.PathPrefix("/images/").Handler(http.StripPrefix("/images/", http.FileServer(http.Dir("html/images"))))
	myRouter.PathPrefix("/fonts/").Handler(http.StripPrefix("/fonts/", http.FileServer(http.Dir("html/fonts"))))

	myRouter.HandleFunc("/", homePage)
	myRouter.HandleFunc("/alldoctori", ReturnAllDoctori)
	myRouter.HandleFunc("/allpacienti", ReturnAllPacienti)
	myRouter.HandleFunc("/doctor/{id}", ReturnSingleDoctor)
	myRouter.HandleFunc("/pacient/{id}", ReturnSinglePacient)
	myRouter.HandleFunc("/login", login)
	myRouter.HandleFunc("/loginDoctor", loginDoctor)

	//myRouter.HandleFunc("/login/{email}/{cnp}", Handler(loginPacient))
	log.Fatal(http.ListenAndServe(":10000", myRouter))
}

const (
	host     = "localhost"
	port     = 5432
	user     = "postgres"
	password = "postgres"
	dbname   = "goblana"
)

type ProductModel struct {
	Db *sql.DB
}

func (productModel ProductModel) FindAllDoctori() ([]class.Doctor, error) {
	rows, err := productModel.Db.Query("SELECT * FROM public.doctori;")
	if err != nil {
		return [] class.Doctor{}, err
	} else {
		var product []class.Doctor
		for rows.Next() {
			var id int
			var nume string
			var prenume string
			var specialitate string
			var mail string
			var cnp string
			err2 := rows.Scan(&id, &nume, &prenume,  &specialitate,&mail, &cnp)
			if err2 != nil {
				return [] class.Doctor{}, err2
			} else {
				product = append(product, class.Doctor{id, nume, prenume, mail, specialitate, cnp})
			}
		}
		return product, nil
	}
}

func (productModel ProductModel) FindAllPacienti() ([]class.Pacient, error) {
	rows, err := productModel.Db.Query("SELECT * FROM public.pacienti;")
	if err != nil {
		return [] class.Pacient{}, err
	} else {
		var product []class.Pacient
		for rows.Next() {
			var id int
			var nume string
			var prenume string
			var mail string
			var cnp string
			err2 := rows.Scan(&id, &nume, &prenume, &mail, &cnp)
			if err2 != nil {
				return [] class.Pacient{}, err2
			} else {
				product = append(product, class.Pacient{id, nume, prenume, mail, cnp})
			}
		}
		return product, nil
	}
}

func init(){
	psqlInfo := fmt.Sprintf("host=%s port=%d user=%s "+"password=%s dbname=%s sslmode=disable", host, port, user, password, dbname)
	db, err := sql.Open("postgres", psqlInfo)
	if err != nil {
		panic(err)
	}
	defer db.Close()

	err = db.Ping()
	if err != nil {
		panic(err)
	}

	fmt.Println("Successfully connected!")


	if err != nil {
		fmt.Println(err)
	} else {
		productModel := ProductModel{
			Db: db,
		}
		doctori, err2 := productModel.FindAllDoctori()
		pacienti, err3 :=productModel.FindAllPacienti()

		if err2 != nil {
			fmt.Println(err2)
		} else if err3 != nil {
			fmt.Println(err3)
		} else {
			Doctori = doctori
			Pacienti = pacienti
		}
	}
}

func main() {
	fmt.Println("Spital blana!")


	//psqlInfo := fmt.Sprintf("host=%s port=%d user=%s "+"password=%s dbname=%s sslmode=disable", host, port, user, password, dbname)
	//db, err := sql.Open("postgres", psqlInfo)
	//if err != nil {
	//	panic(err)
	//}
	//defer db.Close()
	//
	//err = db.Ping()
	//if err != nil {
	//	panic(err)
	//}
	//
	//fmt.Println("Successfully connected!")
	//
	//
	//if err != nil {
	//	fmt.Println(err)
	//} else {
	//	productModel := ProductModel{
	//		Db: db,
	//	}
	//	doctori, err2 := productModel.FindAllDoctori()
	//	pacienti, err3 :=productModel.FindAllPacienti()
	//
	//	if err2 != nil {
	//		fmt.Println(err2)
	//	} else if err3 != nil {
	//		fmt.Println(err3)
	//	} else {
	//		Doctori = doctori
	//		Pacienti = pacienti
	//	}
	//}

	handleRequests()

}
