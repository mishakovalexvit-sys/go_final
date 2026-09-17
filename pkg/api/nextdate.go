package api

import (
	"errors"
	"net/http"
	"strconv"
	"strings"
	"time"
)

const layout = "20060102"

func NextDate(now time.Time, dstart string, repeat string) (string, error) {
	date, err := time.Parse(layout, dstart)
	if err != nil {
		return "", err
	}
	rules := strings.Split(repeat, " ")
	switch {
	case rules[0] == "d" && len(rules) == 2:
		interval, err := strconv.Atoi(rules[1])
		if err != nil {
			return "", err
		}
		if interval > 400 || interval <= 0 {
			return "", errors.New("the interval should be from 1 to 400.")
		}
		for {
			date = date.AddDate(0, 0, interval)
			if date.After(now) {
				break
			}
		}
	case rules[0] == "y" && len(rules) == 1:
		for {
			date = date.AddDate(1, 0, 0)
			if date.After(now) {
				break
			}
		}
	default:
		return "", errors.New("incorrect format")
	}

	return date.Format(layout), nil
}

func nextDayHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		http.Error(w, "Метод не поддерживается", http.StatusMethodNotAllowed)
		return
	}
	nowStr := r.FormValue("now")
	dateStr := r.FormValue("date")
	repeatStr := r.FormValue("repeat")
	var now time.Time
	var err error
	var nextDate string
	if nowStr == "" {
		now = time.Now()
	} else {
		now, err = time.Parse(layout, nowStr)
		if err != nil {
			http.Error(w, "Invalid date format now", http.StatusBadRequest)
			return
		}
	}
	nextDate, err = NextDate(now, dateStr, repeatStr)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	w.Write([]byte(nextDate))
}
