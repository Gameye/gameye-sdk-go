package statistics

import (
	"encoding/json"
	"fmt"
	"log"

	"../patch"
	"github.com/Gameye/gameye-sdk-go/utils"
)

func convertPath(inPath []interface{}) (outPath []string) {
	outPath = []string{}
	for _, segment := range inPath {
		outPath = append(outPath, fmt.Sprintf("%v", segment))
	}
	return outPath
}

func Reduce(state *State, patches *[]patch.Patch) State {

	patchDocument := make(map[string]interface{})
	patchDocument = utils.SetIn(patchDocument, []string{}, state.Statistics)

	for _, singlePatch := range *patches {
		if singlePatch.Value != nil {
			var initializer map[string]interface{}
			err := json.Unmarshal(*singlePatch.Value, &initializer)

			if err == nil {
				path := convertPath(singlePatch.Path)
				patchDocument = utils.SetIn(patchDocument, path, initializer)
			} else {
				log.Printf("could not unmarshal json %v", err)
			}
		}
	}

	return StateWithStatistics(patchDocument)
}
