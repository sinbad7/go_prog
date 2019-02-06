package main

import (
    "fmt"
    "log"
    
    "gopkg.in/mgo.v2"
    "gopkg.in/mgo.v2/bson"
)

type Category struct {
    Id		bson.ObjectId `bson:"_id,omitempty"`
    Name	string
    Description string
}

func main() {
    session, err := mgo.Dial("localhost")
    if err != nil {
	panic(err)
    }
    defer session.Close()
    session.SetMode(mgo.Monotonic, true)
    c := session.DB("taskdb").C("categories")
    
    doc := Category{
	bson.NewObjectId(),
	"Open Source",
	"Task for open-source projects",
    }
    err = c.Insert(&doc)
    if err != nil {
	log.Fatal(err)
    }
    err = c.Insert(&Category{bson.NewObjectId(), "R & D", "R & D Tasks"},
	&Category{bson.NewObjectId(), "Project", "Project Tasks"})
    var count int
    count, err = c.Count()
    if err != nil {
	log.Fatal(err)
    } else {
	fmt.Printf("%d records inserted", count)
    }
}