package daysteps

import (
	"errors"
	"fmt"
	"strconv"
	"strings"
	"time"

	"github.com/Yandex-Practicum/tracker/internal/personaldata"
	"github.com/Yandex-Practicum/tracker/internal/spentenergy"
)

var (
	ErrInvalidFormat   = errors.New("неверный формат данных")
	ErrInvalidSteps    = errors.New("количество шагов должно быть больше нуля")
	ErrInvalidDuration = errors.New("продолжительность должна быть больше нуля")
)

// DaySteps описывает дневную активность (прогулку) пользователя.
type DaySteps struct {
	Steps    int           // количество шагов.
	Duration time.Duration // длительность прогулки.
	personaldata.Personal
}

// Parse разбирает строку формата "678,0h50m" и записывает данные в структуру.
func (ds *DaySteps) Parse(datastring string) (err error) {
	parts := strings.Split(datastring, ",")
	if len(parts) != 2 {
		return fmt.Errorf("%w: ожидалось 2 значения, получено %d", ErrInvalidFormat, len(parts))
	}

	steps, err := strconv.Atoi(parts[0])
	if err != nil {
		return fmt.Errorf("ошибка преобразования количества шагов %q: %w", parts[0], err)
	}
	if steps <= 0 {
		return fmt.Errorf("%w: %d", ErrInvalidSteps, steps)
	}

	duration, err := time.ParseDuration(parts[1])
	if err != nil {
		return fmt.Errorf("ошибка преобразования продолжительности %q: %w", parts[1], err)
	}
	if duration <= 0 {
		return fmt.Errorf("%w: %v", ErrInvalidDuration, duration)
	}

	ds.Steps = steps
	ds.Duration = duration

	return nil
}

// ActionInfo формирует строку с информацией о прогулке.
func (ds DaySteps) ActionInfo() (string, error) {
	distance := spentenergy.Distance(ds.Steps, ds.Height)

	calories, err := spentenergy.WalkingSpentCalories(ds.Steps, ds.Weight, ds.Height, ds.Duration)
	if err != nil {
		return "", fmt.Errorf("ошибка расчета калорий: %w", err)
	}

	return fmt.Sprintf(
		"Количество шагов: %d.\nДистанция составила %.2f км.\nВы сожгли %.2f ккал.\n",
		ds.Steps, distance, calories,
	), nil
}
