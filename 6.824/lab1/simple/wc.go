package main

import (
	"../mr"
	"strconv"
	"strings"
	"unicode"
)

/**
*
* @param filename 文件名
* @param contents 文本内容
* @return a slice of key/value pairs
* @date 2020/9/4
 */
func Map(filename string, contents string) []mr.KeyValue {

	// 检测单词分割符
	ff := func(r rune) bool { return !unicode.IsLetter(r) }

	// 把文本分割为单词数据
	words := strings.FieldsFunc(contents, ff)

	var kva []mr.KeyValue
	for _, w := range words {
		kv := mr.KeyValue{Key: w, Value: "1"}
		kva = append(kva, kv)
	}
	return kva
}

/**
*
* @param key
* @param values
* @return string
* @date 2020/9/4
 */
func Reduce(key string, values []string) string {

	return strconv.Itoa(len(values))
}
