// chapter 3 조건문과 반복문

package main

func main() {
	// 대충 C랑 비슷비슷 ㅎ
	var k int = 1
	if k == 1 { //같은 라인
		println("One")
	}
	k = 3
	if k == 1 {
		println("One")
	} else if k == 2 { //같은 라인
		println("Two")
	} else { //같은 라인
		println("Other")
	}

	// switch 문
	var name string
	var category = 1

	switch category {
	case 1:
		name = "Paper Book"
	case 2:
		name = "eBook"
	case 3, 4:
		name = "Blog"
	default:
		name = "Other"
	}
	println(name)

	sum := 0                    //이런식으로도 선언 가능하네..
	for i := 1; i <= 100; i++ { // 대충 똑같은 구조
		sum += i
	}
	println(sum)

	//반복문 + 조건문
	// while은 안쓰고 그냥 for구문으로 while처럼 쓰는 느낌.
	var a = 1
	for a < 15 {
		if a == 5 {
			a += a
			continue // for루프 시작으로
		}
		a++
		if a > 10 {
			break //루프 빠져나옴
		}
	}
	if a == 11 {
		goto END //goto 사용예
	}
	println(a)
END:
	println("End")

}
