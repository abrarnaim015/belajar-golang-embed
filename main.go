package main

import (
	"embed"
	_ "embed"
	"fmt"
	"io/fs"
	"io/ioutil"
)

//go:embed version.txt
var version string

//go:embed pic.png
var pic []byte

//go:embed files/*.txt
var path embed.FS

func main()  {
	fmt.Println(version)

	err := ioutil.WriteFile("pic_nex.png", pic, fs.ModePerm)
	if err != nil {
		panic(err)
	}

	dirEntries, err := path.ReadDir("files")
	if err != nil {
		panic(err)
	}

	for _, entry := range dirEntries {
		if !entry.IsDir() {
			fmt.Println(entry.Name())
			file, err := path.ReadFile("files/"+ entry.Name())
			if err != nil {
				panic(err)
			}

			fmt.Println(string(file))
		}
	}
}