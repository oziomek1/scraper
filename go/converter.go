package main

func slicesToMap(keys []string, values[]string) map[string]string {
	elementMap := make(map[string]string)
	for idx, el := range keys {
		elementMap[el] = values[idx]
	}
	return elementMap
}
