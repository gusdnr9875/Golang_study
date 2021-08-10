// chapter 1 변수

package main

func main() {
	//var 키워드 뒤에 변수명을 적고, 그 뒤에 변수타입을 적는다.

	var a int
	var f float32 = 11.
	println(a)
	println(f)

	a = 10
	println(a)
	f = 12.0
	println(f)
	// 복수 변수들이 선언된 상황에서 초기값을 지정 가능
	var i, j, k int = 1, 2, 3
	println(i)
	println(j)
	println(k)

	//const 키워드 뒤에 상수명을 적고, 그 뒤에 상수타입, 그리고 상수값을 할당
	const c int = 10
	const s string = "Hi"

	// 여러 개의 상수들 묶어서 지정
	const (
		Visa   = "Visa"
		Master = "MasterCard"
		Amex   = "American Express"
	)

}
