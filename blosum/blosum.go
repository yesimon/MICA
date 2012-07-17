package blosum

var Alphabet62 = "ABCDEFGHIKLMNPQRSTVWXYZ"

var Matrix62 = [][]int{
	{4.0, -2.0, 0.0, -2.0, -1.0, -2.0, 0.0, -2.0, -1.0, -1.0, -1.0, -1.0, -2.0, -1.0, -1.0, -1.0, 1.0, 0.0, 0.0, -3.0, 0.0, -2.0, -1.0, -4.0},
	{-2.0, 4.0, -3.0, 4.0, 1.0, -3.0, -1.0, 0.0, -3.0, 0.0, -4.0, -3.0, 3.0, -2.0, 0.0, -1.0, 0.0, -1.0, -3.0, -4.0, -1.0, -3.0, 1.0, -4.0},
	{0.0, -3.0, 9.0, -3.0, -4.0, -2.0, -3.0, -3.0, -1.0, -3.0, -1.0, -1.0, -3.0, -3.0, -3.0, -3.0, -1.0, -1.0, -1.0, -2.0, -2.0, -2.0, -3.0, -4.0},
	{-2.0, 4.0, -3.0, 6.0, 2.0, -3.0, -1.0, -1.0, -3.0, -1.0, -4.0, -3.0, 1.0, -1.0, 0.0, -2.0, 0.0, -1.0, -3.0, -4.0, -1.0, -3.0, 1.0, -4.0},
	{-1.0, 1.0, -4.0, 2.0, 5.0, -3.0, -2.0, 0.0, -3.0, 1.0, -3.0, -2.0, 0.0, -1.0, 2.0, 0.0, 0.0, -1.0, -2.0, -3.0, -1.0, -2.0, 4.0, -4.0},
	{-2.0, -3.0, -2.0, -3.0, -3.0, 6.0, -3.0, -1.0, 0.0, -3.0, 0.0, 0.0, -3.0, -4.0, -3.0, -3.0, -2.0, -2.0, -1.0, 1.0, -1.0, 3.0, -3.0, -4.0},
	{0.0, -1.0, -3.0, -1.0, -2.0, -3.0, 6.0, -2.0, -4.0, -2.0, -4.0, -3.0, 0.0, -2.0, -2.0, -2.0, 0.0, -2.0, -3.0, -2.0, -1.0, -3.0, -2.0, -4.0},
	{-2.0, 0.0, -3.0, -1.0, 0.0, -1.0, -2.0, 8.0, -3.0, -1.0, -3.0, -2.0, 1.0, -2.0, 0.0, 0.0, -1.0, -2.0, -3.0, -2.0, -1.0, 2.0, 0.0, -4.0},
	{-1.0, -3.0, -1.0, -3.0, -3.0, 0.0, -4.0, -3.0, 4.0, -3.0, 2.0, 1.0, -3.0, -3.0, -3.0, -3.0, -2.0, -1.0, 3.0, -3.0, -1.0, -1.0, -3.0, -4.0},
	{-1.0, 0.0, -3.0, -1.0, 1.0, -3.0, -2.0, -1.0, -3.0, 5.0, -2.0, -1.0, 0.0, -1.0, 1.0, 2.0, 0.0, -1.0, -2.0, -3.0, -1.0, -2.0, 1.0, -4.0},
	{-1.0, -4.0, -1.0, -4.0, -3.0, 0.0, -4.0, -3.0, 2.0, -2.0, 4.0, 2.0, -3.0, -3.0, -2.0, -2.0, -2.0, -1.0, 1.0, -2.0, -1.0, -1.0, -3.0, -4.0},
	{-1.0, -3.0, -1.0, -3.0, -2.0, 0.0, -3.0, -2.0, 1.0, -1.0, 2.0, 5.0, -2.0, -2.0, 0.0, -1.0, -1.0, -1.0, 1.0, -1.0, -1.0, -1.0, -1.0, -4.0},
	{-2.0, 3.0, -3.0, 1.0, 0.0, -3.0, 0.0, 1.0, -3.0, 0.0, -3.0, -2.0, 6.0, -2.0, 0.0, 0.0, 1.0, 0.0, -3.0, -4.0, -1.0, -2.0, 0.0, -4.0},
	{-1.0, -2.0, -3.0, -1.0, -1.0, -4.0, -2.0, -2.0, -3.0, -1.0, -3.0, -2.0, -2.0, 7.0, -1.0, -2.0, -1.0, -1.0, -2.0, -4.0, -2.0, -3.0, -1.0, -4.0},
	{-1.0, 0.0, -3.0, 0.0, 2.0, -3.0, -2.0, 0.0, -3.0, 1.0, -2.0, 0.0, 0.0, -1.0, 5.0, 1.0, 0.0, -1.0, -2.0, -2.0, -1.0, -1.0, 3.0, -4.0},
	{-1.0, -1.0, -3.0, -2.0, 0.0, -3.0, -2.0, 0.0, -3.0, 2.0, -2.0, -1.0, 0.0, -2.0, 1.0, 5.0, -1.0, -1.0, -3.0, -3.0, -1.0, -2.0, 0.0, -4.0},
	{1.0, 0.0, -1.0, 0.0, 0.0, -2.0, 0.0, -1.0, -2.0, 0.0, -2.0, -1.0, 1.0, -1.0, 0.0, -1.0, 4.0, 1.0, -2.0, -3.0, 0.0, -2.0, 0.0, -4.0},
	{0.0, -1.0, -1.0, -1.0, -1.0, -2.0, -2.0, -2.0, -1.0, -1.0, -1.0, -1.0, 0.0, -1.0, -1.0, -1.0, 1.0, 5.0, 0.0, -2.0, 0.0, -2.0, -1.0, -4.0},
	{0.0, -3.0, -1.0, -3.0, -2.0, -1.0, -3.0, -3.0, 3.0, -2.0, 1.0, 1.0, -3.0, -2.0, -2.0, -3.0, -2.0, 0.0, 4.0, -3.0, -1.0, -1.0, -2.0, -4.0},
	{-3.0, -4.0, -2.0, -4.0, -3.0, 1.0, -2.0, -2.0, -3.0, -3.0, -2.0, -1.0, -4.0, -4.0, -2.0, -3.0, -3.0, -2.0, -3.0, 11.0, -2.0, 2.0, -3.0, -4.0},
	{0.0, -1.0, -2.0, -1.0, -1.0, -1.0, -1.0, -1.0, -1.0, -1.0, -1.0, -1.0, -1.0, -2.0, -1.0, -1.0, 0.0, 0.0, -1.0, -2.0, -1.0, -1.0, -1.0, -4.0},
	{-2.0, -3.0, -2.0, -3.0, -2.0, 3.0, -3.0, 2.0, -1.0, -2.0, -1.0, -1.0, -2.0, -3.0, -1.0, -2.0, -2.0, -2.0, -1.0, 2.0, -1.0, 7.0, -2.0, -4.0},
	{-1.0, 1.0, -3.0, 1.0, 4.0, -3.0, -2.0, 0.0, -3.0, 1.0, -3.0, -1.0, 0.0, -1.0, 3.0, 0.0, 0.0, -1.0, -2.0, -3.0, -1.0, -2.0, 4.0, -4.0},
	{-4.0, -4.0, -4.0, -4.0, -4.0, -4.0, -4.0, -4.0, -4.0, -4.0, -4.0, -4.0, -4.0, -4.0, -4.0, -4.0, -4.0, -4.0, -4.0, -4.0, -4.0, -4.0, -4.0, 0.0},
}
