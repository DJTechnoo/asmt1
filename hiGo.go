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
	mainUrl := "http://www.mocky.io/v2/59c6f6f9400000be06afe8a0"
	FirstPart(w, mainUrl)	
}




//		-------------------------------------------------------------------------
//		This function gets a string and returns its body. w is for error printing
//		-------------------------------------------------------------------------
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



//		-------------------------------------------------------------------------
//		This function writes parses and writes "name" and "owner", gives the 
//		"language" url to the SecondPart function as an arg
//		-------------------------------------------------------------------------
func FirstPart(w http.ResponseWriter, mainUrl string){
	body := GetBody(w, mainUrl)
	
	var j Json
	err := json.Unmarshal(body, &j)
	if err != nil{
		fmt.Fprintln(w, "could not unmarshal")
		return 
	}
	fmt.Fprintln(w, "name:", j.Name, "project:", j.Owner.Login)
	
	SecondPart(w, j.Languages_url) //with language url
	ThirdPart(w, j.Contributors_url)					// with commits url
}



//		-------------------------------------------------------------------------
//		This function deals with parsing and writing languages from a map 
//		-------------------------------------------------------------------------
func SecondPart(w http.ResponseWriter, langUrl string){ // with language url
	fmt.Fprintln(w, langUrl)
	var langs map[string]int
	body := GetBody(w, "http://www.mocky.io/v2/59d2150a120000a90a244fdb") 	// 	change to langUrl
	err := json.Unmarshal(body, &langs)
	if err != nil{
		fmt.Fprintln(w, "could not unmarshal")
		return
	}
	for key, _ := range langs {
    fmt.Fprintln(w, key, )
	}
	
}

//		---------------------------------------------
//		This func writes commits and takes w and contributers url
//		---------------------------------------------
func ThirdPart(w http.ResponseWriter, comUrl string){
	var committers []Committer
	fmt.Fprintln(w, comUrl)
	body := GetBody(w, "http://www.mocky.io/v2/59d215a3120000760a244fdd") 	// 	change to comUrl
	err := json.Unmarshal(body, &committers)
	if err != nil{
		fmt.Fprintln(w, "could not unmarshal")
		return
	}
	fmt.Fprintln(w, committers)				// should only list the top contributor
	
}

//		-------------------------------------------------------------------------
//		These structs make up the needed json data
//		-------------------------------------------------------------------------
type Json struct{
	Name string `json: "name"`
	Owner Ownerstr  `json: "owner"`
	Languages_url string `json: "languages_url"`
	Contributors_url string `json: "contributors_url"`
	
}

type Ownerstr struct{
	Login string `json: "login"`
}

//		-------------------------------------------------------
//		Struct for contributions holding commits
//		-------------------------------------------------------
type Committer struct{
	Login string `json: "login"`
	Contributions int `json: "contributions"`
}











