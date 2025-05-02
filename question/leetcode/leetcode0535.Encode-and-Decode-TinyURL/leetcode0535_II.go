package leetcode

import (
	"strconv"
	"strings"
)

const k1, k2 = 1117, 1e9 + 7

type Code struct {
	// (key, longUrl)
	dataBase map[int]string
	// 原始url -> 短码 （去重）
	urlToKey map[string]int
}

func ConstructorII() Code {
	return Code{map[int]string{}, map[string]int{}}

}

// Encodes a URL to a shortened URL.
func (this *Code) encode(longUrl string) string {
	if key, ok := this.urlToKey[longUrl]; ok {
		return "http://tinyurl.com/" + strconv.Itoa(key)
	}
	key, base := 0, 1
	for _, c := range longUrl {
		key = (key*base + int(c)) % k2
		base = (base * k1) % k2
	}
	for this.dataBase[key] != "" {
		key = (key + 1) % k2
	}
	this.dataBase[key] = longUrl
	this.urlToKey[longUrl] = key
	return "http://tinyurl.com/" + strconv.Itoa(key)

}

// Decodes a shortened URL to its original URL.
func (this *Code) decode(shortUrl string) string {
	i := strings.LastIndexByte(shortUrl, '/')
	id, _ := strconv.Atoi(shortUrl[i+1:])
	return this.dataBase[id]
}

/**
 * Your Codec object will be instantiated and called as such:
 * obj := Constructor();
 * url := obj.encode(longUrl);
 * ans := obj.decode(url);
 */
