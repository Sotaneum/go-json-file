package file_test

import (
	"fmt"
	"testing"

	file "github.com/Sotaneum/go-json-file"
)

func TestRunner(t *testing.T) {
	f := &file.File{Path: "./", Name: "test.json"}

	data := map[string]string{}
	data["code"] = "200, 200, 200"

	f.SaveObject(data)
	fmt.Println(f.Load())

	fmt.Println(f.LoadPattern(","))
	f.Save("test")
	f.SavePattern("200, 200, 200", ",")
	fmt.Println(f.Load())
	f.Remove()
}
