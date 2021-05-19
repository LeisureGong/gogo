package main

import (
	"../mr"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"plugin"
	"sort"
)

/**
*	simple sequential MapReduce
 */

// for sorting by key
type ByKey []mr.KeyValue

// 返回当前文本串单词个数
func (a ByKey) Len() int {
	return len(a)
}

// 交换位置
func (a ByKey) Swap(i, j int) {
	a[i], a[j] = a[j], a[i]
}

func (a ByKey) Less(i, j int) bool {
	return a[i].Key < a[j].Key
}

func main() {
	if len(os.Args) < 3 {
		fmt.Fprintf(os.Stderr, "Usage: mrsequential xxx.so inputfiles...\n")
		os.Exit(1)
	}

	mapf, reducef := loadPlugin(os.Args[1])

	// read each input file
	// pass it to map
	// accumulate the intermediate map output
	var intermediate []mr.KeyValue

	for _, filename := range os.Args[2:] {
		file, err := os.Open(filename)
		if err != nil {
			log.Fatalf("cannot open %v", filename)
		}

		content, err := ioutil.ReadAll(file)
		if err != nil {
			log.Fatalf("cannot read %v", filename)
		}
		file.Close()
		kva := mapf(filename, string(content))
		intermediate = append(intermediate, kva...)
	}

	sort.Sort(ByKey(intermediate))

	oname := "map-reduce-output"
	ofile, _ := os.Create(oname)

	i := 0
	for i < len(intermediate) {
		j := i + 1
		for j < len(intermediate) && intermediate[j].Key == intermediate[i].Key {
			j++
		}
		var values []string
		for k := i; k < j; k++ {
			values = append(values, intermediate[k].Value)
		}
		output := reducef(intermediate[i].Key, values)

		// 打印
		fmt.Fprintf(ofile, "%v %v\n", intermediate[i].Key, output)

		i = j
	}

	ofile.Close()
}

//  load functions
func loadPlugin(filename string) (func(string, string) []mr.KeyValue, func(string, []string) string) {

	p, err := plugin.Open(filename)

	if err != nil {
		log.Fatalf("connot load plugin %v", filename)
	}
	xmapf, err := p.Lookup("Map")
	if err != nil {
		log.Fatalf("cannot find Map in %v", filename)
	}
	mapf := xmapf.(func(string, string) []mr.KeyValue)

	xredecef, err := p.Lookup("Reduce")
	if err != nil {
		log.Fatalf("cannot find Reduce in %v", filename)
	}
	reducef := xredecef.(func(string, []string) string)

	return mapf, reducef
}
