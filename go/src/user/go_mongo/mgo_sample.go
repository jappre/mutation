package main

import (
	"fmt"
	"log"
	"time"

	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)

type Person struct {
	Name  string
	Email string
}

type Commit struct {
	Person      Person
	Date        time.Time
	FileChanged int16
	Insertions  int16
	Delettions  int16
}

func main() {
	session, err := mgo.Dial("localhost")
	if err != nil {
		panic(err)
	}
	defer session.Close()

	// Optional. Switch the session to a monotonic behavior.
	session.SetMode(mgo.Monotonic, true)

	c := session.DB("test").C("Commit")
	err = c.Insert(&Commit{Person{"Ale", "+55 53 8116 9639"}, time.Now(), 2, 3, 6})
	if err != nil {
		log.Fatal(err)
	}

	result := Commit{}
	err = c.Find(bson.M{"date": "2016-01-27"}).One(&result)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Deletions:", result.Delettions)
}
