package main 

import (	
	"io/ioutil"
	"fmt"
	"time"
	"strconv"
)

func ReadTimeUsed() int {
	year, month, day := time.Now().Date()
	content, err := ioutil.ReadFile(fmt.Sprintf("%d-%d-%d.txt", year, month, day))
    if (err != nil) {
        return 0
	}
	result, err := strconv.Atoi(string(content))
	if (err != nil) {
		return 0
	} else {
		return result
	}
}

func IncreaseTimeUsed() {
	timeUsed := ReadTimeUsed() + 1
	year, month, day := time.Now().Date()
	ioutil.WriteFile(fmt.Sprintf("%d-%d-%d.txt", year, month, day), []byte(fmt.Sprintf("%d", timeUsed)), 0644)
}