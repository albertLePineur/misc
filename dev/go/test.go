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

  var channel_message = make(chan string)
  var compteur int = 10

  go func(){
    var i int
    for i = 1; i <= compteur; i++ {
      channel_message <- fmt.Sprintf("Message %d", i)
    }
    close(channel_message)
  }()

  for message := range channel_message {
    fmt.Println(message)
  }


  var mon_canal1 = make(chan int)
  var mon_canal2 = make(chan string)

  go func(){
    var liste_entier = [8]int{10,20,30,40,50,60,70,80}
    for i := 0; i < len(liste_entier); i++ {
      mon_canal1 <- liste_entier[i]
    }
    close(mon_canal1)
  }()

  go func(){
    var liste_string = [8]string{"a","b","c","d","e","f","g","h"}
    for i := 0; i < len(liste_string); i++ {
      mon_canal2 <- liste_string[i]
    }
    close(mon_canal2)
  }()


  for msg1 := range mon_canal1 {
    println(msg1)
  }
  for msg2 := range mon_canal2 {
    println(msg2)
  }


}
