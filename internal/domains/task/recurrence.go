package taskdmn

import (
	"errors"
	"strconv"
	"strings"
	"time"
)

const (
	DayKeyLetter   string = "d"
	WeekKeyLetter  string = "w"
	MonthKeyLetter string = "m"
	YearKeyLetter  string = "y"
)

type Recurrence string

func NewRecurrence(pattern string) (Recurrence, error) {
	pattern = strings.TrimSpace(pattern)
	if pattern == "" {
		return "", nil
	}

	parts := strings.Fields(pattern)
	if len(parts) == 0 {
		return "", nil
	}

	validateRules := map[string]func([]string) error{
		DayKeyLetter:   validateDailyRule,
		WeekKeyLetter:  validateWeeklyRule,
		MonthKeyLetter: validateMonthlyRule,
		YearKeyLetter:  validateYearlyRule,
	}

	validator, ok := validateRules[parts[0]]
	if !ok {
		return "", errors.New("unknown rule type")
	}

	if err := validator(parts); err != nil {
		return "", err
	}

	return Recurrence(pattern), nil
}

func validateDailyRule(tokens []string) error {
	if len(tokens) != 2 {
		return errors.New("daily rule: expected 1 argument")
	}
	n, err := strconv.ParseInt(tokens[1], 10, 64)
	if err != nil {
		return errors.New("daily rule: invalid number")
	}
	if n <= 0 || n > 400 {
		return errors.New("daily rule: value must be 1–400")
	}
	return nil
}

func validateWeeklyRule(tokens []string) error {
	if len(tokens) != 2 {
		return errors.New("weekly rule: expected 1 argument")
	}
	return validateDayList(tokens[1], 1, 7)
}

func validateMonthlyRule(tokens []string) error {
	if len(tokens) < 2 || len(tokens) > 3 {
		return errors.New("monthly rule: expected 1 or 2 arguments")
	}
	if err := validateDayList(tokens[1], -2, 31); err != nil {
		return errors.New("monthly rule: invalid days")
	}
	if len(tokens) == 3 {
		if err := validateDayList(tokens[2], 1, 12); err != nil {
			return errors.New("monthly rule: invalid months")
		}
	}
	return nil
}

func validateYearlyRule(tokens []string) error {
	if len(tokens) != 1 {
		return errors.New("yearly rule: no arguments expected")
	}
	return nil
}

func validateDayList(s string, min, max int) error {
	parts := strings.Split(s, ",")
	if len(parts) == 0 {
		return errors.New("empty list")
	}
	for _, p := range parts {
		v, err := strconv.Atoi(strings.TrimSpace(p))
		if err != nil || v < min || v > max {
			return errors.New("value out of range")
		}
	}
	return nil
}

func (r Recurrence) HasNextDate() bool {
	return string(r) == ""
}

func (r Recurrence) NextDate(now, startDate time.Time) time.Time {
	if r.HasNextDate() {
		return time.Time{}
	}

	parts := strings.Fields(string(r))
	if len(parts) == 0 {
		return time.Time{}
	}

	base := now
	if base.Before(startDate) {
		base = startDate
	}

	calcRules := map[string]func(time.Time, time.Time, []string) time.Time{
		DayKeyLetter:   nextDaily,
		WeekKeyLetter:  nextWeekly,
		MonthKeyLetter: nextMonthly,
		YearKeyLetter:  nextYearly,
	}

	calc, ok := calcRules[parts[0]]
	if !ok {
		return time.Time{}
	}

	return calc(base, startDate, parts)
}

func nextDaily(now, startDate time.Time, tokens []string) time.Time {
	n, _ := strconv.Atoi(tokens[1])
	return now.AddDate(0, 0, n)
}

func nextYearly(now, startDate time.Time, tokens []string) time.Time {
	return now.AddDate(1, 0, 0)
}

func nextWeekly(now, startDate time.Time, tokens []string) time.Time {
	days := parseIds(tokens[1])
	for i := 1; i <= 7; i++ {
		next := now.AddDate(0, 0, i)
		w := int(next.Weekday())
		if w == 0 {
			w = 7
		}
		for _, d := range days {
			if w == d && !next.Before(startDate) {
				return next
			}
		}
	}
	return time.Time{}
}

func nextMonthly(now, startDate time.Time, tokens []string) time.Time {
	dayIds := parseIds(tokens[1])
	monthIds := []int{}
	if len(tokens) > 2 {
		monthIds = parseIds(tokens[2])
	}

	for i := 0; i < 12; i++ {
		y, m, _ := now.AddDate(0, i, 1).Date()

		if len(monthIds) > 0 {
			found := false
			for _, mi := range monthIds {
				if int(m) == mi {
					found = true
					break
				}
			}
			if !found {
				continue
			}
		}

		lastDay := time.Date(y, m+1, 0, 0, 0, 0, 0, now.Location()).Day()
		for _, di := range dayIds {
			day := di
			if di == -1 {
				day = lastDay
			} else if di == -2 {
				day = lastDay - 1
			}
			if day < 1 || day > lastDay {
				continue
			}
			candidate := time.Date(y, m, day, now.Hour(), now.Minute(), now.Second(), now.Nanosecond(), now.Location())
			if !candidate.Before(now) && !candidate.Before(startDate) {
				return candidate
			}
		}
	}
	return time.Time{}
}

func parseIds(s string) []int {
	parts := strings.Split(s, ",")
	res := make([]int, 0, len(parts))
	for _, p := range parts {
		if v, err := strconv.Atoi(strings.TrimSpace(p)); err == nil {
			res = append(res, v)
		}
	}
	return res
}
