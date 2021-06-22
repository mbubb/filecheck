package main

import (
  "fmt"
  "flag"
	"errors"
	"io"
	"os"
	"path/filepath"
	"strings"
)


err := filepath.Walk("/Users/mbubb/git/terraforming/projects",
    func(path string, info os.FileInfo, err error) error {
    if err != nil {
        return err
    }
    fmt.Println(path, info.Size())
    return nil
})
if err != nil {
    log.Println(err)
}