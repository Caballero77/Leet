package main

import "fmt"

type TimeMapItem struct {
	Value     string
	Timestamp int
}

type TimeMap struct {
	Map map[string][]TimeMapItem
}

func Constructor() TimeMap {
	return TimeMap{Map: make(map[string][]TimeMapItem)}
}

func (this *TimeMap) Set(key string, value string, timestamp int) {
	items, exist := this.Map[key]
	newItem := TimeMapItem{Value: value, Timestamp: timestamp}
	if exist {
		this.Map[key] = append(items, newItem)
	} else {
		this.Map[key] = []TimeMapItem{newItem}
	}
}

func (this *TimeMap) Get(key string, timestamp int) string {
	items, exist := this.Map[key]
	if !exist {
		return ""
	}
	start, end := 0, len(items)-1
	max := TimeMapItem{Timestamp: 0, Value: ""}
	for start <= end {
		middle := (start + end) / 2
		if items[middle].Timestamp > timestamp {
			end = middle - 1
		} else {
			max = items[middle]
			start = middle + 1
		}
	}
	return max.Value
}

func main() {
	timeMap := Constructor()
	fmt.Println(timeMap.Get("foo", 1))
	timeMap.Set("foo", "bar", 1)
	fmt.Println(timeMap.Get("foo", 1))
	fmt.Println(timeMap.Get("foo", 3))
	timeMap.Set("foo", "bar2", 4)
	fmt.Println(timeMap.Get("foo", 4))
	fmt.Println(timeMap.Get("foo", 5))
	fmt.Println(timeMap.Get("foo", 1))
}
