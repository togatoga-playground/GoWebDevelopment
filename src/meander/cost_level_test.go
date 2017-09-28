package meander_test

import (
	"testing"
	is2 "github.com/cheekybits/is"
	"meander"
)

func TestCostValues(t *testing.T)  {
	is := is2.New(t)
	is.Equal(int(meander.Cost1), 1)
	is.Equal(int(meander.Cost2), 2)
	is.Equal(int(meander.Cost3), 3)
	is.Equal(int(meander.Cost4), 4)
	is.Equal(int(meander.Cost5), 5)
}