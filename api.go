package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"log"
	"net/http"
	"os"

	"gopkg.in/mgo.v2/bson"

	. "build-disk-stat/config"
	. "build-disk-stat/dao"
	. "build-disk-stat/models"
	"build-disk-stat/userip"

	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"
)

var config = Config{}
var dao = DisksDAO{}

// CreateDiskEndPoint POST a new disk object
func CreateDiskEndPoint(w http.ResponseWriter, r *http.Request) {
	defer r.Body.Close()
	//fmt.Printf(r.Method, r.URL.Path)
	var disk DisksJSON
	if err := json.NewDecoder(r.Body).Decode(&disk); err != nil {
		respondWithError(w, http.StatusBadRequest, "Invalid request payload")
		return
	}
	disk.ID = bson.NewObjectId()
	if err := dao.Insert(disk); err != nil {
		respondWithError(w, http.StatusInternalServerError, err.Error())
		return
	}
	respondWithJson(w, http.StatusCreated, map[string]string{"Status": "OK", "ID": disk.ID.Hex()})
}

//FindDiskEndpoint reutrn hash by its ID
func FindDiskEndpoint(w http.ResponseWriter, r *http.Request) {
	userIP, err := userip.FromRequest(r)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	fmt.Println("Request from ", userIP)
	params := mux.Vars(r)
	disk, err := dao.FindById(params["id"])
	if err != nil {
		respondWithError(w, http.StatusBadRequest, "Invalid Disk ID")
		return
	}
	respondWithJson(w, http.StatusOK, disk)
}

//DeleteDiskByID delete disk by ID and return result in JSON format
func DeleteDiskByID(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	err := dao.DeleteByID(params["id"])
	if err != nil {
		respondWithError(w, http.StatusBadRequest, "Invalid Disk ID")
		return
	}
	respondWithJson(w, http.StatusOK, map[string]string{"result": "Delete success"})
}

// respondWithError  return JSON with error message
func respondWithError(w http.ResponseWriter, code int, msg string) {
	respondWithJson(w, code, map[string]string{"error": msg})
}

// respondWithJson return response end code if request success
func respondWithJson(w http.ResponseWriter, code int, payload interface{}) {
	response, _ := json.Marshal(payload)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(code)
	w.Write(response)
}

// loggingHandleToStdout writing handler log to stdout
func loggingHandleToStdout(f func(http.ResponseWriter, *http.Request)) http.Handler {
	return handlers.LoggingHandler(os.Stdout, http.HandlerFunc(f))
}

func init() {
	//getting config file from argument
	var configPath string
	flag.StringVar(&configPath, "c", DefaultConfigPath, "usage -c config.toml")
	flag.Parse()

	config.Read(configPath)

	dao.Server = config.Server
	dao.Database = config.Database
	dao.Connect()
}

// Define HTTP request routes
func main() {
	r := mux.NewRouter()

	//r.Handle("/disks", loggingHandleToStdout(CreateDiskEndPoint)).Methods("POST")
	r.HandleFunc("/disks", CreateDiskEndPoint).Methods("POST")
	r.HandleFunc("/disks/{id}", FindDiskEndpoint).Methods("GET")
	r.HandleFunc("/remove/{id}", DeleteDiskByID).Methods("DELETE")

	loggedRouter := handlers.LoggingHandler(os.Stdout, r)
	if err := http.ListenAndServe(config.Listen, loggedRouter); err != nil {
		log.Fatal(err)
	}
}
