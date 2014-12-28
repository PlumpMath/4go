package aws3

import (
	"fmt"
	"github.com/mitchellh/goamz/aws"
	"github.com/mitchellh/goamz/s3"
	"github.com/tmilewski/goenv"
	"os"
)

func init() {
	err := goenv.Load()

	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

func ListBuckets() {
	accessKey := os.Getenv("AWS_IDENTITY")
	secretKey := os.Getenv("AWS_SECRETKEY")

	auth, err := aws.GetAuth(accessKey, secretKey)

	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	s3 := s3.New(auth, aws.USEast)

	result, err := s3.ListBuckets()

	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	//fmt.Println(result)

	for _, bucket := range result.Buckets {
		fmt.Println(bucket.Name)
	}
}
