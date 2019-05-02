package models

import "gopkg.in/mgo.v2/bson"

type DisksJSON struct {
	ID   bson.ObjectId `bson:"_id" json:"id"`
	Host string        `bson:"host" json:"host"`
	Stat []Stat        `bson:"stat" json:"stat"`
}

type Stat struct {
	Mount []string `bson:"mount" json:"mount"`
	Path  string   `bson:"path" json:"path"`
	Smart *Smart   `bson:"smart" json:"smart"`
}
type Smart struct {
	Info *Info `bson:"info" json:"info"`
}

type Info struct {
	Attributes []Attributes `bson:"attributes" json:"attributes"`
	Capacity   string       `bson:"capacity" json:"capacity"`
	Firmware   string       `bson:"firmware" json:"firmware"`
	Model      string       `bson:"model" json:"model"`
	Serial     string       `bson:"serial" json:"serial"`
}

type Attributes struct {
	AttrName string `bson:"attribute_name" json:"attribute_name"`
	ID       string `bson:"id" json:"id"`
	RawValue string `bson:"raw_value" json:"raw_value"`
	Thresh   string `bson:"thresh" json:"thresh"`
	Value    string `bson:"value" json:"value"`
	Worst    string `bson:"worst" json:"worst"`
}
