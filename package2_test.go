package package2

import (
	"testing"
)

func TestSell(t *testing.T) {
	if Sell(4, 2) != 2 {
		t.Error("sai")
	}

}
