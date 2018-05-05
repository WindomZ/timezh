package timezh

func (t *Time) lookupDay(s string) bool {
	if v, ok := dayNames[s]; ok {
		if t.f.day&t.f.dayPoint > 0 {
			t.f.buf.WriteString(v)
		} else {
			t.f.buf.WriteString(s)
		}
		t.f.dayPoint >>= 1
		return true
	}
	return false
}

func (t *Time) lookupMonth(s string) (ok bool) {
	if v, ok := monthNames[s]; ok {
		if t.f.month&t.f.monthPoint > 0 {
			t.f.buf.WriteString(v)
		} else {
			t.f.buf.WriteString(s)
		}
		t.f.monthPoint >>= 1
		return true
	}
	return false
}

func (t *Time) lookupPM(s string) (ok bool) {
	if v, ok := pmNames[s]; ok {
		if t.f.pm&t.f.pmPoint > 0 {
			t.f.buf.WriteString(v)
		} else {
			t.f.buf.WriteString(s)
		}
		t.f.pmPoint >>= 1
		return true
	}
	return false
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
			if len(s) >= i+6 && t.lookupMonth(s[i:i+6]) {
				j = i + 6
			} else if len(s) >= i+5 && t.lookupMonth(s[i:i+5]) {
				j = i + 5
			} else if len(s) >= i+3 && t.lookupMonth(s[i:i+3]) {
				j = i + 3
			} else if len(s) >= i+2 && t.lookupPM(s[i:i+2]) {
				j = i + 2
			}
		case 'D':
			if len(s) >= i+8 && t.lookupMonth(s[i:i+8]) {
				j = i + 8
			} else if len(s) >= i+3 && t.lookupMonth(s[i:i+3]) {
				j = i + 3
			}
		case 'a', 'P', 'p':
			if len(s) >= i+2 && t.lookupPM(s[i:i+2]) {
				j = i + 2
			}
		case 'F':
			if len(s) >= i+8 && t.lookupMonth(s[i:i+8]) {
				j = i + 8
			} else if len(s) >= i+6 && t.lookupDay(s[i:i+6]) {
				j = i + 6
			} else if len(s) >= i+3 {
				if t.lookupDay(s[i : i+3]) {
					j = i + 3
				} else if t.lookupMonth(s[i : i+3]) {
					j = i + 3
				}
			}
		case 'J':
			if len(s) >= i+7 && t.lookupMonth(s[i:i+7]) {
				j = i + 7
			} else if len(s) >= i+4 && t.lookupMonth(s[i:i+4]) {
				j = i + 4
			} else if len(s) >= i+3 && t.lookupMonth(s[i:i+3]) {
				j = i + 3
			}
		case 'M':
			if len(s) >= i+6 && t.lookupDay(s[i:i+6]) {
				j = i + 6
			} else if len(s) >= i+5 && t.lookupMonth(s[i:i+5]) {
				j = i + 5
			} else if len(s) >= i+3 {
				if t.lookupDay(s[i : i+3]) {
					j = i + 3
				} else if t.lookupMonth(s[i : i+3]) {
					j = i + 3
				}
			}
		case 'N':
			if len(s) >= i+8 && t.lookupMonth(s[i:i+8]) {
				j = i + 8
			} else if len(s) >= i+3 && t.lookupMonth(s[i:i+3]) {
				j = i + 3
			}
		case 'O':
			if len(s) >= i+7 && t.lookupMonth(s[i:i+7]) {
				j = i + 7
			} else if len(s) >= i+3 && t.lookupMonth(s[i:i+3]) {
				j = i + 3
			}
		case 'S':
			if len(s) >= i+9 && t.lookupMonth(s[i:i+9]) {
				j = i + 9
			} else if len(s) >= i+8 && t.lookupDay(s[i:i+8]) {
				j = i + 8
			} else if len(s) >= i+6 && t.lookupDay(s[i:i+6]) {
				j = i + 6
			} else if len(s) >= i+3 {
				if t.lookupDay(s[i : i+3]) {
					j = i + 3
				} else if t.lookupMonth(s[i : i+3]) {
					j = i + 3
				}
			}
		case 'T':
			if len(s) >= i+8 && t.lookupDay(s[i:i+8]) {
				j = i + 8
			} else if len(s) >= i+7 && t.lookupDay(s[i:i+7]) {
				j = i + 7
			} else if len(s) >= i+3 {
				if t.lookupDay(s[i : i+3]) {
					j = i + 3
				}
			}
		case 'W':
			if len(s) >= i+9 && t.lookupDay(s[i:i+9]) {
				j = i + 9
			} else if len(s) >= i+3 && t.lookupDay(s[i:i+3]) {
				j = i + 3
			}
		}
		if j <= i {
			t.f.buf.WriteRune(r)
		}
	}

	return t.f.buf.String()
}
