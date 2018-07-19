package ethapi

import (
	"gopkg.in/mgo.v2/bson"
	"gopkg.in/mgo.v2"
	"container/list"
	"fmt"
)

type Category struct {
	Id          bson.ObjectId `bson:"_id,omitempty"`
	Name        string
	Description string
}

func MongoTest() {
	session, err := mgo.Dial("localhost")
	if err != nil {
		panic(err)
	}
	defer session.Close()

	session.SetMode(mgo.Monotonic, true)
	//获取一个集合
	c := session.DB("taskdb").C("categories")
	doc := Category{
		bson.NewObjectId(),
		"Open Source",
		"Tasks for open-source projects",
	}
	l := list.New()
	l.PushBack(doc)
	//插入一个模型对象
	err = c.Insert(&l)
	if err != nil {
		//log.Fatal(err)
	}

	//插入两个模型对象
	err = c.Insert(&Category{bson.NewObjectId(), "R & D", "R & D Tasks"}, &Category{bson.NewObjectId(), "Project", "Project Tasks"})
	var count int
	count, err = c.Count()
	if err != nil {
		//log.Fatal(err)
	} else {
		fmt.Printf("%d records inserted", count)
	}

}
