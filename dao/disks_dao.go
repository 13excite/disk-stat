package dao

import (
	"log"

	. "build-disk-stat/models"

	mgo "gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)

//DisksDAO describe mongodb acces parametrs
type DisksDAO struct {
	Server   string
	Database string
}

var db *mgo.Database

const (
	COLLECTION = "disks"
)

// Connect establish a connection to database
func (m *DisksDAO) Connect() {
	session, err := mgo.Dial(m.Server)
	if err != nil {
		log.Fatal(err)
	}
	db = session.DB(m.Database)
}

// FindAll find list of disks attributes
func (m *DisksDAO) FindAll() ([]DisksJSON, error) {
	var disks []DisksJSON
	err := db.C(COLLECTION).Find(bson.M{}).All(&disks)
	return disks, err
}

// FindById finding a disk attr by its id
func (m *DisksDAO) FindById(id string) (DisksJSON, error) {
	var disk DisksJSON
	err := db.C(COLLECTION).FindId(bson.ObjectIdHex(id)).One(&disk)
	return disk, err
}

// Insert a disk attr into database
func (m *DisksDAO) Insert(disk DisksJSON) error {
	err := db.C(COLLECTION).Insert(&disk)
	return err
}

// DeleteByID an existing disk stat by ID
func (m *DisksDAO) DeleteByID(id string) error {
	err := db.C(COLLECTION).RemoveId(bson.ObjectIdHex(id))
	return err
}

// Update an existing disk
func (m *DisksDAO) Update(disk DisksJSON) error {
	err := db.C(COLLECTION).UpdateId(disk.ID, &disk)
	return err
}
