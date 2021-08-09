package main

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
	"go.uber.org/zap"
)

var highThroughput = "High Throughput"
var mediumTroughput = "Medium Throughput"
var lowThroughput = "Low Throughput"
var log *zap.SugaredLogger
var port = ":8888"

var roundabout = ControlMethod{
	Method: "Roundabout",
	Efficent: map[string]int{
		highThroughput:  50,
		mediumTroughput: 75,
		lowThroughput:   90,
	},
}
var stopSigns = ControlMethod{
	Method: "Stop Signs",
	Efficent: map[string]int{
		highThroughput:  20,
		mediumTroughput: 30,
		lowThroughput:   40,
	},
}

var trafficLights = ControlMethod{
	Method: "Traffic Lights",
	Efficent: map[string]int{
		highThroughput:  90,
		mediumTroughput: 75,
		lowThroughput:   50,
	},
}

type ControlMethod struct {
	Method   string
	Efficent map[string]int
}

func handleRequests() {
	r := mux.NewRouter()
	r.HandleFunc("/", getControlMethod)

	log.Infof("starting running server on port: %s", port)
	err := http.ListenAndServe(port, r)
	if err != nil {
		log.Fatalf("Failed to serve http: %v", err)
	}
}

func getControlMethod(w http.ResponseWriter, r *http.Request) {
	query := r.URL.Query()
	log.Info("New Request with Query Params: ", query)
	north := query["north"]
	east := query["east"]
	south := query["south"]
	west := query["west"]

	n, err := strconv.Atoi(north[0])
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode("Error")
		log.Error(err)
	}

	e, err := strconv.Atoi(east[0])
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode("Error")
		log.Error(err)
	}

	s, err := strconv.Atoi(south[0])
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode("Error")
		log.Error(err)
	}

	we, err := strconv.Atoi(west[0])
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode("Error")
		log.Error(err)
	}

	totalCPM := calcualteTotalCPM(n, e, s, we)
	cm1 := calculateEffcienctControlMethod(totalCPM)

	log.Info(cm1.Method + " is recommeded")
	json.NewEncoder(w).Encode(cm1.Method + " is recommeded")
}
func main() {
	log = zap.NewExample().Sugar()
	defer log.Sync()

	handleRequests()
}

func calcualteTotalCPM(north int, east int, south int, west int) int {
	return north + east + south + west
}

func calculateEffcienctControlMethod(totalCPM int) ControlMethod {
	var cpmString string
	if totalCPM >= 20 {
		cpmString = highThroughput
	} else if 20 > totalCPM && totalCPM >= 10 {
		cpmString = mediumTroughput
	} else if totalCPM < 10 {
		cpmString = lowThroughput
	}

	controlMethods := []ControlMethod{
		roundabout,
		stopSigns,
		trafficLights,
	}
	log.Info(cpmString)
	return findMostEfficent(controlMethods, cpmString)
}

func findMostEfficent(controlMethods []ControlMethod, s string) ControlMethod {
	temp := 0
	var bestCM ControlMethod
	for _, cm := range controlMethods {
		log.Info(cm.Efficent[s])
		if cm.Efficent[s] > temp {

			temp = cm.Efficent[s]
			bestCM.Efficent = cm.Efficent
			bestCM.Method = cm.Method
		}
	}

	return bestCM
}
