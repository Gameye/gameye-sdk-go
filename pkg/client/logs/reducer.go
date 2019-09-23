package logs

import (
	"encoding/json"
	"fmt"

	"../patch"
	"github.com/Gameye/gameye-sdk-go/utils"
)

func copyPorts(in *map[string]int64) (out map[string]int64) {
	out = make(map[string]int64)
	for k, v := range *in {
		out[k] = v
	}
	return out
}

func convertPath(inPath []interface{}) (outPath []string) {
	outPath = []string{}
	for _, segment := range inPath {
		outPath = append(outPath, fmt.Sprintf("%v", segment))
	}
	return outPath
}

func Reduce(state *State, patches *[]patch.Patch) State {

	patchDocument := make(map[string]interface{})
	patchDocument = utils.SetIn(patchDocument, []string{}, state.Logs)

	for _, patch := range *patches {
		if patch.Value != nil {
			var initializer map[string]interface{}
			json.Unmarshal(*patch.Value, &initializer)
			path := convertPath(patch.Path)
			patchDocument = utils.SetIn(patchDocument, path, initializer)
		}
	}

	return StateWithLogs(patchDocument)
}
