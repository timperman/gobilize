package mobilize

import (
	"strconv"
	"time"
)

type Time time.Time

func (t Time) MarshalJSON() ([]byte, error) {
	return []byte(strconv.FormatInt(time.Time(t).Unix(), 10)), nil
}

func (t *Time) UnmarshalJSON(s []byte) (err error) {
	r := string(s)
	q, err := strconv.ParseInt(r, 10, 64)
	if err != nil {
		return err
	}
	*(*time.Time)(t) = time.Unix(q, 0)
	return nil
}

func (t Time) Unix() int64 {
	return time.Time(t).Unix()
}

func (t Time) Time() time.Time {
	return time.Time(t).Local()
}

func (t Time) String() string {
	return t.Time().String()
}
