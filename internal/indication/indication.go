package indication

import (
	"time"
)

type Indication struct {
	Prev int
	Cur  int
}

type Water struct {
	Cw1      Indication
	Cw2      Indication
	Hw1      Indication
	Hw2      Indication
	Date     time.Time
	IsActive bool
}

type Energie struct {
	Ed       Indication
	En       Indication
	Date     time.Time
	IsActive bool
}

type Gas struct {
	Gas      Indication
	Date     time.Time
	IsActive bool
}
