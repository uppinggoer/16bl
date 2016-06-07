package main

import (
	"fmt"

	"gopkg.in/yaml.v2"
)

func main() {
	// testEmpty()
	// fmt.Println("-----------------------------------")
	// testFlow()
	// fmt.Println("-----------------------------------")
	testInline()
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
