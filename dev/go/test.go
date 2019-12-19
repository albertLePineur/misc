package main

import "fmt"

func main() {
	//var termine = make(chan bool)
	//go func() {
	//  println("salut bande de mongoles, c'est la GoRoutine qui parle !!!")
	//  time.Sleep(30000 * time.Millisecond)
	//  termine <- true
	//}()
	//println("caca pipi prout")
	//<-termine

	//channel2 := make(chan int)
	//go func() {
	//  channel2 <- 1991
	//}()
	//valeur := <-channel2
	//println(valeur)

	var channelMessage = make(chan string)
	var compteur int = 10

	go func() {
		var i int
		for i = 1; i <= compteur; i++ {
			channelMessage <- fmt.Sprintf("Message %d", i)
		}
		close(channelMessage)
	}()

	for message := range channelMessage {
		fmt.Println(message)
	}

	var monCanal1 = make(chan int)
	var monCanal2 = make(chan string)

	go func() {
		var listeEntier = [8]int{10, 20, 30, 40, 50, 60, 70, 80}
		for i := 0; i < len(listeEntier); i++ {
			monCanal1 <- listeEntier[i]
		}
		close(monCanal1)
	}()

	go func() {
		var listeString = [8]string{"a", "b", "c", "d", "e", "f", "g", "h"}
		for i := 0; i < len(listeString); i++ {
			monCanal2 <- listeString[i]
		}
		close(monCanal2)
	}()

	for msg1 := range monCanal1 {
		println(msg1)
	}
	for msg2 := range monCanal2 {
		println(msg2)
	}

}
