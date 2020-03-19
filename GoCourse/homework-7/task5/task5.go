package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

type Racer struct {
	Speed         int
	ReactionSpeed int
	ResultScore   time.Duration
}

func main() {

	s1 := rand.NewSource(time.Now().UnixNano())
	s2 := rand.NewSource(time.Now().UnixNano())
	s3 := rand.NewSource(time.Now().UnixNano())
	s4 := rand.NewSource(time.Now().UnixNano())
	r1 := rand.New(s1)
	r2 := rand.New(s2)
	r3 := rand.New(s3)
	r4 := rand.New(s4)

	startChan := make(chan struct{})
	var wg sync.WaitGroup
	wg.Add(4)

	racer1 := &Racer{
		Speed:         r1.Intn(300),
		ReactionSpeed: r1.Intn(100),
		ResultScore:   0,
	}
	racer2 := &Racer{
		Speed:         r2.Intn(300),
		ReactionSpeed: r2.Intn(100),
		ResultScore:   0,
	}
	racer3 := &Racer{
		Speed:         r3.Intn(300),
		ReactionSpeed: r3.Intn(100),
		ResultScore:   0,
	}
	racer4 := &Racer{
		Speed:         r4.Intn(300),
		ReactionSpeed: r4.Intn(100),
		ResultScore:   0,
	}

	raceDistance := 10000

	go PrepareRacer(racer1, raceDistance, startChan, &wg)
	go PrepareRacer(racer2, raceDistance, startChan, &wg)
	go PrepareRacer(racer3, raceDistance, startChan, &wg)
	go PrepareRacer(racer4, raceDistance, startChan, &wg)

	close(startChan)
	wg.Wait()

	fmt.Println("Results Racer1: ", racer1.ResultScore.String(),
		"\nResults Racer2: ", racer2.ResultScore.String(),
		"\nResults Racer3: ", racer3.ResultScore.String(),
		"\nResults Racer4: ", racer4.ResultScore.String())
}

func PrepareRacer(racer *Racer, raceDistance int, startChan chan struct{}, wg *sync.WaitGroup) {
	<-startChan
	racer.ResultScore = time.Duration(raceDistance/racer.Speed + racer.ReactionSpeed)
	wg.Done()
}
