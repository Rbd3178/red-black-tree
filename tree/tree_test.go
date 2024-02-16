package tree

import (
	"testing"
)

// newFilledTree returns a pointer to a Tree,
// filled with 99 string values with int keys.
// Max key: 9950. Min key: 188.
func newFilledTree() *Tree[int, string] {
	var tr Tree[int, string]
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
	return &tr
}

func TestTreeInsert(t *testing.T) {
	var tr Tree[int, string]
	var tests = []struct {
		name    string
		key     int
		val     string
		wantErr bool
	}{
		{"inserting 100->100", 100, "100", false},
		{"inserting 51->51", 51, "51", false},
		{"inserting 42->42", 42, "42", false},
		{"inserting 65->42", 65, "42", false},
		{"inserting 51 again", 51, "51", true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := tr.Insert(tt.key, tt.val)
			if (err != nil) != tt.wantErr {
				t.Errorf("tree.Insert() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
		})
	}
}

func TestTreeAt(t *testing.T) {
	trFull := newFilledTree()
	var trEmpty Tree[int, string]
	var tests = []struct {
		name    string
		tree    *Tree[int, string]
		key     int
		want    string
		wantErr bool
	}{
		{"3975 should be found", trFull, 3975, "3975", false},
		{"6811 should be found", trFull, 6811, "6811", false},
		{"342 should be found", trFull, 342, "342", false},
		{"9950 (max) should be found", trFull, 9950, "9950", false},
		{"188 (min) should be found", trFull, 188, "188", false},
		{"9999 (larger than max) should not be found", trFull, 9999, "", true},
		{"100 (smaller than min) should not be found", trFull, 100, "", true},
		{"2800 (haven't been inserted) should not be found", trFull, 2800, "", true},
		{"nothing should be found in empty tree", &trEmpty, 42, "", true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := tt.tree.At(tt.key)
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

func TestTreeAssign(t *testing.T) {
	trFull := newFilledTree()
	var trEmpty Tree[int, string]
	var tests = []struct {
		name    string
		tree    *Tree[int, string]
		key     int
		val     string
		want    string
		wantErr bool
	}{
		{"assign \"test\" to key 4606", trFull, 4606, "test", "test", false},
		{"try assigning to non-existing key", trFull, 9999, "test", "test", true},
		{"try assigning in an empty tree", &trEmpty, 10, "test", "test", true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := tt.tree.Assign(tt.key, tt.val)
			if (err != nil) != tt.wantErr {
				t.Errorf("tree.Assign() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if err == nil {
				got, _ := tt.tree.At(tt.key)
				if got != tt.want {
					t.Errorf("tree.At() = %v, want %v", got, tt.want)
				}
			}
		})
	}
}

func TestTreeSize(t *testing.T) {
	trFull := newFilledTree()
	var trEmpty Tree[int, string]
	var tests = []struct {
		name    string
		tree    *Tree[int, string]
		want    int
		wantErr bool
	}{
		{"Size of a filled tree should be 99", trFull, 99, false},
		{"Size of an empty tree should be 0", &trEmpty, 0, false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := tt.tree.Size()
			if got != tt.want {
				t.Errorf("tree.Size() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestTreeMax(t *testing.T) {
	trFull := newFilledTree()
	var trEmpty Tree[int, string]
	var tests = []struct {
		name    string
		tree    *Tree[int, string]
		want    int
		wantErr bool
	}{
		{"Max key should be 9950", trFull, 9950, false},
		{"Error in case of empty tree", &trEmpty, 0, true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, _, err := tt.tree.Max()
			if (err != nil) != tt.wantErr {
				t.Errorf("tree.Max() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("tree.Max() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestTreeMin(t *testing.T) {
	trFull := newFilledTree()
	var trEmpty Tree[int, string]
	var tests = []struct {
		name    string
		tree    *Tree[int, string]
		want    int
		wantErr bool
	}{
		{"Min key should be 188", trFull, 188, false},
		{"Error in case of empty tree", &trEmpty, 0, true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, _, err := tt.tree.Min()
			if (err != nil) != tt.wantErr {
				t.Errorf("tree.Min() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("tree.Min() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestTreeNext(t *testing.T) {
	trFull := newFilledTree()
	var trEmpty Tree[int, string]
	var tests = []struct {
		name    string
		tree    *Tree[int, string]
		key     int
		want    int
		wantErr bool
	}{
		{"First key lagrer than existing 2161 should be 2600", trFull, 2161, 2600, false},
		{"First key lagrer than 9600 should be 9605", trFull, 9600, 9605, false},
		{"Should be no keys larger than existing 9950", trFull, 9950, 0, true},
		{"Should be no keys larger than 10000", trFull, 10000, 0, true},
		{"Should be no larger keys than 42 in empty tree", &trEmpty, 42, 0, true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, _, err := tt.tree.Next(tt.key)
			if (err != nil) != tt.wantErr {
				t.Errorf("tree.Next() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("tree.Next() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestTreePrev(t *testing.T) {
	trFull := newFilledTree()
	var trEmpty Tree[int, string]
	var tests = []struct {
		name    string
		tree    *Tree[int, string]
		key     int
		want    int
		wantErr bool
	}{
		{"First key smaller than existing 1845 should be 1812", trFull, 1845, 1812, false},
		{"First key smaller than 7200 should be 7130", trFull, 7200, 7130, false},
		{"Should be no keys smaller than existing 188", trFull, 188, 0, true},
		{"Should be no keys smaller than 100", trFull, 100, 0, true},
		{"Should be no smaller keys than 42 in empty tree", &trEmpty, 42, 0, true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, _, err := tt.tree.Prev(tt.key)
			if (err != nil) != tt.wantErr {
				t.Errorf("tree.Prev() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("tree.Prev() = %v, want %v", got, tt.want)
			}
		})
	}
}
