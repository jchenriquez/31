package nextperm

import (
	"bufio"
	"encoding/json"
	"fmt"
	"io"
	"os"
	"testing"
)

type Test struct {
	Input  []int `json:"input"`
	Output []int `json:"output"`
}

func TestNextPermutation(testUtil *testing.T) {
	f, err := os.Open("tests.json")

	if err != nil {
		testUtil.Error(err)
		return
	}

	reader := bufio.NewReader(f)
	decoder := json.NewDecoder(reader)
	var tests map[string]Test

	for {
		err = decoder.Decode(&tests)

		if err == nil {
			for name, test := range tests {
				testUtil.Run(name, func(tstUtil *testing.T) {
					NextPermutation(test.Input)
					if fmt.Sprintf("%v", test.Input) != fmt.Sprintf("%v", test.Output) {
						tstUtil.Errorf("use case %v\n", test)
					}
				})
			}
		} else if err == io.EOF {
			break
		} else {
			testUtil.Error(err)
			return
		}
	}
}
