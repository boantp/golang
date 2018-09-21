package main

import (
	"fmt"
	"strconv"
	"time"
)

var frameId = 0
var frameName = ""

func main() {
	framesToCreate := 4
	frameInfoChan := make(chan string)

	for i := 0; i < framesToCreate; i++ {
		go assembleFrame(frameInfoChan)
		go addBody(frameInfoChan)
		go addInterior(frameInfoChan)
		time.Sleep(time.Millisecond * 1000)
	}
}

func assembleFrame(frameInfoChan chan string) {
	frameId++
	frameName = "Frame ID" + strconv.Itoa(frameId)
	fmt.Print("Frame assembly complete", frameName, " Proceed to body")
	frameInfoChan <- frameName
	time.Sleep(time.Millisecond * 5)
}

func addBody(frameInfoChan chan string) {
	body := <-frameInfoChan
	fmt.Println("Add body to", body, "and Proceed to interior")
	frameInfoChan <- frameName
	time.Sleep(time.Millisecond * 5)
}

func addInterior(frameInfoChan chan string) {
	interior := <-frameInfoChan
	fmt.Println("Add interior to", interior, "and Proceed to paint")
	time.Sleep(time.Millisecond * 5)
}
