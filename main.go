package main

import (
    "fmt"
    "os"

    "github.com/urfave/cli"
)

// team represents all necessary attributes of a football-team
type team struct {
    name string
    league string
    att int
    mid int
    def int
    ovr int
    stars float32
}

// draw allows the assignment of a team to each player
type draw map[string]team

// replace in build-process
var Version = "???"

// Enter the cli-program
func main() {
    var stars float64
    var rateMin int
    var rateMax int
    var national bool
    var women bool
    var players int
    var groups int
    var playerFile string

    app := cli.NewApp()
    app.Version = Version
    app.Name = "teamr"
    app.Usage = "football-cup team randomizer"
    app.Flags = []cli.Flag {
        cli.Float64Flag{
            Name:        "stars, s",
            Value:       0,
            Usage:       "star selection mode",
            Destination: &stars,
        },
        cli.IntFlag{
            Name:        "min",
            Value:       76,
            Usage:       "min. rating ovr, att, mid, def",
            Destination: &rateMin,
        },
        cli.IntFlag{
            Name:        "max",
            Value:       80,
            Usage:       "max. rating ovr, att, mid, def",
            Destination: &rateMax,
        },
        cli.IntFlag{
            Name:        "players, p",
            Value:       8,
            Usage:       "number of players",
            Destination: &players,
        },
        cli.IntFlag{
            Name:        "groups, g",
            Value:       2,
            Usage:       "number of groups",
            Destination: &groups,
        },
        cli.StringFlag{
            Name:        "playerfile, f",
            Value:       "",
            Usage:       "read players from file",
            Destination: &playerFile,
        },
        cli.BoolFlag{
            Name:        "nation, n",
            Usage:       "only national teams",
            Destination: &national,
        },
        cli.BoolFlag{
            Name:        "women, w",
            Usage:       "only women national teams",
            Destination: &women,
        },
    }

    app.Action = func(c *cli.Context) error {
        println(fmt.Sprintf("stars %.1f, min %d, max %d, nation %t, women %t", stars, rateMin, rateMax, national, women))
        teamMap := scrapr(stars, rateMin, rateMax, national, women)

        playerMap := playr(players, playerFile)

        // re-test amount of players, as one could enter an empty name
        players = len(playerMap)

        drawn := drawr(playerMap, &teamMap, groups)
        printr(drawn)
        for playerMap = playr(1, ""); len(playerMap) > 0; playerMap = playr(1, "") {
            drawn := drawr(playerMap, &teamMap, 1)
            printr(drawn)
        }

        return nil
    }

    app.Run(os.Args)
}
