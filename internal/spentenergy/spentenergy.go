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

func WalkingSpentCalories(steps int, weight, height float64, duration time.Duration) (float64, error) {
	if steps <= 0 || weight <= 0 || height <= 0 || duration <= 0 {
		return 0, errors.New("invalid input: steps, weight, height and duration must be greater than 0")
	}
	meanSpeed := MeanSpeed(steps, height, duration)
	durationInMinutes := duration.Minutes()
	calories := (weight * meanSpeed * float64(durationInMinutes)) / float64(minInH)
	caloriesWalking := calories * walkingCaloriesCoefficient
	return caloriesWalking, nil
}

func RunningSpentCalories(steps int, weight, height float64, duration time.Duration) (float64, error) {
	if steps <= 0 {
		return 0, errors.New("invalid input: steps must be greater than 0")
	}
	if weight <= 0 {
		return 0, errors.New("invalid input: weight must be greater than 0")
	}
	if height <= 0 {
		return 0, errors.New("invalid input: height must be greater than 0")
	}
	if duration <= 0 {
		return 0, errors.New("invalid input: dauration must be greater than 0")
	}

	meanSpeed := MeanSpeed(steps, height, duration)
	durationInMinutes := duration.Minutes()
	calories := (weight * meanSpeed * float64(durationInMinutes)) / minInH
	return calories, nil
}

func MeanSpeed(steps int, height float64, duration time.Duration) float64 {
	if duration <= 0 {
		return 0
	}
	distance := Distance(steps, height)
	meanSpeed := distance / duration.Hours()
	return meanSpeed
}

func Distance(steps int, height float64) float64 {
	stepLength := height * stepLengthCoefficient
	distance := (stepLength * float64(steps)) / float64(mInKm)
	return distance
}
