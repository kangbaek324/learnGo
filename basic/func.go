package main

func funcEx() {
	// say("hello")
	// say("hello", "world", "!")

	// rs, err := devide(1, 0)
	// if err != nil {
	// 	fmt.Println(err)
	// } else {
	// 	fmt.Println(rs)
	// }
}

// 변수명, 타입
// func say(msg string) {
// 	fmt.Println(msg)
// }

// 값이 여러개일 경우 ...타입
// func say(msg ...string) {
// 	for _, s := range msg {
// 		fmt.Println(s)
// 	}
// }

// 반환값 지정하고 싶을때
// func say() int

// 여러개 지정할때
// func say() int, int, string

// 실제 변수명을 적는다면 자동으로 그 값을 return
// 코드 맨 밑줄 빈값이라도 return을 작성해줘야함 (없을시 에러)
// func say() count int, total int

// 에러 처리시 맨 마지막 인자에 error
// func devide(a int, b int) (int, error) {
// 	if b == 0 {
// 		return 0, fmt.Errorf("0으로 나눌 수 없음")
// 	}
// 	return a / b, nil
// }
