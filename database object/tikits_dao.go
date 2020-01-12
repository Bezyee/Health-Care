package dao

import (
	"log"

	"github.com/abdissabayou/Online-ticket-store/models"
	mgo "gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)

type TikitsDAO struct {
	Server   string
	Database string
}

var tikitdb *mgo.Database

const (
	GLy_COLLECTION = "tikits"
)

// Establish a connection to database
func (m *TikitsDAO) Connect() {
	session, err := mgo.Dial(m.Server)
	if err != nil {
		log.Fatal(err)
	}
	tikitdb = session.DB(m.Database)
}

// Find list of tikits
func (m *TikitsDAO) FindAll() ([]Tikit, error) {
	var tikits []Tikit
	err := tikitdb.C(GLy_COLLECTION).Find(bson.M{}).All(&tikits)
	return tikits, err
}

// Find a tikit by its id
func (m *TikitsDAO) FindById(id string) (Tikit, error) {
	var tikit Tikit
	err := tikitdb.C(GLy_COLLECTION).FindId(bson.ObjectIdHex(id)).One(&tikit)
	return tikit, err
}

func (m *TikitsDAO) FindByName(name string) (Tikit, error) {
	var tikit Tikit
	err := tikitdb.C(GLy_COLLECTION).Find(name).One(&tikit)
	return tikit, err
}

// Insert a tikit into database
func (m *TikitsDAO) Insert(tikit Tikit) error {
	err := tikitdb.C(GLy_COLLECTION).Insert(&tikit)
	return err
}

// Delete an existing tikit
func (m *TikitsDAO) Delete(tikit Tikit) error {
	err := tikitdb.C(GLy_COLLECTION).Remove(&tikit)
	return err
}

// Update an existing tikit
func (m *TikitsDAO) Update(tikit Tikit) error {
	err := tikitdb.C(GLy_COLLECTION).UpdateId(tikit.ID, &tikit)
	return err
}
