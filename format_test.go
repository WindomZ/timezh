package timezh

import (
	"testing"
	"time"

	"github.com/WindomZ/testify/assert"
)

func TestTime_Format(t *testing.T) {
	layout1 := "January, Jan, 2006-01-02, 02-Jan-06, Monday, Mon, 15:04:05PM -0700, 15:04:05Z07:00PM MST"
	assert.Equal(t, time.Now().Format(layout1), Now().Format(layout1))

	layout2 := "January, 一月, Jan, 2006-01-02, 02-一月-06, Monday, 星期一, Mon, 周一, 15:04:05PM -0700, 15:04:05Z07:00下午 MST"
	assert.NotEqual(t, layout2, FormatLayout(layout2))
	assert.NotEqual(t, time.Now().Format(layout2), Now().Format(layout2))

	layout3 := "January, 一月, Jan, 2006-01-02, 02-一月-06, Monday, 星期一, Mon, 周一, 15:04:05PM -0700, 15:04:05.999999999Z07:00下午 MST"
	assert.NotEqual(t, layout3, FormatLayout(layout3))
	assert.NotEqual(t, time.Now().Format(layout3), Now().Format(layout3))

	layout4 := "15:04:05PM party time"
	assert.Equal(t, layout4, FormatLayout(layout4))
	assert.Equal(t, time.Now().Format(layout4), Now().Format(layout4))

	layout5 := "all "
	for month := range monthNames {
		layout5 += month + " "
	}
	for day := range dayNames {
		layout5 += day + " "
	}
	assert.Equal(t, layout5, FormatLayout(layout5))
	assert.Equal(t, time.Now().Format(layout5), Now().Format(layout5))

	layout6 := "15:04:05PM 15:04:05pm 3:04:05PM 3:04:05pm"
	assert.Equal(t, layout6, FormatLayout(layout6))
	assert.Equal(t, time.Now().Format(layout6), Now().Format(layout6))

	layout7 := "15:04:05PM 15:04:05pm 3:04:05PM 3:04:05pm 15:04:05下午 3:04:05下午"
	assert.NotEqual(t, layout7, FormatLayout(layout7))
	if time.Now().Hour() >= 12 {
		assert.Equal(t, time.Now().Format(layout7), Now().Format(layout7))
	} else {
		assert.NotEqual(t, time.Now().Format(layout7), Now().Format(layout7))
	}
}

func TestFormatLayout(t *testing.T) {
	layout1 := "January, Jan, 2006-01-02, 02-Jan-06, Monday, Mon, 15:04:05PM -0700, 15:04:05Z07:00PM MST"
	assert.Equal(t, layout1, FormatLayout(layout1))
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

func BenchmarkTime_Format1(b *testing.B) {
	t := time.Now()
	t.Format("2006年01月02日(January) 3:04:05PM Mon")
	b.StopTimer()
	b.ReportAllocs()
	b.StartTimer()
	b.RunParallel(func(pb *testing.PB) {
		for pb.Next() {
			t.Format("2006年01月02日(January) 3:04:05PM Mon")
		}
	})
}

func BenchmarkTime_Format2(b *testing.B) {
	t := Now()
	t.Format("2006年01月02日(一月) 3:04:05下午 星期一")
	b.StopTimer()
	b.ReportAllocs()
	b.StartTimer()
	b.RunParallel(func(pb *testing.PB) {
		for pb.Next() {
			t.Format("2006年01月02日(一月) 3:04:05下午 星期一")
		}
	})
}

func BenchmarkTime_FormatMix(b *testing.B) {
	t := Now()
	t.FormatMix("2006年01月02日(January, 一月) 下午3:04:05PM 星期一(Mon, 周一)")
	b.StopTimer()
	b.ReportAllocs()
	b.StartTimer()
	b.RunParallel(func(pb *testing.PB) {
		for pb.Next() {
			t.FormatMix("2006年01月02日(January, 一月) 下午3:04:05PM 星期一(Mon, 周一)")
		}
	})
}
