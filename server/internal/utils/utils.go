package utils

import (
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"strconv"
	"time"

	"github.com/google/uuid"
)

func Message(status bool, message string) map[string]interface{} {
	return map[string]interface{}{"status": status, "message": message}
}

func Respond(w http.ResponseWriter, data map[string]interface{}) {
	w.Header().Add("content-type", "application/json")
	json.NewEncoder(w).Encode(data)
}

func GenerateUUID() uuid.UUID {
	uuid := uuid.New()
	return uuid
}

func ParseToUUID(s string) uuid.UUID {
	uuid, _ := uuid.Parse(s)
	return uuid
}

// Method that calculate the hours between dates
func BetweenDateToHour(t1 time.Time, t2 time.Time) (float64, error) {
	var result float64

	// data of time1
	hour1 := float64(t1.Hour())
	minute1 := float64(t1.Minute())
	second1 := float64(t1.Second())

	// data of time2
	hour2 := float64(t2.Hour())
	minute2 := float64(t2.Minute())
	second2 := float64(t2.Second())

	if t1.Before(t2) {
		day := float64(t2.YearDay() - t1.YearDay())
		hour := (hour2 - hour1)
		minute := (minute2 - minute1)
		second := (second2 - second1)

		result = (day * 24) + hour + (minute / 60) + (second / 360)
		return result, nil
	} else {
		result = 0
		return result, errors.New("t1 must be smaller than t2")
	}
}

func CustomFloat(f float64) (float64, error) {
	s := fmt.Sprintf("%.2f", f)
	value, err := strconv.ParseFloat(s, 64)
	if err != nil {
		return float64(0), err
	}
	return value, nil
}

// Average value on a list
func ListAverageValue(values []float64) float64 {
	total := float64(0)
	for _, value := range values {
		total += value
	}

	//result := total / float64(len(values))

	return total
	//return result
}
