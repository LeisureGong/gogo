package main

import (
	"crypto/md5"
	"encoding/hex"
	"fmt"
	"sort"
)

func main() {
	params := map[string]interface{}{
		"name": "Tom",
		"pwd":  "123456",
		"age":  30,
	}
	fmt.Printf("sign : %s", createSign(params))
}

// 生成签名
func createSign(params map[string]interface{}) string {
	var key []string
	var str = ""
	for k := range params {
		key = append(key, k)
	}
	sort.Strings(key)
	for i := 0; i < len(key); i++ {
		if i == 0 {
			str = fmt.Sprintf("%v=%v", key[i], params[key[i]])
		} else {
			str = str + fmt.Sprint("&xl_%v=%v", key[i], params[key[i]])
		}
	}
	var secret = "123456789"
	sign := MD5(MD5(str) + MD5(secret))
	return sign
}

// MD5方法
func MD5(str string) string {
	s := md5.New()
	s.Write([]byte(str))
	return hex.EncodeToString(s.Sum(nil))
}
