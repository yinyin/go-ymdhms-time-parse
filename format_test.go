package ymdhmstimeparse_test

import (
	"testing"
	"time"

	ymdhmstimeparse "github.com/yinyin/go-ymdhms-time-parse"
)

func doFormatTest(t *testing.T, baseTime time.Time, formatString, expectText string) {
	if formatted := baseTime.Format(formatString); formatted != expectText {
		t.Errorf("unexpected result: %v != %v", formatted, expectText)
	}
}

func TestFormat1(t *testing.T) {
	baseTime := time.Date(2018, 2, 27, 23, 56, 9, 0, time.UTC)
	doFormatTest(t, baseTime, ymdhmstimeparse.DashSeparatedYYYYMMDDhhmmss, "2018-02-27 23:56:09")
	doFormatTest(t, baseTime, ymdhmstimeparse.DashSeparatedYYYYMMDDhhmm, "2018-02-27 23:56")
	doFormatTest(t, baseTime, ymdhmstimeparse.DashSeparatedYYYYMMDD, "2018-02-27")
	doFormatTest(t, baseTime, ymdhmstimeparse.ConcatedYYYYMMDDhhmmss, "20180227235609")
	doFormatTest(t, baseTime, ymdhmstimeparse.ConcatedYYYYMMDDhhmm, "201802272356")
	doFormatTest(t, baseTime, ymdhmstimeparse.ConcatedYYYYMMDD, "20180227")
}

func TestFormat2(t *testing.T) {
	baseTime := time.Date(2018, 2, 3, 5, 6, 9, 0, time.UTC)
	doFormatTest(t, baseTime, ymdhmstimeparse.DashSeparatedYYYYMMDDhhmmss, "2018-02-03 05:06:09")
	doFormatTest(t, baseTime, ymdhmstimeparse.DashSeparatedYYYYMMDDhhmm, "2018-02-03 05:06")
	doFormatTest(t, baseTime, ymdhmstimeparse.DashSeparatedYYYYMMDD, "2018-02-03")
	doFormatTest(t, baseTime, ymdhmstimeparse.ConcatedYYYYMMDDhhmmss, "20180203050609")
	doFormatTest(t, baseTime, ymdhmstimeparse.ConcatedYYYYMMDDhhmm, "201802030506")
	doFormatTest(t, baseTime, ymdhmstimeparse.ConcatedYYYYMMDD, "20180203")
}
