package utils

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"rtc_exporter/structure"
)

//处理每一行
func ProcessLine(line []byte) {

}

//读取最后一行信息
func ReadLastLine(line []byte) {

	words := SliceWords(line)
	for i := 0; i < len(words); i++ {
		fmt.Println("word", i+1, ": ", words[i])
	}
}

//读取机器信息
func ReadFirstLine(line []byte) {

	words := SliceWords(line)
	for i := 0; i < len(words); i++ {
		fmt.Println("word", i+1, ": ", words[i])
	}
}

func ReadLine(filePth string, hookfn func([]byte)) error {

	var isFirst bool = true
	f, err := os.Open(filePth)
	if err != nil {
		return err
	}
	defer f.Close()

	bfRd := bufio.NewReader(f)
	for {
		line, err := bfRd.ReadBytes('\n')
		hookfn(line) //放在错误处理前面，即使发生错误，也会处理已经读取到的数据。
		if len(line) > 0 && isFirst {
			ReadFirstLine(line)
			isFirst = false
		} else if isFirst {
			os.Stdout.WriteString("Empty File!")
			break
		}
		if err != nil { //遇到任何错误立即返回，并忽略 EOF 错误信息
			if err == io.EOF {
				ReadLastLine(line)
				return nil
			}
			return err
		}
	}
	return nil
}

func ReadLineJson(filePth string, hookfn func([]byte)) (structure.BasicInfo, error) {

	var isFirst bool = true
	var basInfo structure.BasicInfo

	f, err := os.Open(filePth)
	if err != nil {
		return basInfo, err
	}
	defer f.Close()

	bfRd := bufio.NewReader(f)
	for {
		line, err := bfRd.ReadBytes('\n')
		hookfn(line) //放在错误处理前面，即使发生错误，也会处理已经读取到的数据。
		if len(line) > 0 && isFirst {
			basInfo, err = ReadJson(line)
			isFirst = false
		} else if isFirst {
			os.Stdout.WriteString("Empty File!")
			break
		}
		if err != nil { //遇到任何错误立即返回，并忽略 EOF 错误信息
			if err == io.EOF {
				return basInfo, nil
			}
			return basInfo, err
		}
	}
	return basInfo, nil
}
