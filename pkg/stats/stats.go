package stats

import "github.com/Behzod01/bank/v2/pkg/types"

//Avg рассчитывает среднюю сумму платежа.
func Avg(payments []types.Payment) types.Money {
	var middle, count types.Money
	for _, v := range payments {
		if v.Status != types.StatusFail{
		middle += v.Amount
		count++
		}
	}
	middle /= count
	return middle
}

//TotalInCategory находит сумму покупок в определённой категории.
func TotalInCategory(payments []types.Payment, category types.Category) types.Money {
	var sum types.Money
	for _, v := range payments {
		if v.Category == category && v.Status != types.StatusFail{
			sum += v.Amount
		}		
	}
	return sum
}

//CategoriesAvg считает среднюю сумму платежа по каждой категории.
func CategoriesAvg(payments []types.Payment) map[types.Category]types.Money {
	categories := map[types.Category]types.Money{}
	count := map[types.Category]types.Money{}
	for _, payment := range payments {
		categories[payment.Category] += payment.Amount
		count[payment.Category]++		
	}

	for key := range categories {
		categories[key] /= count[key]		
	}

	return categories
}

//PeriodsDynamic сравнивает расходы по котегориям за два периода.
func PeriodsDynamic(first map[types.Category]types.Money, 
	second map[types.Category]types.Money) map[types.Category]types.Money  {
		result := map[types.Category]types.Money{}

		for key := range second {
			result[key] += second[key]
		}
		for key := range first {
			result[key] -= first[key]			
		}
		return result

}