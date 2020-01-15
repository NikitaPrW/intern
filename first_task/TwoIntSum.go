package sum

// TwoIntSum Фунция получает на вход слайс интов и цель, для которой нужно найти два числа из слайса nums, что бы их сумма была равна target
// Функция возвращает слайс интов, в котором находятся индексы элементов, удовлетворяющих условию.
func TwoIntSum(nums []int, target int) []int {
	var solution []int
	for i, value := range nums {
		for n := i + 1; n < len(nums); n++ {
			if value+nums[n] == target { //проверка, что сумма двух элементов равна target
				solution = append(solution, i, n) // добавление в слайс индексов двух элементов
				return solution                   // возвращается слайс из двух индексов, так как по условию задачи верное решение только одно
			}
		}
	}
	return solution
}
