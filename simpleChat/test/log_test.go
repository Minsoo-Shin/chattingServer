package test

import (
	"errors"
	"fmt"
	"log"
	"sync"
	"testing"
)

func Test_log(t *testing.T) {
	var err = errors.New("error happens")
	if err != nil {
		log.Printf("[log.Printf]: %v \n", err.Error())
	}
	if err != nil {
		fmt.Printf("[fmt.Printf]: %v \n", err.Error())
	}
}

var wg sync.WaitGroup

func Test_goroutineLog(t *testing.T) {
	wg.Add(1)
	go errHappen()
	wg.Wait()
}

func errHappen() {
	defer wg.Done()
	var err = errors.New("error happen")
	if err != nil {
		log.Printf("[log.Printf]: %v \n", err.Error())
	}
	if err != nil {
		fmt.Printf("[fmt.Printf]: %v \n", err.Error())
	}
}
