package dpmgo

import (
	"fmt"
	"github.com/deferpanic/deferclient/deferstats"
	"gopkg.in/mgo.v2"
	"time"
)

var (
	selectThreshold int
)

type DpMgo struct {
	Other *mgo.Collection
}

func NewCollection(m *mgo.Collection) *DpMgo {
	selectThreshold = 100

	return &DpMgo{
		m,
	}
}

func (db *DpMgo) logQuery(startTime time.Time, query string) {

	endTime := time.Now()
	t := int(((endTime.Sub(startTime)).Nanoseconds() / 1000000))

	ddb := deferstats.DeferDB{
		Query: query,
		Time:  t,
	}

	if t >= selectThreshold {
		deferstats.Querylist.Add(ddb)
	}
}

func (db *DpMgo) Insert(docs ...interface{}) error {
	startTime := time.Now()

	// FIXME - need a cap
	query := fmt.Sprintf("%#v", docs)
	defer db.logQuery(startTime, query)
	return db.Other.Insert(docs)
}

func (db *DpMgo) Find(query interface{}) *mgo.Query {
	startTime := time.Now()

	// FIXME - need a cap
	rquery := fmt.Sprintf("%#v", query)
	defer db.logQuery(startTime, rquery)
	return db.Other.Find(query)
}

func (db *DpMgo) EnsureIndexKey(key ...string) error {
	startTime := time.Now()

	defer db.logQuery(startTime, fmt.Sprintf("%#v", key))
	return db.Other.EnsureIndexKey(key...)
}
