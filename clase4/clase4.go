package main

import (
	"fmt"
	"sync"
	"time"
)

// Worker1 gets from msg and response done
func Worker1(msg <-chan string, done chan<- struct{}) {
	time.Sleep(time.Second)
	s := <-msg
	fmt.Println(s)
	done <- struct{}{}
}

func typeSafety() {
	c := make(chan string)
	done := make(chan struct{})
	go Worker1(c, done)
	c <- "Hola mundo!"
	<-done
}

// Worker2 gets from msg1 and msg2
func Worker2(msg1 <-chan string, msg2 <-chan string) {
	timeout := time.After(time.Second * 3)
	for {
		select {
		case s := <-msg1:
			fmt.Println("msg1:", s)
		case s := <-msg2:
			fmt.Println("msg2:", s)
		default:
			select {
			case <-timeout:
				fmt.Println("Timeout!")
				return
			default:
				fmt.Println("Waiting...")
				time.Sleep(time.Millisecond * 400)
			}
		}
	}
}

func multiChannel() {
	c1 := make(chan string)
	c2 := make(chan string)

	go func() {
		for {
			time.Sleep(time.Second)
			c1 <- "Hola desde channel 1"
		}
	}()
	go func() {
		for {
			time.Sleep(time.Second * 2)
			c2 <- "Hola desde channel 2"
		}
	}()

	go Worker2(c1, c2)

	time.Sleep(time.Second * 4)
	fmt.Println("Done multiChannel!")
}

func nonb() {
	c := make(chan string)
	select {
	case s := <-c:
		fmt.Println(s)
	default:
		fmt.Println("Default!")
	}
}

// WorkerCC gets from msg and response done
func WorkerCC(msg <-chan string, done chan<- struct{}) {
	defer func() { done <- struct{}{} }()
	opened := true
	var s string
	// Es preferible utilizar un for range por que realiza esto por detras
	for opened {
		s, opened = <-msg
		fmt.Println("WorkerCC:", opened, s)
	}
}

func closeC() {
	msg := make(chan string)
	done := make(chan struct{})

	go func() {
		for i := 0; i < 5; i++ {
			time.Sleep(time.Second)
			msg <- "Hola desde msg!"
		}
		close(msg)
	}()
	go WorkerCC(msg, done)
	<-done
}

// WorkerWG does something
func WorkerWG() {
	time.Sleep(time.Millisecond * 200)
	fmt.Println("Hola Mundo!")
}

func waitGroup() {
	wg := sync.WaitGroup{}
	wg.Add(5)
	go func() {
		time.Sleep(time.Millisecond * 200)
		wg.Add(-5)
	}()
	wg.Wait()
	fmt.Println("Done first waitGroup!")
	typeSafety()

	wg.Add(1)
	go func() {
		defer wg.Done()
		WorkerWG()
	}()
	wg.Wait()
	fmt.Println("Done second waitGroup!")
}

func main() {
	fmt.Println("C4")

	fmt.Println()
	fmt.Println("TypeSafety")
	fmt.Println("-----------")
	typeSafety()

	fmt.Println()
	fmt.Println("MultiChannel")
	fmt.Println("-----------")
	multiChannel()

	fmt.Println()
	fmt.Println("NonB")
	fmt.Println("-----------")
	nonb()

	fmt.Println()
	fmt.Println("CloseC")
	fmt.Println("-----------")
	closeC()
	// Un channel asignado un nil bloquea siempre: c=nil luego <-c

	fmt.Println()
	fmt.Println("WaitGroup")
	fmt.Println("-----------")
	waitGroup()
}
