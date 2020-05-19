package ban

import (
	"fmt"
	"io/ioutil"
	"log"
)

func readFileList(folderPath string) []string {

	files, err := ioutil.ReadDir(folderPath)
	if err != nil {
		log.Fatal(err)
	}

	var filesLen = len(files)
	fmt.Println(filesLen)
	var result []string = make([]string, filesLen)

	for index, f := range files {
		fmt.Println(f.Name())
		result[index] = f.Name()
	}

	return result
}
