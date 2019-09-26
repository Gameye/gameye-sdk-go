package statistics

// State represents all the raw statistics data
type State struct {
	Statistics map[string]interface{}
}

// StateWithStatistics constructs a State object with a set of raw statistics
func StateWithStatistics(statistics map[string]interface{}) (state State) {
	state = State{
		statistics,
	}
	return
}
