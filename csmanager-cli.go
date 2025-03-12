package main

import (
    "fmt"
    "math/rand"
    "sort"
    "strconv"
    "strings"
    "time"
)

// AIActivity tracks AI team transactions
type AIActivity struct {
    TeamName  string
    Action    string
    Player    string
    Potential int
    Amount    int
}

// Global variables
var aiActivityLog []AIActivity

// Player struct
type Player struct {
    Name        string
    Skills      map[string]int // Aim, Movement, Strategy, Teamwork, Reflexes
    AvgSkill    int
    Potential   int
    GamesPlayed int
}

// Team struct
type Team struct {
    Name       string
    Players    []*Player
    Money      int
    Wins       int
    Losses     int
    Draws      int
    Points     int
    TrophyWins int
}

// ### Player Functions

// NewPlayer creates a new player with random attributes
func NewPlayer() *Player {
    names := []string{"Jake", "Liam", "Noah", "Ethan", "Mason", "Logan", "Lucas", "Aiden", "Caleb", "Owen"}
    name := fmt.Sprintf("%s%d", names[rand.Intn(len(names))], rand.Intn(100))
    player := &Player{
        Name: name,
        Skills: map[string]int{
            "Aim":      rand.Intn(21) + 50, // 50-70
            "Movement": rand.Intn(21) + 50,
            "Strategy": rand.Intn(21) + 50,
            "Teamwork": rand.Intn(21) + 50,
            "Reflexes": rand.Intn(21) + 50,
        },
        Potential:   rand.Intn(19) + 80, // 80-98
        GamesPlayed: 0,
    }
    player.updateAvgSkill()
    return player
}

// updateAvgSkill calculates the player's average skill
func (p *Player) updateAvgSkill() {
    total := 0
    for _, skill := range p.Skills {
        total += skill
    }
    p.AvgSkill = total / len(p.Skills)
}

// Bootcamp improves player skills up to their potential
func (p *Player) Bootcamp() {
    for skill := range p.Skills {
        increase := rand.Intn(5) + 1 // 1-5 improvement
        p.Skills[skill] = min(p.Skills[skill]+increase, p.Potential)
    }
    p.updateAvgSkill()
}

// ImproveAfterGame gives a chance to improve skills post-match
func (p *Player) ImproveAfterGame() {
    for skill, value := range p.Skills {
        if value < p.Potential && rand.Intn(100) < (p.Potential-p.AvgSkill)/2 {
            p.Skills[skill]++
        }
    }
    p.updateAvgSkill()
}

// Decay reduces skills based on games played
func (p *Player) Decay() {
    var decayAmount int
    switch {
    case p.GamesPlayed >= 400:
        decayAmount = 8
    case p.GamesPlayed >= 200:
        decayAmount = 4
    case p.GamesPlayed >= 100:
        decayAmount = 1
    default:
        return
    }
    skills := []string{"Aim", "Movement", "Strategy", "Teamwork", "Reflexes"}
    skillToDecay := skills[rand.Intn(len(skills))]
    p.Skills[skillToDecay] = max(0, p.Skills[skillToDecay]-decayAmount)
    p.updateAvgSkill()
}

// ### Team Functions

// NewTeam creates a new team
func NewTeam(name string, players []*Player, money int) *Team {
    return &Team{
        Name:    name,
        Players: players,
        Money:   money,
    }
}

// AvgSkill calculates the team’s average skill
func (t *Team) AvgSkill() int {
    total := 0
    for _, player := range t.Players {
        total += player.AvgSkill
    }
    return total / len(t.Players)
}

// ### Helper Functions

func min(a, b int) int {
    if a < b {
        return a
    }
    return b
}

func max(a, b int) int {
    if a > b {
        return a
    }
    return b
}

func input(prompt string) string {
    fmt.Print(prompt)
    var response string
    fmt.Scanln(&response)
    return strings.TrimSpace(response)
}

// ### Game Functions

func viewTeam(team *Team) {
    fmt.Printf("\n--- Team: %s ---\n", team.Name)
    fmt.Println("Player       | Aim | Mov | Str | Tmw | Ref | Avg | Pot | GP")
    fmt.Println("-------------+-----+-----+-----+-----+-----+-----+-----+----")
    for i, p := range team.Players {
        fmt.Printf("%d. %-9s | %3d | %3d | %3d | %3d | %3d | %3d | %3d | %3d\n",
            i+1, p.Name, p.Skills["Aim"], p.Skills["Movement"], p.Skills["Strategy"],
            p.Skills["Teamwork"], p.Skills["Reflexes"], p.AvgSkill, p.Potential, p.GamesPlayed)
    }
    fmt.Printf("Team Avg Skill: %d | Money: $%d\n", team.AvgSkill(), team.Money)
}

