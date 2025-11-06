package StringUtil

import (
	"crypto/rand"
	"fmt"
	"strconv"
	"strings"

	JSON "github.com/litecodex/go-common/json_util"
)

func MustToInt64(data string) int64 {
	result, err := strconv.ParseInt(data, 10, 64)
	if err != nil {
		panic(err)
	}
	return result
}

func ToInt64WithDefault(data string, defaultValue int64) int64 {
	result, err := strconv.ParseInt(data, 10, 64)
	if err != nil {
		return defaultValue
	}
	return result
}

func MustToInt(data string) int {
	value, err := strconv.Atoi(data)
	if err != nil {
		panic(fmt.Sprintf("Failed to convert string to int: %s", data))
	}
	return value
}

func MustToBool(str string) bool {
	b, err := strconv.ParseBool(str)
	if err != nil {
		panic(err)
	}
	return b
}

// 生成一个指定长度的随机字符串
const charset = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"

func RandomString(length int) string {
	// 生成一个字节切片来存储随机数据
	byteSlice := make([]byte, length)
	// 从 crypto/rand 生成安全的随机字节
	if _, err := rand.Read(byteSlice); err != nil {
		panic(err)
	}

	// 将随机字节映射到 charset 中的字符
	for i, b := range byteSlice {
		byteSlice[i] = charset[b%byte(len(charset))]
	}

	return string(byteSlice)
}

func MustToString(obj interface{}) string {
	if obj == nil {
		return ""
	}

	switch v := obj.(type) {
	case string:
		return obj.(string)
	case int:
		return fmt.Sprintf("%d", v)
	case int8:
		return fmt.Sprintf("%d", v)
	case int16:
		return fmt.Sprintf("%d", v)
	case int32:
		return fmt.Sprintf("%d", v)
	case int64:
		return fmt.Sprintf("%d", v)
	case uint:
		return fmt.Sprintf("%d", v)
	case uint8:
		return fmt.Sprintf("%d", v)
	case uint16:
		return fmt.Sprintf("%d", v)
	case uint32:
		return fmt.Sprintf("%d", v)
	case uint64:
		return fmt.Sprintf("%d", v)
	case float32:
		return fmt.Sprintf("%f", v)
	case float64:
		return fmt.Sprintf("%f", v)
	case bool:
		return fmt.Sprintf("%t", v)
	default:
		// 对于未知类型，使用 JSON 序列化。序列化失败后会抛异常
		return JSON.MustStringify(obj)
	}
}

// String数组转换为int64数组
func MustToInt64List(str string) []int64 {
	strSlice := strings.Split(str, ",")
	var intSlice []int64
	for _, str := range strSlice {
		num, err := strconv.ParseInt(str, 10, 64)
		if err != nil {
			panic(err)
		}
		intSlice = append(intSlice, num)
	}
	return intSlice
}
