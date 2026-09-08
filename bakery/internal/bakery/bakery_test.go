package bakery

import (
	"context"
	"testing"

	"go.uber.org/goleak"
)

type testCase struct {
	name           string
	k              int
	n              int
	m              int
	t1             int
	t2             int
	expectedResult int
}

func TestBakery(t *testing.T) {
	defer goleak.VerifyNone(t)

	testCases := []testCase{
		{
			name:           "Standard pipeline execution",
			k:              10,
			n:              3,
			m:              2,
			t1:             1,
			t2:             1,
			expectedResult: 10,
		},
		{
			name:           "Zero cakes requested (k=0)",
			k:              0,
			n:              2,
			m:              2,
			t1:             1,
			t2:             1,
			expectedResult: 0,
		},
		{
			name:           "Single worker pipeline (n=1, m=1)",
			k:              5,
			n:              1,
			m:              1,
			t1:             1,
			t2:             1,
			expectedResult: 5,
		},
		{
			name:           "More workers than cakes (n > k, m > k)",
			k:              3,
			n:              10,
			m:              10,
			t1:             1,
			t2:             1,
			expectedResult: 3,
		},
		{
			name:           "Zero delays (t1=0, t2=0)",
			k:              20,
			n:              4,
			m:              4,
			t1:             0,
			t2:             0,
			expectedResult: 20,
		},
		{
			name:           "High concurrency load",
			k:              100,
			n:              20,
			m:              20,
			t1:             1,
			t2:             1,
			expectedResult: 100,
		},
	}

	for _, tt := range testCases {
		t.Run(tt.name, func(t *testing.T) {
			ctx := context.Background()

			outBake := Bake(ctx, tt.k, tt.n, tt.t1)
			outPack := Pack(ctx, tt.m, tt.t2, outBake)
			cakes := Inspect(outPack)

			if len(cakes) != tt.expectedResult {
				t.Errorf("expected %d cakes, got %d", tt.expectedResult, len(cakes))
			}
		})
	}
}
