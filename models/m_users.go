package models

import (
	"gopkg.in/mgo.v2/bson"
	"eternal/iris/models/db"
	"strings"
)

type (
	Users struct {
		Id 		bson.ObjectId 		`bson:"_id,omitempty"`
		Email	string				`bson:"Email" json:"name"`
		FirstName		string		`bson:"FirstName"`
		LastName		string		`bson:"LastName"`
	}
	UsersFilter struct {
		Sort		Sorts
		Records		Records
		Text 		string
	}
)

func (this UsersFilter) GetFilter() (f bson.M) {
	f = bson.M{}
	if strings.Trim(this.Text," ") != "" {
		f["Email"] = bson.RegEx{Pattern:this.Text}
	}
	return
}

func (this *Users) Insert() error {
	this.Id = bson.NewObjectId()

	err := myDb.C(db.UsersTblName).Insert(&this)
	return err
}

func (this *Users) Update() error {
	return myDb.C(db.UsersTblName).Update(bson.M{"_id": this.Id}, &this)
}

func (this *Users) Get() error {
	return myDb.C(db.UsersTblName).Find(bson.M{"_id":this.Id}).One(&this)
}

func (this Users) GetList(filter UsersFilter) ([]Users,error) {
	result := []Users{}
	limit,skip := filter.Records.GetLimit()
	sort := filter.Sort.Get()

	err := myDb.C(db.UsersTblName).Find(filter.GetFilter()).Limit(limit).Skip(skip).Sort(sort...).All(&result)
	return result, err
}

func (this Users) Delete() error {
	return myDb.C(db.UsersTblName).RemoveId(this.Id)
}