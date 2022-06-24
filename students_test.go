package coverage

import (
	"errors"
	"os"
	"reflect"
	"sort"
	"strconv"
	"testing"
	"time"
)

var (
	date = time.Date(2006, 01, 02, 15, 04, 05, 0, time.UTC)
)

// DO NOT EDIT THIS FUNCTION
func init() {
	content, err := os.ReadFile("students_test.go")
	if err != nil {
		panic(err)
	}
	err = os.WriteFile("autocode/students_test", content, 0644)
	if err != nil {
		panic(err)
	}
}

// WRITE YOUR CODE BELOW
func TestPeopleLenOK(t *testing.T) {
	
	tests := map[string]struct {
		arg People
		want int
	}{
		"nil input": {nil, 0},
		"zero people" : {make(People, 0), 0},
		"three people" : {make(People, 3), 3},
	}
	
	for tn, tt := range tests {
		t.Run(tn, func(t *testing.T) {
			got := tt.arg.Len()
			if got != tt.want {
				t.Errorf("want %v but got %v", tt.want, got)
			}
		})
	}
}

func TestPeopleLess(t *testing.T) {
	
	i, j := 0, 1

	tests := map[string]struct{
		arg  People
		want bool		
	}{
		"different lastName" : {
			People{
				Person{lastName: "A"},
				Person{lastName: "Z"},
			},
			true,
		},
		"different firstName" : {
			People{
				Person{firstName: "Z"},
				Person{firstName: "A"},
			},
			false,
		},
		"different birthday" : {
			People{
				Person{birthDay: date},
				Person{birthDay: date.AddDate(20, 0, 0)},
			},
			false,
		},
		"same LastName" : {
			People{
				Person{firstName: "A", birthDay: date.AddDate(20, 0, 0)},
				Person{firstName: "Z", birthDay: date},
			},
			true,
		},
		"same firstName" : {
			People{
				Person{lastName: "A", birthDay: date},
				Person{lastName: "Z", birthDay: date.AddDate(20, 0, 0)},
			},
			false,
		},
		"same birthday" : {
			People{
				Person{firstName: "Z", lastName: "A"},
				Person{firstName: "A", lastName: "Z"},
			},
			false,
		},
	}

	for tn, tt := range tests {
		t.Run(tn, func(t *testing.T) {
			got := tt.arg.Less(i, j)
			if tt.want != got {
				t.Errorf("want %v but got %v\n", tt.want, got)
			}
		})
	}
}

func TestPeopleSwap(t *testing.T) {

	type arg struct {
		People
		i, j int
	}

	pi := Person{"Andrej", "Żuławski", date}
	pj := Person{"小林", "正樹", date.AddDate(20, 0, 0)}
	tests := map[string]struct {
		arg arg		
		want People
	}{
		"swap person" : {
			arg: arg{People{pi, pj}, 0,  1},
			want : People{pj, pi},
		},
	}
	for tn, tt := range tests {
		t.Run(tn, func(t *testing.T) {
			tt.arg.Swap(tt.arg.i, tt.arg.j)
			if !reflect.DeepEqual(tt.want, tt.arg.People) {
				t.Errorf("want %v but got %v\n", tt.want, tt.arg.People)
			}
		})
	}	
}


func TestPeopleSort(t *testing.T) {

	p := Person{"Jean-Luc", "Godard", date}

	tests := map[string]struct {
		arg People
		want People
	}{
		"identical" : {
			want: People{p, p, p}, 
			arg: People{p, p, p},
		},
		"different lastName": {
			want: People{
				Person{lastName: "Kubrick"},
				Person{lastName: "Lynch"},
				Person{lastName: "Żuławski"},
				Person{lastName: "正樹"},
			},
			arg: People{
				Person{lastName: "正樹"},
				Person{lastName: "Kubrick"},
				Person{lastName: "Żuławski"},
				Person{lastName: "Lynch"},
			},
		},
		"different firstName" : {
			want: People{
				Person{firstName: "Andrzej"},
				Person{firstName: "David"},
				Person{firstName: "Stanley"},
				Person{firstName: "小林"},
			},
			arg: People{
				Person{firstName: "Stanley"},
				Person{firstName: "小林"},
				Person{firstName: "Andrzej"},
				Person{firstName: "David" },
			},
		},
		"different birthday" : {
			want: People{
				Person{birthDay: p.birthDay.AddDate(20, 7, 3)},
				Person{birthDay: p.birthDay.AddDate(20, 7, 2)},
				Person{birthDay: p.birthDay.Add(time.Hour * 5)},
				Person{birthDay: p.birthDay.Add(time.Millisecond)},
			},
			arg: People{
				Person{birthDay: p.birthDay.AddDate(20, 7, 2)},
				Person{birthDay: p.birthDay.Add(time.Millisecond)},
				Person{birthDay: p.birthDay.Add(time.Hour * 5)},
				Person{birthDay: p.birthDay.AddDate(20, 7, 3)},
			},
		},
	}

	for tn, tt := range tests {
		t.Run(tn, func(t *testing.T) {
			sort.Stable(tt.arg)
			if !reflect.DeepEqual(tt.want, tt.arg) {
				t.Errorf("want %v but got %v\n", tt.want, tt.arg)
			}
		})
	}

}


