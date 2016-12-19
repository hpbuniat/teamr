package main

import (
    "net/http"

    "github.com/yhat/scrape"
    "golang.org/x/net/html"
    "golang.org/x/net/html/atom"
    "fmt"
    "net/url"
    "strings"
    "strconv"
)

func scrapr(stars float64, rateMin int64, rateMax int64, national bool, women bool) []team {

    var teamMap []team
    var body *html.Node
    page := 1

    for ok := true; ok; ok = more(body) {
        scrapeUrl := urlCreatr(stars, rateMin, rateMax, national, women, page)
        body = readr(scrapeUrl)
        teamMap = append(teamMap, parsr(body)...)

        page++
    }

    return teamMap
}

func more(body *html.Node) bool {
    matcher := func(n *html.Node) bool {
        if n != nil && n.DataAtom == atom.Li && n.Parent != nil && n.Parent.Parent != nil {
            return scrape.Attr(n, "class") == "next"
        }

        return false
    }

    return (body != nil && len(scrape.FindAll(body, matcher)) > 0)
}

func parsr(body *html.Node) []team {
    teamRowMatcher := func(n *html.Node) bool {
        if n.DataAtom == atom.Tr && scrape.Attr(n, "class") == "" && n.Parent != nil && n.Parent.Parent != nil {
            return strings.Contains(scrape.Attr(n.Parent.Parent, "class"), "teams") && n.Parent.Parent.DataAtom == atom.Table
        }

        return false
    }

    teamAttrMatcher := func(n *html.Node) bool {
        if n.DataAtom == atom.Td && n.Parent != nil && n.Parent.Parent != nil {
            return n.Parent.DataAtom == atom.Tr && scrape.Attr(n, "data-title") != ""
        }

        return false
    }

    var teams []team
    teamRows := scrape.FindAll(body, teamRowMatcher)
    for _, teamRow := range teamRows {
        var teamContainer team
        teamAttrs := scrape.FindAll(teamRow, teamAttrMatcher)
        for _, teamAttr := range teamAttrs {

            title := scrape.Attr(teamAttr, "data-title")
            switch title {
                case "Name":
                    teamContainer.name = scrape.Text(teamAttr.FirstChild)

                case "League":
                    teamContainer.league = scrape.Text(teamAttr.FirstChild)

                case "ATT":
                    teamContainer.att, _ = strconv.ParseInt(scrape.Text(teamAttr.FirstChild), 10, 32)

                case "MID":
                    teamContainer.mid, _ = strconv.ParseInt(scrape.Text(teamAttr.FirstChild), 10, 32)

                case "DEF":
                    teamContainer.def, _ = strconv.ParseInt(scrape.Text(teamAttr.FirstChild), 10, 32)

                case "OVR":
                    teamContainer.ovr, _ = strconv.ParseInt(scrape.Text(teamAttr.FirstChild), 10, 32)

                case "Team Rating":
                    teamContainer.stars = float64(len(scrape.FindAll(teamAttr, scrape.ByClass("fa-star")))) + (float64(len(scrape.FindAll(teamAttr, scrape.ByClass("fa-star-half-o")))) * float64(.5))
            }
        }

        // there is team with name "none" in the markup :)
        if (len(teamContainer.name) > 0 && teamContainer.name != "None") {

            teams = append(teams, teamContainer)
        }
    }

    return teams
}

func readr(scrapeUrl *url.URL) *html.Node {
    println(fmt.Sprintf("Fetching: %s", scrapeUrl.String()))
    resp, err := http.Get(scrapeUrl.String())
    if err != nil {
        panic(err)
    }

    body, err := html.Parse(resp.Body)
    if err != nil {
        panic(err)
    }

    return body
}

func urlCreatr(stars float64, rateMin int64, rateMax int64, national bool, women bool, page int) *url.URL {

    baseUrl := fmt.Sprintf("https://www.fifaindex.com/teams/%d/?", page)
    if (rateMin > 0 && rateMax > 0) {
        baseUrl += fmt.Sprintf("overallrating_0=%d&overallrating_1=%d&", rateMin, rateMax)
        baseUrl += fmt.Sprintf("attackrating_0=%d&attackrating_1=%d&", rateMin, rateMax)
        baseUrl += fmt.Sprintf("midfieldrating_0=%d&midfieldrating_1=%d&", rateMin, rateMax)
        baseUrl += fmt.Sprintf("defenserating_0=%d&defenserating_1=%d&", rateMin, rateMax)
    }

    if national == true {
        baseUrl += "type=1&"
    } else if women == true {
        baseUrl += "type=2&"
    }

    if (stars > 0) {
        if (stars == float64(int64(stars))) {
            baseUrl += fmt.Sprintf("stars=%.0f", stars)
        } else {
            baseUrl += fmt.Sprintf("stars=%.1f", stars)
        }
    }

    returnUrl, _ := url.Parse(baseUrl);
    return returnUrl
}
