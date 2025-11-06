package JSON

import (
	"bytes"
	"fmt"
	"sort"

	"github.com/goccy/go-json"
)

// 序列化
func MustStringify(data interface{}) string {
	result, err := Stringify(data)
	if err != nil {
		panic(err)
	}
	return result
}

func Stringify(data interface{}) (string, error) {
	if str, isString := data.(string); isString {
		return str, nil
	}
	jsonBytes, err := json.Marshal(data)
	if err != nil {
		return "", err
	}
	return string(jsonBytes), nil
}

// 反序列化
func MustParseToMap(obj interface{}) map[string]interface{} {
	var result map[string]interface{}
	if jsonString, ok := obj.(string); ok {
		// 使用 json.Unmarshal 将 JSON 字符串解析到 map 中
		err := json.Unmarshal([]byte(jsonString), &result)
		if err != nil {
			panic(err)
		}
		return result
	} else {
		MustParse(MustStringify(obj), &result)
		return result
	}
}

func MustParse(jsonString string, result interface{}) {
	err := json.Unmarshal([]byte(jsonString), &result)
	if err != nil {
		panic(err)
	}
}

func Parse(jsonString string, result interface{}) error {
	return json.Unmarshal([]byte(jsonString), &result)
}

func ParseBytes(byteArr []byte, result interface{}) error {
	return json.Unmarshal(byteArr, &result)
}

// 序列化 - 按 ASCII 顺序排序键（高效、稳定、兼容复杂对象）
func MustStringifySortByASCII(data interface{}) string {
	result, err := StringifySortByASCII(data)
	if err != nil {
		panic(err)
	}
	return result
}

func StringifySortByASCII(data interface{}) (string, error) {
	var buf bytes.Buffer
	err := marshalOrderedJSON(&buf, data)
	return buf.String(), err
}

func marshalOrderedJSON(buf *bytes.Buffer, data interface{}) error {
	switch v := data.(type) {
	case nil:
		buf.WriteString("null")

	case bool:
		fmt.Fprintf(buf, "%v", v)

	case float64:
		// 处理数字
		b, err := json.Marshal(v)
		if err != nil {
			return err
		}
		buf.Write(b)

	case string:
		// 处理字符串，使用标准 JSON 转义
		b, err := json.Marshal(v)
		if err != nil {
			return err
		}
		buf.Write(b)

	case []interface{}:
		// 处理数组
		buf.WriteString("[")
		for i, item := range v {
			if i > 0 {
				buf.WriteString(",")
			}
			if err := marshalOrderedJSON(buf, item); err != nil {
				return err
			}
		}
		buf.WriteString("]")

	case map[string]interface{}:
		// 处理对象 - 按 ASCII 排序键
		buf.WriteString("{")
		keys := make([]string, 0, len(v))
		for k := range v {
			keys = append(keys, k)
		}
		sort.Strings(keys)

		for i, k := range keys {
			if i > 0 {
				buf.WriteString(",")
			}
			// 序列化键名
			keyBytes, err := json.Marshal(k)
			if err != nil {
				return err
			}
			buf.Write(keyBytes)
			buf.WriteString(":")
			// 递归序列化值
			if err := marshalOrderedJSON(buf, v[k]); err != nil {
				return err
			}
		}
		buf.WriteString("}")

	default:
		// 处理其他类型（结构体、自定义类型等）
		// 先转换为 map[string]interface{}
		b, err := json.Marshal(v)
		if err != nil {
			return err
		}
		var obj map[string]interface{}
		if err := json.Unmarshal(b, &obj); err != nil {
			return err
		}
		return marshalOrderedJSON(buf, obj)
	}
	return nil
}