func TestMatrixNew(t *testing.T) {
	
	var (
		atoiStr = "zero and one"
		_, atoiErr = strconv.Atoi("zero")
		_, atoiErrEmptyString = strconv.Atoi("")
		shapeErr = errors.New("Rows need to be the same length")
	)
	
	type result struct {
		m *Matrix
		err error
	}

	tests := map[string]struct {
		arg string
		want result
	}{
		"ok" : {
			arg : "1 0 0\n0 1 0\n0 0 1",
			want: result{
				&Matrix{3, 3, 
					[]int{
						1, 0, 0,
						0, 1, 0,
						0, 0, 1,
					}},
				nil,
			},
		},
		"not square matrix" : {
			arg : "1\n0 1\n0 0 1\n",
			want: result{
				nil,
				shapeErr,
			},
		},
		"empty string" : {
			arg : *new(string),
			want: result{
				nil,
				atoiErrEmptyString,
			},
		},
		"want atoi error" : {
			arg : atoiStr,
			want: result{
				nil,
				atoiErr,
			},
		},
	}

	for tn, tt := range tests {
		t.Run(tn, func(t *testing.T) {
			m, err := New(tt.arg)
			if err != nil && err.Error() != tt.want.err.Error(){
					t.Errorf("err: want %v but got %v", tt.want.err, err)
			} else {
				if m != nil && tt.want.m != nil {
					if !reflect.DeepEqual(*tt.want.m, *m) {
						t.Errorf("Matrix.data: want %v but got %v", *tt.want.m, *m)
					}
				}
			}
		})
	}
}



func TestMatrixRows(t *testing.T) {
	tests := map[string]struct {
		arg Matrix
		want [][]int
	}{
		"ok" : {
			Matrix{3, 3, 
				[]int{
					1, 0, 0,
					0, 1, 0,
					0, 0, 1,
				}},
			[][]int{
				{1, 0, 0},
				{0, 1, 0},
				{0, 0, 1},
			},	
		},
	}

	for tn, tt := range tests {
		t.Run(tn, func(t *testing.T) {
			got := tt.arg.Rows()
			if !reflect.DeepEqual(tt.want, got) {
				t.Errorf("want %v but got %v", tt.want, got)
			}
		})		
	}
}

func TestMatrixCols(t *testing.T) {
	tests := map[string]struct {
		arg Matrix
		want [][]int
	}{
		"ok" : {
			Matrix{3, 3, 
				[]int{
					1, 2, 3,
					4, 5, 6,
					7, 8, 9,
				}},
			[][]int{
				{1, 4, 7},
				{2, 5, 8},
				{3, 6, 9},
			},	
		},
	}

	for tn, tt := range tests {
		t.Run(tn, func(t *testing.T) {
			got := tt.arg.Cols()
			if !reflect.DeepEqual(tt.want, got) {
				t.Errorf("want %v but got %v", tt.want, got)
			}
		})		
	}
}

func TestMatrixSet(t *testing.T) {
	
	type (
		arg struct {
			m Matrix
			i, j, v int
		}
		result struct {
			flag bool
			m Matrix
		}
	)

	m := Matrix {
		3, 3,
		[]int {
			1, 2, 3,
			4, 5, 6,
			7, 8, 9,
		},
	}

	tests := map[string]struct {
		data arg
		want result
	}{
		"ok" : {
			data : arg{m, 0, 1, -1},
			want : result{
				true,
				Matrix {
					3, 3,
					[]int {
					1, -1, 3,
					4, 5, 6,
					7, 8, 9,
					},
				},
			},
		},
		"out of bounds 1" : {
			data : arg{m, 100, 100, -1},
			want : result{
				false,
				Matrix {
					3, 3,
					[]int {
					1, -1, 3,
					4, 5, 6,
					7, 8, 9,
					},
				},
			},
		},
		"out of bounds 2" : {
			data : arg{m, -100, -100, -1},
			want : result{
				false,
				Matrix {
					3, 3,
					[]int {
					1, -1, 3,
					4, 5, 6,
					7, 8, 9,
					},
				},
			},
		},
	}

	for tn, tt := range tests {
		t.Run(tn, func(t *testing.T) {
			got := tt.data.m.Set(tt.data.i, tt.data.j, tt.data.v)
			if tt.want.flag != got {
				t.Errorf("flag: got %v but want %v", got, tt.want.flag)
			}
			if !reflect.DeepEqual(tt.want.m, m) {
				t.Errorf("Matrix: got %v but want %v", m, tt.want.m)
			}
		})
	}
}