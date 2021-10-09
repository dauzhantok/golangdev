package main

import (
	"fmt"
	"math"
	"net/http"
	"time"
)

type User struct {
	Name           string
	Kg, Hight, Imt float64
}

var da = User{"Dauzhan", 60, 1.8, 0}

func home(page http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(page, "%v have IMT %v", da.Name, da.Imt)
	fmt.Fprintln(page, "")

}
func main() {
	da.Imt = da.Kg / math.Pow(da.Hight, 2)
	time.Sleep(1 * time.Second)
	http.HandleFunc("/", home)
	http.ListenAndServe(":9000", nil)
}
