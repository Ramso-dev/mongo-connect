package main

import (
	"fmt"
	"time"

	"gopkg.in/mgo.v2"
)

//const MongoDb details
const (
	hosts      = "https://mydb-cw-portal-plg.playground.itandtel.at:27017"
	database   = "sampledb"
	username   = "ramsopanzer"
	password   = "P30AWo6lDWTKbT6l"
	collection = "messages"
)

func main() {

	info := &mgo.DialInfo{
		Addrs:    []string{hosts},
		Timeout:  60 * time.Second,
		Database: database,
		Username: username,
		Password: password,
	}

	session, err1 := mgo.DialWithInfo(info)
	if err1 != nil {
		panic(err1)
	}

	col := session.DB(database).C(collection)

	count, err2 := col.Count()
	if err2 != nil {
		panic(err2)
	}
	fmt.Println(fmt.Sprintf("Messages count: %d", count))
}
