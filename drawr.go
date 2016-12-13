package main

import (
    "fmt"
    "time"
    "math"
    "math/rand"
)

func playr(number int64) []string {

    var players []string

    println("Enter player names (confirm each with enter)")
    for i := int64(0); i < number; i++ {
        var t string
        fmt.Printf("Player %d:", i+1)
        fmt.Scanln(&t)
        players = append(players, t)
    }

    return players
}

func drawr(players []string, teams []team, groups int64) map[int]draw {

    rand.Seed(time.Now().Unix())

    group := 1
    groupMax := int(math.Ceil(float64(int64(len(players)) / groups)))

    var drawn map[int]draw
    drawn = make(map[int]draw)
    drawn[group] = make(draw)

    for len(players) > 0 {
        if (len(drawn[group]) >= groupMax) {
            group++
            drawn[group] = make(draw)
        }

        player := rand.Int() % len(players)
        team := rand.Int() % len(teams)

        drawn[group][players[player]] = teams[team]

        // remove picked elements
        players[player] = players[len(players)-1]
        players = players[:len(players)-1]

        teams[team] = teams[len(teams)-1]
        teams = teams[:len(teams)-1]
    }

    return drawn
}