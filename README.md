# Defer Panic gorp add-on

[![GoDoc](https://godoc.org/github.com/deferpanic/dpmgo?status.svg)](https://godoc.org/github.com/deferpanic/dpmgo)

### Defer Panic mgo add-on pkg

Many [deferpanic](https://deferpanic.com "deferpanic") users use [mgo](https://github.com/go-mgo/mgo "mgo").
This is the first attempt to start adding support. Much is left to do and this will probably change
dramatically in the near future. 

Currently the DeferPanic wrapper operates at the collection level.

```
package main

import (
	"fmt"
	"github.com/deferpanic/deferclient/deferstats"
	"github.com/deferpanic/dpmgo"
	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
	"log"
	"time"
)

type Person struct {
	Name  string
	Phone string
}

func main() {
	dps := deferstats.NewClient("v00L0K6CdKjE4QwX5DL1iiODxovAHUfo")
	go dps.CaptureStats()

	session, err := mgo.Dial("127.0.0.1")
	if err != nil {
		panic(err)
	}
	defer session.Close()

	// Optional. Switch the session to a monotonic behavior.
	session.SetMode(mgo.Monotonic, true)

	_c := session.DB("test").C("people")
	c := dpmgo.NewCollection(_c)

	err = c.Insert(&Person{"Ale", "+55 53 8116 9639"},
		&Person{"Cla", "+55 53 8402 8510"})
	if err != nil {
		log.Fatal(err)
	}

	result := Person{}
	err = c.Find(bson.M{"name": "Ale"}).One(&result)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Phone:", result.Phone)

	time.Sleep(200 * time.Second)
}
```
