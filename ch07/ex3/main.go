package main

import (
	"fmt"
	"io"
	"math/rand"
	"os"
	"sort"
)

type Team struct {
	Name    string
	Players []string
}

type League struct {
	Name  string
	Teams map[string]Team
	Wins  map[string]int
}

var (
	arizonaDiamondbacks  = Team{Name: "Arizona Diamondbacks", Players: []string{"p1", "p2", "p3"}}
	atlantaBraves        = Team{Name: "Atlanta Braves", Players: []string{"p1", "p2", "p3"}}
	baltimoreOrioles     = Team{Name: "Baltimore Orioles", Players: []string{"p1", "p2", "p3"}}
	bostonRedSox         = Team{Name: "Boston Red Sox", Players: []string{"p1", "p2", "p3"}}
	chicagoCubs          = Team{Name: "Chicago Cubs", Players: []string{"p1", "p2", "p3"}}
	chicagoWhiteSox      = Team{Name: "Chicago WhiteSox", Players: []string{"p1", "p2", "p3"}}
	cincinnatiReds       = Team{Name: "Cincinnati Reds", Players: []string{"p1", "p2", "p3"}}
	clevelandGuardians   = Team{Name: "Cleveland Guardians", Players: []string{"p1", "p2", "p3"}}
	coloradoRockies      = Team{Name: "Colorado Rockies", Players: []string{"p1", "p2", "p3"}}
	detroitTigers        = Team{Name: "Detroit Tigers", Players: []string{"p1", "p2", "p3"}}
	houstonAstros        = Team{Name: "Houston Astros", Players: []string{"p1", "p2", "p3"}}
	kansasCityRoyals     = Team{Name: "Kansas City Royals", Players: []string{"p1", "p2", "p3"}}
	losAngelesAngels     = Team{Name: "Los Angeles Angels", Players: []string{"p1", "p2", "p3"}}
	losAngelesDodgers    = Team{Name: "Los Angeles Dodgers", Players: []string{"p1", "p2", "p3"}}
	miamiMarlins         = Team{Name: "Miami Marlins", Players: []string{"p1", "p2", "p3"}}
	milwaukeeBrewers     = Team{Name: "Milwaukee Brewers", Players: []string{"p1", "p2", "p3"}}
	minnesotaTwins       = Team{Name: "Minnesota Twins", Players: []string{"p1", "p2", "p3"}}
	newYorkMets          = Team{Name: "New York Mets", Players: []string{"p1", "p2", "p3"}}
	newYorkYankees       = Team{Name: "New York Yankees", Players: []string{"p1", "p2", "p3"}}
	oaklandAthletics     = Team{Name: "Oakland Athletics", Players: []string{"p1", "p2", "p3"}}
	philadelphiaPhillies = Team{Name: "Philadelphia Phillies", Players: []string{"p1", "p2", "p3"}}
	pittsburghPirates    = Team{Name: "Pittsburgh Pirates", Players: []string{"p1", "p2", "p3"}}
	sanDiegoPadres       = Team{Name: "San Diego Padres", Players: []string{"p1", "p2", "p3"}}
	sanFranciscoGiants   = Team{Name: "San Francisco Giants", Players: []string{"p1", "p2", "p3"}}
	seattleMariners      = Team{Name: "Seattle Mariners", Players: []string{"p1", "p2", "p3"}}
	stLouisCardinals     = Team{Name: "St. Louis Cardinals", Players: []string{"p1", "p2", "p3"}}
	tampaBayRays         = Team{Name: "Tampa Bay Rays", Players: []string{"p1", "p2", "p3"}}
	texasRangers         = Team{Name: "Texas Rangers", Players: []string{"p1", "p2", "p3"}}
	torontoBlueJays      = Team{Name: "Toronto Blue Jays", Players: []string{"p1", "p2", "p3"}}
	washingtonNationals  = Team{Name: "Washington Nationals", Players: []string{"p1", "p2", "p3"}}

	league = League{
		Name: "MLB",
		Teams: map[string]Team{
			arizonaDiamondbacks.Name:  arizonaDiamondbacks,
			atlantaBraves.Name:        atlantaBraves,
			baltimoreOrioles.Name:     baltimoreOrioles,
			bostonRedSox.Name:         bostonRedSox,
			chicagoCubs.Name:          chicagoCubs,
			chicagoWhiteSox.Name:      chicagoWhiteSox,
			cincinnatiReds.Name:       cincinnatiReds,
			clevelandGuardians.Name:   clevelandGuardians,
			coloradoRockies.Name:      coloradoRockies,
			detroitTigers.Name:        detroitTigers,
			houstonAstros.Name:        houstonAstros,
			kansasCityRoyals.Name:     kansasCityRoyals,
			losAngelesAngels.Name:     losAngelesAngels,
			losAngelesDodgers.Name:    losAngelesDodgers,
			miamiMarlins.Name:         miamiMarlins,
			milwaukeeBrewers.Name:     milwaukeeBrewers,
			minnesotaTwins.Name:       minnesotaTwins,
			newYorkMets.Name:          newYorkMets,
			newYorkYankees.Name:       newYorkYankees,
			oaklandAthletics.Name:     oaklandAthletics,
			philadelphiaPhillies.Name: philadelphiaPhillies,
			pittsburghPirates.Name:    pittsburghPirates,
			sanDiegoPadres.Name:       sanDiegoPadres,
			sanFranciscoGiants.Name:   sanFranciscoGiants,
			seattleMariners.Name:      seattleMariners,
			stLouisCardinals.Name:     stLouisCardinals,
			tampaBayRays.Name:         tampaBayRays,
			texasRangers.Name:         texasRangers,
			torontoBlueJays.Name:      torontoBlueJays,
			washingtonNationals.Name:  washingtonNationals,
		},
		Wins: map[string]int{},
	}
)

