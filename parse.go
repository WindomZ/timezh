package timezh

import "time"

// Parse parses a formatted string and returns the time value it represents.
// The layout defines the format by showing how the reference time,
// defined to be
//	星期一 一月 2 下午15:04:05 -0700 MST 2006
// Same as time.Parse, but support Chinese string.
func Parse(layout, value string) (t Time, err error) {
	t.Time, err = time.Parse(t.parseLayout(layout), t.parseValue(value))
	return
}

// ParseLayout returns a standard textual formatted layout.
// Such as
//	星期一 一月 2 下午15:04:05 -0700 MST 2006
// returns
//	Mon Jan 2 PM15:04:05 -0700 MST 2006
func ParseLayout(layout string) string {
	return new(Time).parseLayout(layout)
}

func (t *Time) parseLayout(layout string) string {
	t.f.reset()
	j := 0
	for i, r := range layout {
		if i < j {
			continue
		}
		switch r {
		case '星':
			if len(layout) >= i+9 && layout[i:i+9] == "星期一" {
				j = i + 9
				t.f.buf.WriteString("Monday")
			}
		case '周':
			if len(layout) >= i+6 && layout[i:i+6] == "周一" {
				j = i + 6
				t.f.buf.WriteString("Mon")
			}
		case '一':
			if len(layout) >= i+6 && layout[i:i+6] == "一月" {
				j = i + 6
				t.f.buf.WriteString("Jan")
			}
		case '下':
			if len(layout) >= i+6 && layout[i:i+6] == "下午" {
				j = i + 6
				t.f.buf.WriteString("PM")
			}
		}
		if j <= i {
			t.f.buf.WriteRune(r)
		}
	}

	return t.f.buf.String()
}

// ParseValue returns a standard textual from
// a textual representation of the time value formatted.
// Such as
//	星期二 二月 2 下午15:04:05 -0700 MST 2006
// returns
//	Tue Feb 2 PM15:04:05 -0700 MST 2006
func ParseValue(value string) string {
	return new(Time).parseValue(value)
}

func (t *Time) parseValue(value string) string {
	t.f.reset()
	j := 0
	for i, r := range value {
		if i < j {
			continue
		}
		switch r {
		case '星':
			if len(value) >= i+9 {
				s := value[i : i+9]
				for k, day := range longDayMap {
					if s == day {
						j = i + 9
						t.f.buf.WriteString(longDayNames[k])
						break
					}
				}
			}
		case '周':
			if len(value) >= i+6 {
				s := value[i : i+6]
				for k, day := range shortDayMap {
					if s == day {
						j = i + 6
						t.f.buf.WriteString(shortDayNames[k])
						break
					}
				}
			}
		case '一', '二', '三', '四', '五', '六', '七', '八', '九', '十':
			if len(value) >= i+6 {
				s := value[i : i+6]
				for k, day := range shortMonthMap {
					if s == day {
						j = i + 6
						t.f.buf.WriteString(shortMonthNames[k])
						break
					}
				}
			}
		case '上', '下':
			if len(value) >= i+6 {
				s := value[i : i+6]
				for k, day := range longPMMap {
					if s == day {
						j = i + 6
						t.f.buf.WriteString(longPMNames[k])
						break
					}
				}
			}
		}
		if j <= i {
			t.f.buf.WriteRune(r)
		}
	}

	return t.f.buf.String()
}
