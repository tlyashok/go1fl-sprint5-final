package spentenergy

import (
	"errors"
	"time"
)

// Основные константы, необходимые для расчетов.
const (
	mInKm                      = 1000 // количество метров в километре.
	minInH                     = 60   // количество минут в часе.
	stepLengthCoefficient      = 0.45 // коэффициент для расчета длины шага на основе роста.
	walkingCaloriesCoefficient = 0.5  // коэффициент для расчета калорий при ходьбе.
)

// Ошибки проверки входных параметров для расчета калорий.
var (
	ErrInvalidSteps    = errors.New("количество шагов должно быть больше нуля")
	ErrInvalidWeight   = errors.New("вес должен быть больше нуля")
	ErrInvalidHeight   = errors.New("рост должен быть больше нуля")
	ErrInvalidDuration = errors.New("продолжительность должна быть больше нуля")
)

// WalkingSpentCalories возвращает количество калорий, потраченных при ходьбе.
func WalkingSpentCalories(steps int, weight, height float64, duration time.Duration) (float64, error) {
	if err := validate(steps, weight, height, duration); err != nil {
		return 0, err
	}

	meanSpeed := MeanSpeed(steps, height, duration)
	durationInMinutes := duration.Minutes()
	calories := (weight * meanSpeed * durationInMinutes) / minInH

	return calories * walkingCaloriesCoefficient, nil
}

// RunningSpentCalories возвращает количество калорий, потраченных при беге.
func RunningSpentCalories(steps int, weight, height float64, duration time.Duration) (float64, error) {
	if err := validate(steps, weight, height, duration); err != nil {
		return 0, err
	}

	meanSpeed := MeanSpeed(steps, height, duration)
	durationInMinutes := duration.Minutes()

	return (weight * meanSpeed * durationInMinutes) / minInH, nil
}

// MeanSpeed возвращает среднюю скорость в км/ч.
// При неположительной продолжительности или количестве шагов возвращает 0,
// что также защищает от деления на ноль.
func MeanSpeed(steps int, height float64, duration time.Duration) float64 {
	if duration <= 0 || steps <= 0 {
		return 0
	}

	distance := Distance(steps, height)

	return distance / duration.Hours()
}

// Distance возвращает дистанцию в километрах.
func Distance(steps int, height float64) float64 {
	stepLength := height * stepLengthCoefficient

	return float64(steps) * stepLength / mInKm
}

// validate проверяет входные параметры для расчета калорий.
func validate(steps int, weight, height float64, duration time.Duration) error {
	switch {
	case steps <= 0:
		return ErrInvalidSteps
	case weight <= 0:
		return ErrInvalidWeight
	case height <= 0:
		return ErrInvalidHeight
	case duration <= 0:
		return ErrInvalidDuration
	}

	return nil
}
