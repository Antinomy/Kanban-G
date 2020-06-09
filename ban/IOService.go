package ban

import (
	"encoding/json"
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
	fmt.Println(folderPath, "total", filesLen)
	var result []string = make([]string, filesLen)

	for index, f := range files {
		fmt.Println("FileName", f.Name())
		result[index] = f.Name()
	}

	return result
}

func loadConfig() kc.Jconf {
	var configPath = ".././conf/conf.json"

	_, err := os.Stat(configPath)
	if os.IsNotExist(err) {
		configPath = "./conf/conf.json"
	}

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
	var result error = os.Rename(rootPath+"/"+existPath, rootPath+"/"+newPath)
	return result
}
