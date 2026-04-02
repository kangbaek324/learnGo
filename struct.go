package main

type person struct {
	name string
	age  int
}

type dict struct {
	data map[int]string
}

// struct == 필드의 집합체

// 포인터를 반환하지 않으면 만든 객체가 아닌 객체를 복사해서 반환하게 됨
func newDict() *dict {
	d := dict{}
	d.data = map[int]string{}

	return &d
}

func structEx() {
	// p := person{}

	// p.name = "kang"
	// p.age = 19

	// p := person{name: "kang", age: 19}
	// p := person{"kang", 19}

	// 포인터 반환
	// p := new(person)
	// p.name = "kang"
	// p.age = 19

	// fmt.Printf("p: %v\n", p)

	// dic := newDict()
	// dic.data[0] = "A"

	// fmt.Println(dic)
}
