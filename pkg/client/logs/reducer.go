package logs

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

func reduce(state *State, patches *[]patch.Patch) State {

	patchDocument := make(map[string]interface{})
	patchDocument = utils.SetIn(patchDocument, []string{}, state.Logs)

	for _, singlePatch := range *patches {
		if singlePatch.Value != nil {
			var initializer interface{}
			err := json.Unmarshal(*singlePatch.Value, &initializer)
			path := convertPath(singlePatch.Path)
			if err == nil {
				patchDocument = utils.SetIn(patchDocument, path, initializer)
			} else{
				log.Printf("logs.reduce could not unmarshal %v; %v", path, err)
			}
		}
	}

	return StateWithLogs(patchDocument)
}
