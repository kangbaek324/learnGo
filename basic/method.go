package main

// type Rect struct {
// 	width, hight int
// }

// method의 종류
// value reciver, pointer reciver
// value는 데이터를 복사하여 전달하기때문에 데이터 변경이 일어나도 원본 영향 X
// pointer는 포인터 주소만 전달하기때문에 데이터 변경시 원본 데이터도 함께 변경됨

// func (r *Rect) area() int {
// 	r.hight++
// 	return r.width * r.hight
// }

// func (r Rect) area2() int {
// 	r.hight++
// 	return r.width * r.hight
// }

func methodEx() {
	// rect := Rect{10, 20}
	// area := rect.area()
	// area2 := rect.area2()

	// fmt.Println(area)
	// fmt.Println(rect)
}
