package collectroes

type Days struct {
	collector map[int]map[int][]float64
}

func NewDays() *Days {
	return &Days{
		collector: make(map[int]map[int][]float64),
	}
}

func (d *Days)Push(day, hour int, temp float64) {
	if _, ok := d.collector[day]; !ok {
		d.collector[day] = make(map[int][]float64)
	}
	dayMap := d.collector[day]
	if _, ok := dayMap[hour]; !ok {
		dayMap[hour] = make([]float64, 0)
	}
	dayMap[hour] = append(dayMap[hour], temp)
}
