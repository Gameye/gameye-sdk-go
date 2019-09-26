package logs

import (
	"fmt"
	"sort"
	"strconv"
)

// SelectAllLogs returns all logs from the store
func SelectAllLogs(state State) (logLines []LogLine) {
	rawLogs := state.Logs["line"]
	if rawLogs != nil {
		for _, v := range rawLogs.(map[string]interface{}) {
			rawLine := v.(map[string]interface{})
			if rawLine != nil {
				lineKeyInt, _ := strconv.Atoi(fmt.Sprintf("%v", rawLine["lineKey"]))
				logLine := LogLine{
					LineKey: lineKeyInt,
					Payload: rawLine["payload"].(string),
				}
				logLines = append(logLines, logLine)
			}
		}
	}

	sort.Slice(logLines, func(p, q int) bool {
		return (logLines[q].LineKey - logLines[p].LineKey) > 0
	})

	return logLines
}

// SelectLogsSince returns logs after the given line number
func SelectLogsSince(state State, lineNumber int) (logLines []LogLine) {
	rawLogs := state.Logs["line"]
	if rawLogs != nil {
		for _, v := range rawLogs.(map[string]interface{}) {
			rawLine := v.(map[string]interface{})
			if rawLine != nil {
				lineKeyInt, _ := strconv.Atoi(fmt.Sprintf("%v", rawLine["lineKey"]))
				logLine := LogLine{
					LineKey: lineKeyInt,
					Payload: rawLine["payload"].(string),
				}
				if logLine.LineKey > lineNumber {
					logLines = append(logLines, logLine)
				}
			}
		}
	}

	sort.Slice(logLines, func(p, q int) bool {
		return (logLines[q].LineKey - logLines[p].LineKey) > 0
	})

	return logLines
}
