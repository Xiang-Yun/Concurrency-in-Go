package main

import (
	"fmt"
	"sync"
	"time"
)

type ChopS struct {
	sync.Mutex
}

type Philo struct {
	Id              int
	leftCS, rightCS *ChopS
}

func (p Philo) eat(wg *sync.WaitGroup, host chan bool) {
	defer wg.Done()
	for i := 0; i < 3; i++ {
		host <- true //Request permission to eat

		p.leftCS.Lock()
		p.rightCS.Lock()

		fmt.Println("starting to eat", p.Id)
		time.Sleep(time.Millisecond * 100)
		fmt.Println("finishing eating", p.Id)

		p.rightCS.Unlock()
		p.leftCS.Unlock()

		<-host // Release permission

		time.Sleep(time.Millisecond * 100)
	}
}

func main() {
	// create chopsticks
	CSticks := make([]*ChopS, 5)
	for i := 0; i < 5; i++ {
		CSticks[i] = new(ChopS)
	}

	// create philosophers
	philos := make([]*Philo, 5)
	for i := 0; i < 5; i++ {
		philos[i] = &Philo{
			Id:      i + 1,
			leftCS:  CSticks[i],
			rightCS: CSticks[(i+1)%5],
		}
	}

	// Host channel to allow no more than 2 philosophers to eat concurrently
	host := make(chan bool, 2)

	// start philosophers
	var wg sync.WaitGroup
	for i := 0; i < 5; i++ {
		wg.Add(1)
		go philos[i].eat(&wg, host)
	}

	// wait for all philosophers to finish
	wg.Wait()
}
