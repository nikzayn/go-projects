package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/chenjiandongx/go-echarts/charts"
)

func handler(w http.ResponseWriter, r *http.Request) {
	items := []string{"Jan", "Feb", "Mar", "Apr", "May", "Jun"}
	bar := charts.NewBar()
	bar.SetGlobalOptions(charts.TitleOpts{Title: "Bar-Graph"})
	bar.AddXAxis(items).
		AddYAxis("axis-A", []int{20, 30, 40, 50, 60, 70}).
		AddYAxis("axis-B", []int{35, 14, 25, 60, 44, 23})
	f, err := os.Create("bar.html")
	if err != nil {
		log.Println(err)
	}
	bar.Render(w, f)
}

func main() {
	fmt.Println("Rendering the graph on route /bar")
	http.HandleFunc("/bar", handler)
	http.ListenAndServe(":8080", nil)
}
