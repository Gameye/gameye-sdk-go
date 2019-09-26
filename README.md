# Gameye Sdk for Go

Create eSport and competitive matches for Counter-Strike: Global Offensive, Team Fortress 2, Left 4 Dead 2, Killing Floor 2, Insurgency and Day of Infamy for your platform without fixed monthly costs or any need for your own server infrastructure. Simply implement the Gameye API to kick off online matches when you need them; you will even be able to implement the scores/statistics directly on your website. How cool is that!

## API Key
You need an API key to use this Sdk, to obtain a free Gameye API key, please send us [an email](mailto:support@gameye.com)


## Example!
Watch bots kill each other and get acquainted with our Sdk and real-time statistics.

First, get an API key!
Then checkout this repo on your computer.

Export your api key as an environment variable, like this
```
export GAMEYE_API_TOKEN=mysupersecretkey
```
And now, run
```
go run ./examples/bots.go
```
to see bots kill eachother!


## Contributing
We encourage everyone to help us improve our public packages. If you want to
contribute please submit a [pull request](https://github.com/Gameye/gameye-sdk-go/pulls).

But, never commit something that breaks the build! You may prevent this a
little bit by linking the `test.sh` script as a git `pre-commit` hook!

like this:
```bash
ln test.sh .git/hooks/pre-commit
```

Now, just before every commit, your code will be compiled and tested!


## License
[BSD (Berkeley Software Distribution) License](https://opensource.org/licenses/bsd-license.php). 2017-2019 Gameye B.V.


## Support
Contact: [gameye.com](https://gameye.com) — support@gameye.com