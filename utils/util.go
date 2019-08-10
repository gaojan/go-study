package main

import (
	"crypto/md5"
	"crypto/rand"
	"encoding/base64"
	"encoding/hex"
	"encoding/json"
	"fmt"
	"io"
	"time"
)

// 小工具封装

// 生成32位MD5字符串
func GenerateMd5Str(s string) string {
	h := md5.New()
	h.Write([]byte(s))
	return hex.EncodeToString(h.Sum(nil))
}

// 生成随机字符串
func GenerateSaltStr() string {
	b := make([]byte, 48)
	if _, err := io.ReadFull(rand.Reader, b); err != nil {
		return ""
	}
	return GenerateMd5Str(base64.URLEncoding.EncodeToString(b))
}

// 生成（随机字符串+时间戳）字符串
func GenerateTokenStr() string {
	timestamp := string(time.Now().Unix())
	s := GenerateSaltStr()
	return GenerateMd5Str(s + timestamp)
}

// 字符串转时间戳
func Str2TimeStamp(s string) (time.Time, error) {
	formatTime, err := time.Parse("2006-01-02 15:04:05", s)
	if err != nil {
		fmt.Println("string to time error:", err)
		return time.Time{}, err
	}
	return formatTime, nil
}

// 获取当前时间
func GetCurrentTime() string {
	UnixTime := time.Now().Unix()
	dataTimeStr := time.Unix(UnixTime, 0).Format("2006-01-02 15:04:05")
	return dataTimeStr
}

// 结构体转字典
func StructToMap(obj interface{}) (map[string]interface{}, error) {
	jsonBytes, err := json.Marshal(obj)
	if err != nil {
		return nil, err
	}
	dict := make(map[string]interface{})
	jsonErr := json.Unmarshal(jsonBytes, &dict)
	if jsonErr != nil {
		return nil, jsonErr
	}
	return dict, jsonErr
}

func main() {
	fmt.Println("test")
}
