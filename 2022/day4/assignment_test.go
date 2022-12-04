package main

import "testing"

func Test_assigment_Contains(t *testing.T) {
	type fields struct {
		rangeLeft  int
		rangeRight int
	}
	type args struct {
		pair assigment
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   bool
	}{{
		name:   "should return true if assignment contains provider that is a single value",
		fields: fields{rangeLeft: 4, rangeRight: 6},
		args:   args{pair: assigment{rangeLeft: 6, rangeRight: 6}},
		want:   true,
	}, {
		name:   "should return true if assignment contains provider that is equal in range",
		fields: fields{rangeLeft: 4, rangeRight: 5},
		args:   args{pair: assigment{rangeLeft: 4, rangeRight: 5}},
		want:   true,
	}, {
		name:   "should return true if assignment contains provider that is a sub section",
		fields: fields{rangeLeft: 4, rangeRight: 8},
		args:   args{pair: assigment{rangeLeft: 5, rangeRight: 7}},
		want:   true,
	}, {
		name:   "should return false if entirely out of range (higher)",
		fields: fields{rangeLeft: 4, rangeRight: 8},
		args:   args{pair: assigment{rangeLeft: 10, rangeRight: 12}},
		want:   false,
	}, {
		name:   "should return false if entirely out of range (lower)",
		fields: fields{rangeLeft: 4, rangeRight: 8},
		args:   args{pair: assigment{rangeLeft: 1, rangeRight: 2}},
		want:   false,
	}, {
		name:   "should return false if partially out of range (higher)",
		fields: fields{rangeLeft: 4, rangeRight: 8},
		args:   args{pair: assigment{rangeLeft: 6, rangeRight: 9}},
		want:   false,
	}, {
		name:   "should return false if partially out of range (lower)",
		fields: fields{rangeLeft: 4, rangeRight: 8},
		args:   args{pair: assigment{rangeLeft: 2, rangeRight: 5}},
		want:   false,
	}}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			a := assigment{
				rangeLeft:  tt.fields.rangeLeft,
				rangeRight: tt.fields.rangeRight,
			}
			if got := a.Contains(tt.args.pair); got != tt.want {
				t.Errorf("Contains() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_assigment_Overlap(t *testing.T) {
	type fields struct {
		rangeLeft  int
		rangeRight int
	}
	type args struct {
		pair assigment
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   bool
	}{{
		name:   "should return true if entirely out of range (higher)",
		fields: fields{rangeLeft: 4, rangeRight: 8},
		args:   args{pair: assigment{rangeLeft: 10, rangeRight: 12}},
		want:   false,
	}, {
		name:   "should return true if entirely out of range (lower)",
		fields: fields{rangeLeft: 4, rangeRight: 8},
		args:   args{pair: assigment{rangeLeft: 1, rangeRight: 2}},
		want:   false,
	}, {
		name:   "should return true if partially out of range (higher)",
		fields: fields{rangeLeft: 4, rangeRight: 8},
		args:   args{pair: assigment{rangeLeft: 6, rangeRight: 9}},
		want:   true,
	}, {
		name:   "should return true if partially out of range (lower)",
		fields: fields{rangeLeft: 4, rangeRight: 8},
		args:   args{pair: assigment{rangeLeft: 2, rangeRight: 5}},
		want:   true,
	}}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			a := assigment{
				rangeLeft:  tt.fields.rangeLeft,
				rangeRight: tt.fields.rangeRight,
			}
			if got := a.Overlap(tt.args.pair); got != tt.want {
				t.Errorf("Contains() = %v, want %v", got, tt.want)
			}
		})
	}
}
