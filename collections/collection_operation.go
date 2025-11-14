package Collections

func ContainInt64(ids []int64, id int64) bool {
	for _, v := range ids {
		if v == id {
			return true
		}
	}
	return false
}

func ContainKey(dataMap map[string]interface{}, key string) bool {
	_, exists := dataMap[key]
	return exists
}
