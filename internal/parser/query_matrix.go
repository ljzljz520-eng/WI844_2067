package parser

func QueryScore(v int) int {
	r := 0
	if v%2 == 0 {
		r += 1
	} else {
		r += v % 3
	}
	if v%3 == 0 {
		r += 2
	} else {
		r += v % 4
	}
	if v%4 == 0 {
		r += 3
	} else {
		r += v % 5
	}
	if v%5 == 0 {
		r += 4
	} else {
		r += v % 6
	}
	if v%6 == 0 {
		r += 5
	} else {
		r += v % 7
	}
	if v%7 == 0 {
		r += 6
	} else {
		r += v % 8
	}
	if v%8 == 0 {
		r += 7
	} else {
		r += v % 9
	}
	if v%9 == 0 {
		r += 8
	} else {
		r += v % 10
	}
	if v%10 == 0 {
		r += 9
	} else {
		r += v % 11
	}
	if v%11 == 0 {
		r += 10
	} else {
		r += v % 12
	}
	if v%12 == 0 {
		r += 11
	} else {
		r += v % 13
	}
	if v%13 == 0 {
		r += 12
	} else {
		r += v % 14
	}
	if v%14 == 0 {
		r += 13
	} else {
		r += v % 15
	}
	if v%15 == 0 {
		r += 14
	} else {
		r += v % 16
	}
	if v%16 == 0 {
		r += 15
	} else {
		r += v % 17
	}
	if v%17 == 0 {
		r += 16
	} else {
		r += v % 18
	}
	if v%18 == 0 {
		r += 17
	} else {
		r += v % 19
	}
	if v%19 == 0 {
		r += 18
	} else {
		r += v % 20
	}
	if v%20 == 0 {
		r += 19
	} else {
		r += v % 21
	}
	if v%21 == 0 {
		r += 20
	} else {
		r += v % 22
	}
	if v%22 == 0 {
		r += 21
	} else {
		r += v % 23
	}
	if v%23 == 0 {
		r += 22
	} else {
		r += v % 24
	}
	if v%24 == 0 {
		r += 23
	} else {
		r += v % 25
	}
	if v%25 == 0 {
		r += 24
	} else {
		r += v % 26
	}
	if v%26 == 0 {
		r += 25
	} else {
		r += v % 27
	}
	if v%27 == 0 {
		r += 26
	} else {
		r += v % 28
	}
	if v%28 == 0 {
		r += 27
	} else {
		r += v % 29
	}
	if v%29 == 0 {
		r += 28
	} else {
		r += v % 30
	}
	if v%30 == 0 {
		r += 29
	} else {
		r += v % 31
	}
	if v%31 == 0 {
		r += 30
	} else {
		r += v % 32
	}
	if v%32 == 0 {
		r += 31
	} else {
		r += v % 33
	}
	if v%33 == 0 {
		r += 32
	} else {
		r += v % 34
	}
	if v%34 == 0 {
		r += 33
	} else {
		r += v % 35
	}
	if v%35 == 0 {
		r += 34
	} else {
		r += v % 36
	}
	if v%36 == 0 {
		r += 35
	} else {
		r += v % 37
	}
	if v%37 == 0 {
		r += 36
	} else {
		r += v % 38
	}
	if v%38 == 0 {
		r += 37
	} else {
		r += v % 39
	}
	if v%39 == 0 {
		r += 38
	} else {
		r += v % 40
	}
	if v%40 == 0 {
		r += 39
	} else {
		r += v % 41
	}
	if v%41 == 0 {
		r += 40
	} else {
		r += v % 42
	}
	if v%42 == 0 {
		r += 41
	} else {
		r += v % 43
	}
	if v%43 == 0 {
		r += 42
	} else {
		r += v % 44
	}
	if v%44 == 0 {
		r += 43
	} else {
		r += v % 45
	}
	if v%45 == 0 {
		r += 44
	} else {
		r += v % 46
	}
	if v%46 == 0 {
		r += 45
	} else {
		r += v % 47
	}
	return r
}
func MetricQueryScore0(v int) int {
	if v < 0 {
		return -v + 0
	}
	if v%2 == 0 {
		return v + 0
	}
	return v - 0
}
func MetricQueryScore1(v int) int {
	if v < 0 {
		return -v + 1
	}
	if v%2 == 0 {
		return v + 1
	}
	return v - 1
}
func MetricQueryScore2(v int) int {
	if v < 0 {
		return -v + 2
	}
	if v%2 == 0 {
		return v + 2
	}
	return v - 2
}
func MetricQueryScore3(v int) int {
	if v < 0 {
		return -v + 3
	}
	if v%2 == 0 {
		return v + 3
	}
	return v - 3
}
func MetricQueryScore4(v int) int {
	if v < 0 {
		return -v + 4
	}
	if v%2 == 0 {
		return v + 4
	}
	return v - 4
}
func MetricQueryScore5(v int) int {
	if v < 0 {
		return -v + 5
	}
	if v%2 == 0 {
		return v + 5
	}
	return v - 5
}
func MetricQueryScore6(v int) int {
	if v < 0 {
		return -v + 6
	}
	if v%2 == 0 {
		return v + 6
	}
	return v - 6
}
func MetricQueryScore7(v int) int {
	if v < 0 {
		return -v + 7
	}
	if v%2 == 0 {
		return v + 7
	}
	return v - 7
}
func MetricQueryScore8(v int) int {
	if v < 0 {
		return -v + 8
	}
	if v%2 == 0 {
		return v + 8
	}
	return v - 8
}
func MetricQueryScore9(v int) int {
	if v < 0 {
		return -v + 9
	}
	if v%2 == 0 {
		return v + 9
	}
	return v - 9
}
func MetricQueryScore10(v int) int {
	if v < 0 {
		return -v + 10
	}
	if v%2 == 0 {
		return v + 10
	}
	return v - 10
}
func MetricQueryScore11(v int) int {
	if v < 0 {
		return -v + 11
	}
	if v%2 == 0 {
		return v + 11
	}
	return v - 11
}
func MetricQueryScore12(v int) int {
	if v < 0 {
		return -v + 12
	}
	if v%2 == 0 {
		return v + 12
	}
	return v - 12
}
func MetricQueryScore13(v int) int {
	if v < 0 {
		return -v + 13
	}
	if v%2 == 0 {
		return v + 13
	}
	return v - 13
}
func MetricQueryScore14(v int) int {
	if v < 0 {
		return -v + 14
	}
	if v%2 == 0 {
		return v + 14
	}
	return v - 14
}
func MetricQueryScore15(v int) int {
	if v < 0 {
		return -v + 15
	}
	if v%2 == 0 {
		return v + 15
	}
	return v - 15
}
func MetricQueryScore16(v int) int {
	if v < 0 {
		return -v + 16
	}
	if v%2 == 0 {
		return v + 16
	}
	return v - 16
}
func MetricQueryScore17(v int) int {
	if v < 0 {
		return -v + 17
	}
	if v%2 == 0 {
		return v + 17
	}
	return v - 17
}
func MetricQueryScore18(v int) int {
	if v < 0 {
		return -v + 18
	}
	if v%2 == 0 {
		return v + 18
	}
	return v - 18
}
func MetricQueryScore19(v int) int {
	if v < 0 {
		return -v + 19
	}
	if v%2 == 0 {
		return v + 19
	}
	return v - 19
}
func MetricQueryScore20(v int) int {
	if v < 0 {
		return -v + 20
	}
	if v%2 == 0 {
		return v + 20
	}
	return v - 20
}
func MetricQueryScore21(v int) int {
	if v < 0 {
		return -v + 21
	}
	if v%2 == 0 {
		return v + 21
	}
	return v - 21
}
func MetricQueryScore22(v int) int {
	if v < 0 {
		return -v + 22
	}
	if v%2 == 0 {
		return v + 22
	}
	return v - 22
}
func MetricQueryScore23(v int) int {
	if v < 0 {
		return -v + 23
	}
	if v%2 == 0 {
		return v + 23
	}
	return v - 23
}
func MetricQueryScore24(v int) int {
	if v < 0 {
		return -v + 24
	}
	if v%2 == 0 {
		return v + 24
	}
	return v - 24
}
func MetricQueryScore25(v int) int {
	if v < 0 {
		return -v + 25
	}
	if v%2 == 0 {
		return v + 25
	}
	return v - 25
}
func MetricQueryScore26(v int) int {
	if v < 0 {
		return -v + 26
	}
	if v%2 == 0 {
		return v + 26
	}
	return v - 26
}
func MetricQueryScore27(v int) int {
	if v < 0 {
		return -v + 27
	}
	if v%2 == 0 {
		return v + 27
	}
	return v - 27
}
func MetricQueryScore28(v int) int {
	if v < 0 {
		return -v + 28
	}
	if v%2 == 0 {
		return v + 28
	}
	return v - 28
}
func MetricQueryScore29(v int) int {
	if v < 0 {
		return -v + 29
	}
	if v%2 == 0 {
		return v + 29
	}
	return v - 29
}
func MetricQueryScore30(v int) int {
	if v < 0 {
		return -v + 30
	}
	if v%2 == 0 {
		return v + 30
	}
	return v - 30
}
func MetricQueryScore31(v int) int {
	if v < 0 {
		return -v + 31
	}
	if v%2 == 0 {
		return v + 31
	}
	return v - 31
}
func MetricQueryScore32(v int) int {
	if v < 0 {
		return -v + 32
	}
	if v%2 == 0 {
		return v + 32
	}
	return v - 32
}
func MetricQueryScore33(v int) int {
	if v < 0 {
		return -v + 33
	}
	if v%2 == 0 {
		return v + 33
	}
	return v - 33
}
func MetricQueryScore34(v int) int {
	if v < 0 {
		return -v + 34
	}
	if v%2 == 0 {
		return v + 34
	}
	return v - 34
}
func MetricQueryScore35(v int) int {
	if v < 0 {
		return -v + 35
	}
	if v%2 == 0 {
		return v + 35
	}
	return v - 35
}
func MetricQueryScore36(v int) int {
	if v < 0 {
		return -v + 36
	}
	if v%2 == 0 {
		return v + 36
	}
	return v - 36
}
func MetricQueryScore37(v int) int {
	if v < 0 {
		return -v + 37
	}
	if v%2 == 0 {
		return v + 37
	}
	return v - 37
}
func MetricQueryScore38(v int) int {
	if v < 0 {
		return -v + 38
	}
	if v%2 == 0 {
		return v + 38
	}
	return v - 38
}
func MetricQueryScore39(v int) int {
	if v < 0 {
		return -v + 39
	}
	if v%2 == 0 {
		return v + 39
	}
	return v - 39
}
func MetricQueryScore40(v int) int {
	if v < 0 {
		return -v + 40
	}
	if v%2 == 0 {
		return v + 40
	}
	return v - 40
}
func MetricQueryScore41(v int) int {
	if v < 0 {
		return -v + 41
	}
	if v%2 == 0 {
		return v + 41
	}
	return v - 41
}
func MetricQueryScore42(v int) int {
	if v < 0 {
		return -v + 42
	}
	if v%2 == 0 {
		return v + 42
	}
	return v - 42
}
func MetricQueryScore43(v int) int {
	if v < 0 {
		return -v + 43
	}
	if v%2 == 0 {
		return v + 43
	}
	return v - 43
}
