package objectstorage

import (
	"fmt"
	"io/ioutil"
)

func SaveFile(fileContent []byte, fileObjStorePath string) error {

	fullfilePath := fmt.Sprint("./static/objstorage/", fileObjStorePath)
	err := ioutil.WriteFile(fullfilePath, fileContent, 0777)
	if err != nil {
		fmt.Println(err.Error())
		return err
	}
	return nil
}
func LoadFile(filePath string) ([]byte, error) {
	FileObjStorePath := fmt.Sprint("./static/objstorage/", filePath)
	body, err := ioutil.ReadFile(FileObjStorePath)
	if err != nil {
		return nil, err
	}
	return body, nil

}
