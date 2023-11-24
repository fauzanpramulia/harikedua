package main

// import "fmt"


//channel
// func main() {
// 	c:= make(chan string)
// 	go introduce("Airell", c)
// 	go introduce("Nanda", c)
// 	go introduce("Mailo", c)
// 	msg1 := <-c
// 	fmt.Println(msg1)
// 	msg2 := <-c
// 	fmt.Println(msg2)
// 	msg3 := <-c
// 	fmt.Println(msg3)

// 	close(c)
// }

// func introduce(student string, c chan string) {
// 	result := fmt.Sprintf("Hai, my name is %s", student)
// 	c <- result
// }


// func main(){
// 	c := make (chan string)
// 	students := []string{"Airell", "Mailo", "Indah"}
// 	for _, v := range students{
// 		go func (student string) {
// 			fmt.Println ("Student", student)
// 			result := fmt.Sprintf ("Hai, my name is %s", student)
// 			c <- result
// 		}(v)
// 	}

// 	for i := 1; i <= 3; i++{
// 		print(c)
// 	}
// 	close (c)
// }

// func print(c chan string) {
// 	fmt.Println(<-c)
// }

//======================================================================

// import (
// 	"fmt"
// 	"net/http"
// )

//  var PORT = ":8080"

// func main(){
// 	http.HandleFunc (("/api/hello"), greet)
// 	http.ListenAndServe (PORT, nil)
// }

// func greet(w http.ResponseWriter, r *http.Request) {
// 	msg := "Hello World"
// 	fmt.Fprint(w, msg)
// }