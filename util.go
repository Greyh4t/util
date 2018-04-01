package util

import (
	"bufio"
	"crypto/md5"
	"encoding/hex"
	"log"
	"os"
	"os/exec"
	"path"
	"path/filepath"
	"runtime"
	"strconv"
	"strings"
	"syscall"
	"time"
	"unicode/utf8"

	"golang.org/x/text/encoding/simplifiedchinese"
)

const (
	FullTime = 0
	DateTime = 1
	TimeTime = 2
)

var LF string

func Init() {
	switch runtime.GOOS {
	case "windows":
		LF = "\r\n"
	case "darwin":
		LF = "\r"
	default:
		LF = "\n"
	}
}

func Now(timeType int) string {
	switch timeType {
	case DateTime:
		return time.Now().Format("2006-01-02")
	case TimeTime:
		return time.Now().Format("03:04:05")
	default:
		return time.Now().Format("2006-01-02 03:04:05")
	}
}

func Md5(t []byte) string {
	h := md5.New()
	h.Write(t)
	cipherStr := h.Sum(nil)
	return hex.EncodeToString(cipherStr)
}

func Min(first int, rest ...int) int {
	min := first
	for _, v := range rest {
		if v < min {
			min = v
		}
	}
	return min
}

func Max(first int, rest ...int) int {
	max := first
	for _, v := range rest {
		if v > max {
			max = v
		}
	}
	return max
}

func ReadLines(file string) ([]string, error) {
	f, err := os.Open(file)
	if err != nil {
		return nil, err
	}
	var lines []string
	scanner := bufio.NewScanner(f)
	for scanner.Scan() {
		lines = append(lines, strings.TrimRight(scanner.Text(), "\r\n"))
	}
	f.Close()
	return lines, scanner.Err()
}

func Write2File(file, text string) (int, error) {
	f, err := os.OpenFile(file, os.O_RDWR|os.O_CREATE|os.O_TRUNC, 0666)
	if err != nil {
		return 0, err
	}
	defer f.Close()
	return f.WriteString(text)
}

func Append2File(file, text string) (int, error) {
	f, err := os.OpenFile(file, os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666)
	if err != nil {
		return 0, err
	}
	defer f.Close()
	return f.WriteString(text)
}

func LeftPad(str, pad string, length int) string {
	return strings.Repeat(pad, length-len(str)) + str
}

func RightPad(str, pad string, length int) string {
	return str + strings.Repeat(pad, length-len(str))
}

func SplitList(list []string, per int) [][]string {
	var (
		rlines [][]string
		count  = len(list) / per
	)
	if len(list)%per != 0 {
		count += 1
	}
	for i := 0; i < count; i++ {
		start := i * per
		end := (i + 1) * per
		if end > len(list) {
			end = len(list)
		}
		rlines = append(rlines, list[start:end])
	}
	return rlines
}

func GBK2Utf8(in []byte) []byte {
	if !utf8.Valid(in) {
		out, err := simplifiedchinese.GB18030.NewDecoder().Bytes(in)
		if err == nil {
			return out
		}
	}
	return in
}

func CheckIdCard(IdCard string) bool {
	var idCardArr = []byte(strings.ToUpper(strings.TrimSpace(IdCard)))
	if len(idCardArr) < 18 {
		return false
	}
	var weight = [...]int{7, 9, 10, 5, 8, 4, 2, 1, 6, 3, 7, 9, 10, 5, 8, 4, 2}
	var validate = [...]byte{49, 48, 88, 57, 56, 55, 54, 53, 52, 51, 50}
	var sum int
	for i := 0; i < len(idCardArr)-1; i++ {
		b, err := strconv.Atoi(string(idCardArr[i]))
		if err != nil {
			return false
		}
		sum += b * weight[i]
	}
	return validate[sum%11] == idCardArr[17]
}

func TimeCost(now time.Time) {
	pc, _, _, _ := runtime.Caller(1)
	log.Printf("%s took %s", runtime.FuncForPC(pc).Name(), time.Since(now))
}

func SelfPid() int {
	return syscall.Getpid()
}

func SelfName() string {
	file, _ := exec.LookPath(os.Args[0])
	absFile, _ := filepath.Abs(file)
	_, name := path.Split(strings.Replace(absFile, `\`, "/", -1))
	return name
}
