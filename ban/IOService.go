package ban

import (
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	kc "kanban/conf"
	"log"
	"os"
)

func readFileList(folderPath string) []string {

	files, err := ioutil.ReadDir(folderPath)
	if err != nil {
		log.Fatal(err)
	}

	var filesLen = len(files)
	// fmt.Println(folderPath, "total", filesLen)
	var result []string = make([]string, filesLen)

	for index, f := range files {
		// fmt.Println("FileName", f.Name())
		result[index] = f.Name()
	}

	return result
}

func existFile(filePath string) bool {

	// println("checking path: " + filePath)

	if _, err := os.Stat(filePath); err == nil {
		// path/to/whatever exists
		return true
	} else if errors.Is(err, os.ErrNotExist) {
		// path/to/whatever does *not* exist
		return false
	} else {
		println(err)
	}

	return false
}

func findConfPath(folderPath string) string {
	var configPath = ".././conf/conf.json"

	var banConfPath = folderPath + "/conf.json"
	if existFile(banConfPath) {
		return banConfPath
	}

	if existFile(configPath) {
		return configPath
	}

	configPath = "./conf/conf.json"
	return configPath
}

func loadConfig(folderPath string) kc.Jconf {

	var configPath = findConfPath(folderPath)
	println("loading config: " + configPath)

	var config kc.Jconf = readJsonConfig(configPath)
	return config
}

func readJsonConfig(filePath string) kc.Jconf {

	data, err := ioutil.ReadFile(filePath)
	if err != nil {
		log.Fatal(err)
	}

	// json data
	var result kc.Jconf

	// unmarshall it
	err = json.Unmarshal(data, &result)
	if err != nil {
		log.Fatal("error:", err)
	}

	return result
}

func moveFile(rootPath string, existPath string, newPath string) error {
	existFile := rootPath + "/" + existPath
	changeFile := rootPath + "/" + newPath

	fmt.Println("Changing file : [", existFile, "]")
	fmt.Println("           To : [", changeFile, "]")

	var result error = os.Rename(existFile, changeFile)

	return result
}

func createFile(fullFilePath string) error {

	fmt.Println("Creating file : [", fullFilePath, "]")

	_, err := os.Create(fullFilePath)

	return err
}
