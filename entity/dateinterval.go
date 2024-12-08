package entity

import (
	"applicationDesignTest/pkg"
	"time"
)

type DateInterval struct {
	From time.Time `json:"from"`
	To   time.Time `json:"to"`
}

func (d *DateInterval) Validate() error {
	if d.From.IsZero() || d.To.IsZero() {
		return pkg.ErrInvalidDataValue
	}

	return nil
}

func (d *DateInterval) GetInterval() []time.Time {
	i := make([]time.Time, 0)
	for nd := d.From; d.To.After(nd); nd = nd.AddDate(0, 0, 1) {
		i = append(i, nd)
	}
	i = append(i, d.To)

	return i
}