func main() {
	league.MatchResult(arizonaDiamondbacks, atlantaBraves, rand.Intn(20), rand.Intn(20))
	league.MatchResult(baltimoreOrioles, bostonRedSox, rand.Intn(20), rand.Intn(20))
	league.MatchResult(chicagoCubs, chicagoWhiteSox, rand.Intn(20), rand.Intn(20))
	league.MatchResult(cincinnatiReds, clevelandGuardians, rand.Intn(20), rand.Intn(20))
	league.MatchResult(coloradoRockies, detroitTigers, rand.Intn(20), rand.Intn(20))
	league.MatchResult(houstonAstros, kansasCityRoyals, rand.Intn(20), rand.Intn(20))
	league.MatchResult(losAngelesAngels, losAngelesDodgers, rand.Intn(20), rand.Intn(20))
	league.MatchResult(miamiMarlins, milwaukeeBrewers, rand.Intn(20), rand.Intn(20))
	league.MatchResult(minnesotaTwins, newYorkMets, rand.Intn(20), rand.Intn(20))
	league.MatchResult(newYorkYankees, oaklandAthletics, rand.Intn(20), rand.Intn(20))
	league.MatchResult(philadelphiaPhillies, pittsburghPirates, rand.Intn(20), rand.Intn(20))
	league.MatchResult(sanDiegoPadres, sanFranciscoGiants, rand.Intn(20), rand.Intn(20))
	league.MatchResult(seattleMariners, stLouisCardinals, rand.Intn(20), rand.Intn(20))
	league.MatchResult(tampaBayRays, texasRangers, rand.Intn(20), rand.Intn(20))
	league.MatchResult(torontoBlueJays, washingtonNationals, rand.Intn(20), rand.Intn(20))
	RankPrinter(&league, os.Stdout)
}

func (l *League) MatchResult(team1, team2 Team, score1, score2 int) {
	if _, ok := l.Teams[team1.Name]; !ok {
		return
	}
	if _, ok := l.Teams[team2.Name]; !ok {
		return
	}
	if score1 == score2 {
		return
	}
	if score1 > score2 {
		l.Wins[team1.Name]++
	} else {
		l.Wins[team2.Name]++
	}
}

func (l *League) Ranking() []string {
	names := make([]string, 0, len(l.Teams))
	for k := range l.Teams {
		names = append(names, k)
	}
	sort.Slice(names, func(i, j int) bool {
		return l.Wins[names[i]] > l.Wins[names[j]]
	})
	return names
}

type Ranker interface {
	Ranking() []string
}

func RankPrinter(r Ranker, w io.Writer) {
	results := r.Ranking()
	for i, v := range results {
		io.WriteString(w, fmt.Sprintf("%d: %s", i+1, v))
		w.Write([]byte("\n"))
	}
}
