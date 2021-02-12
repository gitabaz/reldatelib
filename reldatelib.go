package datelib

import "time"
import "strings"
import "errors"
import "regexp"
import "math"
import "strconv"

// TODO: Rewrite this to avoid having to duplicate the keys
var relTimeOrdKeys = []string{"LAST", "THIS", "NEXT", "FIRST", "THIRD", "FOURTH", "FIFTH", "SIXTH", "SEVENTH", "EIGHTH", "NINTH", "TENTH", "ELEVENTH", "TWELFTH"}
var relTimeOrdinal = map[string]int{
	"LAST":  -1,
	"THIS":  0,
	"NEXT":  1,
	"FIRST": 1,
	"THIRD":    3,
	"FOURTH":   4,
	"FIFTH":    5,
	"SIXTH":    6,
	"SEVENTH":  7,
	"EIGHTH":   8,
	"NINTH":    9,
	"TENTH":    10,
	"ELEVENTH": 11,
	"TWELFTH":  12,
}

var relTimeUnit = []string{"SEC", "SECOND", "MIN", "MINUTE", "HOUR", "DAY", "WEEK", "FORTNIGHT", "MONTH", "YEAR"}

// TODO: Rewrite this to avoid having to duplicate the keys
var relTimeAgoKeys = []string{"AGO", "HENCE"}
var relTimeAgo = map[string]int{
    "AGO": -1,
    "HENCE": 1,
}

var relTimeShift = []string{"NOW", "TODAY", "YESTERDAY", "TOMORROW"}

var weekdayKeys = []string{"MONDAY", "TUESDAY", "WEDNESDAY", "THURSDAY", "FRIDAY", "SATURDAY", "SUNDAY"}
var weekdays = map[string]int{
    "MONDAY": 0,
    "TUESDAY": 1,
    "WEDNESDAY": 2,
    "THURSDAY": 3,
    "FRIDAY": 4,
    "SATURDAY": 5,
    "SUNDAY": 6}

var regexStr = "(?P<num>-?\\d+|" + strings.Join(relTimeOrdKeys, "|") +
               ")?\\s?(?P<unit>" + strings.Join(relTimeUnit, "|") + "|" + strings.Join(weekdayKeys, "|") +
               ")S?\\b\\s?(?P<ago>" + strings.Join(relTimeAgoKeys, "|") + ")?|^(" + strings.Join(relTimeShift, "|") + ")$"
var r, _ = regexp.Compile(regexStr)

func Parse(str string) (string, error) {
	str = strings.ToUpper(str)
	format := "Mon 02 Jan 2006 03:04:05 PM MST"

    res := r.FindAllStringSubmatch(str, -1)

    var relativeTime = make(map[string]int)
    t := time.Now().UTC()
    if len(res) == 0 {
        return "", errors.New("Could not parse string")
    }
    for _, m := range res {
        if m[0] == m[4] {
            relativeTime[m[4]] = 1
        } else {
            numUnits := 1
            var err error
            if m[1] != "" {
                numUnits, err = strconv.Atoi(m[1])
                if err != nil {
                    var ok = false
                    numUnits, ok = relTimeOrdinal[m[1]]
                    if ok == false {
                        return "", errors.New("Could not parse string")
                    }
                }
            }
            prefactor, ok := relTimeAgo[m[3]]
            if (ok == false) {
                prefactor = 1
            }
            wd, ok := weekdays[m[2]]
            if ok == true {
                relWeekday := weekdays[strings.ToUpper(time.Now().UTC().Weekday().String())] - wd
                offset := 0
                if (numUnits * relWeekday >= 0) {
                    offset = prefactor * numUnits * 7 - relWeekday
                } else {
                    offset = prefactor * (int(math.Abs(float64(numUnits))) - 1) * 7 - relWeekday
                }

                return time.Date(t.Year(), t.Month(), t.Day() + offset, 0, 0, 0, 0, time.Local).Format(format), nil
            } else {
                unit := m[2]
                relativeTime[unit] += prefactor * numUnits
            }
        }
    }
    return t.
    AddDate(relativeTime["YEAR"], relativeTime["MONTH"], relativeTime["TOMORROW"] -1 * relativeTime["YESTERDAY"] + relativeTime["DAY"] + relativeTime["WEEK"] * 7 + relativeTime["FORTNIGHT"] * 14).
    Add(time.Duration(relativeTime["HOUR"]) * time.Hour + time.Duration(relativeTime["MINUTE"] + relativeTime["MIN"]) * time.Minute + time.Duration(relativeTime["SECOND"] + relativeTime["SEC"]) * time.Second).Local().Format(format), nil
}
