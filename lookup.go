package timezh

func (t *Time) lookupDay() bool {
	if t.f.mixed {
		ok := t.f.day&t.f.dayPoint != 0
		t.f.dayPoint >>= 1
		return ok
	}
	return true
}

func (t *Time) lookupMonth() bool {
	if t.f.mixed {
		ok := t.f.month&t.f.monthPoint != 0
		t.f.monthPoint >>= 1
		return ok
	}
	return true
}

func (t *Time) lookupPM() bool {
	if t.f.mixed {
		ok := t.f.pm&t.f.pmPoint != 0
		t.f.pmPoint >>= 1
		return ok
	}
	return true
}

func (t *Time) lookup(s string) string {
	t.f.reset()
	j := 0
	for i, r := range s {
		if i < j {
			continue
		}
		switch r {
		case 'A':
			if len(s) >= i+6 && s[i:i+6] == "August" {
				j = i + 6
				if t.lookupMonth() {
					t.f.buf.WriteString(longMonthMap[August])
				} else {
					t.f.buf.WriteString("August")
				}
			} else if len(s) >= i+5 && s[i:i+5] == "April" {
				j = i + 5
				if t.lookupMonth() {
					t.f.buf.WriteString(longMonthMap[April])
				} else {
					t.f.buf.WriteString("April")
				}
			} else if len(s) >= i+3 && s[i:i+3] == "Apr" {
				j = i + 3
				if t.lookupMonth() {
					t.f.buf.WriteString(shortMonthMap[April])
				} else {
					t.f.buf.WriteString("Apr")
				}
			} else if len(s) >= i+3 && s[i:i+3] == "Aug" {
				j = i + 3
				if t.lookupMonth() {
					t.f.buf.WriteString(shortMonthMap[August])
				} else {
					t.f.buf.WriteString("Aug")
				}
			} else if len(s) >= i+2 && s[i:i+2] == "AM" {
				j = i + 2
				if t.lookupPM() {
					t.f.buf.WriteString(longPMMap[AM])
				} else {
					t.f.buf.WriteString("AM")
				}
			}
		case 'D':
			if len(s) >= i+8 && s[i:i+8] == "December" {
				j = i + 8
				if t.lookupMonth() {
					t.f.buf.WriteString(longMonthMap[December])
				} else {
					t.f.buf.WriteString("December")
				}
			} else if len(s) >= i+3 && s[i:i+3] == "Dec" {
				j = i + 3
				if t.lookupMonth() {
					t.f.buf.WriteString(shortMonthMap[December])
				} else {
					t.f.buf.WriteString("Dec")
				}
			}
		case 'F':
			if len(s) >= i+8 && s[i:i+8] == "February" {
				j = i + 8
				if t.lookupMonth() {
					t.f.buf.WriteString(longMonthMap[February])
				} else {
					t.f.buf.WriteString("February")
				}
			} else if len(s) >= i+6 && s[i:i+6] == "Friday" {
				j = i + 6
				if t.lookupDay() {
					t.f.buf.WriteString(longDayMap[Friday])
				} else {
					t.f.buf.WriteString("Friday")
				}
			} else if len(s) >= i+3 {
				switch s[i : i+3] {
				case "Feb":
					j = i + 3
					if t.lookupMonth() {
						t.f.buf.WriteString(shortMonthMap[February])
					} else {
						t.f.buf.WriteString("Feb")
					}
				case "Fri":
					j = i + 3
					if t.lookupDay() {
						t.f.buf.WriteString(shortDayMap[Friday])
					} else {
						t.f.buf.WriteString("Fri")
					}
				}
			}
		case 'J':
			if len(s) >= i+7 && s[i:i+7] == "January" {
				j = i + 7
				if t.lookupMonth() {
					t.f.buf.WriteString(longMonthMap[January])
				} else {
					t.f.buf.WriteString("January")
				}
			} else if len(s) >= i+4 && s[i:i+4] == "June" {
				j = i + 4
				if t.lookupMonth() {
					t.f.buf.WriteString(longMonthMap[June])
				} else {
					t.f.buf.WriteString("June")
				}
			} else if len(s) >= i+4 && s[i:i+4] == "July" {
				j = i + 4
				if t.lookupMonth() {
					t.f.buf.WriteString(longMonthMap[July])
				} else {
					t.f.buf.WriteString("July")
				}
			} else if len(s) >= i+3 {
				switch s[i : i+3] {
				case "Jan":
					j = i + 3
					if t.lookupMonth() {
						t.f.buf.WriteString(shortMonthMap[January])
					} else {
						t.f.buf.WriteString("Jan")
					}
				case "Jun":
					j = i + 3
					if t.lookupMonth() {
						t.f.buf.WriteString(shortMonthMap[June])
					} else {
						t.f.buf.WriteString("Jun")
					}
				case "Jul":
					j = i + 3
					if t.lookupMonth() {
						t.f.buf.WriteString(shortMonthMap[July])
					} else {
						t.f.buf.WriteString("Jul")
					}
				}
			}
		case 'M':
			if len(s) >= i+6 && s[i:i+6] == "Monday" {
				j = i + 6
				if t.lookupDay() {
					t.f.buf.WriteString(longDayMap[Monday])
				} else {
					t.f.buf.WriteString("Monday")
				}
			} else if len(s) >= i+5 && s[i:i+5] == "March" {
				j = i + 5
				if t.lookupMonth() {
					t.f.buf.WriteString(longMonthMap[March])
				} else {
					t.f.buf.WriteString("March")
				}
			} else if len(s) >= i+3 {
				switch s[i : i+3] {
				case "Mar":
					j = i + 3
					if t.lookupMonth() {
						t.f.buf.WriteString(shortMonthMap[March])
					} else {
						t.f.buf.WriteString("Mar")
					}
				case "May":
					j = i + 3
					if t.lookupMonth() {
						t.f.buf.WriteString(longMonthMap[May])
					} else {
						t.f.buf.WriteString("May")
					}
				case "Mon":
					j = i + 3
					if t.lookupDay() {
						t.f.buf.WriteString(shortDayMap[Monday])
					} else {
						t.f.buf.WriteString("Mon")
					}
				}
			}
		case 'N':
			if len(s) >= i+8 && s[i:i+8] == "November" {
				j = i + 8
				if t.lookupMonth() {
					t.f.buf.WriteString(longMonthMap[November])
				} else {
					t.f.buf.WriteString("November")
				}
			} else if len(s) >= i+3 && s[i:i+3] == "Nov" {
				j = i + 3
				if t.lookupMonth() {
					t.f.buf.WriteString(shortMonthMap[November])
				} else {
					t.f.buf.WriteString("Nov")
				}
			}
		case 'O':
			if len(s) >= i+7 && s[i:i+7] == "October" {
				j = i + 7
				if t.lookupMonth() {
					t.f.buf.WriteString(longMonthMap[October])
				} else {
					t.f.buf.WriteString("October")
				}
			} else if len(s) >= i+3 && s[i:i+3] == "Oct" {
				j = i + 3
				if t.lookupMonth() {
					t.f.buf.WriteString(shortMonthMap[October])
				} else {
					t.f.buf.WriteString("Oct")
				}
			}
		case 'P':
			if len(s) >= i+2 {
				switch s[i : i+2] {
				case "PM":
					j = i + 2
					if t.lookupPM() {
						t.f.buf.WriteString(longPMMap[PM])
					} else {
						t.f.buf.WriteString("PM")
					}
				}
			}
		case 'S':
			if len(s) >= i+9 && s[i:i+9] == "September" {
				j = i + 9
				if t.lookupMonth() {
					t.f.buf.WriteString(longMonthMap[September])
				} else {
					t.f.buf.WriteString("September")
				}
			} else if len(s) >= i+8 && s[i:i+8] == "Saturday" {
				j = i + 8
				if t.lookupDay() {
					t.f.buf.WriteString(longDayMap[Saturday])
				} else {
					t.f.buf.WriteString("Saturday")
				}
			} else if len(s) >= i+6 && s[i:i+6] == "Sunday" {
				j = i + 6
				if t.lookupDay() {
					t.f.buf.WriteString(longDayMap[Sunday])
				} else {
					t.f.buf.WriteString("Sunday")
				}
			} else if len(s) >= i+3 {
				switch s[i : i+3] {
				case "Sep":
					j = i + 3
					if t.lookupMonth() {
						t.f.buf.WriteString(shortMonthMap[September])
					} else {
						t.f.buf.WriteString("Sep")
					}
				case "Sun":
					j = i + 3
					if t.lookupDay() {
						t.f.buf.WriteString(shortDayMap[Sunday])
					} else {
						t.f.buf.WriteString("Sun")
					}
				case "Sat":
					j = i + 3
					if t.lookupDay() {
						t.f.buf.WriteString(shortDayMap[Saturday])
					} else {
						t.f.buf.WriteString("Sat")
					}
				}
			}
		case 'T':
			if len(s) >= i+8 && s[i:i+8] == "Thursday" {
				j = i + 8
				if t.lookupDay() {
					t.f.buf.WriteString(longDayMap[Thursday])
				} else {
					t.f.buf.WriteString("Thursday")
				}
			} else if len(s) >= i+7 && s[i:i+7] == "Tuesday" {
				j = i + 7
				if t.lookupDay() {
					t.f.buf.WriteString(longDayMap[Tuesday])
				} else {
					t.f.buf.WriteString("Tuesday")
				}
			} else if len(s) >= i+3 {
				switch s[i : i+3] {
				case "Tue":
					j = i + 3
					if t.lookupDay() {
						t.f.buf.WriteString(shortDayMap[Tuesday])
					} else {
						t.f.buf.WriteString("Tue")
					}
				case "Thu":
					j = i + 3
					if t.lookupDay() {
						t.f.buf.WriteString(shortDayMap[Thursday])
					} else {
						t.f.buf.WriteString("Thu")
					}
				}
			}
		case 'W':
			if len(s) >= i+9 && s[i:i+9] == "Wednesday" {
				j = i + 9
				if t.lookupDay() {
					t.f.buf.WriteString(longDayMap[Wednesday])
				} else {
					t.f.buf.WriteString("Wednesday")
				}
			} else if len(s) >= i+3 && s[i:i+3] == "Wed" {
				j = i + 3
				if t.lookupDay() {
					t.f.buf.WriteString(shortDayMap[Wednesday])
				} else {
					t.f.buf.WriteString("Wed")
				}
			}
		}
		if j <= i {
			t.f.buf.WriteRune(r)
		}
	}

	return t.f.buf.String()
}
