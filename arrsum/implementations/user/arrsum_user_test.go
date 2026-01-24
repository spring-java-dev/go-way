package user

import (
	"testing"

	"github.com/helios16/go-way/arrsum"
)

func TestArrSum(t *testing.T) {
	for _, tc := range arrsum.BaseTestCases {
		t.Run(tc.Name, func(t *testing.T) {
			got := ArrSum(tc.Input...)
			if got != tc.Want {
				t.Errorf("ArrSum(%v) = %d, want %d", tc.Input, got, tc.Want)
			}
		})
	}
}
