package main

import (
	"fmt"
)

func readInt(prompt string) int {
	var num int

	for true {
		fmt.Printf(prompt)
		_, err := fmt.Scanf("%d", &num)
		if err != nil {
			fmt.Printf("Invalid number: %s\n", err)
			continue
		}
		break
	}

	return num
}

type IntArr1d []int
type IntMat2d []IntArr1d

type IntStructure interface {
	fill()
	process()
	print()
}

func (r IntArr1d) fill() {
	size := len(r)

	for i := 0; i < size; i++ {
		r[i] = readInt(fmt.Sprintf("Element %d: ", i))
	}
}

func (r IntMat2d) fill() {
	size := len(r)

	for i := 0; i < size; i++ {
		for j := 0; j < size; j++ {
			r[i][j] = readInt(fmt.Sprintf("Element [%d][%d]: ", i, j))
		}
	}
}

func createArray1d() IntArr1d {
	arrSize := readInt("Array length: ")
	return make(IntArr1d, arrSize)
}

func createMat2d() IntMat2d {
	matSize := readInt("Square matrix size: ")

	mat := make(IntMat2d, matSize)
	for i := 0; i < matSize; i++ {
		mat[i] = make(IntArr1d, matSize)
	}
	return mat
}

func (r IntArr1d) process() {
	size := len(r)

	for i := 0; i < size; i += 2 {
		for j := i + 2; j < size; j += 2 {
			if r[i] > r[j] {
				r[i], r[j] = r[j], r[i]
			}
		}
	}
}

func (r IntMat2d) process() {
	size := len(r)
	for row := 0; row < size; row += 2 {
		for i := 0; i < size/2; i++ {
			r[row][i], r[row][size-i-1] = r[row][size-i-1], r[row][i]
		}
	}
}

func (r IntArr1d) print() {
	fmt.Printf("Array: %v\n", r)
}

func (r IntMat2d) print() {
	size := len(r)

	fmt.Println("Matrix:")
	for i := 0; i < size; i++ {
		for j := 0; j < size; j++ {
			fmt.Printf("%d ", r[i][j])
		}
		fmt.Println("")
	}
}

/* Введення, виведення вихідних даних та їх обробку здійснити у вигляді функцій. Врахувати ситуації помилок і панік */
/* Написати програму обробки одно- та двовимірного масивів. Використовувати методи та функції.
а) Написати перевизначені методи. Написати демонстраційну програму для виклику цих методів.
б) Написати інтерфейс замість перевизначених методів із завдання а. Написати демонстраційну програму для виклику інтерфейсу */
/* одновимірний масив: Відсортувати за зростанням лише парні елементи масиву */
/* двовимірний масив: Перевернути всі парні рядки матриці */
func main() {
	var s IntStructure
	s = createArray1d()
	s.fill()

	fmt.Printf("Array before sorting: ")
	s.print()

	s.process()

	fmt.Printf("Array after sorting: ")
	s.print()

	/* ************************* */

	s = createMat2d()
	s.fill()
	fmt.Printf("Matrix before reversing: ")
	s.print()

	s.process()

	fmt.Printf("Matrix after reversing: ")
	s.print()
}
