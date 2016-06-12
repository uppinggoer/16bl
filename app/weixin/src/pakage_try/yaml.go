package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
	"util"

	"gopkg.in/yaml.v2"
)

func main() {
	// testEmpty()
	// fmt.Println("-----------------------------------")
	// testFlow()
	// fmt.Println("-----------------------------------")
	// testInline()
	// fmt.Println("-----------------------------------")
	// testDumplicate()
	// fmt.Println("-----------------------------------")
	// testMap()
	fmt.Println("-----------------------------------")
	// testFile()
	testFileStruct()
	// fmt.Println("-----------------------------------")
	// testList()
}

func testFileStruct() {
	// 界面上区块按钮  如 banner/nav/func
	type button struct {
		Icon    string `json:"icon"`
		Name    string `json:"name"`
		Url     string `json:"url"`
		Trigger string `json:"trigger"`
	}

	// 下方列表
	type class struct {
		Name        string `json:"name"`
		Img         string `json:"img"`
		GoodsIdList []int  `yaml:"goodsIdList" json:"-"`
	}

	// 下方列表
	type Home struct {
		Banner []button `json:"banner"`
		Nav    []button `json:"nav"`
		Func   []button `json:"func"`
		Class  []class  `json:"class"`
	}

	file, _ := os.Open("/Users/yanghongzhi/work/16bl/app/weixin/conf/app/home.yaml")
	confData, _ := ioutil.ReadAll(file)

	t := &Home{}
	yaml.Unmarshal(confData, t)
	// fmt.Println(t)
	b, _ := json.Marshal(t)
	fmt.Println(string(b))
}
func testList() {
	var data = `
a: Test!
b:
  c: 
    w: 2
    d: 4
  d: [3, 4]
m:
  a:
    k: 
    p: 
  b:
    w: 6
    i: 5
`
	m := make(map[interface{}]interface{})
	yaml.Unmarshal([]byte(data), &m)
	foo := m["b"].(map[interface{}]interface{})["c"].(map[interface{}]interface{})

	t := make(map[string]interface{})
	t["A"] = "Sa"
	t["B"] = "Sb"
	t["C"] = "Sc"

	var list []interface{}
	list = append(list, t)
	list = append(list, t)
	list = append(list, t)
	foo["w"] = list
	d, _ := yaml.Marshal(&m)

	fmt.Println(string(d))
}
func testEmpty() {
	type T struct {
		F int "a,omitempty"
		B int
	}
	var txt []byte
	txt, _ = yaml.Marshal(&T{B: 2}) // Returns "b: 2\n"
	fmt.Printf("%v\n", string(txt))

	txt, _ = yaml.Marshal(&T{F: 1}) // Returns "a: 1\nb: 0\n"
	fmt.Printf("%v\n", string(txt))
}

func testMap() {
	var data = `
a: Test!
b:
  c: 
    w: 2
    d: 4
  d: [3, 4]
m:
  a:
    k: 5
    p: 6
  b:
    w: 6
    i: 5
`
	m := make(map[interface{}]interface{})
	yaml.Unmarshal([]byte(data), &m)
	foo := m["b"].(map[interface{}]interface{})["c"].(map[interface{}]interface{})["w"]
	fmt.Printf("--- map:\n%v\n", foo)

	fmt.Println("--- mapv2 ----------")

	var inter interface{}
	err := util.Get(&m, &inter, "b", "c", "w")
	if nil != err {
		fmt.Println(err)
	}
	fmt.Printf("--- map v2:\n%v\n", inter)
	// k := map[interface{}]interface{}(m["b"])
	// fmt.Printf("--- map:\n%v\n", k)
	// n := k["c"]
	// fmt.Printf("--- map:\n%v\n", n)
	// yaml.Unmarshal([]byte(m["b"]), &m)
	// fmt.Printf("--- map b:\n%v\n", m)
}

func testDumplicate() {
	var data = `
a: Test!
b:
  c: 
    m:2
    d:4
  d: [3, 4]
m:
  a:5
  b:4
`
	var data2 = `
a: Test!
b:
  d: [2, 4]
m:
  d:5
`
	type T struct {
		A string
		B struct {
			RenamedC int   `yaml:"c"`
			D        []int `yaml:",flow"`
		}
	}
	t := T{}
	yaml.Unmarshal([]byte(data+data2), &t)
	fmt.Printf("--- map:\n%v\n", t)

	m := make(map[interface{}]interface{})
	yaml.Unmarshal([]byte(data+data2), &m)
	fmt.Printf("--- map:\n%v\n", m["a"])
}

// 序列化可以转回 map/struct
// 字符串到结构(map/struct)
func testFlow() {
	var data = `
a: Test!
b:
  c: 2
  d: [3, 4]
`
	type T struct {
		A string
		B struct {
			RenamedC int   `yaml:"c"`
			D        []int `yaml:",flow"`
		}
	}

	t := T{}

	yaml.Unmarshal([]byte(data), &t)
	fmt.Printf("--- struct:\n%v", t)

	d, _ := yaml.Marshal(&t)
	fmt.Printf("--- struct text:\n%s\n", string(d))

	m := make(map[interface{}]interface{})

	yaml.Unmarshal([]byte(data), &m)
	fmt.Printf("--- map:\n%v", m)

	d, _ = yaml.Marshal(&m)
	fmt.Printf("--- map text:\n%s\n", string(d))

	yaml.Unmarshal([]byte(data), &t)
	fmt.Printf("--- struct text:\n%v\n", t)
}

// 序列化可以转回 map/struct
// 字符串到结构(map/struct)
// 嵌入
func testInline() {
	var data = `
a: Test!
b:
  c: 2
  d: [3, 4]
`

	type InB struct {
		B struct {
			RenamedC int   `yaml:"c"`
			D        []int `yaml:",flow"`
		}
	}
	type OuA struct {
		A  string
		TT InB `yaml:",inline"`
	}

	t := OuA{}

	yaml.Unmarshal([]byte(data), &t)
	fmt.Printf("--- struct:\n%v", t)

	d, _ := yaml.Marshal(&t)
	fmt.Printf("--- struct text:\n%s\n", string(d))

	m := make(map[interface{}]interface{})

	yaml.Unmarshal([]byte(data), &m)
	fmt.Printf("--- map:\n%v", m)

	d, _ = yaml.Marshal(&m)
	fmt.Printf("--- map text:\n%s\n", string(d))

	yaml.Unmarshal([]byte(data), &t)
	fmt.Printf("--- struct text:\n%v\n", t)
}
