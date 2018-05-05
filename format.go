package timezh

import "bytes"

type format struct {
	layout     string
	day        uint64
	dayPoint   uint64
	month      uint64
	monthPoint uint64
	pm         uint64
	pmPoint    uint64
	buf        bytes.Buffer
}

func (f *format) init(s string) {
	f.layout = s
	f.day = 0
	f.dayPoint = 1
	f.month = 0
	f.monthPoint = 1
	f.pm = 0
	f.pmPoint = 1
	f.buf.Reset()
}

func (f *format) reset() {
	f.buf.Reset()
}

// Format returns a textual representation of the time value formatted
// according to layout, which defines the format by showing how the reference
// time, defined to be
//	Mon Jan 2 15:04:05 -0700 MST 2006
// would be displayed if it were the value; it serves as an example of the
// desired output. The same display rules will then be applied to the time
// value.
//
// A fractional second is represented by adding a period and zeros
// to the end of the seconds section of layout string, as in "15:04:05.000"
// to format a time stamp with millisecond precision.
//
// Predefined layouts ANSIC, UnixDate, RFC3339 and others describe standard
// and convenient representations of the reference time. For more information
// about the formats and the definition of the reference time, see the
// documentation for ANSIC and the other constants defined by this package.
func (t Time) Format(layout string) string {
	return t.lookup(t.Time.Format(t.formatLayout(layout)))
}

// FormatLayout returns a standard textual formatted layout.
// Such as
//	星期一 一月 2 下午15:04:05 -0700 MST 2006
// returns
//	Mon Jan 2 PM15:04:05 -0700 MST 2006
func FormatLayout(layout string) string {
	return new(Time).formatLayout(layout)
}

func (t *Time) formatLayout(layout string) string {
	t.f.init(layout)
	j := 0
	for i, r := range layout {
		if i < j {
			continue
		}
		switch r {
		case 'M':
			if len(layout) >= i+6 && layout[i:i+6] == "Monday" {
				j = i + 6
				t.f.day <<= 1
				t.f.dayPoint <<= 1
				t.f.buf.WriteString("Monday")
			} else if len(layout) >= i+3 && layout[i:i+3] == "Mon" {
				j = i + 3
				t.f.day <<= 1
				t.f.dayPoint <<= 1
				t.f.buf.WriteString("Mon")
			}
		case 'J':
			if len(layout) >= i+7 && layout[i:i+7] == "January" {
				j = i + 7
				t.f.month <<= 1
				t.f.monthPoint <<= 1
				t.f.buf.WriteString("January")
			} else if len(layout) >= i+3 && layout[i:i+3] == "Jan" {
				j = i + 3
				t.f.month <<= 1
				t.f.monthPoint <<= 1
				t.f.buf.WriteString("Jan")
			}
		case 'P', 'p':
			if len(layout) >= i+2 {
				s := layout[i : i+2]
				if s == "PM" || s == "pm" {
					j = i + 2
					t.f.pm <<= 1
					t.f.pmPoint <<= 1
					t.f.buf.WriteString(s)
				}
			}
		case '星':
			if len(layout) >= i+9 && layout[i:i+9] == "星期一" {
				j = i + 9
				t.f.day = (t.f.day + 1) << 1
				t.f.dayPoint <<= 1
				t.f.buf.WriteString("Monday")
			}
		case '周':
			if len(layout) >= i+6 && layout[i:i+6] == "周一" {
				j = i + 6
				t.f.day = (t.f.day + 1) << 1
				t.f.dayPoint <<= 1
				t.f.buf.WriteString("Mon")
			}
		case '一':
			if len(layout) >= i+6 && layout[i:i+6] == "一月" {
				j = i + 6
				t.f.month = (t.f.month + 1) << 1
				t.f.monthPoint <<= 1
				t.f.buf.WriteString("Jan")
			}
		case '下':
			if len(layout) >= i+6 && layout[i:i+6] == "下午" {
				j = i + 6
				t.f.pm = (t.f.pm + 1) << 1
				t.f.pmPoint <<= 1
				t.f.buf.WriteString("PM")
			}
		}
		if j <= i {
			t.f.buf.WriteRune(r)
		}
	}

	return t.f.buf.String()
}
