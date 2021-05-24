//package main
//
//import (
//	"encoding/json"
//	"fmt"
//	_ "github.com/lib/pq"
//	"log"
//	"net/http"
//)
//
//const (
//	host     = "localhost"
//	port     = 5432
//	user     = "postgres"
//	password = "postgres"
//	dbname   = "examen"
//)
//
//type Article struct {
//	Title string `json:"Title"`
//	Desc string `json:"desc"`
//	Content string `json:"content"`
//}
//
//var Articles []Article
//
//func homePage(w http.ResponseWriter, r *http.Request){
//	fmt.Fprintf(w, "Welcome to the HomePage!")
//	fmt.Println("Endpoint Hit: homePage")
//}
//
//func handleRequests() {
//
//	http.HandleFunc("/", homePage)
//	// add our articles route and map it to our
//	// returnAllArticles function like so
//	http.HandleFunc("/articles", returnAllArticles)
//	log.Fatal(http.ListenAndServe(":8080 ", nil))
//}
//
//func returnAllArticles(w http.ResponseWriter, r *http.Request){
//	fmt.Println("Endpoint Hit: returnAllArticles")
//	json.NewEncoder(w).Encode(Articles)
//}
//
//func main() {
//	Articles = []Article{
//		Article{Title: "Hello", Desc: "Article Description", Content: "Article Content"},
//		Article{Title: "Hello 2", Desc: "Article Description", Content: "Article Content"},
//	}
//	handleRequests()
//}
////func main() {
////	//
////	//psqlInfo := fmt.Sprintf("host=%s port=%d user=%s "+
////	//	"password=%s dbname=%s sslmode=disable",
////	//	host, port, user, password, dbname)
////	//db, err := sql.Open("postgres", psqlInfo)
////	//if err != nil {
////	//	panic(err)
////	//}
////	//defer db.Close()
////	//
////	//err = db.Ping()
////	//if err != nil {
////	//	panic(err)
////	//}
////	//
////	//fmt.Println("Successfully connected!")
////
////	//db.Query("SELECT echipa_id, denumire\n\tFROM public.echipe;")
////	Articles = []Article{
////		Article{Title: "Hello", Desc: "Article Description", Content: "Article Content"},
////		Article{Title: "Hello 2", Desc: "Article Description", Content: "Article Content"},
////	}
////	handleRequests()
////}