![Gameye](https://gameye.com/img/logo_blue.png)

# Gameye SDK for Go

Create eSport and competitive matches for Counter-Strike: Global Offensive, Team Fortress 2, Left 4 Dead 2, Killing Floor 2, Insurgency and Day of Infamy for your platform without fixed monthly costs or any need for your own server infrastructure. Simply implement the Gameye API to kick off online matches when you need the, - you will even be able to implement the scores/statistics directly on your website. How cool is that!


## Installation
You need an API key to use this SDK, to obtain a free Gameye API key, please send us [an email](mailto:support@gameye.com)

## Contributing
We encourage everyone to help us improve our public packages. If you want to
contribute please submit a [pull request](https://github.com/Gameye/gameye-sdk-go/pulls).


## License
[BSD (Berkeley Software Distribution) License](https://opensource.org/licenses/bsd-license.php). 2017-2018 Gameye B.V.


## Support
Contact: [gameye.com](https://gameye.com) — support@gameye.com

## Example code
This example will print out the total number of kills in an match.

Set `GAMEYE_API_ENDPOINT` and `GAMEYE_API_TOKEN` and pass the `matchKey` as 
the first command line argument to the program.

```go
package main

import (
	"fmt"
	"io"
	"os"
	"os/signal"
	"syscall"

	"github.com/Gameye/gameye-sdk-go/models"

	"github.com/Gameye/gameye-sdk-go/clients"
	"github.com/Gameye/gameye-sdk-go/selectors"
)

func main() {
	var err error

	matchKey := os.Args[1]

	client := clients.NewGameyeClient(clients.GameyeClientConfig{})
	var sub *clients.StatisticQuerySubscription
	sub, err = client.SubscribeStatistic(matchKey)
	if err != nil {
		panic(err)
	}

	go func() {
		lastKillTotal := -1
		for {
			var err error
			var state *models.StatisticQueryState
			state, err = sub.NextState()
			if err == io.EOF {
				err = nil
				break
			}
			if err != nil {
				panic(err)
			}
			killTotal := selectKillTotal(state)
			if killTotal != lastKillTotal {
				fmt.Printf("Kill total %d\n", killTotal)
			}
		}
	}()

	signalChannel := make(chan os.Signal, 1)
	signal.Notify(
		signalChannel,
		syscall.SIGINT,
		syscall.SIGTERM,
	)

	<-signalChannel

	sub.Cancel()
}

func selectKillTotal(state *models.StatisticQueryState) (
	total int,
) {
	playerList := selectors.SelectPlayerList(state)
	for _, playerItem := range playerList {
		if value, exists := playerItem.Statistic["kill"]; exists {
			total += value
		}
	}
	return
}
```