func bootcampPlayers(team *Team) {
    costPerPlayer := 50
    fmt.Printf("\nBootcamp costs $%d per player. Money: $%d\n", costPerPlayer, team.Money)
    viewTeam(team)
    choice := input("Enter player numbers (e.g., '1 3') or 'all': ")
    var selected []*Player
    if choice == "all" {
        selected = team.Players
    } else {
        nums := strings.Split(choice, " ")
        for _, numStr := range nums {
            num, err := strconv.Atoi(numStr)
            if err == nil && num > 0 && num <= len(team.Players) {
                selected = append(selected, team.Players[num-1])
            }
        }
    }
    totalCost := len(selected) * costPerPlayer
    if team.Money < totalCost {
        fmt.Println("Not enough money!")
        return
    }
    team.Money -= totalCost
    for _, p := range selected {
        p.Bootcamp()
    }
    fmt.Printf("Sent %d players to bootcamp for $%d!\n", len(selected), totalCost)
    viewTeam(team)
}

func playLeagueMatch(team *Team, opponents []*Team, marketPlayers *[]*Player) {
    opponent := opponents[rand.Intn(len(opponents))]
    fmt.Printf("\nLeague Match: %s vs %s\n", team.Name, opponent.Name)
    winner := simulateLeagueMatch(team, opponent)
    if winner == team {
        fmt.Println("Victory!")
        team.Wins++
        team.Points += 3
        opponent.Losses++
    } else if winner == opponent {
        fmt.Println("Defeat!")
        team.Losses++
        opponent.Wins++
        opponent.Points += 3
    } else {
        fmt.Println("Draw!")
        team.Draws++
        team.Points++
        opponent.Draws++
        opponent.Points++
    }
    // Simulate AI matches and give income
    simulateLeagueRound(opponents)
    for _, opp := range opponents {
        opp.Money += 500 // Income per match cycle
        aiTransferDecision(opp, marketPlayers)
    }
    team.Money += 500 // Player team income
}

func simulateLeagueMatch(team1, team2 *Team) *Team {
    team1Skill := team1.AvgSkill() + rand.Intn(11) - 5
    team2Skill := team2.AvgSkill() + rand.Intn(11) - 5
    for _, p := range team1.Players {
        p.GamesPlayed++
        p.ImproveAfterGame()
        p.Decay()
    }
    for _, p := range team2.Players {
        p.GamesPlayed++
        p.ImproveAfterGame()
        p.Decay()
    }
    if team1Skill > team2Skill {
        return team1
    } else if team2Skill > team1Skill {
        return team2
    }
    return nil // Draw
}

func simulateLeagueRound(opponents []*Team) {
    rand.Shuffle(len(opponents), func(i, j int) {
        opponents[i], opponents[j] = opponents[j], opponents[i]
    })
    if len(opponents) >= 4 {
        fmt.Println("\nAI Matches:")
        winner1 := simulateLeagueMatch(opponents[0], opponents[1])
        if winner1 == opponents[0] {
            opponents[0].Wins++
            opponents[0].Points += 3
            opponents[1].Losses++
            fmt.Printf("%s beats %s\n", opponents[0].Name, opponents[1].Name)
        } else if winner1 == opponents[1] {
            opponents[1].Wins++
            opponents[1].Points += 3
            opponents[0].Losses++
            fmt.Printf("%s beats %s\n", opponents[1].Name, opponents[0].Name)
        } else {
            opponents[0].Draws++
            opponents[0].Points++
            opponents[1].Draws++
            opponents[1].Points++
            fmt.Printf("%s draws %s\n", opponents[0].Name, opponents[1].Name)
        }
        winner2 := simulateLeagueMatch(opponents[2], opponents[3])
        if winner2 == opponents[2] {
            opponents[2].Wins++
            opponents[2].Points += 3
            opponents[3].Losses++
            fmt.Printf("%s beats %s\n", opponents[2].Name, opponents[3].Name)
        } else if winner2 == opponents[3] {
            opponents[3].Wins++
            opponents[3].Points += 3
            opponents[2].Losses++
            fmt.Printf("%s beats %s\n", opponents[3].Name, opponents[2].Name)
        } else {
            opponents[2].Draws++
            opponents[2].Points++
            opponents[3].Draws++
            opponents[3].Points++
            fmt.Printf("%s draws %s\n", opponents[2].Name, opponents[3].Name)
        }
    }
}

