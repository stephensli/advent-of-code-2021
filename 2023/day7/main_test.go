package main

import (
	"reflect"
	"testing"
)

func Test_parser(t *testing.T) {
	type args struct {
	}
	tests := []struct {
		name      string
		input     []string
		partTwo   bool
		wantHands []cardHand
	}{{
		name:  "should correctly rank same types based on card values",
		input: []string{"77888 1", "77788 2"},
		wantHands: []cardHand{{
			values:   []string{"7", "7", "7", "8", "8"},
			cardType: fullHouseType,
			value:    2,
		}, {
			values:   []string{"7", "7", "8", "8", "8"},
			cardType: fullHouseType,
			value:    1,
		}},
	}, {
		name:    "should correctly handle p2",
		input:   []string{"7777J 1", "QJJQ2 1", "QQQQ2 1"},
		partTwo: true,
		wantHands: []cardHand{{
			values:   []string{"Q", "J", "J", "Q", "2"},
			cardType: fourOfAKindType,
			value:    1,
		},
			{
				values:   []string{"Q", "Q", "Q", "Q", "2"},
				cardType: fourOfAKindType,
				value:    1,
			},
			{
				values:   []string{"7", "7", "7", "7", "J"},
				cardType: fiveOfAKindType,
				value:    1,
			}},
	}}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if gotHands := parser(tt.input, tt.partTwo); !reflect.DeepEqual(gotHands, tt.wantHands) {
				t.Errorf("parser() = %v, want %v", gotHands, tt.wantHands)
			}
		})
	}
}
