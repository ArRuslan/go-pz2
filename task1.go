package main

import "fmt"

type IntArr1d []int

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

func (r *IntArr1d) print() {
	fmt.Printf("Array: %v\n", r)
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
}
