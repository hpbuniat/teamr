package main

import (
    "testing"
    "golang.org/x/net/html"
    "os"
    "reflect"
)

func prepare(file string) *html.Node {
    fixture, _ := os.Open("testdata/" + file + ".html")
    body, _ := html.Parse(fixture)
    return body
}

func TestParsr(t *testing.T) {
    fixture := prepare("p1")
    teams := []team {
        {"Olym. Lyonnais", "Ligue 1", 80, 78, 78, 78, 4.0},
        {"Milan", "Serie A TIM", 80, 78, 77, 78, 4.0},
        {"Zenit", "Russian League", 79, 80, 77, 78, 4.0},
        {"Lazio", "Serie A TIM", 79, 79, 78, 78, 4.0},
        {"Celta Vigo", "Liga BBVA", 79, 79, 78, 78, 4.0},
        {"West Ham", "Barclays PL", 79, 79, 76, 78, 4.0},
        {"Leicester City", "Barclays PL", 79, 77, 77, 78, 4.0},
        {"Fiorentina", "Serie A TIM", 79, 76, 79, 78, 4.0},
        {"Beşiktaş", "Süper Lig", 78, 78, 77, 78, 4.0},
        {"Real Sociedad", "Liga BBVA", 78, 78, 76, 78, 4.0},
        {"Southampton", "Barclays PL", 77, 77, 78, 78, 4.0},
        {"Shakhtar Donetsk", "Rest of World", 75, 79, 77, 78, 4.0},
        {"Crystal Palace", "Barclays PL", 81, 77, 75, 77, 4.0},
        {"Fenerbahçe", "Süper Lig", 79, 77, 78, 77, 4.0},
        {"OGC Nice", "Ligue 1", 79, 77, 78, 77, 4.0},
        {"Swansea City", "Barclays PL", 79, 77, 74, 77, 4.0},
        {"Real Betis", "Liga BBVA", 79, 76, 75, 77, 4.0},
        {"Torino", "Serie A TIM", 78, 78, 76, 77, 4.0},
        {"Watford", "Barclays PL", 77, 78, 77, 77, 4.0},
        {"Galatasaray", "Süper Lig", 77, 78, 74, 77, 4.0},
        {"Hertha BSC", "Bundesliga", 77, 77, 77, 77, 4.0},
        {"RCD Espanyol", "Liga BBVA", 77, 77, 77, 77, 4.0},
        {"1899 Hoffenheim", "Bundesliga", 77, 76, 76, 77, 4.0},
        {"Málaga CF", "Liga BBVA", 76, 77, 76, 77, 4.0},
        {"UD Las Palmas", "Liga BBVA", 75, 78, 76, 77, 4.0},
        {"Werder Bremen", "Bundesliga", 79, 75, 75, 76, 4.0},
        {"West Brom", "Barclays PL", 78, 76, 75, 76, 4.0},
        {"1. FC Köln", "Bundesliga", 78, 76, 74, 76, 4.0},
        {"Rubin Kazan", "Russian League", 78, 75, 73, 76, 4.0},
        {"Spartak Moscow", "Russian League", 77, 76, 76, 76, 4.0},
    }

    s := parsr(fixture)
    if (reflect.DeepEqual(s, teams) != true) {
        t.Errorf("parsr(%q) => %q, want %q", fixture, s, teams)
    }
}

func TestMore(t *testing.T) {
    var fixture = []struct {
        in  *html.Node
        out bool
    } {
        {
            prepare("p1"),
            true,
        }, {
            prepare("p2"),
            false,
        },
    }

    for _, sut := range fixture {
        s := more(sut.in)
        if s != sut.out {
            t.Errorf("more(%q) => %q, want %q", sut.in, s, sut.out)
        }
    }
}

func TestCreatr(t *testing.T) {
    var fixture = []struct {
        stars float64
        rateMin int64
        rateMax int64
        national bool
        women bool
        page int
        out string
    } {
        {
            0.0,
            76,
            80,
            false,
            false,
            1,
            "https://www.fifaindex.com/teams/1/?overallrating_0=76&overallrating_1=80&attackrating_0=76&attackrating_1=80&midfieldrating_0=76&midfieldrating_1=80&defenserating_0=76&defenserating_1=80&",
        }, {
            5.0,
            0,
            0,
            true,
            false,
            2,
            "https://www.fifaindex.com/teams/2/?type=1&stars=5",
        }, {
            4.5,
            0,
            0,
            false,
            true,
            2,
            "https://www.fifaindex.com/teams/2/?type=2&stars=4.5",
        },
    }

    for _, sut := range fixture {
        s := creatr(sut.stars, sut.rateMin, sut.rateMax, sut.national, sut.women, sut.page)
        if s.String() != sut.out {
            t.Errorf("creatr(%q, %q, %q, %q, %q, %q) => %q, want %q", sut.stars, sut.rateMin, sut.rateMax, sut.national, sut.women, sut.page, s.String(), sut.out)
        }
    }
}
