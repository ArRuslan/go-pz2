package main

import (
	"fmt"
	"math"
)

type Person struct {
	fullName string
	address  string
	birthday string
}

func printPersonInfo(name string, address string, birthDate string) {
	fmt.Printf("Name: %s, Address: %s, Birth Date: %s\n", name, address, birthDate)
}

func (s Person) PrintInfo() {
	fmt.Printf("Name: %s, Address: %s, Birth Date: %s\n", s.fullName, s.address, s.birthday)
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
func solveTs(eqFunc func(float64) float64, a float64, b float64, precision float64, maxIters int) float64 {
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
	fmt.Printf("Result: %.4f", solveTs(equationFunc, 2, 3, 0.0001, 25))
}
