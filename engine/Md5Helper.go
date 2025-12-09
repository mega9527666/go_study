package Md5Helper

import (
	"crypto/md5"
	"encoding/hex"
)

var private_key string = "jhao"

func Init(key string) {
	private_key = key
}

// 计算字符串的 MD5
func GetMd5_default(input string) string {
	hash := md5.Sum([]byte(input))
	return hex.EncodeToString(hash[:])
}

func GetMd5_encrypt(input string) string {
	var md5_value string = GetMd5_default("{" + input + private_key + "}")
	return md5_value
}
