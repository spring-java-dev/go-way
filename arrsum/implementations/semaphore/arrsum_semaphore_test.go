package semaphore

import (
	"testing"

	"github.com/helios16/go-way/arrsum"
)

func TestArrSumSemaphore(t *testing.T) {
	for _, tt := range arrsum.BaseTestCases {
		t.Run(tt.Name, func(t *testing.T) {
			got := ArrSum(tt.Input...)
			if got != tt.Want {
				t.Errorf("ArrSum(%v) = %d, want %d", tt.Input, got, tt.Want)
			}
		})
	}
}
