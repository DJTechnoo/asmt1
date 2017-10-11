package main

import (
	"testing"
	"net/http"
	"encoding/json"
	"io/ioutil"
)

func Test_FirstTest(t *testing.T){
	
	mainUrl := "http://www.mocky.io/v2/59c6f6f9400000be06afe8a0"
	res, err := http.Get(mainUrl)
	if err != nil{
		t.Error("error")
		return
	}
	body, err := ioutil.ReadAll(res.Body)
	if err != nil{
		t.Error("error2")
		  return
	}
	
	var j Json
	err = json.Unmarshal(body, &j)
	if err != nil{
		
		return 
	}
	//fmt.Fprintln(w, "name:", j.Name, "project:", j.Owner.Login)
	if j.Name != "go"{
		t.Error("name should be go")
	}
	//SecondPart(w, j.Languages_url) //with language url
	//ThirdPart(w, j.Contributors_url)					// with commits url
	
}