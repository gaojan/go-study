package main

import (
	"crypto/sha256"
	"encoding/hex"
	"fmt"
	"sort"
	"strings"
)

func SortParams(params map[string]string) string {
	keyList := make([]string, 0)
	for key, value := range params {
		if value != "" && value != " " && key != "sign" {
			keyList = append(keyList, key)
		}
	}
	sort.Strings(keyList)
	stringA := ""
	for _, s := range keyList {
		stringA += s + "=" + params[s] + "&"
	}
	if strings.HasSuffix(stringA, "&") {
		stringA = stringA[:len(stringA)-1]
	}
	return stringA
}

func createSha256Str(stringSignTemp string, key string) string {
	sum := sha256.Sum256([]byte(stringSignTemp + key))
	str := strings.ToUpper(hex.EncodeToString(sum[:]))
	return str
}

func GenerateSha256SignByParamsMap(m map[string]string, key string) string {
	return createSha256Str(SortParams(m), key)
}

func main() {
	var m = make(map[string]string)
	m["msgId"] = "31646261383236342d356330362d346665612d393536622d6235383362643832"
	m["msgSrc"] = "xinRuiTai"
	m["requestTimestamp"] = "2019-07-18 15:54:42"
	m["merchantId"] = "898000000000099"
	m["terminalId"] = "00000002"
	m["merOrderDate"] = "20190612"
	m["merOrderId"] = "23453245"
	key := "11112222333344445555666677778888"
	fmt.Println(m)
	sign := GenerateSha256SignByParamsMap(m, key)
	fmt.Println(sign)

}
