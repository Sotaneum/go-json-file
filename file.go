package file

import (
	"encoding/json"
	"io/ioutil"
	"os"
	"strings"
)

// File : 파일 정보를 저장합니다.
type File struct {
	Path string
	Name string
}

func (f *File) getFilePath() string {
	return f.Path + "/" + f.Name
}

// SaveObject : 파일을 저장합니다.
func (f *File) SaveObject(data interface{}) error {
	str, err := json.Marshal(data)
	if err != nil {
		return err
	}
	return f.SavePattern(string(str), "")
}

// SavePatternObject : 파일을 저장합니다.
func (f *File) SavePatternObject(data interface{}, pattern string) error {
	str, err := json.Marshal(data)
	if err != nil {
		return err
	}
	return f.SavePattern(string(str), pattern)
}

// Save : 파일을 저장합니다.
func (f *File) Save(data string) error {
	return f.SavePattern(data, "")
}

// SavePattern : 파일을 저장합니다.
func (f *File) SavePattern(data, pattern string) error {
	file, err := os.OpenFile(f.getFilePath(), os.O_WRONLY|os.O_CREATE, 0600)

	if err != nil {
		return err
	}

	defer file.Close()

	if _, writeErr := file.WriteString(data + pattern); writeErr != nil {
		return writeErr
	}

	return nil
}

// Remove : 파일을 삭제합니다.
func (f *File) Remove() error {
	return os.Remove(f.getFilePath())
}

// Load : 파일을 저장합니다.
func (f *File) Load() string {
	file, err := os.Open(f.getFilePath())

	if err != nil {
		return ""
	}

	d, _ := ioutil.ReadAll(file)

	return string(d)
}

// LoadPattern : pattern으로 split를 해서 array로 반환합니다.
func (f *File) LoadPattern(pattern string) []string {
	return strings.Split(f.Load(), pattern)
}
