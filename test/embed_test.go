package test

import (
	_ "embed"
)

// //go:embed version.txt
// var version string

// //go:embed ../pic.png
// var pic []byte

// func TestString(t *testing.T)  {
// 	fmt.Println(version)
// }

// func TestByte(t *testing.T)  {
// 	err := ioutil.WriteFile("pic_new.png", pic, fs.ModePerm)
// 	if err != nil {
// 		panic(err)
// 	}
// }

// //go:embed files/a.txt
// //go:embed files/b.txt
// //go:embed files/c.txt
// var files embed.FS
// func TestMultipleFiles(t *testing.T)  {
// 	a, err := files.ReadFile("files/a.txt")
// 	if err != nil {
// 		panic(err)
// 	}
// 	b, _ := files.ReadFile("files/b.txt")
// 	c, _ := files.ReadFile("files/c.txt")

// 	fmt.Println(string(a))
// 	fmt.Println(string(b))
// 	fmt.Println(string(c))
// }

// //go:embed files/*.txt
// var path embed.FS

// func TestPathMatcher(t *testing.T)  {
// 	dirEntries, err := path.ReadDir("files")
// 	if err != nil {
// 		panic(err)
// 	}

// 	for _, entry := range dirEntries {
// 		if !entry.IsDir() {
// 			fmt.Println(entry.Name())
// 			file, err := path.ReadFile("files/"+ entry.Name())
// 			if err != nil {
// 				panic(err)
// 			}

// 			fmt.Println(string(file))
// 		}
// 	}
// }