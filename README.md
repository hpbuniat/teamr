# teamr
You're having a cup for EA's popular football simulation, but no clue how to fairly assign teams to all players?
Sick of keeping notes up to date, to ensure a proper & fair team-roster?

Here comes **teamr**!

**teamr** takes the number of players, groups (e.g. when playing with many people) and the desired team-rating (absolute values and/or stars), giving you nice & fair groups for your tournament.

## Install

### Pre-compiled executables
Get them [here](http://github.com/hpbuniat/teamr/releases).

## Usage manual
```console
Usage: teamr [global options]

global options:
   --stars value, -s value    star selection mode (default: 0)
   --min value                min. rating ovr, att, mid, def (default: 76)
   --max value                max. rating ovr, att, mid, def (default: 80)
   --nation, -n               only national teams
   --women, -w                only women national teams
   --players value, -p value  number of players (default: 8)
   --groups value, -g value   number of groups (default: 2)
   --help, -h                 show help
   --version, -v              print the version
```

### Example
```console
teamr -p 8 -g 2 --stars 4 --min 0 --max 0
```
Will create 2 groups with 4 players each, using 4 star teams only
```console
====== Group 1
+-----------+----------------+-------------+-----+-----+-----+-----+-------+
|  PLAYER   |      TEAM      |   LEAGUE    | ATT | MID | DEF | OVR | STARS |
+-----------+----------------+-------------+-----+-----+-----+-----+-------+
| Player F  | Celta Vigo     | Liga BBVA   |  79 |  79 |  78 |  78 |   4.0 |
| Player A  | Olym. Lyonnais | Ligue 1     |  82 |  78 |  78 |  78 |   4.0 |
| Player D  | Chievo Verona  | Serie A TIM |  74 |  77 |  75 |  76 |   4.0 |
| Player H  | Milan          | Serie A TIM |  80 |  79 |  77 |  78 |   4.0 |
+-----------+----------------+-------------+-----+-----+-----+-----+-------+

====== Group 2
+-----------+-----------------+-------------+-----+-----+-----+-----+-------+
|  PLAYER   |      TEAM       |   LEAGUE    | ATT | MID | DEF | OVR | STARS |
+-----------+-----------------+-------------+-----+-----+-----+-----+-------+
| Player E  | FC Augsburg     | Bundesliga  |  75 |  76 |  74 |  76 |   4.0 |
| Player C  | Eint. Frankfurt | Bundesliga  |  77 |  73 |  75 |  75 |   4.0 |
| Player B  | West Brom       | Barclays PL |  78 |  76 |  75 |  76 |   4.0 |
| Player G  | 1. FC Köln      | Bundesliga  |  78 |  76 |  74 |  76 |   4.0 |
+-----------+-----------------+-------------+-----+-----+-----+-----+-------+
```

## Credits
- https://www.fifaindex.com/de/ for their incredible work of providing all the information teamr uses.
