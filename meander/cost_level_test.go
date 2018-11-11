package meander_test

import (
	"testing"

	"github.com/aikuma0130/goWeb/meander"
	"github.com/cheekybits/is"
)

func TestCostValues(t *testing.T) {
	is := is.New(t)
	is.Equal(int(meander.Cost1), 1)
	is.Equal(int(meander.Cost2), 2)
	is.Equal(int(meander.Cost3), 3)
	is.Equal(int(meander.Cost4), 4)
	is.Equal(int(meander.Cost5), 5)
}
