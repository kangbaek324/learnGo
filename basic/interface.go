package main

// interface == 메서드의 집합체

// type Shape interface {
// 	area() float64
// 	perimeter() float64
// }

// type Rect struct {
// 	width, height float64
// }

// type Circle struct {
// 	radius float64
// }

// func (r Rect) area() float64 { return r.width * r.height }
// func (r Rect) perimeter() float64 {
// 	return 2 * (r.width + r.height)
// }

// func (c Circle) area() float64 {
// 	return math.Pi * c.radius * c.radius
// }
// func (c Circle) perimeter() float64 {
// 	return 2 * math.Pi * c.radius
// }

// // shapes 인터페이스에 있는 메서드가 구현되어 있으면 다 ㄱㅊ
// func showArea(shapes ...Shape) {
// 	for _, s := range shapes {
// 		a := s.area() //인터페이스 메서드 호출
// 		println(a)
// 	}
// }

func intercaeEx() {
	// r := Rect{10., 20.}
	// c := Circle{10}

	// showArea(r, c)
}

// // 빈 인터페이스로 any 구현가능
// func blabla(v interface{}) {}

// // Go 1.18+ 부터 any 타입 지원
// func blabla2(v any) {}

// ex

// type error interface {
// 	Error() string // 이 메서드만 있으면 error 취급
// }

// type MyError struct{ msg string }

// func (e MyError) Error() string { return e.msg } // error 인터페이스 구현!

// ---------------------

// // 파일 읽기
// f, _ := os.Open("test.txt")
// // f는 io.Reader 구현함

// // HTTP 응답 읽기
// resp, _ := http.Get("https://google.com")
// // resp.Body도 io.Reader 구현함

// // 버퍼 읽기
// buf := strings.NewReader("hello")
// // buf도 io.Reader 구현함

// io.Reader를 받는 함수 하나로
// Go가 알아서 맞는 타입의 read를 호출함
// func readAll(r io.Reader) {
//     data, _ := io.ReadAll(r)
//     fmt.Println(string(data))
// }

// // 파일도 OK
// readAll(file)

// // HTTP 응답도 OK
// readAll(resp.Body)

// // 문자열도 OK
// readAll(strings.NewReader("hello"))
