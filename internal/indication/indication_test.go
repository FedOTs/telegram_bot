package indication

import (
	"fmt"
	"testing"
	"time"
)

func TestStruct(t *testing.T) {
	ind1 := Indication{10, 12}
	ind2 := Indication{12, 14}
	ind3 := Indication{102, 122}
	ind4 := Indication{1, 5}
	fmt.Println(ind1)
	date := time.Now().UTC()
	water := Water{ind1, ind2, ind3, ind4, date, true}
	energie := Energie{ind1, ind2, date, true}
	gas := Gas{ind1, date, true}
	fmt.Println(water, energie, gas)
}
