package test

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"testing"
)

func Test_ReadInput(t *testing.T) {
	file, err := os.OpenFile(
		"hello.txt",
		os.O_CREATE|os.O_RDWR|os.O_TRUNC, // 파일이 없으면 생성,
		// 읽기/쓰기, 파일을 연 뒤 내용 삭제
		os.FileMode(0644), // 파일 권한은 644
	)
	if err != nil {
		fmt.Println(err)
		return
	}
	defer file.Close() // main 함수가 끝나기 직전에 파일을 닫음

	w := bufio.NewWriter(file)
	w.WriteString("hello i'm ms")
	w.Flush()
	r := bufio.NewReader(file)

	f, err := file.Stat()
	if err != nil {
		fmt.Println(err)
		return
	}
	b := make([]byte, f.Size())
	file.Seek(0, io.SeekStart)
	r.Read(b)

	fmt.Println(string(b))
}
