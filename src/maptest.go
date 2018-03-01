package main

import "fmt"

//var map_variable map[key_data_type]value_data_type
//map_variable := make(map[key_data_type]value_data_type)
func maptest() {

	var countryCapitalMap map[string]string
	countryCapitalMap = make(map[string]string)
	/* map 插入 key-value 对，各个国家对应的首都 */
	countryCapitalMap["France"] = "Paris"
	countryCapitalMap["Italy"] = "Rome"
	countryCapitalMap["Japan"] = "Tokyo"
	countryCapitalMap["India"] = "New Delhi"

	for country := range countryCapitalMap {
		fmt.Println("Capital of", country, "is", countryCapitalMap[country])
	}
	/* 查看元素在集合中是否存在 */
	captial, ok := countryCapitalMap["United States"]
	if ok {
		fmt.Println("Capital of United States is", captial)
	} else {
		fmt.Println("Capital of United States is not present")
	}

	fmt.Println("test delete")
	/* 删除元素 */
	delete(countryCapitalMap, "France")
	fmt.Println("after delete map")
	/* 打印 map */
	for country := range countryCapitalMap {
		fmt.Println("Capital of", country, "is", countryCapitalMap[country])
	}
}

func Factorial(n uint64) (result uint64) {
	if n > 0 {
		result = n * Factorial(n-1)
		return result
	}
	return 1
}
func fibonacci(n int) int {
	if n < 2 {
		return n
	}
	return fibonacci(n-2) + fibonacci(n-1)
}
