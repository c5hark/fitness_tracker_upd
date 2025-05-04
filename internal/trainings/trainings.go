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

type Training struct {
	Steps        int
	TrainingType string
	Duration     time.Duration
	personaldata.Personal
}

func (t *Training) Parse(datastring string) (err error) {
	sliceData := strings.Split(datastring, ",")
	if len(sliceData) != 3 {
		return errors.New("invalid string format: must be 3 elements")
	}
	t.Steps, err = strconv.Atoi(sliceData[0])
	if err != nil {
		return errors.New("invalid integer format: falied to convert the number of steps to an integer")
	}
	if t.Steps <= 0 {
		return errors.New("invalid count of steps: steps must be greater than zero")
	}

	t.TrainingType = sliceData[1]

	t.Duration, err = time.ParseDuration(sliceData[2])
	if err != nil {
		return errors.New("invalid duration format: falied to parse duration as time.Duration")
	}
	if t.Duration <= 0 {
		return errors.New("invalid duration format: time of time.Duration must be greater than zero")
	}
	return nil
}

func (t Training) ActionInfo() (string, error) {
	distance := spentenergy.Distance(t.Steps, t.Height)
	meanSpeed := spentenergy.MeanSpeed(t.Steps, t.Height, t.Duration)

	var err error
	var calories float64
	var trainingName string

	switch t.TrainingType {
	case "Ходьба":
		trainingName = "Ходьба"
		calories, err = spentenergy.WalkingSpentCalories(t.Steps, t.Weight, t.Height, t.Duration)
		if err != nil {
			return "", err
		}
	case "Бег":
		trainingName = "Бег"
		calories, err = spentenergy.RunningSpentCalories(t.Steps, t.Weight, t.Height, t.Duration)
		if err != nil {
			return "", err
		}
	default:
		return "", errors.New("неизвестный тип тренировки: " + t.TrainingType)
	}
	result := fmt.Sprintf(
		"Тип тренировки: %s\nДлительность: %.2f ч.\nДистанция: %.2f км.\nСкорость: %.2f км/ч\nСожгли калорий: %.2f\n",
		trainingName, t.Duration.Hours(), distance, meanSpeed, calories,
	)
	return result, nil
}
