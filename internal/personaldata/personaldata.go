package personaldata

import "fmt"

// Personal хранит персональные данные пользователя.
type Personal struct {
	Name   string  // имя пользователя.
	Weight float64 // вес пользователя в килограммах.
	Height float64 // рост пользователя в метрах.
}

// Print выводит персональные данные пользователя на экран.
func (p Personal) Print() {
	fmt.Printf("Имя: %s\n", p.Name)
	fmt.Printf("Вес: %.2f\n", p.Weight)
	fmt.Printf("Рост: %.2f\n", p.Height)
}
