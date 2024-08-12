package tax

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestTax(t *testing.T) {

	price := Tax(1000)
	if price != 80 {
		t.Error("Expected 80, got ", price)
	}
}

func TestTax2(t *testing.T) {
	price := Tax(1000)
	assert.Equal(t, 80.0, price, "they should be equal")

	price = Tax(1001)
	assert.Equal(t, 72.072, price, "they should be equal")
}
