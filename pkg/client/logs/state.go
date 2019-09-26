package logs

// State represents all the raw log data
type State struct {
	Logs map[string]interface{}
}

// StateWithLogs constructs a State object with a set of raw logs
func StateWithLogs(logs map[string]interface{}) (state State) {
	state = State{
		logs,
	}
	return
}
