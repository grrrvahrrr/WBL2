package main

import (
	"bufio"
	"bytes"
	"io"
	"strings"
	"testing"
)

func TestReadLines(t *testing.T) {
	var tests = []struct {
		testStr        string
		expectedResult []string
	}{
		{"test1\ntest2\ntest3\ntest4\ntest5\n", []string{"test1\n", "test2\n", "test3\n", "test4\n", "test5\n"}},
		{"test6\ntest7\ntest8\ntest9\ntest0\n", []string{"test6\n", "test7\n", "test8\n", "test9\n", "test0\n"}},
	}

	for _, v := range tests {
		result, err := MockReadLines(v.testStr)
		if err != nil {
			t.Error("error reading: ", err)
		}
		if len(result) != len(v.expectedResult) {
			t.Errorf("expectedResult array(%s) and result(%s) array are of different lengths", v.expectedResult, result)
		}
		for i, vv := range result {
			if vv != v.expectedResult[i] {
				t.Errorf("expectedResult value(%s) and result array value(%s) are different", v.expectedResult[i], vv)
			}
		}
	}

}

func MockReadLines(str string) ([]string, error) {
	var lines []string
	reader := strings.NewReader(str)
	r := bufio.NewReader(reader)
	for {
		const delim = '\n'
		line, err := r.ReadString(delim)
		if err == nil || len(line) > 0 {
			if err != nil {
				line += string(delim)
			}
			lines = append(lines, line)
		}
		if err != nil {
			if err == io.EOF {
				break
			}
			return nil, err
		}
	}
	return lines, nil
}

func TestWriteLines(t *testing.T) {
	var tests = []struct {
		testSlice            []string
		expectedResult       string
		expectedResultRev    string
		expectedResultUnique string
		expectedResultRU     string
	}{
		{[]string{"test1\n", "test2\n", "test3\n", "test4\n", "test5\n", "test5\n"},
			"test1\ntest2\ntest3\ntest4\ntest5\ntest5\n",
			"test5\ntest5\ntest4\ntest3\ntest2\ntest1\n",
			"test1\ntest2\ntest3\ntest4\ntest5\n",
			"test1\ntest2\ntest3\ntest4\ntest5\n"},
		{[]string{"test6\n", "test7\n", "test8\n", "test9\n", "test0\n", "test0\n"},
			"test6\ntest7\ntest8\ntest9\ntest0\ntest0\n",
			"test0\ntest0\ntest9\ntest8\ntest7\ntest6\n",
			"test6\ntest7\ntest8\ntest9\ntest0\n",
			"test6\ntest7\ntest8\ntest9\ntest0\n"},
	}

	for _, v := range tests {
		result, err := MockWriteLines(v.testSlice, false, false)
		if err != nil {
			t.Error("error reading: ", err)
		}
		if string(result) != v.expectedResult {
			t.Errorf("expectedResult value(%s) and result value(%s) are different", v.expectedResult, string(result))
		}
	}

	for _, v := range tests {
		result, err := MockWriteLines(v.testSlice, false, true)
		if err != nil {
			t.Error("error reading: ", err)
		}
		if string(result) != v.expectedResultUnique {
			t.Errorf("expectedResultUnique value(%s) and result value(%s) are different", v.expectedResultUnique, string(result))
		}
	}

	for _, v := range tests {
		result, err := MockWriteLines(v.testSlice, true, false)
		if err != nil {
			t.Error("error reading: ", err)
		}
		if string(result) != v.expectedResultRev {
			t.Errorf("expectedResultRev value(%s) and result value(%s) are different", v.expectedResultRev, string(result))
		}
	}

	for _, v := range tests {
		result, err := MockWriteLines(v.testSlice, true, true)
		if err != nil {
			t.Error("error reading: ", err)
		}
		if string(result) != v.expectedResultRU {
			t.Errorf("expectedResultRU value(%s) and result value(%s) are different", v.expectedResultRU, string(result))
		}
	}
}

func MockWriteLines(lines []string, revflag bool, uflag bool) (string, error) {
	var buff bytes.Buffer

	if revflag {
		for i, j := 0, len(lines)-1; i < j; i, j = i+1, j-1 {
			lines[i], lines[j] = lines[j], lines[i]
		}
	}

	if uflag {
		inResult := make(map[string]bool)
		var result []string
		for _, str := range lines {
			if _, ok := inResult[str]; !ok {
				inResult[str] = true
				result = append(result, str)
			}
		}
		lines = result
	}

	for _, line := range lines {
		_, err := buff.WriteString(line)
		if err != nil {
			return "", err
		}
	}

	return buff.String(), nil
}

func TestNumSort(t *testing.T) {
	var tests = []struct {
		testSlice      []string
		expectedResult []string
	}{
		{[]string{"5\n", "1\n", "3\n", "2\n", "4\n"}, []string{"1\n", "2\n", "3\n", "4\n", "5\n"}},
		{[]string{"6\n", "10\n", "8\n", "9\n", "7\n"}, []string{"6\n", "7\n", "8\n", "9\n", "10\n"}},
	}
	for _, v := range tests {
		result, err := numSort(v.testSlice)
		if err != nil {
			t.Error("error reading: ", err)
		}
		if len(result) != len(v.expectedResult) {
			t.Errorf("expectedResult array(%s) and result(%s) array are of different lengths", v.expectedResult, result)
		}
		for i, vv := range result {
			if vv != v.expectedResult[i] {
				t.Errorf("expectedResult value(%s) and result array value(%s) are different", v.expectedResult[i], vv)
			}
		}
	}
}

func TestColumnSort(t *testing.T) {
	var test = []struct {
		testSlice          []string
		expectedResultCol0 []string
		expectedResultCol1 []string
	}{
		{[]string{"manager c\n", "clerk d\n", "employee b\n", "peon a\n", "director f\n", "guard e\n"},
			[]string{"clerk d\n", "director f\n", "employee b\n", "guard e\n", "manager c\n", "peon a\n"},
			[]string{"peon a\n", "employee b\n", "manager c\n", "clerk d\n", "guard e\n", "director f\n"}},
	}

	for _, v := range test {
		result := columnSort(v.testSlice, 0)
		if len(result) != len(v.expectedResultCol0) {
			t.Errorf("expectedResult array(%s) and result(%s) array are of different lengths", v.expectedResultCol0, result)
		}
		for i, vv := range result {
			if vv != v.expectedResultCol0[i] {
				t.Errorf("expectedResult value(%s) and result array value(%s) are different", v.expectedResultCol0[i], vv)
			}
		}
		result = columnSort(v.testSlice, 1)
		if len(result) != len(v.expectedResultCol1) {
			t.Errorf("expectedResult array(%s) and result(%s) array are of different lengths", v.expectedResultCol1, result)
		}
		for i, vv := range result {
			if vv != v.expectedResultCol1[i] {
				t.Errorf("expectedResult value(%s) and result array value(%s) are different", v.expectedResultCol1[i], vv)
			}
		}
	}
}
