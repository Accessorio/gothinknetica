package people

import (
	"io"
)

type Client struct {
	ID   uint
	NAME string
	AGE  uint
}

type Employee struct {
	ID   uint
	NAME string
	AGE  uint
}

func (c *Client) Age() uint {
	return c.AGE
}

func (e *Employee) Age() uint {
	return e.AGE
}

type Ager interface {
	Age() uint
}

func MaxAge(age ...Ager) (uint, string) {
	if len(age) != 0 {
		max := age[0].Age()
		for i := 0; i < len(age); i++ {
			if age[i].Age() > max {
				max = age[i].Age()
			}
		}
		return max, "OK"
	} else {
		return 0, "Empty list"
	}
}

func MaxAgeAny(age ...any) any {
	if len(age) != 0 {
		max := age[0]
		for i := 0; i < len(age); i++ {
			switch people := age[i].(type) {
			case Employee:
				switch Max := max.(type) {
				case Employee:
					if people.AGE > Max.AGE {
						max = people
					}
				case Client:
					if people.AGE > Max.AGE {
						max = people
					}
				}
				continue
			case Client:
				switch Max := max.(type) {
				case Employee:
					if people.AGE > Max.AGE {
						max = people
					}
				case Client:
					if people.AGE > Max.AGE {
						max = people
					}
				}
				continue
			}
		}
		return max
	} else {
		return "Empty list"
	}
}

func OnlyStrings(w io.Writer, a ...any) []string {
	if len(a) != 0 {
		s := []string{}
		for _, v := range a {
			switch t := v.(type) {
			case string:
				s = append(s, t)
				_, _ = w.Write([]byte(t))
			}
		}
		return s
	} else {
		return nil
	}
}
