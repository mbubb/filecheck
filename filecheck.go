package main

import (
	"bufio"
	"flag"
	"fmt"
	"log"
	"os"
	"path/filepath"
)

func main() {
	pathPtr := flag.String("basepath", "/Users/mbubb/git/terraforming", "path to search for files")
	flag.Parse()
	
	fmt.Printf(*pathPtr)

	list := getTerraScript(*pathPtr)
	//for i, p := range list {
	//	fmt.Printf("[%d:%s===%s]\n", i, path.Dir(p), path.Base(p))
	// }
	for _, tfile := range list {
		fmt.Printf("%s\n", tfile)
		readFile(tfile)
	}

}

func getTerraScript(rootpath string) []string {

	list := make([]string, 0, 10)

	err := filepath.Walk(rootpath, func(path string, info os.FileInfo, err error) error {
		if info.IsDir() {
			return nil
		}
		if filepath.Ext(path) == ".tf" {
			list = append(list, path)
		}
		return nil
	})
	if err != nil {
		log.Fatal(err)
	}
	return list
}

func readFile(fname string) {
    f, err := os.Open(fname)
    if err != nil {
        log.Fatal(err)
    }
    defer f.Close()

		scanner := bufio.NewScanner(f)
    for scanner.Scan() {
        fmt.Println(scanner.Text())
    }

    if err := scanner.Err(); err != nil {
        log.Fatal(err)
    }

}