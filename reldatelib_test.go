package datelib

import "bytes"
import "testing"
import "os/exec"
import "strings"

func runCmd(param string) (string) {
    cmd := exec.Command("date", "-d", param)
    var out bytes.Buffer
    cmd.Stdout = &out
    cmd.Run()
    return strings.TrimSuffix(out.String(), "\n")
}

func cmpRes(t *testing.T, act string, exp string) {
    if act != exp {
        t.Errorf("output was incorrect, got: %q, want %q.", act, exp)
    }
}

func TestNow(t *testing.T) {
    param := "now"
    exp := runCmd(param)
    act, _ := Parse(param)
    cmpRes(t, act, exp)
}

func TestToday(t *testing.T) {
    param := "today"
    exp := runCmd(param)
    act, _ := Parse(param)
    cmpRes(t, act, exp)
}

func TestYesterday(t *testing.T) {
    param := "yesterday"
    exp := runCmd(param)
    act, _ := Parse(param)
    cmpRes(t, act, exp)
}

func TestTomorrow(t *testing.T) {
    param := "tomorrow"
    exp := runCmd(param)
    act, _ := Parse(param)
    cmpRes(t, act, exp)
}

func TestWeek(t *testing.T) {
    param := "week"
    exp := runCmd(param)
    act, _ := Parse(param)
    cmpRes(t, act, exp)
}

func TestCombo(t *testing.T) {
    param := "20 DAYS 4 HOURS AGO 3 WEEKS"
    exp := runCmd(param)
    act, _ := Parse(param)
    cmpRes(t, act, exp)
}

func TestNDaysAgo(t *testing.T) {
    param := "2 days ago"
    exp := runCmd(param)
    act, _ := Parse(param)
    cmpRes(t, act, exp)
}

func TestNWeeksAgo(t *testing.T) {
    param := "20 weeks ago"
    exp := runCmd(param)
    act, _ := Parse(param)
    cmpRes(t, act, exp)
}

func TestNHoursAgo(t *testing.T) {
    param := "2 hours ago"
    exp := runCmd(param)
    act, _ := Parse(param)
    cmpRes(t, act, exp)
}

func TestNSecsAgo(t *testing.T) {
    param := "2 secs ago"
    exp := runCmd(param)
    act, _ := Parse(param)
    cmpRes(t, act, exp)
}

func TestNSecondsAgo(t *testing.T) {
    param := "20 seconds ago"
    exp := runCmd(param)
    act, _ := Parse(param)
    cmpRes(t, act, exp)
}

func TestNMinAgo(t *testing.T) {
    param := "15 min ago"
    exp := runCmd(param)
    act, _ := Parse(param)
    cmpRes(t, act, exp)
}

func TestNMinutesAgo(t *testing.T) {
    param := "15 minutes ago"
    exp := runCmd(param)
    act, _ := Parse(param)
    cmpRes(t, act, exp)
}

func TestNFortnightsAgo(t *testing.T) {
    param := "1 fortnight ago"
    exp := runCmd(param)
    act, _ := Parse(param)
    cmpRes(t, act, exp)
}

func TestNextHour(t *testing.T) {
    param := "next hour"
    exp := runCmd(param)
    act, _ := Parse(param)
    cmpRes(t, act, exp)
}

func TestNextHourAgo(t *testing.T) {
    param := "next hour ago"
    exp := runCmd(param)
    act, _ := Parse(param)
    cmpRes(t, act, exp)
}

func TestNextFriday(t *testing.T) {
    param := "next friday"
    exp := runCmd(param)
    act, _ := Parse(param)
    cmpRes(t, act, exp)
}

func TestNextThursday(t *testing.T) {
    param := "next thursday"
    exp := runCmd(param)
    act, _ := Parse(param)
    cmpRes(t, act, exp)
}

func TestNextSaturday(t *testing.T) {
    param := "next saturday"
    exp := runCmd(param)
    act, _ := Parse(param)
    cmpRes(t, act, exp)
}

func TestNextSunday(t *testing.T) {
    param := "next sunday"
    exp := runCmd(param)
    act, _ := Parse(param)
    cmpRes(t, act, exp)
}

func TestFirstFriday(t *testing.T) {
    param := "first friday"
    exp := runCmd(param)
    act, _ := Parse(param)
    cmpRes(t, act, exp)
}

func TestThirdFriday(t *testing.T) {
    param := "third friday"
    exp := runCmd(param)
    act, _ := Parse(param)
    cmpRes(t, act, exp)
}

func TestLastFriday(t *testing.T) {
    param := "last friday"
    exp := runCmd(param)
    act, _ := Parse(param)
    cmpRes(t, act, exp)
}

func TestLastThursday(t *testing.T) {
    param := "last thursday"
    exp := runCmd(param)
    act, _ := Parse(param)
    cmpRes(t, act, exp)
}

func TestLastSaturday(t *testing.T) {
    param := "last saturday"
    exp := runCmd(param)
    act, _ := Parse(param)
    cmpRes(t, act, exp)
}

func TestFail(t *testing.T) {
    param := "potato"
    cmd := exec.Command("date", "-d", param)
    var out bytes.Buffer
    cmd.Stdout = &out
    cmd.Run()

    res, err := Parse(param)
    if err == nil {
        t.Errorf("expected error but got success")
    }
    if res != strings.TrimSuffix(out.String(), "\n") {
        t.Errorf("output was incorrect, got: %q, want %q.", res, out.String())
    }
}
