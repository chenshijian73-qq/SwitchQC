package file

import (
	"bufio"
	"fmt"
	log "github.com/sirupsen/logrus"
	"io/ioutil"
	"os"
	"regexp"
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

func IsExist(filepath string) bool {
	if _, err := os.Stat(filepath); os.IsNotExist(err) {
		return false
	}
	return true
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

func DeleteMatchRow(filePath, key string) (err error) {
	// 1. 打开文件并读取内容
	file, err := os.OpenFile(filePath, os.O_RDWR, 0777)
	if err != nil {
		log.Error(err)
		return
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	var lines []string
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}

	// 2. 将文件内容按行分割成字符串切片
	if err = scanner.Err(); err != nil {
		log.Error(err)
		return
	}

	// 3. 遍历切片并删除匹配特定模式的行
	var filteredLines []string
	re := regexp.MustCompile(key)
	for _, line := range lines {
		if !re.MatchString(line) {
			filteredLines = append(filteredLines, line)
		}
	}

	// 4. 将切片重新组合成一个字符串，并将其写回到原始文件中
	output := []byte(fmt.Sprintf("%s\n", filteredLines[0]))
	for _, line := range filteredLines[1:] {
		output = append(output, []byte(fmt.Sprintf("%s\n", line))...)
	}

	err = file.Truncate(0)
	if err != nil {
		log.Error(err)
		return
	}

	_, err = file.Seek(0, 0)
	if err != nil {
		log.Error(err)
		return
	}

	_, err = file.Write(output)
	if err != nil {
		log.Error(err)
		return
	}

	return
}
