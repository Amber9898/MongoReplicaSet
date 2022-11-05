package utils

import (
	"context"
	"errors"
	"fmt"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

const(
	rs0_0_ip = "192.168.45.3:27017"
	rs0_1_ip = "192.168.45.3:27018"
	rs0_2_ip = "192.168.45.3:27019"
	replicaSetName = "rs0"
	mongoUser = "admin"
	mongoPassword = "123456"
)

func ConnectToMongo() *mongo.Client{
	//"mongodb://localhost:27017,localhost:27018/?replicaSet=replset"
	ipList := []string{rs0_1_ip, rs0_0_ip, rs0_2_ip}
	uri, err := generateUri(ipList, replicaSetName)
	if err != nil{
		fmt.Println(err)
		return nil
	}
	clientOpts := options.Client().ApplyURI(uri)
	client, err := mongo.Connect(context.TODO(), clientOpts)
	if err != nil {
		fmt.Println(err)
		return nil
	}
	return client
}

func generateUri(ipList []string, replicaSetName string) (string, error){
	if len(ipList) == 0{
		return "", errors.New("ip list's lenth is 0")
	}
	ipStr := ""
	for _, ip := range ipList{
		ipStr += ip+","
	}
	arr := []rune(ipStr)
	if len(arr) > 0{
		if arr[len(arr)-1] == ','{
			arr = arr[:len(arr)-1]
		}
	}
	ipStr = string(arr)

	uri := fmt.Sprintf("mongodb://%v:%v@%v/?replicaSet=%v",
		mongoUser, mongoPassword, ipStr, replicaSetName)
	return uri, nil
}
