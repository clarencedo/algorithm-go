package leetcode

import (
	"strconv"
	"strings"
)

// 自增
type Codec struct {
	dataBase map[int]string
	id       int
}

func Constructor() Codec {
	return Codec{map[int]string{}, 0}

}

// Encodes a URL to a shortened URL.
func (this *Codec) encode(longUrl string) string {
	this.id++
	this.dataBase[this.id] = longUrl
	return "http://tinyurl.com/" + strconv.Itoa(this.id)

}

// Decodes a shortened URL to its original URL.
func (this *Codec) decode(shortUrl string) string {
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
