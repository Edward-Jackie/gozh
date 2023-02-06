package oss

import (
	"fmt"
	"github.com/aliyun/aliyun-oss-go-sdk/oss"
	"os"
)

func GetToken() (string, error) {
	client, err := oss.New(os.Getenv("oss.endPoint"), os.Getenv("oss.accessKeyId"), os.Getenv("oss.accessKeySecret"))
	if err != nil {
		fmt.Println("Error:", err)
		return "", err
	}
	result, err := client.GetBucketACL(os.Getenv("oss.bucket"))
	if err != nil {
		return "", err
	}
	return result.ACL, nil
}
