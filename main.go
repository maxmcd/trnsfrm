package main

import (
	"fmt"
	"sort"
	"strings"
)

type dataPosition struct {
	dataNum  int
	position int
	length   int
}

// sort by position
type byPosition []dataPosition

func (a byPosition) Len() int           { return len(a) }
func (a byPosition) Swap(i, j int)      { a[i], a[j] = a[j], a[i] }
func (a byPosition) Less(i, j int) bool { return a[i].position < a[j].position }

func Trnsfrm(input, guide string) string {
	lines := strings.Split(input, "\n")
	data := extractData(lines)
	return outputFromData(data, guide)
}

func extractData(lines []string) (data [][]string) {
	for _, line := range lines {
		datas := strings.Split(line, ", ")
		data = append(data, datas)
	}
	return
}

func outputFromData(data [][]string, guide string) string {
	if len(data) < 0 {
		return guide
	}
	firstData := data[0]

	var replacementRules []dataPosition

	for dataNum, dataPoint := range firstData {
		position := strings.Index(guide, dataPoint)
		if position > -1 {
			replacementRules = append(
				replacementRules,
				dataPosition{
					dataNum:  dataNum,
					position: position,
					length:   len(dataPoint),
				},
			)
		}
	}

	var output []string
	for _, dataRow := range data {
		forOutput := guide
		var outPutArray []string
		var previousIndex int
		sort.Sort(byPosition(replacementRules))
		for _, rule := range replacementRules {
			outPutArray = append(
				outPutArray,
				forOutput[previousIndex:rule.position],
			)
			outPutArray = append(
				outPutArray,
				dataRow[rule.dataNum],
			)
			previousIndex = rule.length + rule.position
		}
		output = append(output, strings.Join(outPutArray, ""))
	}
	return strings.Join(output, "\n")
}

func main() {
	input := [][]string{
		[]string{
			"3",
			"Roberto Carlos",
			"soccer",
			"Brazil",
		},
		[]string{
			"35",
			"Michael Jordan",
			"basketball",
			"USA",
		},
	}
	output := outputFromData(input, "I am so good at soccer when 33333 people are playing with Roberto Carlos")
	fmt.Println(output)
}
