package mysort

import "fmt"

type Journey struct {
	Mode     string
	Duration float32
	Stops    int
	Price    float64
}

func (p Journey) Display() string {
	return fmt.Sprintf("%s:%d", p.Mode, p.Duration)
}


type ByStops []Journey

func (a ByStops) Len() int {
	return len(a)
}

func (a ByStops) Less(i, j int) bool {
	return a[i].Stops < a[j].Stops
}

func (a ByStops) Swap(i, j int) {
	a[i], a[j] = a[j], a[i]
}


type ByPrice []Journey

func (a ByPrice) Len() int {
	return len(a)
}

func (a ByPrice) Less(i, j int) bool {
	return a[i].Price < a[j].Price
}

func (a ByPrice) Swap(i, j int) {
	a[i], a[j] = a[j], a[i]
}



