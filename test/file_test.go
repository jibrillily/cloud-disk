package test

import (
	"crypto/md5"
	"fmt"
	"io/ioutil"
	"math"
	"os"
	"strconv"
	"testing"
)

// 分片大小

const chunkSize = 100 * 1024 * 1024 //100M

// 文件分片
func TestGenerateChunkFile(t *testing.T) {
	//拿到文件总大小
	fileinfo, err := os.Stat("./img/eva.jpg")
	if err != nil {
		t.Fatal(err)
	}

	//math.Ceil向上取整
	chunkNum := int(math.Ceil(float64(fileinfo.Size()) / float64(chunkSize)))
	//0777表示：创建了一个普通文件，所有人拥有所有的读、写、执行权限
	//0666表示：创建了一个普通文件，所有人拥有对该文件的读、写权限，但是都不可执行
	//0644表示：创建了一个普通文件，文件所有者对该文件有读写权限，用户组和其他人只有读权限， 都没有执行权限
	myFile, err := os.OpenFile("./img/eva.jpg", os.O_RDONLY, 0666)
	if err != nil {
		t.Fatal(err)
	}

	b := make([]byte, chunkSize)
	for i := 0; i < chunkNum; i++ {
		//指定读取文件的起始位置
		myFile.Seek(int64(i*chunkSize), 0)
		//当最后一个分片时不需要读取chunkSize这么大
		if chunkSize > fileinfo.Size()-int64(i*chunkSize) {
			b = make([]byte, fileinfo.Size()-int64(i*chunkSize))
		}
		myFile.Read(b)
		f, err := os.OpenFile("./"+strconv.Itoa(i)+".chunk", os.O_CREATE|os.O_WRONLY, os.ModePerm)
		if err != nil {
			t.Fatal(err)
		}
		f.Write(b)
		f.Close()
	}
	myFile.Close()
}

// 分片文件合并
func TestMergeChunkFile(t *testing.T) {
	myFile, err := os.OpenFile("test2.mp4", os.O_CREATE|os.O_WRONLY|os.O_APPEND, os.ModePerm)
	if err != nil {
		t.Fatal(err)
	}

	//取得分片个数
	fileinfo, err := os.Stat("./img/test.mp4")
	if err != nil {
		t.Fatal(err)
	}
	chunkNum := int(math.Ceil(float64(fileinfo.Size()) / float64(chunkSize)))

	for i := 0; i < chunkNum; i++ {
		f, err := os.OpenFile("./"+strconv.Itoa(i)+".chunk", os.O_RDONLY, os.ModePerm)
		if err != nil {
			t.Fatal(err)
		}
		b, err := ioutil.ReadAll(f)
		if err != nil {
			return
		}

		myFile.Write(b)
		f.Close()
	}
	myFile.Close()
}

//文件一致性
func TestCheck(t *testing.T) {
	//获取源文件信息
	file1, err := os.OpenFile("./img/eva.jpg", os.O_RDONLY, os.ModePerm)
	if err != nil {
		t.Fatal(err)
	}
	b1, err := ioutil.ReadAll(file1)
	if err != nil {
		t.Fatal(err)
	}
	//获取第二个文件的信息
	file2, err := os.OpenFile("test2.mp4", os.O_RDONLY, os.ModePerm)
	if err != nil {
		t.Fatal(err)
	}
	b2, err := ioutil.ReadAll(file2)
	if err != nil {
		t.Fatal(err)
	}
	s1 := fmt.Sprintf("%x", md5.Sum(b1))
	s2 := fmt.Sprintf("%x", md5.Sum(b2))

	fmt.Println(s1)
	fmt.Println(s2)
	fmt.Println(s1 == s2)

}
