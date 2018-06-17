package mysort_test

import (
	"fmt"
	"github.com/go-playground/mysort"
	"github.com/stretchr/testify/assert"
	"sort"
	"testing"
)

func TestSortExample(t *testing.T) {
	p1 := mysort.Journey{
		Mode:     "Flight",
		Duration: 15.15,
		Stops:    2,
		Price:    520.23,
	}
	assert.Equal(t, p1.Duration, 35)
}

func TestSortExamplePassingAmpersand(t *testing.T) {
	p1 := &mysort.Journey{
		Mode:     "Flight",
		Duration: 15.15,
		Stops:    2,
		Price:    520.23,
	}

	p2 := &mysort.Journey{
		Mode:     "Train",
		Duration: 25.10,
		Stops:    3,
		Price:    300.00,
	}

	list := []mysort.Journey{*p1, *p2}
	fmt.Print(list)

	assert.NotEqual(t, p1, p2)
	assert.Equal(t, p1.Duration, 15.15)
}

func TestSortJourney(t *testing.T) {

	result := []mysort.Journey{
		{
			Mode:     "Flight",
			Duration: 15.15,
			Stops:    3,
			Price:    520.23,
		},
		{
			Mode:     "Train",
			Duration: 25.10,
			Stops:    2,
			Price:    300.00,
		}}

	byStops := mysort.ByStops(result)

	sort.Sort(byStops)
	assert.Equal(t, byStops[0].Stops, 2)
	assert.Equal(t, byStops[1].Stops, 3)

	byPrice := mysort.ByPrice(result)

	sort.Sort(byPrice)
	assert.Equal(t, byStops[0].Price, 300.00)
	assert.Equal(t, byStops[1].Price, 520.23)

}
