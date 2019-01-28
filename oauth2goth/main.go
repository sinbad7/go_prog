package main

import (
	"encoding/json"
	"fmt"
	"html/template"
	"log"
	"net/http"
	"os"

	"github.com/gorilla/pat"
	"github.com/markbates/goth"
	"github.com/markbates/goth/gothic"
	"github.com/markbates/goth/providers/facebook"
	"github.com/markbates/goth/providers/twitter"
)

type Configuration struct {
	TwitterKey		string
	TwitterSecret	string
	FacebookKey		string
	FacebookSecret	string
}

var config Configuration

func init() {
	file, _ := os.Open("config.json")
	decoder := json.NewDecoder(file)
	config = Configuration{}
	err := decoder.Decode(&config)
	if err != nil {
		log.Fatal(err)
	}
}

func callbackAuthHandler(res http.ResponseWriter, req *http.Request) {
	user, err := gothic.CompleteUserAuth(res, req)
	if err != nil {
		fmt.Println(res, err)
		return
	}	
	t, _ :template.New("userinfo").Parse(userTemplate)
	t.Execute(res, user)
}

func indexHandler(res http.ResponseWriter, req *http.Request) {
	t, _ := template.New("index").Parse(indexTemplate)
	t.Execute(res, nil)
}

func main() {
	goth.UseProviders(
		twitter.New(config.TwitterKey, config.TwitterSecret, "http://iktc.ru:8080/auth/twitter/callback"),
		facebook.New(config.FacebookKey, config.FacebookSecret, "http://iktc.ru:8080/auth/facebook/callback"),
	)
}	