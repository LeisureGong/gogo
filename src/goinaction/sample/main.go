package main

import (
	_ "github.com/LeisureGong/gogo/src/goinaction/sample/matchers"
	"github.com/LeisureGong/gogo/src/goinaction/sample/search"
	"log"
	"os"
)

func init() {
	log.SetOutput(os.Stdout)
}

func main() {
	search.Run("president")
}
