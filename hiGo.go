package main
import (
	"fmt"
	"net/http"
	//"strings"
	"io/ioutil"
	"encoding/json"
)

func main(){
	http.HandleFunc("/", HandlerTest )
	http.ListenAndServe(":8080", nil)
}

func HandlerTest(w http.ResponseWriter, r *http.Request){
	
	//parts := strings.Split(r.URL.Path, "/")
	FirstPart(w)
	
	
	
}

func GetBody(w http.ResponseWriter, url string) []byte{
	
	res, err := http.Get(url)
	if err != nil{
		fmt.Fprintln(w, "error or something")
		// must break out 
	}
	body, err := ioutil.ReadAll(res.Body)
	if err != nil{
		fmt.Fprintln(w, "could not get content")
		  // must break out
	}
	return body
}




func FirstPart(w http.ResponseWriter){
	body := GetBody(w, "http://www.mocky.io/v2/59c6f6f9400000be06afe8a0")
	
	var j Json
	err := json.Unmarshal(body, &j)
	if err != nil{
		fmt.Fprintln(w, "could not unmarshal")
		return 
	}
	fmt.Fprintln(w, "name:", j.Name, "project:", j.Owner.Login)
	
	SecondPart(w) //with language url
}




func SecondPart(w http.ResponseWriter){ // with language url
	var langs map[string]int
	body := GetBody(w, "http://www.mocky.io/v2/59d2150a120000a90a244fdb")
	err := json.Unmarshal(body, &langs)
	if err != nil{
		fmt.Fprintln(w, "could not unmarshal")
		return
	}
	for key, _ := range langs {
    fmt.Fprintln(w, key, )
	}
	
}

type Json struct{
	Name string `json: "name"`
	Owner Ownerstr  `json: "owner"`
	//language url
	
}

type Ownerstr struct{
	Login string `json: "login"`
}











