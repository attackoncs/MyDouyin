/*
 * minio 测试
 */

package minio

import (
	"fmt"
	"os"
	"strings"
	"testing"
)

func TestCreateBucket(t *testing.T) {
	CreateBucket("TiktokTest")
}

func TestUploadLocalFile(t *testing.T) {
	info, err := UploadLocalFile("TiktokTest", "test.mp4", "./test.mp4", "video/mp4")
	fmt.Println(info, err)
}

func TestUploadFile(t *testing.T) {
	file, _ := os.Open("./test.mp4")
	defer file.Close()
	fi, _ := os.Stat("./test.mp4")
	err := UploadFile("TiktokTest", "ceshi2", file, fi.Size())
	fmt.Println(err)
}

func TestGetFileUrl(t *testing.T) {
	url, err := GetFileUrl("TiktokTest", "test.mp4", 0)
	fmt.Println(url, err, strings.Split(url.String(), "?")[0])
	fmt.Println(url.Path, url.RawPath)
}
