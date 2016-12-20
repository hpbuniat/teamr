package main

import (
    "os"
    "fmt"

    "github.com/olekukonko/tablewriter"
)

func printr(drawn map[int]draw) {
    for number, group := range drawn {
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

        fmt.Printf("\n====== Group %d\n", number)
        table.Render()

        fmt.Print("\n====== Teams are ready!\n")
        fmt.Print("Confirm with enter or type a name to assign a new team (e.g. invalid teams according to more current stats\n\n")
    }
}
