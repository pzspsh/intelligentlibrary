/*
@File   : main.go
@Author : pan
@Time   : 2024-11-21 16:29:45
*/
package main

import (
	"context"
	"fmt"

	"github.com/minio/minio-go/v7"
	"github.com/minio/minio-go/v7/pkg/credentials"
)

type Minio struct {
	Endpoint        string
	AccessKeyId     string
	SecretAccessKey string
	UseSSL          bool
	Client          *minio.Client
}

func (m *Minio) UploadFile() error {
	var err error
	// m.Client.ListBuckets(context.Background())
	// m.Client.ListObjects(context.Background(), "", minio.ListObjectsOptions{})
	// m.Client.FGetObject(context.Background(), "桶", "要下载的文件路径", "要存储的本地路径", minio.GetObjectOptions{})
	// m.Client.FPutObject(context.Background(), "桶对象名称", "要上传到的minio路径", "要上传本地的文件路径", minio.PutObjectOptions{})
	return err
}

func (m *Minio) DownloadFile() error {
	var err error
	if ok, _ := m.Client.BucketExists(context.Background(), "桶"); ok {
		objects, err := m.Client.GetObject(context.Background(), "桶", "对象名/", minio.GetObjectOptions{})
		if err != nil {
			fmt.Println("get object error: ", err)
		}
		fmt.Println(objects)
	}
	return err
}

func (m *Minio) MinioConn() (*Minio, error) {
	var err error
	var client *minio.Client
	client, err = minio.New(m.Endpoint, &minio.Options{
		Creds:  credentials.NewStaticV4(m.AccessKeyId, m.SecretAccessKey, ""),
		Secure: m.UseSSL,
	})
	if err != nil {
		fmt.Println("连接minio 失败: ", err)
		return m, err
	}
	fmt.Println("连接minio成功：", client)
	m.Client = client
	return m, err
}

func main() {
	minio := &Minio{
		Endpoint:        "targetip:port",
		AccessKeyId:     "",
		SecretAccessKey: "",
	}
	minio.MinioConn()
}
