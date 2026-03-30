package main

func sliceEx() {
	// 슬라이스 생성
	// var a = []int{1, 2, 3}
	// fmt.Println(a)

	// make를 이용한 슬라이스 생성
	// 슬라이스 타입, 길이, 용량
	// 길이 -> 실제 접근 가능한 공간
	// 용량 -> 메모리에 미리 잡아둔 공간
	//  L> Cap 범위내에서 append하면 실제 복사 X,
	//  L> Cap 범위를 넘으면 그때 실제 메모리 할당
	// var s = make([]int, 5, 10)
	// fmt.Println(len(s), cap(s))
	// fmt.Println(s)

	// Slice에 길이와 용량을 정하지 않으면 Nil Slice로 Nil과 같은 값을 가짐
	// var s []int
	// fmt.Println(s == nil) // true

	// 부분 슬라이스
	// 2:5 << 2번째 인덱스는 inclusive 5번째 인덱스는 exclusive
	// 앞뒤 생략시 첫번째, 마지막 인덱스가 기본값으로 들어감
	// s := []int{1, 2, 3, 4, 5}
	// fmt.Println(s[1:2])

	// append
	// s := []int{1, 2, 3, 4, 5}
	// s = append(s, 6)
	// fmt.Println(s)

	// cap 늘어나는거 체크 하기
	// sliceA := make([]int, 0, 3)
	// for i := 1; i <= 15; i++ {
	// 	fmt.Println(len(sliceA), cap(sliceA))
	// 	sliceA = append(sliceA, i)
	// }

	// -----------------------------------

	// sliceA := []int{1, 2, 3}
	// sliceB := []int{4, 5, 6}

	// sliceA = append(sliceA, sliceB...)

	// fmt.Println(sliceA) // [1 2 3 4 5 6]

	// -----------------------------------

	// target := []int{1, 2, 3}
	// source := []int{4, 5, 6}

	// fmt.Println(target)
	// fmt.Println(source)

	// copy(target, source)

	// fmt.Println(target)
	// fmt.Println(source)

	// target[0] = 7

	// fmt.Println(target)
	// fmt.Println(source)

	// Orginal
	// [1 2 3]
	// [4 5 6]

	// Copy (deep)
	// [4 5 6]
	// [4 5 6]

	// Result
	// [7 5 6]
	// [4 5 6]

	// -----------------------------------

	// 슬라이스로 짜른다면 앞에 짤린만큼 Cap도 줄어듬
	// s := []int{0, 1, 2, 3, 4, 5, 6}
	// fmt.Println(cap(s)) // 7
	// s = s[2:5]          // [X(0, 1), cap(len(2, 3, 4), 5, 6)] 길이는 3 (2, 3, 4) 앞전에꺼는 짤렸으니 빼고 뒤에 원래 있던거 5, 6 포함해서
	// fmt.Println(cap(s)) // 5가 나옴

	// 만약 슬라이스 한 상태에서 append를 하게된다면 다음 인덱스 값이 있어도 (현재 5)
	// 덮어씌워짐
}
