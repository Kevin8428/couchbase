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

	a := couchBaseResponse{}
	bucket.Get("User:"+os.Getenv("USER"), &a)
	fmt.Printf("User: %+v\n", a)
}
