package main

import (
	"bufio"
	"errors"
	"fmt"
	"math"
	"math/rand"
	"os"
)

type IntArr []int
type IntMat []IntArr

type Person struct {
	fullName [3]string
	address  string
	birthday string
}

var argErr = errors.New("expected at least one argument")

func printPersonInfoExplicit(name, address, birthDate string) {
	fmt.Printf("Name: %s, Address: %s, Birth Date: %s\n", name, address, birthDate)
}

func printPersonInfoImplicit(info ...string) {
	if len(info) == 0 {
		return
	}
	if len(info) > 0 {
		fmt.Printf("Name: \"%s\", ", info[0])
	}
	if len(info) > 1 {
		fmt.Printf("Address: \"%s\", ", info[1])
	}
	if len(info) > 2 {
		fmt.Printf("Birth Date: \"%s\", ", info[2])
	}
	fmt.Print("\n")
}

func printPersonInfoExplicitAndImplicit(name, address string, birthDate ...string) {
	fmt.Printf("Name: %s, Address: %s", name, address)
	if len(birthDate) > 0 {
		fmt.Printf(", Birth Date: %s", birthDate[0])
	}
	fmt.Println()
}

func printPersonFullName(fullName ...string) {
	lastName := ""
	if len(fullName) > 0 {
		lastName = fullName[0]
	}
	firstName := ""
	if len(fullName) > 1 {
		firstName = fullName[1]
	}
	middleName := ""
	if len(fullName) > 2 {
		middleName = fullName[2]
	}

	if lastName == "" {
		lastName = "DefaultLastName"
	}
	if firstName == "" {
		firstName = "DefaultFirstName"
	}
	if middleName == "" {
		middleName = "DefaultMiddleName"
	}

	fmt.Printf("Full name: \"%s\" \"%s\" \"%s\"\n", lastName, firstName, middleName)
}

func minFunc(nums ...int) (int, error) {
	if len(nums) == 0 {
		return 0, argErr
	}

	minNum := nums[0]
	for _, num := range nums {
		if num < minNum {
			minNum = num
		}
	}

	return minNum, nil
}

func (r IntArr) average() (int, error) {
	if len(r) == 0 {
		return 0, argErr
	}

	sum := 0
	for _, num := range r {
		sum += num
	}

	return sum / len(r), nil
}

func (r IntMat) average() (int, error) {
	if len(r) == 0 || len(r[0]) == 0 {
		return 0, argErr
	}

	sum := 0
	count := 0
	for _, row := range r {
		count += len(row)
		for _, num := range row {
			sum += num
		}
	}

	return sum / count, nil
}

/*
	input: Function f,
	      endpoint values a, b,
	      tolerance TOL,
	      maximum iterations NMAX

conditions: a < b,

	either f(a) < 0 and f(b) > 0 or f(a) > 0 and f(b) < 0

output: value which differs from a root of f(x) = 0 by less than TOL

N ← 1
while N ≤ NMAX do // limit iterations to prevent infinite loop

	c ← (a + b)/2 // new midpoint
	if f(c) = 0 or (b – a)/2 < TOL then // solution found
	    Output(c)
	    Stop
	end if
	N ← N + 1 // increment step counter
	if sign(f(c)) = sign(f(a)) then a ← c else b ← c // new interval

end while
Output("Method failed.") // max number of steps exceeded
*/
func solveEquation(eqFunc func(float64) float64, a, b, precision float64, maxIters int) float64 {
	if !(a < b) {
		panic("Check failed: a < b")
	}
	if !((eqFunc(a) < 0 && eqFunc(b) > 0) || (eqFunc(a) > 0 && eqFunc(b) < 0)) {
		panic("Check failed: either f(a) < 0 and f(b) > 0 or f(a) > 0 and f(b) < 0")
	}

	for n := 0; n < maxIters; n++ {
		c := (a + b) / 2
		res := eqFunc(c)
		if (res == 0) || ((b-a)/2 < precision) {
			return c
		}
		if res*eqFunc(a) < 0 {
			b = c
		} else {
			a = c
		}
	}

	panic("Exceeded maximum number of iterations")
}

