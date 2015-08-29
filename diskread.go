package main

import (
	"flag"
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"os"
	"path/filepath"
	"time"
)

var root string

func init() {
	flag.StringVar(&root, "root", ".", "Root of directory to start in")
}

func main() {
	flag.Parse()
	var sz int64
	var files int64
	t := time.Now()
	pt := t
	err := filepath.Walk(root, func(p string, info os.FileInfo, err error) error {
		if info == nil {
			log.Printf("nil info for [%s]", p)
		} else if !info.IsDir() {
			f, err := os.Open(p)
			if err != nil {
				log.Print(err)
				return nil
			}
			_, err = io.Copy(ioutil.Discard, f)
			if err != nil {
				log.Print(err)
			}
			err = f.Close()
			if err != nil {
				log.Print(err)
			}
			sz += info.Size()
			files++
		}
		if time.Since(pt) > time.Second {
			fmt.Printf("%d files\r", files)
			pt = time.Now()
		}
		return nil
	})
	if err != nil {
		log.Print(err)
	}
	d := time.Since(t)
	fmt.Printf("%d files\n", files)
	fmt.Printf("%.2f KB in %s\n", float64(sz)/1024.0, d.String())
	fmt.Printf("%.2f KB/sec\n", float64(sz)/1024.0/d.Seconds())
}
