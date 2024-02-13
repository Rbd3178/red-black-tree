package tree

import (
	"reflect"
	"testing"
)

func (tr *tree) fill() {
	tr.Insert(8215, "8215")
	tr.Insert(9676, "9676")
	tr.Insert(8191, "8191")
	tr.Insert(7130, "7130")
	tr.Insert(2114, "2114")
	tr.Insert(6613, "6613")
	tr.Insert(8369, "8369")
	tr.Insert(7381, "7381")
	tr.Insert(4557, "4557")
	tr.Insert(7673, "7673")
	tr.Insert(4842, "4842")
	tr.Insert(9301, "9301")
	tr.Insert(5699, "5699")
	tr.Insert(7735, "7735")
	tr.Insert(245, "245")
	tr.Insert(7127, "7127")
	tr.Insert(4793, "4793")
	tr.Insert(2005, "2005")
	tr.Insert(8301, "8301")
	tr.Insert(2649, "2649")
	tr.Insert(6473, "6473")
	tr.Insert(9605, "9605")
	tr.Insert(4606, "4606")
	tr.Insert(2161, "2161")
	tr.Insert(4505, "4505")
	tr.Insert(2984, "2984")
	tr.Insert(7833, "7833")
	tr.Insert(5809, "5809")
	tr.Insert(4756, "4756")
	tr.Insert(8129, "8129")
	tr.Insert(9129, "9129")
	tr.Insert(9832, "9832")
	tr.Insert(9000, "9000")
	tr.Insert(5616, "5616")
	tr.Insert(188, "188")
	tr.Insert(5602, "5602")
	tr.Insert(6276, "6276")
	tr.Insert(6650, "6650")
	tr.Insert(9292, "9292")
	tr.Insert(912, "912")
	tr.Insert(7966, "7966")
	tr.Insert(5322, "5322")
	tr.Insert(5669, "5669")
	tr.Insert(4175, "4175")
	tr.Insert(7779, "7779")
	tr.Insert(1845, "1845")
	tr.Insert(9006, "9006")
	tr.Insert(3981, "3981")
	tr.Insert(3267, "3267")
	tr.Insert(2038, "2038")
	tr.Insert(5764, "5764")
	tr.Insert(2718, "2718")
	tr.Insert(3975, "3975")
	tr.Insert(6516, "6516")
	tr.Insert(5639, "5639")
	tr.Insert(1812, "1812")
	tr.Insert(4098, "4098")
	tr.Insert(2949, "2949")
	tr.Insert(2600, "2600")
	tr.Insert(1020, "1020")
	tr.Insert(7882, "7882")
	tr.Insert(7620, "7620")
	tr.Insert(3538, "3538")
	tr.Insert(4850, "4850")
	tr.Insert(1073, "1073")
	tr.Insert(7435, "7435")
	tr.Insert(342, "342")
	tr.Insert(8433, "8433")
	tr.Insert(4231, "4231")
	tr.Insert(4766, "4766")
	tr.Insert(6811, "6811")
	tr.Insert(6495, "6495")
	tr.Insert(6658, "6658")
	tr.Insert(5097, "5097")
	tr.Insert(9950, "9950")
	tr.Insert(8433, "8433")
	tr.Insert(7944, "7944")
	tr.Insert(2086, "2086")
	tr.Insert(6721, "6721")
	tr.Insert(3810, "3810")
	tr.Insert(7204, "7204")
	tr.Insert(9058, "9058")
	tr.Insert(2011, "2011")
	tr.Insert(671, "671")
	tr.Insert(7889, "7889")
	tr.Insert(3161, "3161")
	tr.Insert(3669, "3669")
	tr.Insert(1082, "1082")
	tr.Insert(5630, "5630")
	tr.Insert(6738, "6738")
	tr.Insert(786, "786")
	tr.Insert(7235, "7235")
	tr.Insert(5929, "5929")
	tr.Insert(9171, "9171")
	tr.Insert(5570, "5570")
	tr.Insert(3616, "3616")
	tr.Insert(578, "578")
	tr.Insert(8823, "8823")
	tr.Insert(5290, "5290")
	tr.Insert(4311, "4311")
	// max: 9950
	// min: 188
}
func TestTreeAt(t *testing.T) {
	tr := New(reflect.TypeOf(""))
	tr.fill()
	var tests = []struct {
		name    string
		key     int
		want    string
		wantErr bool
	}{
		{"3975 should be found", 3975, "3975", false},
		{"6811 should be found", 6811, "6811", false},
		{"342 should be found", 342, "342", false},
		{"9950 (max) should be found", 9950, "9950", false},
		{"188 (min) should be found", 188, "188", false},
		{"9999 (larger than max) should not be found", 9999, "", true},
		{"100 (smaller than min) should not be found", 100, "", true},
		{"2800 (haven't been inserted) should not be found", 2800, "", true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := tr.At(tt.key)
			if (err != nil) != tt.wantErr {
				t.Errorf("tree.At() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("tree.At() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestTreeNext(t *testing.T) {
	tr := New(reflect.TypeOf(""))
	tr.fill()
	var tests = []struct {
		name    string
		key     int
		want    string
		wantErr bool
	}{
		{"3975 should be found", 3975, "3975", false},
		{"6811 should be found", 6811, "6811", false},
		{"342 should be found", 342, "342", false},
		{"9950 (max) should be found", 9950, "9950", false},
		{"188 (min) should be found", 188, "188", false},
		{"9999 (larger than max) should not be found", 9999, "", true},
		{"100 (smaller than min) should not be found", 100, "", true},
		{"2800 (haven't been inserted) should not be found", 2800, "", true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := tr.At(tt.key)
			if (err != nil) != tt.wantErr {
				t.Errorf("tree.At() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("tree.At() = %v, want %v", got, tt.want)
			}
		})
	}
}
