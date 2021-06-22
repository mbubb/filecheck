package main

import (
	"flag"
	"fmt"
	"os"
	"path"
	"path/filepath"
)

func main() {
	pathPtr := flag.String("basepath", "/Users/mbubb/git/terraforming", "string - path to search")
	flag.Parse()
	
	fmt.Printf(*pathPtr)

	list := getTerraScript(*pathPtr)
	for i, p := range list {
		fmt.Printf("[%d:%s===%s]\n", i, path.Dir(p), path.Base(p))
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
		fmt.Printf("walk error [%v]\n", err)
	}
	return list
}