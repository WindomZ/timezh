package timezh

import (
	"testing"
	"time"

	"github.com/WindomZ/testify/assert"
)

func TestTime_Format(t *testing.T) {
	testTime, _ := time.Parse("2006-01-02 15:04:05", "2009-01-03 18:15:05")

	assert.Equal(t, "2009-01-03(一月) 下午6:15:05 星期六(周六)",
		T(testTime).Format("2006-01-02(January) PM3:04:05 Monday(Mon)"))
	assert.Equal(t, "2009年01月03日(一月) 下午6:15:05 星期六(周六)",
		T(testTime).Format("2006年01月02日(January) PM3:04:05 Monday(Mon)"))
	assert.Equal(t, "2009年01月03日(一月) 下午6:15:05 星期六(周六)",
		T(testTime).Format("2006年01月02日(一月) 下午3:04:05 星期一(周一)"))
	assert.Equal(t, "2009年01月03日(一月, 一月) 下午6:15:05下午 星期六(周六, 周六)",
		T(testTime).Format("2006年01月02日(January, 一月) 下午3:04:05PM 星期一(Mon, 周一)"))
}

func TestTime_FormatMix(t *testing.T) {
	testTime, _ := time.Parse("2006-01-02 15:04:05", "2009-01-03 18:15:05")
	assert.Equal(t, "2009年01月03日(January, 一月) 下午6:15:05PM 星期六(Sat, 周六)",
		T(testTime).FormatMix("2006年01月02日(January, 一月) 下午3:04:05PM 星期一(Mon, 周一)"))

	testTime, _ = time.Parse("2006-01-02 15:04:05", "2009-01-04 02:15:05")
	assert.Equal(t, "2009年01月04日(January, 一月) 上午2:15:05AM 星期日(Sun, 周日)",
		T(testTime).FormatMix("2006年01月02日(January, 一月) 下午3:04:05PM 星期一(Mon, 周一)"))
}

func TestTime_Format_All(t *testing.T) {
	layout1 := "January, Jan, 2006-01-02, 02-Jan-06, Monday, Mon, 15:04:05PM -0700, 15:04:05Z07:00PM MST"
	assert.NotEqual(t, time.Now().Format(layout1), Now().Format(layout1))
	assert.NotEqual(t, Now().Format(layout1), Now().FormatMix(layout1))

	layout2 := "January, 一月, Jan, 2006-01-02, 02-一月-06, Monday, 星期一, Mon, 周一, 15:04:05PM -0700, 15:04:05Z07:00下午 MST"
	assert.NotEqual(t, layout2, FormatLayout(layout2))
	assert.NotEqual(t, time.Now().Format(layout2), Now().Format(layout2))
	assert.NotEqual(t, Now().Format(layout2), Now().FormatMix(layout2))

	layout3 := "January, 一月, Jan, 2006-01-02, 02-一月-06, Monday, 星期一, Mon, 周一, 15:04:05PM -0700, 15:04:05.999999999Z07:00下午 MST"
	assert.NotEqual(t, layout3, FormatLayout(layout3))
	assert.NotEqual(t, time.Now().Format(layout3), Now().Format(layout3))
	assert.NotEqual(t, Now().Format(layout3), Now().FormatMix(layout3))

	layout4 := "15:04:05PM party time"
	assert.Equal(t, layout4, FormatLayout(layout4))
	assert.NotEqual(t, time.Now().Format(layout4), Now().Format(layout4))
	assert.NotEqual(t, Now().Format(layout4), Now().FormatMix(layout4))

	layout5 := "all "
	for _, month := range shortMonthNames {
		layout5 += month + " "
	}
	for _, month := range longMonthNames {
		layout5 += month + " "
	}
	for _, day := range longDayNames {
		layout5 += day + " "
	}
	for _, day := range shortDayNames {
		layout5 += day + " "
	}
	assert.Equal(t, layout5, FormatLayout(layout5))
	assert.NotEqual(t, time.Now().Format(layout5), Now().Format(layout5))
	assert.NotEqual(t, Now().Format(layout5), Now().FormatMix(layout5))

	layout6 := "15:04:05PM 15:04:05pm 3:04:05PM 3:04:05pm"
	layout7 := "15:04:05PM 15:04:05pm 3:04:05PM 3:04:05pm 15:04:05下午 3:04:05下午"

	testTime, _ := time.Parse("2006-01-02 15:04:05", "2009-01-03 18:15:05")

	assert.NotEqual(t, layout6, FormatLayout(layout6))
	assert.NotEqual(t, testTime.Format(layout6), T(testTime).Format(layout6))
	assert.Equal(t, "18:15:05PM 18:15:05PM 6:15:05PM 6:15:05PM",
		T(testTime).FormatMix(layout6))

	assert.NotEqual(t, layout7, FormatLayout(layout7))
	assert.NotEqual(t, testTime.Format(layout7), T(testTime).Format(layout7))
	assert.Equal(t, "18:15:05PM 18:15:05PM 6:15:05PM 6:15:05PM 18:15:05下午 6:15:05下午",
		T(testTime).FormatMix(layout7))

	testTime, _ = time.Parse("2006-01-02 15:04:05", "2009-01-04 02:15:05")

	assert.NotEqual(t, layout6, FormatLayout(layout6))
	assert.NotEqual(t, testTime.Format(layout6), T(testTime).Format(layout6))
	assert.Equal(t, "02:15:05AM 02:15:05AM 2:15:05AM 2:15:05AM",
		T(testTime).FormatMix(layout6))

	assert.NotEqual(t, layout7, FormatLayout(layout7))
	assert.NotEqual(t, testTime.Format(layout7), T(testTime).Format(layout7))
	assert.Equal(t, "02:15:05AM 02:15:05AM 2:15:05AM 2:15:05AM 02:15:05上午 2:15:05上午",
		T(testTime).FormatMix(layout7))
}

