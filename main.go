package main

import (
	"TestReplicaSet/utils"
	"context"
	"fmt"
	"go.mongodb.org/mongo-driver/bson"
)

type Test struct {
	Info string `bson:"info"`
}
func main(){
	client := utils.ConnectToMongo()
	if client == nil{
		return
	}

	cur, err := client.Database("test").Collection("test").Find(context.TODO(), bson.D{})
	if err != nil{
		fmt.Println(err)
		return
	}
	defer cur.Close(context.TODO())
	for cur.Next(context.TODO()){
		test := &Test{}
		if err := cur.Decode(test);err ==nil{
			fmt.Println("info---->", test.Info)
		}
	}

}