func aiTransferDecision(team *Team, marketPlayers *[]*Player) {
    if rand.Intn(10) < 7 { // 70% chance
        market := *marketPlayers
        if team.Money < 1000 && len(team.Players) > 0 {
            // Sell best player
            bestIdx := 0
            for i, p := range team.Players {
                if p.AvgSkill > team.Players[bestIdx].AvgSkill {
                    bestIdx = i
                }
            }
            player := team.Players[bestIdx]
            price := player.AvgSkill * 20
            team.Money += price
            market = append(market, player)
            team.Players[bestIdx] = NewPlayer()
            aiActivityLog = append(aiActivityLog, AIActivity{
                TeamName:  team.Name,
                Action:    "sold",
                Player:    player.Name,
                Potential: player.Potential,
                Amount:    price,
            })
        } else if team.Money > 3000 && len(market) > 0 {
            // Buy better player
            worstIdx := 0
            for i, p := range team.Players {
                if p.AvgSkill < team.Players[worstIdx].AvgSkill {
                    worstIdx = i
                }
            }
            worstSkill := team.Players[worstIdx].AvgSkill
            for i, p := range market {
                price := p.AvgSkill * 20
                if p.AvgSkill > worstSkill && price <= team.Money {
                    team.Money -= price
                    team.Players[worstIdx] = p
                    market = append(market[:i], market[i+1:]...)
                    aiActivityLog = append(aiActivityLog, AIActivity{
                        TeamName:  team.Name,
                        Action:    "bought",
                        Player:    p.Name,
                        Potential: p.Potential,
                        Amount:    price,
                    })
                    break
                }
            }
        }
        *marketPlayers = market
    }
}

func enterTournament(team *Team, opponents []*Team) {
    entryFee := 1000
    if team.Money < entryFee {
        fmt.Printf("\nNeed $%d to enter tournament. Current: $%d\n", entryFee, team.Money)
        return
    }
    team.Money -= entryFee
    fmt.Println("\n--- Tournament Begins! ---")
    tourneyTeams := []*Team{team}
    for len(tourneyTeams) < 4 {
        opp := opponents[rand.Intn(len(opponents))]
        duplicate := false
        for _, t := range tourneyTeams {
            if t == opp {
                duplicate = true
                break
            }
        }
        if !duplicate {
            tourneyTeams = append(tourneyTeams, opp)
        }
    }
    fmt.Println("Semifinals:")
    semi1Winner := simulateTournamentMatch(tourneyTeams[0], tourneyTeams[1])
    semi2Winner := simulateTournamentMatch(tourneyTeams[2], tourneyTeams[3])
    fmt.Println("Final:")
    winner := simulateTournamentMatch(semi1Winner, semi2Winner)
    winner.TrophyWins++
    if winner == team {
        prize := 5000
        team.Money += prize
        fmt.Printf("You won the tournament and $%d!\n", prize)
    } else {
        fmt.Printf("%s won the tournament.\n", winner.Name)
    }
}

func simulateTournamentMatch(team1, team2 *Team) *Team {
    team1Skill := team1.AvgSkill() + rand.Intn(11) - 5
    team2Skill := team2.AvgSkill() + rand.Intn(11) - 5
    for _, p := range team1.Players {
        p.GamesPlayed++
        p.ImproveAfterGame()
        p.Decay()
    }
    for _, p := range team2.Players {
        p.GamesPlayed++
        p.ImproveAfterGame()
        p.Decay()
    }
    fmt.Printf("%s vs %s: ", team1.Name, team2.Name)
    if team1Skill > team2Skill {
        fmt.Println(team1.Name, "wins!")
        return team1
    } else if team2Skill > team1Skill {
        fmt.Println(team2.Name, "wins!")
        return team2
    }
    fmt.Println("Tie, defaulting to first team!")
    return team1
}

func visitTransferMarket(team *Team, marketPlayers *[]*Player) {
    market := *marketPlayers
    fmt.Println("\n--- Transfer Market ---")
    fmt.Printf("Your Money: $%d\n", team.Money)
    fmt.Println("\nAvailable Players:")
    for i, p := range market {
        price := p.AvgSkill * 20
        fmt.Printf("%d. %s - Avg: %d, Pot: %d, Price: $%d\n", i+1, p.Name, p.AvgSkill, p.Potential, price)
    }
    fmt.Println("\nYour Team:")
    for i, p := range team.Players {
        value := p.AvgSkill * 20
        fmt.Printf("%d. %s - Avg: %d, Pot: %d, Value: $%d\n", i+1, p.Name, p.AvgSkill, p.Potential, value)
    }
    action := input("\n(buy/sell/exit): ")
    if action == "buy" && len(market) > 0 {
        num, _ := strconv.Atoi(input("Player number to buy: "))
        if num < 1 || num > len(market) {
            fmt.Println("Invalid number.")
            return
        }
        player := market[num-1]
        price := player.AvgSkill * 20
        if team.Money < price {
            fmt.Println("Not enough money!")
            return
        }
        replaceNum, _ := strconv.Atoi(input("Replace player number: "))
        if replaceNum < 1 || replaceNum > len(team.Players) {
            fmt.Println("Invalid number.")
            return
        }
        team.Money -= price
        market = append(market[:num-1], market[num:]...)
        team.Players[replaceNum-1] = player
        fmt.Printf("Bought %s for $%d!\n", player.Name, price)
    } else if action == "sell" {
        num, _ := strconv.Atoi(input("Player number to sell: "))
        if num < 1 || num > len(team.Players) {
            fmt.Println("Invalid number.")
            return
        }
        player := team.Players[num-1]
        value := player.AvgSkill * 20
        team.Money += value
        market = append(market, player)
        team.Players[num-1] = NewPlayer()
        fmt.Printf("Sold %s for $%d, new player added.\n", player.Name, value)
    }
    *marketPlayers = market
}

