package stats

import "github.com/Behzod01/bank/pkg/types"

//Avg рассчитывает среднюю сумму платежа.
func Avg(payments []types.Payment) types.Money {
	var middle, count types.Money
	for _, v := range payments {
		middle += v.Amount
		count++
	}
	middle /= count
	return middle
}

//TotalInCategory находит сумму покупок в определённой категории.
func TotalInCategory(payments []types.Payment, category types.Category) types.Money {
	var sum types.Money
	for _, v := range payments {
		if v.Category == category{
			sum += v.Amount
		}		
	}
	return sum
}
