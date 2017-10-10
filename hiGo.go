package main
import (
	"fmt"
	"net/http"
	"strings"
)

func add(a, b int) int{
	return a + b
}

func main(){
	http.HandleFunc("/no/", handlerTest )
	http.ListenAndServe(":8080", nil)
}

func handlerTest(w http.ResponseWriter, r *http.Request){
	
	parts := strings.Split(r.URL.Path, "/")
	
	fmt.Fprintln(w, parts)
	fmt.Fprintln(w, len(parts))
	name := parts[len(parts)-1]
	fmt.Fprintf(w, name)
	fmt.Fprintf(w, "nothing")
}
