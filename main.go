package main

import "fmt"

var highThroughput = "High Throughput"
var mediumTroughput = "Medium Throughput"
var lowThroughput = "Low Throughput"

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
		highThroughput:  50,
		mediumTroughput: 75,
		lowThroughput:   90,
	},
}

type ControlMethod struct {
	Method   string
	Efficent map[string]int
}

func main() {
	north := 5
	east := 5
	south := 5
	west := 0
	totalCPM := calcualteTotalCPM(north, east, south, west)
	cm1, cm2 := calculateEffcienctControlMethod(totalCPM)
	fmt.Println(cm1.Method)
	if cm2 != nil {
		fmt.Println(cm2.Method)
	}

}

func calcualteTotalCPM(north int, east int, south int, west int) int {
	return north + east + south + west
}

func calculateEffcienctControlMethod(totalCPM int) (*ControlMethod, *ControlMethod) {
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

	return findMostEfficent(controlMethods, cpmString)
}

func findMostEfficent(controlMethods []ControlMethod, s string) (*ControlMethod, *ControlMethod) {
	var temp int
	bestCM := &ControlMethod{}
	secondCM := &ControlMethod{}
	for _, cm := range controlMethods {
		if cm.Efficent[s] > temp {
			temp = cm.Efficent[s]
			bestCM.Method = cm.Method
			bestCM.Efficent = cm.Efficent
		} else if cm.Efficent[s] == temp {
			secondCM.Method = cm.Method
			secondCM.Efficent = cm.Efficent
		}
	}

	return bestCM, secondCM
}
