package main

import "bytes"

func getKey(str string) string {
	// 记录 26 个英文字母出现的次数
	cnt := [26]int{}
	for _, c := range str {
		cnt[int(c-'a')]++
	}
	// 各个字母出现次数相同即为一组，拥有相同的 key byte.Buffer 缓冲区
	var key bytes.Buffer
	for _, n := range cnt {

		key.WriteByte(byte(n))

	}

	return key.String()
}

func groupAnagrams(strs []string) [][]string {
	result := [][]string{}
	kv := map[string][]string{}

	for _, s := range strs {
		key := getKey(s)
		kv[key] = append(kv[key], s)
	}


	for _, v := range kv {

		result = append(result, v)
	}
	return result
}