func equationFunc(x float64) float64 {
	return 3*math.Sin(math.Sqrt(x)) + 0.35*x - 3.8
}

/* Для 1, 2 та 3 стовчика із сформованого масиву структур (поля: ПІБ, адреса, дата народження) вона повинна наслідуватися структурою з варіанту (поля структури варіанту визначаються студентом) виконати завдання за варіантом */
/* Написати функції:
1) з параметрами заданими явно;
2) з параметрами заданими не явно;
3) частина параметрів задана явно, а частина не явно */
/* Продемонструйте різні способи виклику функції */
/* Для 4-го стовчика вирішити рівняння зазначеним у варіанті методом. Рівняння передати у функцію як параметр за допомогою покажчика, використати інтерфейс */
/* Функція з параметрами, що замовчуються: Друк прізвища, імені та по батькові */
/* Функція зі змінним числом параметрів: Мінімальний елемент у списку параметрів */
/* Перевизначити методи: Середнє арифметичне масиву */
/* Передача функції як параметра іншої функції за допомогою вказівника:
Метод ітерацій 3*sin(sqrt(x)) + 0.35*x - 3.8 = 0
Відрізок, що містить корінь: [2;3]
Точное значение: 2.2985 */

func main() {
	var err error
	var person Person

	scanner := bufio.NewScanner(os.Stdin)

	fmt.Printf("Person first name: ")
	if !scanner.Scan() && scanner.Err() != nil {
		panic(scanner.Err())
	}
	person.fullName[0] = scanner.Text()

	fmt.Printf("Person last name: ")
	if !scanner.Scan() && scanner.Err() != nil {
		panic(scanner.Err())
	}
	person.fullName[1] = scanner.Text()

	fmt.Printf("Person middle name: ")
	if !scanner.Scan() && scanner.Err() != nil {
		panic(scanner.Err())
	}
	person.fullName[2] = scanner.Text()

	fmt.Printf("Person address: ")
	if !scanner.Scan() && scanner.Err() != nil {
		panic(scanner.Err())
	}
	person.address = scanner.Text()

	fmt.Printf("Person birthday: ")
	if !scanner.Scan() && scanner.Err() != nil {
		panic(scanner.Err())
	}
	person.birthday = scanner.Text()

	fmt.Print("Person info (явно): ")
	printPersonInfoExplicit(person.fullName[0], person.address, person.birthday)

	fmt.Print("Person info (неявно): ")
	printPersonInfoImplicit(person.fullName[0], person.address, person.birthday)

	fmt.Print("Person info (явно/неявно): ")
	printPersonInfoExplicitAndImplicit(person.fullName[0], person.address)

	fmt.Print("Person info (явно/неявно): ")
	printPersonInfoExplicitAndImplicit(person.fullName[0], person.address, person.birthday)

	fmt.Print("Person name (default arguments): ")
	printPersonFullName(person.fullName[:]...)

	min_, _ := minFunc(54, 26, 87, 11, 99)
	fmt.Printf("Minimal number between (54, 26, 87, 11, 99): %d\n", min_)

	arr := make(IntArr, 10)
	for i := 0; i < len(arr); i++ {
		arr[i] = rand.Intn(128)
	}

	fmt.Printf("Array: %v\n", arr)

	avg, err := arr.average()
	if err == nil {
		fmt.Printf("Average of array is %d\n", avg)
	} else {
		fmt.Printf("Failed to calculate average of array: %s\n", err)
	}

	mat := make(IntMat, 3)
	for i := 0; i < len(mat); i++ {
		mat[i] = make(IntArr, len(mat))
		for j := 0; j < len(mat[i]); j++ {
			mat[i][j] = rand.Intn(128)
		}
	}

	fmt.Printf("Matrix: %v\n", mat)

	avg, err = mat.average()
	if err == nil {
		fmt.Printf("Average of matrix is %d\n", avg)
	} else {
		fmt.Printf("Failed to calculate average of matrix: %s\n", err)
	}

	fmt.Printf("Result of equation: %.4f", solveEquation(equationFunc, 2, 3, 0.0001, 25))
}
