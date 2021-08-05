// chapter 2 연산자

// chapter 1 변수

package main

func main() {
	// 대충 C랑 비슷비슷 ㅎ
	var a, b, c int = 1, 2, 0
	c = (a + b) * 5
	println(c)
	c++ // 근데 증감연산은 print안에 못 포함시킴;;
	println(c)

	println(a == b)
	println(a != c)
	println(a >= b)

	c = (a & b) << 5 // 비트 연산도 똑같
	println(c)

	//요새 언어 답지 않게 포인터도 있다!

	var k int = 10
	var p = &k  //k의 주소를 할당
	println(*p) //p가 가리키는 주소에 있는 실제 내용을 출력
}
