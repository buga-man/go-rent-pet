package main

import (
	"errors"
	"fmt"
	"net/http"
)

const port = ":8080"

func main() {
	// fmt.Println("Hello")
	fmt.Println("Start server")
	http.HandleFunc("/", Home)
	http.HandleFunc("/second", Second)
	http.HandleFunc("/divide", Divide)
	http.ListenAndServe(port, nil)
}

func Home(w http.ResponseWriter, r *http.Request) {
	n, err := fmt.Fprintf(w, "Hello")
	if err != nil {
		panic("AAAA")
	} else {
		fmt.Printf("%d bytes written\n", n)
	}
}

func Second(w http.ResponseWriter, r *http.Request) {
	n, err := fmt.Fprintf(w, "Hello from second")
	if err != nil {
		panic("AAAA")
	} else {
		fmt.Printf("%d bytes written\n", n)
	}
}

func DivideNumbers(x, y float32) (float32, error) {
	if y == 0.0 {
		err := errors.New("Zerodevisionerror")
		return 0, err
	}
	res := x / y
	return res, nil
}

func Divide(w http.ResponseWriter, r *http.Request) {
	result, error := DivideNumbers(2.3, 0)
	if error != nil {
		fmt.Fprintf(w, "Oh my godness, %s", error)
		return
	}
	_, err := fmt.Fprintf(w, "devision is: %f", result)
	if err != nil {
		panic("Oh my godness")

	}

}
