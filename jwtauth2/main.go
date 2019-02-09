package main

import (
    "net/http"
    "encoding/json"
    "os"
    "github.com/gorilla/handlers"
    "github.com/gorilla/mux"
    
)

type Product struct {
    Id		int
    Name	string
    Slug	string
    Description string
}

var products = []Product{
    Product{Id: 1, Name: "Hover Shooters", Slug: "hover-shooters", Description: "Shoot your way to the top on 14 different hoverboards"},
    Product{Id: 2, Name: "Ocean Explorer", Slug: "ocean-explorer", Description: "Explore the depths of the sea in this one of a kind"},
    Product{Id: 3, Name: "Dinosaur Park", Slug: "dinosaur-park", Description: "Go back 65 million years in the past and ride a T-Rex"},
    Product{Id: 4, Name: "Cars VR", Slug: "cars-vr", Description: "Get behind the wheel of the fastest cars in the world."},
    Product{Id: 5, Name: "Robin Hood", Slug: "robin-hood", Description: "Pick up the bow and arrow and master the art of archery"},
    Product{Id: 6, Name: "Real world VR", Slug: "real-world-vr", Description: "Explore the seven wonders of the world in VR"},
}

var StatusHandler = http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
    w.Write([]byte("API is up and running"))
})

var ProductsHandler = http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
    payload, _ := json.Marshal(products)
    w.Header().Set("Content-Type", "application/json")
    w.Write([]byte(payload))
})

var AddFeedbackHandler = http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
    var product Product
    vars := mux.Vars(r)
    slug := vars["slug"]
    
    for _, p := range products {
	if p.Slug == slug {
	    product = p
	}
    }
    
    w.Header().Set("Content-Type", "application/json")
    if product.Slug != "" {
	payload, _ := json.Marshal(product)
	w.Write([]byte(payload))
    } else {
	w.Write([]byte("Product Not Found"))
    }
})


func main() {
    r := mux.NewRouter()
    r.Handle("/", http.FileServer(http.Dir("./views/")))
    r.Handle("/status", StatusHandler).Methods("GET")
    r.Handle("/products", ProductsHandler).Methods("GET")
    r.Handle("/products/{slug}/feedback", AddFeedbackHandler).Methods("POST")
    r.PathPrefix("/static/").Handler(http.StripPrefix("/static/", http.FileServer(http.Dir("./static/"))))
    http.ListenAndServe(":8080", handlers.LoggingHandler(os.Stdout, r))

}

var NotImplemented = http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
    w.Write([]byte("Not Implemented"))
})