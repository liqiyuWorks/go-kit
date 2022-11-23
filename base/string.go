package base

import (
	"bytes"
	"log"
	"strconv"
	"unicode"
)

// 内嵌bytes.Buffer，支持连写
type Buffer struct {
	*bytes.Buffer
}

func NewBuffer() *Buffer {
	return &Buffer{Buffer: new(bytes.Buffer)}
}

func (b *Buffer) Append(i interface{}) *Buffer {
	switch val := i.(type) {
	case int:
		b.append(strconv.Itoa(val))
	case int64:
		b.append(strconv.FormatInt(val, 10))
	case uint:
		b.append(strconv.FormatUint(uint64(val), 10))
	case uint64:
		b.append(strconv.FormatUint(val, 10))
	case string:
		b.append(val)
	case []byte:
		b.Write(val)
	case rune:
		b.WriteRune(val)
	}
	return b
}
func (b *Buffer) append(s string) *Buffer {
	defer func() {
		if err := recover(); err != nil {
			log.Println("*****内存不够了！******")
		}
	}()
	b.WriteString(s)
	return b
}

/**
 * @description: 首字母大写
 * @return {*}
 * @author: liqiyuWorks
 */
func Ucfirst(str string) string {
	for i, v := range str {
		return string(unicode.ToUpper(v)) + str[i+1:]
	}
	return ""
}

/**
 * @description: 首字母小写
 * @param {string} str
 * @return {*}
 * @author: liqiyuWorks
 */
func Lcfirst(str string) string {
	for i, v := range str {
		return string(unicode.ToLower(v)) + str[i+1:]
	}
	return ""
}

/**
 * @description: 截取字符串
 * @param {string} str
 * @param {int} n
 * @param {bool} reverse: true: 反向切片, false: 正向切片
 * @return {*}
 * @author: liqiyuWorks
 */
func StringSub(str string, n int, reverse bool) string {
	if !reverse {
		return string([]byte(str)[:n])
	} else {
		return string([]rune(str)[:len([]rune(str))-n])
	}
}
