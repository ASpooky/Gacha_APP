package utils

//1~5のレアリティとして
//1:40%
//2:30%
//3:20%
//4:8%
//5:2%

func NumToRarity(num int) int {
	if num < 40 {
		return 1
	} else if num < 70 {
		return 2
	} else if num < 90 {
		return 3
	} else if num < 98 {
		return 4
	} else if num <= 100 {
		return 5
	} else {
		return 0
	}
}