func TestFormatLayout(t *testing.T) {
	layout1 := "January, Jan, 2006-01-02, 02-Jan-06, Monday, Mon, 15:04:05PM -0700, 15:04:05Z07:00PM MST"
	assert.Equal(t, layout1, FormatLayout(layout1))
}

func TestFormatChinese(t *testing.T) {
	assert.Equal(t, "一月, 一月, 02-一月-06, 星期一, 周一, 15:04:05下午",
		FormatChinese("January, Jan, 02-Jan-06, Monday, Mon, 15:04:05PM"))
}

func BenchmarkFormatLayout(b *testing.B) {
	b.StopTimer()
	b.ReportAllocs()
	b.StartTimer()
	b.RunParallel(func(pb *testing.PB) {
		for pb.Next() {
			FormatLayout("2006年01月02日(January, 一月) 下午3:04:05PM 星期一(Mon, 周一)")
		}
	})
}

func BenchmarkFormatChinese(b *testing.B) {
	b.StopTimer()
	b.ReportAllocs()
	b.StartTimer()
	b.RunParallel(func(pb *testing.PB) {
		for pb.Next() {
			FormatLayout("January, Jan, 02-Jan-06, Monday, Mon, 15:04:05PM")
		}
	})
}

func BenchmarkTime_Format1(b *testing.B) {
	b.StopTimer()
	t := time.Now()
	b.ReportAllocs()
	b.StartTimer()
	b.RunParallel(func(pb *testing.PB) {
		for pb.Next() {
			t.Format("2006年01月02日 January 3:04:05PM Mon")
		}
	})
}

func BenchmarkTime_Format2(b *testing.B) {
	b.StopTimer()
	t := Now()
	b.ReportAllocs()
	b.StartTimer()
	b.RunParallel(func(pb *testing.PB) {
		for pb.Next() {
			t.Format("2006年01月02日 一月 3:04:05下午 星期一")
		}
	})
}

func BenchmarkTime_FormatMix(b *testing.B) {
	b.StopTimer()
	t := Now()
	t.FormatMix("2006年01月02日(January, 一月) 下午3:04:05PM 星期一(Mon, 周一)")
	b.ReportAllocs()
	b.StartTimer()
	b.RunParallel(func(pb *testing.PB) {
		for pb.Next() {
			t.FormatMix("2006年01月02日(January, 一月) 下午3:04:05PM 星期一(Mon, 周一)")
		}
	})
}
