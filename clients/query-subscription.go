package clients

import (
	"bufio"
	"encoding/json"
	"io"
	"strings"

	"github.com/Gameye/gameye-sdk-go/utils"
)

type queryPatch struct {
	Path  []string    `json:"path"`
	Value interface{} `json:"value"`
}

func newQuerySubscription(
	reader io.ReadCloser,
	cancelFunc func(),
) *querySubscription {
	return &querySubscription{
		reader,
		cancelFunc,
		make(map[string]interface{}),
		bufio.NewReader(reader),
	}
}

type querySubscription struct {
	reader     io.ReadCloser
	cancelFunc func()
	lastState  map[string]interface{}
	lineReader *bufio.Reader
}

func (qs querySubscription) Cancel() {
	qs.cancelFunc()
	qs.reader.Close()
}

func (qs querySubscription) NextState() (
	state map[string]interface{},
	err error,
) {
	state = qs.lastState

	line := ""
	for line == "" {
		line, err = qs.lineReader.ReadString('\n')
		if err == io.EOF {
			return
		}
		if err != nil {
			return
		}
		line = strings.TrimSpace(line)
	}

	lineReader := strings.NewReader(line)
	decoder := json.NewDecoder(lineReader)

	var patches []queryPatch
	err = decoder.Decode(&patches)
	if err != nil {
		return
	}

	for _, patch := range patches {
		state = utils.SetIn(state, patch.Path, patch.Value)
	}

	qs.lastState = state
	return
}
