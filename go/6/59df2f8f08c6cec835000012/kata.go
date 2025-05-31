package kata

import (
	"fmt"
	"sort"
	"strings"
)

func Meeting(s string) string {
	upper := strings.ToUpper(s)
	split := strings.Split(upper, ";")

	persons := make([]string, len(split))
	for index, person := range split {
		name, surname := strings.Split(person, ":")[0], strings.Split(person, ":")[1]
		persons[index] = fmt.Sprintf("(%s, %s)", surname, name)
	}

	sort.Strings(persons)
	return strings.Join(persons, "")
}
