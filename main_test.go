package main

import (
	"fmt"
	"testing"
)

type Test struct {
	input  string
	guide  string
	output string
}

var tests []Test

func init() {

	tests = []Test{
		Test{
			input:  "3, Roberto Carlos, soccer, Brazil\n35, Michael Jordan, baseball, USA\n6, Lebron James, basketball, USA\n10, Shinji Kagawa, soccer, Japan",
			guide:  "Roberto Carlos, jersey number 3",
			output: "Roberto Carlos, jersey number 3\nMichael Jordan, jersey number 35\nLebron James, jersey number 6\nShinji Kagawa, jersey number 10",
		},
		Test{
			input:  "3, Roberto Carlos, soccer, Brazil\n35, Michael Jordan, baseball, USA\n6, James Dean Lebron, basketball, USA\n10, Shinji Kagawa, soccer, Japan",
			guide:  "CARLOS, jersey number 3",
			output: "CARLOS, jersey number 3\nJORDAN, jersey number 35\nLEBRON, jersey number 6\nKAGAWA, jersey number 10",
		},
		Test{
			input:  "\"class\": \"101\", \"students\": 101}\n{\"class\": \"201\", \"students\": 80}\n{\"class\": \"202\", \"students\": 50}\n{\"class\": \"301\", \"students\": 120}",
			guide:  "Class 101 has 101 students",
			output: "Class 101 has 101 students\nClass 201 has 80 students\nClass 202 has 50 students\nClass 301 has 120 students",
		},
		Test{
			input:  "2015-12-07T09:15:17-05:00\n1998-11-05T08:15:21-05:00\n1999-01-03T04:33:30-05:00\n2000-11-05T09:16:00-05:00",
			guide:  "07/12/2015 at 09h15",
			output: "07/12/2015 at 09h15\n05/11/1998 at 08h15\n03/01/1999 at 04h33\n05/11/2000 at 09h16",
		},
		Test{
			input:  "Joel, 1949, piano\nLennon, 1940, guitar\nGrohl, 1969, drums\nBurton, 1962, bass",
			guide:  "{name: 'Joel', instrument: 'piano'}",
			output: "{name: 'Joel', instrument: 'piano'}\n{name: 'Lennon', instrument: 'guitar'}\n{name: 'Grohl', instrument: 'drums'}\n{name: 'Burton', instrument: 'bass'}",
		},
		Test{
			input:  "http://example1.org/path/index.html\nhttp://www.example2.org/path/index.html\nhttp://www.example3.org/\nhttps://www.example4.org/a/b/c/d/e/f/g/hijklmnop",
			guide:  "example1.org, path/index.html",
			output: "example1.org, path/index.html\nwww.example2.org, path/index.html\nwww.example3.org, \nwww.example4.org, a/b/c/d/e/f/g/hijklmnop",
		},
		Test{
			input:  "Feature: 37, threshold: 4386, +\nFeature: 11, threshold: 1, +\nFeature: 10, threshold: 13, +\nFeature: 0, threshold: 34, +\nFeature: 39, threshold: 44, +",
			guide:  "x[37] >= 4386",
			output: "x[37] >= 4386\nx[11] >= 1\nx[10] >= 13\nx[0] >= 34\nx[39] >= 44",
		},
		Test{
			input:  "{name: \"James\", age:\"30\", hobby: \"running\", genderWord: \"his\"}\n{name: \"Erin\", age:\"28, hobby: \"cooking\", genderWord: \"her\"}\n{name: \"Owen\", age:\"3\", hobby: \"playing chase\", genderWord: \"his\"}\n{name: \"Luke\", age:\"1\", hobby: \"reading\", genderWord: \"his\"}",
			guide:  "James is 30 years old and his favorite hobby is running",
			output: "James is 30 years old and his favorite hobby is running\nErin is 28 years old and her favorite hobby is cookin\nOwen is 3 years old and his favorite hobby is playing chas\nLuke is 1 years old and his favorite hobby is reading",
		},
	}
}

func TestTrnsfrm(t *testing.T) {
	for _, value := range tests {
		output := Trnsfrm(value.input, value.guide)
		// fmt.Println(output)
		// fmt.Println(value.output)
		if output != value.output {
			t.Errorf("test with guide \"%s\" did not pass", value.guide)
		} else {
			fmt.Printf("test with guide \"%s\" did pass\n", value.guide)
		}
	}
}

/*
  Input:
  {"message":"hello there","id":1}
  {"message":"why hello there","id":2}

  Example:
  {"id":1,"message":"hello there"}

  Output:
  {"id":1,"message":"hello there"}
  {"there":id,"message":"why hello"}
-----




nput data:
  2015-09-15T09:15:17-05:00
  1998-11-05T08:15:21-05:00
  1999-01-03T04:33:30-05:00
  2000-11-05T09:16:00-05:00
pattern:
  09:15 on 09-15
result:
  09:15 on 09-15
  11:05 on 11-05
  01:03 on 01-03
  11:05 on 11-05
What I was expecting:
  09:15 on 09-15
  08:15 on 11-05
  04:33 on 01-03
  09:16 on 11-05


Input:
 Bogdan, "Yucca"
 Josy, "Orange County"
 Bill, "San Diego"

 Example:
 Bogdan lives in Yucca

 Output:
 Bogdan lives in Yucca
 Josy lives in Orange
 Bill lives in San
*/
