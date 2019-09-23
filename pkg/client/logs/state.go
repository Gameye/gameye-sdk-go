package logs

type State struct {
	Logs map[string]interface{}
}

func StateWithLogs(logs map[string]interface{}) (state State) {
	state = State{
		logs,
	}
	return
}
