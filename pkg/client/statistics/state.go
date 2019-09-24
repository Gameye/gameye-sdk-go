package statistics

type State struct {
	Statistics map[string]interface{}
}

func StateWithStatistics(statistics map[string]interface{}) (state State) {
	state = State{
		statistics,
	}
	return
}
