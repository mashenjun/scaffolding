package template

const MgoDBFile = `package model

import (
	"github.com/globalsign/mgo/bson"
	"time"
)

const (
	DefaultCollSessionPoolLimit = 100
)

// base model
type BaseModel struct {
	IsDeleted int           `+ "`bson:\"is_deleted\"`"+ "\n"+
	`CreatedAt time.Time     `+ "`bson:\"created_at\"`"+ "\n"+
	`UpdatedAt time.Time     `+ "`bson:\"updated_at\"`"+ "\n"+
	`DeletedAt time.Time     `+ "`bson:\"deleted_at\"`"+ "\n"+
`}

// the coll tag will be used as collection name
type Collections struct {
	// define collection here
}

var collections *Collections

func Init() error {
	panic("not implemented")
}

func inject(base bson.M, key string, operator string, val interface{}) {
	if _, ok := base[key]; !ok {
		base[key] = bson.M{operator: val}
	} else {
		base[key].(bson.M)[operator] = val
	}
}`
