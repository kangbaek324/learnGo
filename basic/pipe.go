package main

// import (
// 	"fmt"
// 	"time"
// )

// // 각 단계 처리 시간 시뮬레이션
// func stage1(n int) int {
// 	time.Sleep(10 * time.Millisecond) // 파일 읽기
// 	return n
// }

// func stage2(n int) int {
// 	time.Sleep(10 * time.Millisecond) // 리사이징
// 	return n * 2
// }

// func stage3(n int) int {
// 	time.Sleep(10 * time.Millisecond) // 저장
// 	return n
// }

// // ========== 동기 처리 ==========
// func syncProcess(items []int) {
// 	start := time.Now()

// 	for _, n := range items {
// 		r1 := stage1(n)
// 		r2 := stage2(r1)
// 		stage3(r2)
// 	}

// 	fmt.Printf("동기 처리 시간: %v\n", time.Since(start))
// }

// // ========== 파이프라인 처리 ==========
// func gen(items []int) <-chan int {
// 	ch := make(chan int)
// 	go func() {
// 		for _, n := range items {
// 			ch <- stage1(n)
// 		}
// 		close(ch)
// 	}()
// 	return ch
// }

// func resize(in <-chan int) <-chan int {
// 	ch := make(chan int)
// 	go func() {
// 		for n := range in {
// 			ch <- stage2(n)
// 		}
// 		close(ch)
// 	}()
// 	return ch
// }

// func save(in <-chan int) <-chan int {
// 	ch := make(chan int)
// 	go func() {
// 		for n := range in {
// 			ch <- stage3(n)
// 		}
// 		close(ch)
// 	}()
// 	return ch
// }

// func pipelineProcess(items []int) {
// 	start := time.Now()

// 	// 파이프라인 연결
// 	// gen → resize → save
// 	ch1 := gen(items)
// 	ch2 := resize(ch1)
// 	ch3 := save(ch2)

// 	// 결과 수집
// 	for range ch3 {
// 	}

// 	fmt.Printf("파이프라인 처리 시간: %v\n", time.Since(start))
// }

func pipeEx() {
	// 	items := make([]int, 20) // 20개 데이터
	// 	for i := range items {
	// 		items[i] = i + 1
}

// 	fmt.Println("========== 20개 데이터 처리 ==========")
// 	fmt.Println("각 단계: 10ms 소요")
// 	fmt.Println("--------------------------------------")
// 	syncProcess(items)
// 	pipelineProcess(items)
// 	fmt.Println("--------------------------------------")

// 	expected := 20 * 3 * 10
// 	fmt.Printf("동기 예상: %dms\n", expected)
// 	fmt.Printf("파이프라인 예상: ~%dms (단계수 * 10ms + 데이터수 * 10ms)\n", 3*10+20*10)
// }
