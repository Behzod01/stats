package stats

import (
	"reflect"
	"testing"

	"github.com/Behzod01/bank/v2/pkg/types"
)

func TestCategoriesAvg(t *testing.T) {
	payments := []types.Payment{
		{Amount: 1_000, Category: "Ресторан"},
		{Amount: 250_000, Category: "Машина"},
		{Amount: 1_500, Category: "Телефон"},
		{Amount: 500_000, Category: "Машина"},
		{Amount: 450, Category: "Ресторан"},
		{Amount: 5_000, Category: "Телефон"},
	}
	expected := map[types.Category]types.Money{
		"Ресторан": 725,
		"Машина": 375_000,
		"Телефон": 3_250,
	}
	result := CategoriesAvg(payments)

	if !reflect.DeepEqual(expected, result){
		t.Errorf("invalid result, expected: %v, actual: %v", expected, result)
	}
}