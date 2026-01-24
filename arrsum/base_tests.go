package arrsum

// BaseTestCases contains common test cases for all ArrSum implementations
var BaseTestCases = []struct {
	Name  string
	Input [][]int
	Want  int
}{
	{"single array positive", [][]int{{1, 2, 3, 4, 5}}, 15},
	{"single array negative", [][]int{{-1, -2, -3}}, -6},
	{"single zero", [][]int{{0}}, 0},
	{"single large number", [][]int{{100}}, 100},
	{"canceling numbers", [][]int{{-5, 5}}, 0},
	{"two small arrays", [][]int{{1, 2}, {3, 4}}, 10},
	{"nil input", nil, 0},
	{"large numbers", [][]int{{1000000, 2000000, 3000000}}, 6000000},
	{"three arrays", [][]int{{100, 200}, {300, 400}, {500, 600}}, 2100},
	{"four arrays sequential", [][]int{{1, 2, 3}, {4, 5, 6}, {7, 8, 9}, {10, 11, 12}}, 78},
	{"mixed large numbers zero sum", [][]int{{-1000000}, {1000000}, {500000}, {-500000}}, 0},
	{"five arrays 1-25", [][]int{{1, 2, 3, 4, 5}, {6, 7, 8, 9, 10}, {11, 12, 13, 14, 15}, {16, 17, 18, 19, 20}, {21, 22, 23, 24, 25}}, 325},
	{"five arrays mixed values", [][]int{{-1, -2, -3, -4, -5}, {10, 20, 30, 40, 50}, {100, 200, 300, 400, 500}, {-10, -20, -30, -40, -50}, {5, 15, 25, 35, 45}}, 1610},
}
