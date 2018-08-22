package clients

import (
	"bufio"
	"encoding/json"
	"io"
	"strings"

	"github.com/Gameye/gameye-sdk-go/src/utils"
)

type queryPatch struct {
	Path  []string    `json:"path"`
	Value interface{} `json:"value"`
}

type QuerySubscription interface {
	Cancel()
	NextState() (
		state map[string]interface{},
		err error,
	)
}

func NewQuerySubscription(
	reader *bufio.Reader,
	cancelFunc func(),
) QuerySubscription {
	return &querySubscription{
		reader,
		cancelFunc,
		make(map[string]interface{}),
	}
}

type querySubscription struct {
	reader     *bufio.Reader
	cancelFunc func()
	lastState  map[string]interface{}
}

func (qs querySubscription) Cancel() {
	qs.cancelFunc()
}

func (qs querySubscription) NextState() (
	state map[string]interface{},
	err error,
) {
	state = qs.lastState
	for {
		var line string
		line, err = qs.reader.ReadString('\n')
		if err == io.EOF {
			return
		}
		if err != nil {
			return
		}
		if len(line) == 0 {
			continue
		}

		lineReader := strings.NewReader(line)
		decoder := json.NewDecoder(lineReader)

		var patches []queryPatch
		err = decoder.Decode(&patches)
		if err != nil {
			panic(err)
		}

		for _, patch := range patches {
			state = utils.SetIn(state, patch.Path, patch.Value)
		}
		break
	}
	qs.lastState = state
	return
}
