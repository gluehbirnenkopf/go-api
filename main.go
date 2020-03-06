package main

import (
  "fmt"
  "log"
  "net/http"
  "github.com/gorilla/mux"
  "io/ioutil"
)

func homePage(w http.ResponseWriter, r *http.Request){
    fmt.Fprintf(w, "Welcome to the free stock screener!")
    fmt.Println("Endpoint Hit: homePage")
}

func handleRequests() {
    myRouter := mux.NewRouter().StrictSlash(true)
    myRouter.HandleFunc("/", homePage)
    myRouter.HandleFunc("/getStockValue", getStockValue).Methods("POST")
    log.Fatal(http.ListenAndServe(":10000", myRouter))
}

func getStockValue(w http.ResponseWriter, r *http.Request) {
    // get the body of our POST request
    // return the string response containing the request body
    reqBody, _ := ioutil.ReadAll(r.Body)
    fmt.Fprintf(w, "%+v", string(reqBody))
}

func main() {
    handleRequests()
}
