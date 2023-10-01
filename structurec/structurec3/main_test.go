//	на интерфейсы: Есть структура которая представляет магазин - Shop, у нее есть метод AddModel который на вход принимает

//

// интерфейс Transport и добавляет машину на склад в магазин, склад это мапа или слайс.

// В магазин можно добавить машину, мотоцикл, мопед и должна быть возможность в будущем добавить любой транспорт.

// Это разные структуры у которых есть поле кол-во колес и мощность.

// В структуре Shop сделать метод который возвращает транспорт у которого колес больше чем n, n передается на вход функции

// и метод который возвращает транспорт с самой большой мощностью.

// И еще тест написать на две функции

package main

import (
	"testing"


)

func TestAvto_GetWheel(t *testing.T) {
	 
	
	tests := []struct {
		name string
		Transport map[string]Transport
		want int
	}{
		{
			name: "sucsess",
			Transport: map[string]Transport{
				"one": mock,
				"second" : &Scooter{
					Wheel: 2,
				},
			},
			want: 3,
		},// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.a.GetWheel(); got != tt.want {
				t.Errorf("Avto.GetWheel() = %v, want %v", got, tt.want)
			}
		})
	}
}
