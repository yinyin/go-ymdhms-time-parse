package ymdhmstimeparse_test

import (
	"testing"
	"time"

	ymdhmstimeparse "github.com/yinyin/go-ymdhms-time-parse"
)

func doParseTest(t *testing.T, value string, year, month, day, hour, minute, second int) {
	parsed := ymdhmstimeparse.Parse(value, time.UTC)
	target := time.Date(year, time.Month(month), day, hour, minute, second, 0, time.UTC)
	if !parsed.Equal(target) {
		t.Errorf("unexpected result for [%v]: %#v != %#v", value, parsed, target)
	}
}

func TestParse20090520(t *testing.T) {
	doParseTest(t, "20090520", 2009, 5, 20, 0, 0, 0)
	doParseTest(t, "2009-05-20", 2009, 5, 20, 0, 0, 0)
	doParseTest(t, "09-05-20", 2009, 5, 20, 0, 0, 0)
}

func TestParse20090520010203(t *testing.T) {
	doParseTest(t, "20090520010203", 2009, 5, 20, 1, 2, 3)
	doParseTest(t, "2009-05-20 01:02:03", 2009, 5, 20, 1, 2, 3)
	doParseTest(t, "09-05-20 1;2:3", 2009, 5, 20, 1, 2, 3)
	doParseTest(t, "09-05-20 T 1;2:3", 2009, 5, 20, 1, 2, 3)
}

func TestParse19970520010203(t *testing.T) {
	doParseTest(t, "19970520010203", 1997, 5, 20, 1, 2, 3)
	doParseTest(t, "1997-05-20 01:02:03", 1997, 5, 20, 1, 2, 3)
	doParseTest(t, "97-05-20 1;2:3", 1997, 5, 20, 1, 2, 3)
	doParseTest(t, "97-05-20 T 1;2:3", 1997, 5, 20, 1, 2, 3)
}