func checkFinances(team *Team) {
    fmt.Printf("\n%s Finances: $%d\n", team.Name, team.Money)
}

func viewLeagueStandings(team *Team, opponents []*Team) {
    fmt.Println("\n--- League Standings ---")
    allTeams := append([]*Team{team}, opponents...)
    sort.Slice(allTeams, func(i, j int) bool {
        return allTeams[i].Points > allTeams[j].Points
    })
    fmt.Println("Team             | W | L | D | Pts")
    fmt.Println("-----------------+---+---+---+----")
    for _, t := range allTeams {
        fmt.Printf("%-16s | %2d | %2d | %2d | %3d\n", t.Name, t.Wins, t.Losses, t.Draws, t.Points)
    }
}

func viewAIActivity() {
    fmt.Println("\n--- AI Team Activity (Last 10) ---")
    start := max(0, len(aiActivityLog)-10)
    for _, a := range aiActivityLog[start:] {
        fmt.Printf("%s %s %s (Pot: %d) for $%d\n", a.TeamName, a.Action, a.Player, a.Potential, a.Amount)
    }
}

func viewTrophyRanking(team *Team, opponents []*Team) {
    fmt.Println("\n--- Trophy Rankings ---")
    allTeams := append([]*Team{team}, opponents...)
    sort.Slice(allTeams, func(i, j int) bool {
        return allTeams[i].TrophyWins > allTeams[j].TrophyWins
    })
    fmt.Println("Team             | Trophies")
    fmt.Println("-----------------+---------")
    for _, t := range allTeams {
        fmt.Printf("%-16s | %2d\n", t.Name, t.TrophyWins)
    }
}

// ### Main Game Loop

func main() {
    rand.Seed(time.Now().UnixNano())
    fmt.Println("Welcome to Counter-Strike Manager GO!")
    teamName := input("Enter your team name: ")
    players := make([]*Player, 5)
    for i := range players {
        players[i] = NewPlayer()
    }
    team := NewTeam(teamName, players, 5000)

    teamNames := []string{"ThunderHub", "BlazeSquad", "IceWolves", "ShadowPeak", "SteelVipers"}
    opponents := make([]*Team, 5)
    for i := range opponents {
        oppPlayers := make([]*Player, 5)
        for j := range oppPlayers {
            oppPlayers[j] = NewPlayer()
        }
        opponents[i] = NewTeam(teamNames[i], oppPlayers, 1000+rand.Intn(2000))
    }

    marketPlayers := make([]*Player, 10)
    for i := range marketPlayers {
        marketPlayers[i] = NewPlayer()
    }

    for {
        fmt.Println("\n--- Menu ---")
        fmt.Println("1. View Team")
        fmt.Println("2. Send Players to Bootcamp")
        fmt.Println("3. Play League Match")
        fmt.Println("4. Enter Tournament")
        fmt.Println("5. Visit Transfer Market")
        fmt.Println("6. Check Finances")
        fmt.Println("7. View League Standings")
        fmt.Println("8. View AI Team Activity")
        fmt.Println("9. View Trophy Ranking")
        fmt.Println("10. Exit")
        choice := input("Choose an option: ")

        switch choice {
        case "1":
            viewTeam(team)
        case "2":
            bootcampPlayers(team)
        case "3":
            playLeagueMatch(team, opponents, &marketPlayers)
            marketPlayers = append(marketPlayers, NewPlayer(), NewPlayer())
            if len(marketPlayers) > 15 {
                marketPlayers = marketPlayers[len(marketPlayers)-15:]
            }
        case "4":
            enterTournament(team, opponents)
        case "5":
            visitTransferMarket(team, &marketPlayers)
        case "6":
            checkFinances(team)
        case "7":
            viewLeagueStandings(team, opponents)
        case "8":
            viewAIActivity()
        case "9":
            viewTrophyRanking(team, opponents)
        case "10":
            fmt.Println("Thanks for playing, boss!")
            return
        default:
            fmt.Println("Invalid choice, try again.")
        }
    }
}
