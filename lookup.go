package timezh

func (t *Time) lookupDay(s string) bool {
	if v, ok := dayFormat[s]; ok {
		if t.f.mixed {
			if t.f.day&t.f.dayPoint > 0 {
				t.f.buf.WriteString(v)
			} else {
				t.f.buf.WriteString(s)
			}
			t.f.dayPoint >>= 1
			return true
		}
		t.f.buf.WriteString(v)
		return true
	}
	return false
}

func (t *Time) lookupMonth(s string) (ok bool) {
	if v, ok := monthFormat[s]; ok {
		if t.f.mixed {
			if t.f.month&t.f.monthPoint > 0 {
				t.f.buf.WriteString(v)
			} else {
				t.f.buf.WriteString(s)
			}
			t.f.monthPoint >>= 1
			return true
		}
		t.f.buf.WriteString(v)
		return true
	}
	return false
}

func (t *Time) lookupPM(s string) (ok bool) {
	if v, ok := pmFormat[s]; ok {
		if t.f.mixed {
			if t.f.pm&t.f.pmPoint > 0 {
				t.f.buf.WriteString(v)
			} else {
				t.f.buf.WriteString(s)
			}
			t.f.pmPoint >>= 1
			return true
		}
		t.f.buf.WriteString(v)
		return true
	}
	return false
}

func (t *Time) lookup(s string) string {
	t.f.reset()
	j, l := 0, len(s)
	for i, r := range s {
		if i < j {
			continue
		}
		switch r {
		case 'A':
			if l >= i+6 && t.lookupMonth(s[i:i+6]) {
				j = i + 6
			} else if l >= i+5 && t.lookupMonth(s[i:i+5]) {
				j = i + 5
			} else if l >= i+3 && t.lookupMonth(s[i:i+3]) {
				j = i + 3
			} else if l >= i+2 && t.lookupPM(s[i:i+2]) {
				j = i + 2
			}
		case 'D':
			if l >= i+8 && t.lookupMonth(s[i:i+8]) {
				j = i + 8
			} else if l >= i+3 && t.lookupMonth(s[i:i+3]) {
				j = i + 3
			}
		case 'a', 'P', 'p':
			if l >= i+2 && t.lookupPM(s[i:i+2]) {
				j = i + 2
			}
		case 'F':
			if l >= i+8 && t.lookupMonth(s[i:i+8]) {
				j = i + 8
			} else if l >= i+6 && t.lookupDay(s[i:i+6]) {
				j = i + 6
			} else if l >= i+3 {
				if t.lookupDay(s[i : i+3]) {
					j = i + 3
				} else if t.lookupMonth(s[i : i+3]) {
					j = i + 3
				}
			}
		case 'J':
			if l >= i+7 && t.lookupMonth(s[i:i+7]) {
				j = i + 7
			} else if l >= i+4 && t.lookupMonth(s[i:i+4]) {
				j = i + 4
			} else if l >= i+3 && t.lookupMonth(s[i:i+3]) {
				j = i + 3
			}
		case 'M':
			if l >= i+6 && t.lookupDay(s[i:i+6]) {
				j = i + 6
			} else if l >= i+5 && t.lookupMonth(s[i:i+5]) {
				j = i + 5
			} else if l >= i+3 {
				if t.lookupDay(s[i : i+3]) {
					j = i + 3
				} else if t.lookupMonth(s[i : i+3]) {
					j = i + 3
				}
			}
		case 'N':
			if l >= i+8 && t.lookupMonth(s[i:i+8]) {
				j = i + 8
			} else if l >= i+3 && t.lookupMonth(s[i:i+3]) {
				j = i + 3
			}
		case 'O':
			if l >= i+7 && t.lookupMonth(s[i:i+7]) {
				j = i + 7
			} else if l >= i+3 && t.lookupMonth(s[i:i+3]) {
				j = i + 3
			}
		case 'S':
			if l >= i+9 && t.lookupMonth(s[i:i+9]) {
				j = i + 9
			} else if l >= i+8 && t.lookupDay(s[i:i+8]) {
				j = i + 8
			} else if l >= i+6 && t.lookupDay(s[i:i+6]) {
				j = i + 6
			} else if l >= i+3 {
				if t.lookupDay(s[i : i+3]) {
					j = i + 3
				} else if t.lookupMonth(s[i : i+3]) {
					j = i + 3
				}
			}
		case 'T':
			if l >= i+8 && t.lookupDay(s[i:i+8]) {
				j = i + 8
			} else if l >= i+7 && t.lookupDay(s[i:i+7]) {
				j = i + 7
			} else if l >= i+3 {
				if t.lookupDay(s[i : i+3]) {
					j = i + 3
				}
			}
		case 'W':
			if l >= i+9 && t.lookupDay(s[i:i+9]) {
				j = i + 9
			} else if l >= i+3 && t.lookupDay(s[i:i+3]) {
				j = i + 3
			}
		}
		if j <= i {
			t.f.buf.WriteRune(r)
		}
	}

	return t.f.buf.String()
}
