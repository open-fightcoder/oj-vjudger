package store

import (
	"testing"

	"strings"

	"github.com/minio/minio-go"
	"github.com/open-fightcoder/oj-vjudger/common/g"
)

func TestMinio(t *testing.T) {
	g.LoadConfig("../../cfg/cfg.toml.debug")
	InitMinio()
	//_, err := MinioClient.PutObject("ttt", "test.txt", strings.NewReader("aaaaaassssssdddd"), -1, minio.PutObjectOptions{ContentType: "application/octet-stream"})
	_, err := MinioClient.PutObject("code", "vjudger.c", strings.NewReader("/*Author：LuwenjingDate：Fri May 04 2018 16:34:37 GMT+0800 (CST)*/\n#include <stdio.h> \nint main (void) {printf(\"Hello world!\");return 0;}"), -1, minio.PutObjectOptions{ContentType: "application/octet-stream"})
	if err != nil {
		panic(err)
	}
}
