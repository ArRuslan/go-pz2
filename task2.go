package main

import (
	"errors"
	"fmt"
	"math"
)

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
		fmt.Printf("Name: %s", info[0])
	}
	if len(info) > 1 {
		fmt.Printf("Address: %s", info[1])
	}
	if len(info) > 2 {
		fmt.Printf("Birth Date: %s", info[2])
	}
	fmt.Print("\n")
}

// TODO: частина параметрів задана явно, а частина не явно

func printPersonFullName(p Person) {
	lastName := p.fullName[0]
	firstName := p.fullName[1]
	middleName := p.fullName[2]

	if lastName == "" {
		lastName = "DefaultLastName"
	}
	if firstName == "" {
		firstName = "DefaultFirstName"
	}
	if middleName == "" {
		middleName = "DefaultMiddleName"
	}

	fmt.Printf("Full name: %s %s %s\n", lastName, firstName, middleName)
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

type IAverage interface {
	average() (int, error)
}

func (r IntArr1d) average() (int, error) {
	if len(r) == 0 {
		return 0, argErr
	}

	sum := 0
	for _, num := range r {
		sum += num
	}

	return sum / len(r), nil
}

func (e Person) average() (int, error) {
	return -1, nil // TODO
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
	fmt.Printf("Result: %.4f", solveEquation(equationFunc, 2, 3, 0.0001, 25))
}
