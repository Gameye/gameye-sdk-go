package logs

import "sort"

func SelectAllLogs(state State) (logLines []LogLine) {
	rawLogs := state.Logs["line"]
	if rawLogs != nil {
		for _, v := range rawLogs.(map[string]interface{}) {
			rawLine := v.(map[string]interface{})
			if rawLine != nil {
				logLine := LogLine{
					LineKey: int(rawLine["lineKey"].(float64)),
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

func SelectLogsSince(state State, lineNumber int) (logLines []LogLine) {
	rawLogs := state.Logs["line"]
	if rawLogs != nil {
		for _, v := range rawLogs.(map[string]interface{}) {
			rawLine := v.(map[string]interface{})
			if rawLine != nil {
				logLine := LogLine{
					LineKey: int(rawLine["lineKey"].(float64)),
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
