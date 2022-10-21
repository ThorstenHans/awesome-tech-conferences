package template

import (
	"sort"
	"time"

	"github.com/ThorstenHans/awesome-tech-conferences/tools/awesome-readme-generator/pkg/confs"
)

func FormatDate(t time.Time) string {
	return t.Format("02-Jan-2006")
}

func HasCallForPaper(c confs.Conference) bool {
	return c.HasCallForPapers()
}

func GetMapLength(m map[int][]confs.Conference) int {
	return len(m)
}

func GetKeys(m map[int][]confs.Conference) []int {
	keys := make([]int, 0, len(m))
	for k := range m {
		keys = append(keys, k)
	}
	sort.Slice(keys, func(i, j int) bool {
		return keys[i] > keys[j]
	})
	return keys
}

func GetValues(m map[int][]confs.Conference, year int) []confs.Conference {
	return m[year]
}
