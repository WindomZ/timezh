package timezh

import (
	"testing"
	"time"

	"github.com/WindomZ/testify/assert"
)

func TestTime_Format(t *testing.T) {
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
	for month := range monthNames {
		layout5 += month + " "
	}
	for day := range dayNames {
		layout5 += day + " "
	}
	assert.Equal(t, layout5, FormatLayout(layout5))
	assert.NotEqual(t, time.Now().Format(layout5), Now().Format(layout5))
	assert.NotEqual(t, Now().Format(layout5), Now().FormatMix(layout5))

	layout6 := "15:04:05PM 15:04:05pm 3:04:05PM 3:04:05pm"
	assert.NotEqual(t, layout6, FormatLayout(layout6))
	assert.NotEqual(t, time.Now().Format(layout6), Now().Format(layout6))
	assert.NotEqual(t, Now().Format(layout6), Now().FormatMix(layout6))

	layout7 := "15:04:05PM 15:04:05pm 3:04:05PM 3:04:05pm 15:04:05下午 3:04:05下午"
	assert.NotEqual(t, layout7, FormatLayout(layout7))
	assert.NotEqual(t, time.Now().Format(layout7), Now().Format(layout7))
	assert.NotEqual(t, Now().Format(layout7), Now().FormatMix(layout7))
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
