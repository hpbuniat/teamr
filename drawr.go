package main

import (
    "fmt"
    "time"
    "math"
    "math/rand"
    "bufio"
    "os"
    "strings"
)

func playr(number int, file string) []string {

    var players []string

    if (file != "") {
        file, err := os.Open(file)
        if err != nil {
            panic(err)
        }

        defer file.Close()

        scanner := bufio.NewScanner(file)
        for scanner.Scan() {
            t := strings.TrimSpace(scanner.Text())
            if (t != "") {
                players = append(players, t)
            }
        }
    } else {
        println("Enter player names (confirm with enter)")
        for i := int(0); i < number; i++ {
            reader := bufio.NewReader(os.Stdin)
            fmt.Printf("Player %d: ", i + 1)
            t, _ := reader.ReadString('\n')
            t = strings.TrimSpace(t)

            _ = "breakpoint"
            if (t != "") {
                players = append(players, t)
            }
        }
    }

    return players
}

func drawr(players []string, teams []team, groups int) map[int]draw {

    if (len(teams) < len(players)) {
        fmt.Printf("Not enough teams found with given parameters (%d teams for %d players", len(teams), len(players))
        os.Exit(1);
    }

    group := 1
    groupMax := int(math.Ceil(float64(len(players)) / float64(groups)))

    var drawn map[int]draw
    drawn = make(map[int]draw)
    drawn[group] = make(draw)

    for len(players) > 0 {
        rand.Seed(time.Now().UnixNano())
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
