package main
//app_id 536032116885838
import (
	"encoding/json"
	"fmt"
	"html/template"
	"log"
	"net/http"
	"os"

	"github.com/gorilla/pat"
	"github.com/gorilla/sessions"
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
var indexTemplate = `
<p><a href="/auth/twitter">Log in with Twitter</a></p>
<p><a href="/auth/facebook">Log in with Facebook</a></p>
`

var userTemplate = `
<p>Name: {{.Name}}</p>
<p>Email: {{.Email}}</p>
<p>NickName: {{.NickName}}</p>
<p>Location: {{.Location}}</p>
<p>AvatarURL: {{.AvatarURL}}</p>
<p>UserID: {{.UserID}}</p>
<p>AccessToken: {{.AccessToken}}</p>
`
func init() {
	gothic.Store = sessions.NewCookieStore([]byte("sk;aldjsdm,ahtrqwiouehr1348yy4y29834"))
	file, _ := os.Open("config.json")
	decoder := json.NewDecoder(file)
	config = Configuration{}
	err := decoder.Decode(&config)
	if err != nil {
		fmt.Println(config)
		log.Fatal(err)
	}
}

func callbackAuthHandler(res http.ResponseWriter, req *http.Request) {
	user, err := gothic.CompleteUserAuth(res, req)
	if err != nil {
		fmt.Println(res, err)
		return
	}	
	t, _ := template.New("userinfo").Parse(userTemplate)
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
	r := pat.New()
	r.Get("/auth/{provider}/callback", callbackAuthHandler)
	r.Get("/auth/{provider}", gothic.BeginAuthHandler)
	r.Get("/", indexHandler)

	server := &http.Server{
		Addr:		":8080",
		Handler:	r,
	}
	log.Println("Listening...")
	server.ListenAndServe()
}	