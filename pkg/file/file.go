package file

import (
	"fmt"
	log "github.com/sirupsen/logrus"
	"io/ioutil"
	"os"
)

func CreateFile(filepath string, content string) (err error) {
	err = ioutil.WriteFile(filepath, []byte(content), 0644)
	if err != nil {
		return
	}
	log.Info(fmt.Sprintf("File %s created successfully!", filepath))
	fmt.Println("File created successfully!")
	return
}

func SaveFile(filepath string, content string) (err error) {
	// 打开文件
	f, err := os.OpenFile(filepath, os.O_CREATE|os.O_RDWR|os.O_TRUNC, 0644)
	if err != nil {
		return err
	}
	defer f.Close()

	if err != nil {
		return err
	}
	// 写入文件内容
	_, err = f.WriteString(content)
	if err != nil {
		return err
	}

	log.Info(fmt.Sprintf("create file %s", filepath))
	return nil
}

func SaveFileIfModified(filepath string, content string) (err error) {
	// 读取文件内容
	fileContent, err := ioutil.ReadFile(filepath)
	if err != nil {
		fmt.Println(err)
		return err
	}
	// 将文件内容与 content 进行比较
	if string(fileContent) != content {
		// 如果不一致，则将 content 写入文件
		err = ioutil.WriteFile(filepath, []byte(content), 0644)
		if err != nil {
			return err
		}
	}

	return nil
}

func DeleteFile(filepath string) (err error) {
	err = os.Remove(filepath)
	if err != nil {
		return err
	}
	log.Info(fmt.Sprintf("delete file %s", filepath))
	return nil
}
