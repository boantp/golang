package main

import (
	"fmt"
	"strconv"
	"time"
)

var frameId = 0
var frameName = ""
var assemblyArrangement [3]string

func main() {
	assemblyArrangement[0] = "frame"
	assemblyArrangement[1] = "body"
	assemblyArrangement[2] = "interior"

	framesToCreate := len(assemblyArrangement)
	frameInfoChan := make(chan string)

	numberOfCars := 3
	for i := 1; i < numberOfCars+1; i++ {
		for stageNumber := 0; stageNumber < framesToCreate; stageNumber++ {
			go assemblyStage(frameInfoChan, assemblyArrangement[stageNumber], stageNumber, framesToCreate, i)
			time.Sleep(time.Millisecond * 1000)
		}
	}
}

func assemblyStage(frameInfoChan chan string, stage string, stageNumber int, framesToCreate int, carId int) {
	nextStage := "paint"
	if stageNumber < framesToCreate {
		frameName = "Frame ID" + strconv.Itoa(carId)
		if stageNumber != framesToCreate-1 {
			nextStage = assemblyArrangement[stageNumber+1]
		}
	}
	fmt.Println("Add", stage, "to", frameName, "and proceed to", nextStage)
	frameInfoChan <- frameName
	time.Sleep(time.Millisecond * 10)
}
