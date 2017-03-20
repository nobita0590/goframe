package models

import (
	"eternal/iris/models/db"
)

var myDb = &db.MyDb

type (
	Records struct {
		Number		int
		Page 		int
	}
	Sorts map[string]bool
)

func (this Sorts) Get() (val []string) {
	for key,flag := range this {
		if flag {
			val = append(val,key)
		}else{
			val = append(val,"-"+key)
		}
	}
	return
}

func (this Records) GetLimit() (limit int,skip int) {
	if this.Page < 1 {
		this.Page = 1
	}
	if this.Number < 10 {
		limit = 10
	}else{
		limit = this.Number
	}

	skip = (this.Page - 1) * limit
	return
}
