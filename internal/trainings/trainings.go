package trainings

import (
	"errors"
	"fmt"
	"strconv"
	"strings"
	"time"

	"github.com/Yandex-Practicum/tracker/internal/personaldata"
	"github.com/Yandex-Practicum/tracker/internal/spentenergy"
)

// Поддерживаемые типы тренировок.
const (
	trainingTypeRunning = "Бег"
	trainingTypeWalking = "Ходьба"
)

// Ошибки разбора данных и формирования информации о тренировке.
var (
	ErrInvalidFormat       = errors.New("неверный формат данных")
	ErrInvalidSteps        = errors.New("количество шагов должно быть больше нуля")
	ErrInvalidDuration     = errors.New("продолжительность должна быть больше нуля")
	ErrUnknownTrainingType = errors.New("неизвестный тип тренировки")
)

// Training описывает тренировку пользователя.
type Training struct {
	Steps        int           // количество шагов, проделанных за тренировку.
	TrainingType string        // тип тренировки (бег или ходьба).
	Duration     time.Duration // длительность тренировки.
	personaldata.Personal
}

// Parse разбирает строку формата "3456,Ходьба,3h00m" и записывает данные в структуру.
// Ресивер — указатель, чтобы изменения полей сохранялись в исходной структуре.
// Поля заполняются только после успешной проверки всех значений.
func (t *Training) Parse(datastring string) (err error) {
	parts := strings.Split(datastring, ",")
	if len(parts) != 3 {
		return fmt.Errorf("%w: ожидалось 3 значения, получено %d", ErrInvalidFormat, len(parts))
	}

	steps, err := strconv.Atoi(parts[0])
	if err != nil {
		return fmt.Errorf("ошибка преобразования количества шагов %q: %w", parts[0], err)
	}
	if steps <= 0 {
		return fmt.Errorf("%w: %d", ErrInvalidSteps, steps)
	}

	duration, err := time.ParseDuration(parts[2])
	if err != nil {
		return fmt.Errorf("ошибка преобразования продолжительности %q: %w", parts[2], err)
	}
	if duration <= 0 {
		return fmt.Errorf("%w: %v", ErrInvalidDuration, duration)
	}

	t.Steps = steps
	t.TrainingType = parts[1]
	t.Duration = duration

	return nil
}

// ActionInfo формирует строку с информацией о тренировке.
func (t Training) ActionInfo() (string, error) {
	distance := spentenergy.Distance(t.Steps, t.Height)
	meanSpeed := spentenergy.MeanSpeed(t.Steps, t.Height, t.Duration)

	var calories float64
	var err error

	switch t.TrainingType {
	case trainingTypeRunning:
		calories, err = spentenergy.RunningSpentCalories(t.Steps, t.Weight, t.Height, t.Duration)
	case trainingTypeWalking:
		calories, err = spentenergy.WalkingSpentCalories(t.Steps, t.Weight, t.Height, t.Duration)
	default:
		return "", fmt.Errorf("%w: %q", ErrUnknownTrainingType, t.TrainingType)
	}
	if err != nil {
		return "", fmt.Errorf("ошибка расчета калорий: %w", err)
	}

	return fmt.Sprintf(
		"Тип тренировки: %s\nДлительность: %.2f ч.\nДистанция: %.2f км.\nСкорость: %.2f км/ч\nСожгли калорий: %.2f\n",
		t.TrainingType, t.Duration.Hours(), distance, meanSpeed, calories,
	), nil
}
