package main

import (
    "fmt"
    "time"
)

type Person struct {
	FirstName, LastName string
	Dob 		    time.Time
	Email, Location     string
}


func (p Person) PrintName() {
    fmt.Printf("\n%s %s\n", p.FirstName, p.LastName)
}    

func (p Person) PrintDetails(){
    fmt.Printf("[Date of Birth: %s, Email: %s, Location: %s]\n", p.Dob.String(), p.Email, p.Location )
}

func (p *Person) ChangeLocation(newLocation string){
    p.Location= newLocation
}

type Admin struct {
    Person
    Roles []string
}

type Member struct {
    Person
    Skills []string
}

func main() {
    alex :=Admin {
	   Person{
	          "Alex",
	          "John",
	          time.Date(1970, time.January, 0,0,0,0,0, time.UTC),
	          "alex@email.com",
	          "New York",
	          },
	          []string("Manage Team", "Manage Tasks"),
	   }
    shiju := Member {
        Person{
	"Shiju",
	"Varghese",
	time.Date(1979, time.February, 0,0,0,0,0, time.UTC),
	"shiju@email.com",
	"Kochi"},
	[]string("Go", "Docker", "Kubernetes"),
    }
    alex.PrintName()
    alex.PrintDetails()
    shiju.PrintName()
    shiju.PrintDetails()
}