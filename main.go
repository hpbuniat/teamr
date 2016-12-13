package main

import (
    "fmt"
    "os"

    "github.com/urfave/cli"
    "github.com/olekukonko/tablewriter"
)

type team struct {
    name string
    league string
    att int64
    mid int64
    def int64
    ovr int64
    stars float64
}

type draw map[string]team

// replace in build-process
var Version = "???"

func main() {
    var stars float64
    var rateMin int64
    var rateMax int64
    var national bool
    var women bool
    var players int64
    var groups int64

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
        cli.Int64Flag{
            Name:        "min",
            Value:       76,
            Usage:       "min. rating ovr, att, mid, def",
            Destination: &rateMin,
        },
        cli.Int64Flag{
            Name:        "max",
            Value:       80,
            Usage:       "max. rating ovr, att, mid, def",
            Destination: &rateMax,
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
        cli.Int64Flag{
            Name:        "players, p",
            Value:       8,
            Usage:       "number of players",
            Destination: &players,
        },
        cli.Int64Flag{
            Name:        "groups, g",
            Value:       2,
            Usage:       "number of groups",
            Destination: &groups,
        },
    }

    app.Action = func(c *cli.Context) error {
        println(fmt.Sprintf("stars %.1f, min %d, max %d, nation %t, women %t", stars, rateMin, rateMax, national, women))
        teamMap := scrapr(stars, rateMin, rateMax, national, women)
        if (int64(len(teamMap)) < players) {
            fmt.Printf("Not enough teams found with given parameters (%d teams for %d players", len(teamMap), players)
            os.Exit(1);
        }

        playerMap := playr(players)
        draw := drawr(playerMap, teamMap, groups)

        for number, group := range draw {
            table := tablewriter.NewWriter(os.Stdout)
            table.SetHeader([]string{"Player", "Team", "League", "Att", "Mid", "Def", "Ovr", "Stars"})

            for player, team := range group {
                row := []string{
                    player,
                    team.name,
                    team.league,
                    fmt.Sprintf("%d", team.att),
                    fmt.Sprintf("%d", team.mid),
                    fmt.Sprintf("%d", team.def),
                    fmt.Sprintf("%d", team.ovr),
                    fmt.Sprintf("%.1f", team.stars),
                }

                table.Append(row)
            }

            fmt.Print("\n")
            fmt.Printf("====== Group %d\n", number)
            table.Render()
        }

        return nil
    }

    app.Run(os.Args)
}
