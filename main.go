package main

import (
	"encoding/json"
	"fmt"
	// "time"
)

// 1. 简单 json 解析

// func main() {
// 	type FruitBasket struct {
// 		Name    string    `json:"name"`
// 		Fruit   []string  `json:"fruit"`
// 		Id      int64     `json:"id"`
// 		Created time.Time `json:"created"`
// 	}

// 	// json.UnMarshal()方法接收的是字节切片，
// 	// 所以首先需要把JSON字符串转换成字节切片c := []byte(s)
// 	jsonData := []byte(`
//     {
//         "name": "Standard",
//         "fruit": [
//              "Apple",
//             "Banana",
//             "Orange"
//         ],
//         "id": 999,
//         "created": "2018-04-09T23:00:00Z"
//     }
// 	`)

// 	var basket FruitBasket
// 	if err := json.Unmarshal(jsonData, &basket); err != nil {
// 		fmt.Printf("json unmarshal failed, err: %v\n", err)
// 		return
// 	}

// 	// err := json.Unmarshal(jsonData, &basket)
// 	// if err != nil {
// 	// 	fmt.Println(err)
// 	// }

// 	fmt.Println(basket.Name, basket.Fruit, basket.Id)
// 	fmt.Println(basket.Created)
// }

// 2. 解析内嵌对象的JSON

// func main() {
// 	type Fruit struct {
// 		Name     string `json":name"`
// 		PriceTag string `json:"priceTag"`
// 	}

// 	type FruitBasket struct {
// 		Name    string    `json:"name"`
// 		Fruit   Fruit     `json:"fruit"`
// 		Id      int64     `json:"id"`
// 		Created time.Time `json:"created"`
// 	}

// 	jsonData := []byte(`
//     {
//         "name": "Standard",
//         "fruit" : {"name": "Apple", "priceTag": "$1"},
//         "def": 999,
//         "created": "2018-04-09T23:00:00Z"
//     }`)

// 	var basket FruitBasket
//     if err := json.Unmarshal(jsonData, &basket); err != nil {
// 		fmt.Printf("json unmarshal failed, err: %v\n", err)
// 		return
//     }

//     fmt.Println(basket.Name, basket.Fruit, basket.Id)
//     fmt.Println(basket.Created)
// }

// 3. 解析内嵌对象数组的JSON

// func main() {
// 	type Fruit struct {
// 		Name     string `json:"name"`
// 		PriceTag string `json:"priceTag"`
// 	}

// 	type FruitBasket struct {
// 		Name    string    `json:"name"`
// 		Fruit   []Fruit   `json:"fruit"`
// 		Id      int64     `json:"id"`
// 		Created time.Time `json:"created"`
// 	}

// 	jsonData := []byte(`
// 	{
// 		"name": "Standard",
// 		"fruit" : [
// 			{
// 				"name": "Apple",
// 				"priceTag": "$1"
// 			},
// 			{
// 				"name": "Pear",
// 				"priceTag": "$1.5"
// 			}
// 		],
// 		"id": 999,
// 		"created": "2018-04-09T23:00:00Z"
// 	}
// 	`)

// 	var basket FruitBasket
// 	if err := json.Unmarshal(jsonData, &basket); err != nil {
// 		fmt.Printf("json unmarshal failed, err: %v\n", err)
// 		return
// 	}

// 	fmt.Println(basket.Name, basket.Fruit, basket.Id)
// 	fmt.Println(basket.Created)
// }

// 4. 解析具有动态Key的对象
// func main() {
// 	type Fruit struct {
// 		Name     string `json:"name"`
// 		PriceTag string `json:"priceTag"`
// 	}

// 	type FruitBasket struct {
// 		Name    string           `json:"name"`
// 		Fruit   map[string]Fruit `json:"fruit"`
// 		Id      int64            `json:"id"`
// 		Created time.Time        `json:"created"`
// 	}

// 	jsonData := []byte(`
// 	{
// 		"name": "Standard",
// 		"Fruit" : {
// 			"1": {
// 				"Name": "Apple",
// 				"PriceTag": "$1"
// 			},
// 			"2": {
// 				"Name": "Pear",
// 				"PriceTag": "$1.5"
// 			}
// 		},
// 		"id": 999,
// 		"created": "2018-04-09T23:00:00Z"
// 	}
// 	`)

// 	var basket FruitBasket
// 	if err := json.Unmarshal(jsonData, &basket); err != nil {
// 		fmt.Printf("json unmarshal failed, err: %v\n", err)
// 		return
// 	}

// 	for _, item := range basket.Fruit {
// 		fmt.Println(item.Name, item.PriceTag)
// 	}
// }

// 5. 解析包含任意层级的数组和对象的JSON数据
/*
针对包含任意层级的JSON数据，encoding/json包使用：

map[string]interface{} 存储JSON对象
[]interface 存储JSON数组
json.Unmarshl 将会把任何合法的JSON数据存储到一个interface{}类型的值，通过使用空接口类型我们可以存储任意值，但是使用这种类型作为值时需要先做一次类型断言。
*/

// func main() {
// 	jsonData := []byte(`{"Name":"Eve","Age":6,"Parents":["Alice","Bob"]}`)

// 	var v interface{}
// 	json.Unmarshal(jsonData, &v)
// 	data := v.(map[string]interface{})

// 	for k, v := range data {
// 		switch v := v.(type) {
// 		case string:
// 			fmt.Println(k, v, "(string)")
// 		case float64:
// 			fmt.Println(k, v, "(float64)")
// 		case []interface{}:
// 			fmt.Println(k, "(array):")
// 			for i, u := range v {
// 				fmt.Println("    ", i, u)
// 			}
// 		default:
// 			fmt.Println(k, v, "(unknown)")
// 		}
// 	}
// }

// 6. 用 Decoder解析数据流
/*
上面都是使用的UnMarshall解析的JSON数据，如果JSON数据的载体是打开的文件或者HTTP请求体这种
数据流（他们都是io.Reader的实现），我们不必把JSON数据读取出来后再去调用encode/json包的
UnMarshall方法，包提供的Decode方法可以完成读取数据流并解析JSON数据最后填充变量的操作。
*/


