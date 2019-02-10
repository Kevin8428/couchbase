package main

import (
	"fmt"
	"os"

	"github.com/gocb"
)

// bucket reference - reuse as bucket reference in the application
var bucket *gocb.Bucket

type couchBaseResponse struct {
	ID            string            `json:"id"`
	Active        bool              `json:"active"`
	EntityId      int               `json:"entityId"`
	Devices       map[string]string `json:"devices"`
	EntityType    string            `json:"entityType"`
	CreatedOn     string            `json:"createdOn"`
	LastUpdatedOn string            `json:"lastUpdatedOn"`
}

func main() {
	cluster, err := gocb.Connect("http://couchbase-dev.dispatch.me:8091")
	if err != nil {
		fmt.Println("ERROR CONNECTING TO CLUSTER:", err)
	}

	cluster.Authenticate(gocb.PasswordAuthenticator{
		Username: os.Getenv("BUCKET"),
		Password: os.Getenv("SASL"),
	})

	bucket, err := cluster.OpenBucket("device-token-dev", "")
	if err != nil {
		fmt.Println("ERROR OPENING BUCKET:", err)
	}

	// a := map[string]interface{}{}
	// bucket.Get("User:"+os.Getenv("USER"), &a)
	// fmt.Printf("User: %v\n", a)

	a := couchBaseResponse{}
	bucket.Get("User:"+os.Getenv("USER"), &a)
	fmt.Printf("User: %+v\n", a)
}

// 1. unmarshall to real struct

/*
{
	"id": "User:19",
	"active": true,
	"devices": {
	  "2e73d0ed62438ba6550c8e4728d83ac85dad351c3532d6aaadbd88e0ce3ae063": "arn:aws:sns:us-east-1:176506442715:endpoint/APNS/communications-hub-apple-push-dev-release/67de8fda-0804-3275-9e04-a63dc05fd8a5"
	},
	"entityId": 19,
	"entityType": "User",
	"createdOn": "2018-08-02T18:43:08.052818899Z",
	"lastUpdatedOn": "2019-01-17T22:18:09Z"
}
*/
