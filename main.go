package main

import (
	"fmt"
	"html/template"
	"net/http"
	"runtime"
)

const (
	Port = ":8080"
)

func serveStatic(w http.ResponseWriter, r *http.Request) {
	t, err := template.ParseFiles("main.html")
	if err != nil {
		fmt.Println(err)
	}
	items := struct {
		Country  string
		City     string
		message1 string
		message2 string
		message3 string
		message4 string
		message5 string
		message6 string
		message7 string
		Test     string
	}{
		Country:  "Australia",
		City:     "Paris1",
		message1: "Hello Brad",
		message2: "Can You Buy Me",
		message3: "could I get it for 60 - 60 =",
		message4: "DOLLARS or PESOS",
		message5: "please buy me discord nitro one year plessssssssss",
		message6: "PESOS",
		message7: "If my gpu's size is 12 inches and matthew's ps4 gpu's size is at least 1 inch. Then who has the bigger size?",
		Test:     "Testing if this message gets printed",
	}

	t.Execute(w, items)
}

func main() {

	//https
	http.HandleFunc("/", serveStatic)
	http.ListenAndServe(Port, nil)

	//Variables
	var game string = "DEAD ISLAND 2"
	var face string = "please please please :)"
	var price int = 60
	var moneysymbol string = "$$$$$$"
	var nitromes string = "please buy me discord nitro one year"
	var nitroprice float32 = 2639.99
	var MyGPUSize int = 12
	var matthewGPUSize int = 10
	var DhanGPU int = 11

	//Strings
	// var message1 string = "Hello Brad"
	// var message2 string = "Can You Buy Me"
	// var message3 string = "could I get it for 60 - 60 ="
	// var message4 string = "DOLLARS or PESOS"
	// var message5 string = "please buy me discord nitro one year plessssssssss"
	// var message6 string = "PESOS"
	// var message7 string = "If my gpu's size is 12 inches and matthew's ps4 gpu's size is at least 1 inch. Then who has the bigger size?"

	//constants
	const gameprize = 60

	//print line
	println("Hello Brad")
	println("Can You Buy Me")
	println(game)
	println(face)
	println(price)
	println(moneysymbol)
	println("could I get it for 60 - 60 =")
	println(price - gameprize)
	println("DOLLARS or PESOS")
	println("please buy me discord nitro one year plessssssssss")
	println(nitroprice)
	println("PESOS")
	println(nitromes)

	println("If my gpu's size is 12 inches and matthew's ps4 gpu's size is at least 1 inch. Then who has the bigger size?")

	// if MyGPUSize >= matthewGPUSize {
	// 	println("My DIck Is WaY bigGEr than matthew")
	// } else if MyGPUSize == matthewGPUSize {
	// 	println("WE BOTH HAVE the same size LOL :)")
	// } else {
	// 	println("matthew's dick is way bigger than aldrien")
	// }

	println("NOW this is Switch Cases")

	switch {
	case MyGPUSize >= matthewGPUSize || DhanGPU >= matthewGPUSize:
		println("Aldriens GPU Is bigger")
	case MyGPUSize >= 9:
		println("GPU BIG???")
	default:
		println("An Error Has Occured :)))))")

	}

	switch os := runtime.GOOS; os {
	case "darwin":
		println("Your RUnning at Gay MackoOSSSASS")
	case "linux":
		println("Linux Hacking Machine :)")
	default:
		println("You Could be running at Windows 10")
	}

	//An Array Of Best Game Consoles
	GameConsole := [6]string{"PS4", "PS5", "PS3", "XBOXONE", "XSX", "Xbox360"}
	fmt.Println(GameConsole)

}
