package number

import (
	"strings"
	"testing"
)

func TestParseRanges(t *testing.T) {
	testCases := []struct {
		ranges string
		set    map[int]bool
	}{
		{
			ranges: "5",
			set: map[int]bool{
				1: false,
				2: false,
				3: false,
				4: false,
				5: true,
				6: false,
				7: false,
				8: false,
				9: false,
			},
		},
		{
			ranges: "5~",
			set: map[int]bool{
				1: false,
				2: false,
				3: false,
				4: false,
				5: true,
				6: true,
				7: true,
				8: true,
				9: true,
			},
		},
		{
			ranges: "~5",
			set: map[int]bool{
				1: true,
				2: true,
				3: true,
				4: true,
				5: true,
				6: false,
				7: false,
				8: false,
				9: false,
			},
		},
		{
			ranges: "3~6",
			set: map[int]bool{
				1: false,
				2: false,
				3: true,
				4: true,
				5: true,
				6: true,
				7: false,
				8: false,
				9: false,
			},
		},
		{
			ranges: "even",
			set: map[int]bool{
				1: false,
				2: true,
				3: false,
				4: true,
				5: false,
				6: true,
				7: false,
				8: true,
				9: false,
			},
		},
		{
			ranges: "odd",
			set: map[int]bool{
				1: true,
				2: false,
				3: true,
				4: false,
				5: true,
				6: false,
				7: true,
				8: false,
				9: true,
			},
		},
		{
			ranges: "2n",
			set: map[int]bool{
				1: false,
				2: true,
				3: false,
				4: true,
				5: false,
				6: true,
				7: false,
				8: true,
				9: false,
			},
		},
		{
			ranges: "2,odd,3~5,4n",
			set: map[int]bool{
				1: true,
				2: true,
				3: true,
				4: true,
				5: true,
				6: false,
				7: true,
				8: true,
				9: true,
			},
		},
	}

	for _, tc := range testCases {
		checkers := parseRanges(strings.Split(tc.ranges, ","))
		for i, expected := range tc.set {
			actual := inRanges(i, checkers)
			if actual != expected {
				msg := "Did not parse ranges correctory."
				t.Errorf("%s Parse '%s'.\nExpected: %d=>%v\nActual  : %d=>%v\n",
					msg, tc.ranges, i, expected, i, actual)
			}
		}
	}
}
