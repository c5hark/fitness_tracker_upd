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

type DaySteps struct {
	Steps    int
	Duration time.Duration
	personaldata.Personal
}

func (ds *DaySteps) Parse(datastring string) (err error) {
	sliceData := strings.Split(datastring, ",")
	if len(sliceData) != 2 {
		return errors.New("invalid string format: must be 2 elements")
	}

	ds.Steps, err = strconv.Atoi(sliceData[0])
	if err != nil {
		return errors.New("invalid integer format: falied to convert the number of steps to an integer")
	}
	if ds.Steps <= 0 {
		return errors.New("invalid count of steps: steps must be greater than zero")
	}

	ds.Duration, err = time.ParseDuration(sliceData[1])
	if err != nil {
		return errors.New("invalid duration format: falied to parse duration as time.Duration")
	}
	if ds.Duration <= 0 {
		return errors.New("invalid duration format: time of time.Duration must be greater than zero")
	}
	return nil
}

func (ds DaySteps) ActionInfo() (string, error) {
	if ds.Steps <= 0 {
		return "", errors.New("invalid count of steps: steps must be greater than zero")
	}

	distance := spentenergy.Distance(ds.Steps, ds.Height)
	var calories float64
	var err error

	calories, err = spentenergy.WalkingSpentCalories(ds.Steps, ds.Weight, ds.Height, ds.Duration)
	if err != nil {
		return "", err
	}
	result := fmt.Sprintf(
		"Количество шагов: %d.\nДистанция составила %.2f км.\nВы сожгли %.2f ккал.\n",
		ds.Steps, distance, calories,
	)
	return result, nil
}
