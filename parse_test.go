package timezh

import (
	"testing"
	"time"

	"github.com/WindomZ/testify/assert"
)

func TestParseLayout(t *testing.T) {
	layout1 := "January, Jan, 2006-01-02, 02-Jan-06, Monday, Mon, 15:04:05PM -0700, 15:04:05Z07:00PM MST"
	assert.Equal(t, layout1, ParseLayout(layout1))
	assert.Equal(t, layout1, ParseLayout(FormatLayout(layout1)))

	layout21 := "一月, 2006-01-02, 02-Jan-06, 星期一, 周一, Mon, 15:04:05下午 -0700, 15:04:05Z07:00PM MST"
	layout22 := "Jan, 2006-01-02, 02-Jan-06, Monday, Mon, Mon, 15:04:05PM -0700, 15:04:05Z07:00PM MST"
	assert.Equal(t, layout22, ParseLayout(layout21))
	assert.Equal(t, layout22, ParseLayout(FormatLayout(layout21)))
}

func TestParseValue(t *testing.T) {
	value11 := "一月, 2006-01-02, 02-Jan-06, 星期一, 周一, Mon, 15:04:05下午 -0700, 15:04:05Z07:00PM MST"
	value12 := "Jan, 2006-01-02, 02-Jan-06, Monday, Mon, Mon, 15:04:05PM -0700, 15:04:05Z07:00PM MST"
	assert.Equal(t, value12, ParseValue(value11))

	value21 := "二月, 2006-01-02, 02-Tue-06, 星期二, 周二, Tue, 15:04:05下午 -0700, 15:04:05Z07:00PM MST"
	value22 := "Feb, 2006-01-02, 02-Tue-06, Tuesday, Tue, Tue, 15:04:05PM -0700, 15:04:05Z07:00PM MST"
	assert.Equal(t, value22, ParseValue(value21))
}

func TestParse(t *testing.T) {
	testTime, _ := time.Parse("2006-01-02 15:04:05", "2009-01-03 18:15:05")

	layout1 := "Jan, 2006-01-02, 02-Jan-06, Mon, 15:04:05PM -0700, 15:04:05Z07:00PM MST"
	if result, err := Parse(layout1, testTime.Format(layout1)); assert.NoError(t, err) {
		assert.Equal(t, testTime.UnixNano(), result.UnixNano())
	}
	if result, err := Parse(layout1, T(testTime).Format(layout1)); assert.NoError(t, err) {
		assert.Equal(t, testTime.UnixNano(), result.UnixNano())
	}

	layout2 := "一月, 2006-01-02, 02-Jan-06, 星期一, 周一, Mon, 15:04:05下午 -0700, 15:04:05Z07:00PM MST"
	if result, err := Parse(layout2, testTime.Format(layout2)); assert.NoError(t, err) {
		assert.Equal(t, testTime.UnixNano(), result.UnixNano())
	}
	if result, err := Parse(layout2, T(testTime).Format(layout2)); assert.NoError(t, err) {
		assert.Equal(t, testTime.UnixNano(), result.UnixNano())
	}
}
