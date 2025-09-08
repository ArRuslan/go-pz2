package main

import "fmt"

type IntArr1d []int
type IntMat2d []IntArr1d

func createAndFillArray1d() IntArr1d {
	var arrSize uint
	var err error

	fmt.Printf("Array length: ")
	_, err = fmt.Scanf("%d", &arrSize)
	if err != nil {
		panic(err)
	}

	arr := make(IntArr1d, arrSize)

	var i uint
	for i = 0; i < arrSize; i++ {
		fmt.Printf("Element %d: ", i)
		_, err = fmt.Scanf("%d", &arr[i])
		if err != nil {
			panic(err)
		}
	}

	return arr
}

func createAndFillMat2d() IntMat2d {
	var matSize uint
	var err error

	fmt.Printf("Square matrix size: ")
	_, err = fmt.Scanf("%d", &matSize)
	if err != nil {
		panic(err)
	}

	mat := make(IntMat2d, matSize)

	var i uint
	for i = 0; i < matSize; i++ {
		mat[i] = make(IntArr1d, matSize)

		var j uint
		for j = 0; j < matSize; j++ {
			fmt.Printf("Element [%d][%d]: ", i, j)
			_, err = fmt.Scanf("%d", &mat[i][j])
			if err != nil {
				panic(err)
			}
		}
	}

	return mat
}

func (r *IntArr1d) sortEven() {
	size := len(*r)

	for i := 0; i < size; i += 2 {
		for j := i + 2; j < size; j += 2 {
			if (*r)[i] > (*r)[j] {
				(*r)[i], (*r)[j] = (*r)[j], (*r)[i]
			}
		}
	}
}

func (r *IntMat2d) reverseEven() {
	size := len(*r)
	for row := 0; row < size; row += 2 {
		for i := 0; i < size/2; i++ {
			(*r)[row][i], (*r)[row][size-i-1] = (*r)[row][size-i-1], (*r)[row][i]
		}
	}
}

func (r *IntArr1d) print() {
	fmt.Printf("Array: %v\n", r)
}

func (r *IntMat2d) print() {
	size := len(*r)

	fmt.Println("Matrix:")
	for i := 0; i < size; i++ {
		for j := 0; j < size; j++ {
			fmt.Printf("%d ", (*r)[i][j])
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
	arr := createAndFillArray1d()

	fmt.Printf("Array before sorting: ")
	arr.print()

	arr.sortEven()

	fmt.Printf("Array after sorting: ")
	arr.print()

	mat := createAndFillMat2d()
	fmt.Printf("Matrix before reversing: ")
	mat.print()

	mat.reverseEven()

	fmt.Printf("Matrix after reversing: ")
	mat.print()
}
