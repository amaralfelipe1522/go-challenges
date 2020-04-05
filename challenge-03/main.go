package main

/*
Acceptance Criteria
In order to successfully complete this challenge, your project will have to:

Use Go Modules as a means of managing your dependencies
Contain a main.go file which references a module stats which is contained within a stats/stats.go file within your project.
This stats/stats.go file must contain the logic for collecting the hardware utilization stats from your machine and exposing them as a HTTP function.
*/
import (
	"log"
	"net/http"

	stats "github.com/amaralfelipe1522/status-pc"
)

func main() {
	//fmt.Println(stats.PrintStatus())
	http.HandleFunc("/", stats.PrintStatusHTTP)
	log.Println("Executando na porta 3000")
	log.Fatal(http.ListenAndServe(":3000", nil))
}
