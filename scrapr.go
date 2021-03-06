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

// scrapr wraps all internal functions to collect all teams matching the search by querying fifaindex.com.
// The page is increased internally until there are no pages left.
func scrapr(stars float64, rateMin int, rateMax int, national bool, women bool) []team {

    var teamMap []team
    var body *html.Node
    page := 1

    for ok := true; ok; ok = more(body) {
        scrapeUrl := creatr(stars, rateMin, rateMax, national, women, page)
        body = readr(scrapeUrl)
        teamMap = append(teamMap, parsr(body)...)

        page++
    }

    return teamMap
}

// more is a internal helper to determine, if there are more pages to query
func more(body *html.Node) bool {
    matcher := func(n *html.Node) bool {
        if n != nil && n.DataAtom == atom.Li && n.Parent != nil && n.Parent.Parent != nil {
            return scrape.Attr(n, "class") == "next"
        }

        return false
    }

    return (body != nil && len(scrape.FindAll(body, matcher)) > 0)
}

// parsr does the actual parsing of teams & attributes.
// The result is a list of teams.
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
                    t, _ := strconv.ParseInt(scrape.Text(teamAttr.FirstChild), 10, 32)
                    teamContainer.att = int(t)

                case "MID":
                    t, _ := strconv.ParseInt(scrape.Text(teamAttr.FirstChild), 10, 32)
                    teamContainer.mid = int(t)

                case "DEF":
                    t, _ := strconv.ParseInt(scrape.Text(teamAttr.FirstChild), 10, 32)
                    teamContainer.def =  int(t)

                case "OVR":
                    t, _ := strconv.ParseInt(scrape.Text(teamAttr.FirstChild), 10, 32)
                    teamContainer.ovr = int(t)

                case "Team Rating":
                    teamContainer.stars = float32(len(scrape.FindAll(teamAttr, scrape.ByClass("fa-star")))) + (float32(len(scrape.FindAll(teamAttr, scrape.ByClass("fa-star-half-o")))) * float32(.5))
            }
        }

        // there is team with name "none" in the markup :)
        if (len(teamContainer.name) > 0 && teamContainer.name != "None") {

            teams = append(teams, teamContainer)
        }
    }

    return teams
}

// readr does the actual querying of fifaindex.com and parses the response-body into a html.Node
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

// creatr builds the fully-qualified URL to query fifaindex.com
func creatr(stars float64, rateMin int, rateMax int, national bool, women bool, page int) *url.URL {

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
        if (stars == float64(int(stars))) {
            baseUrl += fmt.Sprintf("stars=%.0f", stars)
        } else {
            baseUrl += fmt.Sprintf("stars=%.1f", stars)
        }
    }

    returnUrl, _ := url.Parse(baseUrl);
    return returnUrl
}
