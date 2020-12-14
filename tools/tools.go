package tools

import (
	"log"
	"runtime"
	"strconv"
	"strings"
	"time"
)

// VerifyIDCard 校验身份证号是否合法
func VerifyIDCard(IDCard string) bool {
	var idCardBytes = []byte(strings.ToUpper(IDCard))
	if len(idCardBytes) != 18 {
		return false
	}

	var weight = [...]int{7, 9, 10, 5, 8, 4, 2, 1, 6, 3, 7, 9, 10, 5, 8, 4, 2}
	var validate = [...]byte{49, 48, 88, 57, 56, 55, 54, 53, 52, 51, 50}
	var sum int
	for i := 0; i < len(idCardBytes)-1; i++ {
		b, err := strconv.Atoi(string(idCardBytes[i]))
		if err != nil {
			return false
		}
		sum += b * weight[i]
	}
	return validate[sum%11] == idCardBytes[17]
}

// TimeStep 按步分割时间
func TimeStep(start, end time.Time, step time.Duration) func() (time.Time, time.Time, bool) {
	asc := start.Before(end)
	if (asc && step > 0) || (!asc && step < 0 && !start.Equal(end)) {
		start = start.Add(-1 * step)
		return func() (time.Time, time.Time, bool) {
			start = start.Add(step)
			newend := start.Add(step)
			if newend.Equal(end) || (asc && newend.After(end)) || (!asc && newend.Before(end)) {
				return start, end, false
			}
			return start, newend, true
		}
	}

	return func() (time.Time, time.Time, bool) {
		return start, end, false
	}
}

// TimeCost 记录函数执行时间
// 在函数开头加入 defer TimeCost(time.Now()) 即可
func TimeCost(now time.Time) {
	pc, _, _, _ := runtime.Caller(1)
	log.Printf("%s took %s", runtime.FuncForPC(pc).Name(), time.Since(now))
}
