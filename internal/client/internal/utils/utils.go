package utils

func CloneMap(sourceMap map[string]string) map[string]string {
	newMap := make(map[string]string)
	for key, value := range sourceMap {
		newMap[key] = value
	}

	return newMap
}
