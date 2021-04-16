package hw03_frequency_analysis //nolint:golint,stylecheck
import (
	"sort"
	"strings"
)

func Top10(str string) []string {
	// str = strings.ToLower(str)
	// разделили строку, сложили по словам в слайс
	slStr := strings.Fields(str)

	// проверили на нулевость (дальше считаем, что ненулевая строка)
	if len(slStr) == 0 {
		return nil
	}

	// создаем и заполняем мапу
	mStr := make(map[string]int)
	for _, val := range slStr {
		if _, ok := mStr[val]; !ok {
			mStr[val] = 1
		} else {
			mStr[val]++
		}
	}

	// усечённый слайс строк (без повторов)
	slKey := make([]string, 0)
	for i := range mStr {
		slKey = append(slKey, i)
	}

	// сортировка (slKey пересортирован)
	sort.Slice(slKey, func(i, j int) bool {
		return mStr[slKey[i]] > mStr[slKey[j]]
	})

	var Cnt, beg = mStr[slKey[0]], 0
	for i := range slKey {
		var end int
		if Cnt != mStr[slKey[i]] {
			Cnt = mStr[slKey[i]]
			end = i
			sort.Strings(slKey[beg:end])
			beg = i
		}
		if i == len(slKey)-1 {
			sort.Strings(slKey[beg:])
		}
	}

	// подготовка первых 10 (или меньше)
	slTop10 := make([]string, 0)
	for i := 0; i < len(slKey); i++ {
		if i < 10 {
			slTop10 = append(slTop10, slKey[i])
		}
	}

	return slTop10
}
