# CSManager GO

CSManager GO is a text-based simulation game where you manage a Counter-Strike team. Train your players, compete in leagues and tournaments, manage finances, and make strategic transfers to lead your team to victory.

## Installation

To get started with CSManager GO, follow these steps:

1. **Install Go**: If you don't have Go installed, download and install it from [golang.org](https://golang.org/dl/).
2. **Clone the Repository**: Clone this repository to your local machine.
```bash
git clone https://github.com/sneaky2x/CSManager-GO.git
```
3. **Navigate to the Directory**: Change to the project directory.
```bash
cd CSManager-GO
```
4. **Build the Game**: Compile the game using Go.
```bash
go build csmanager-cli.go 
```
5. **Run the Game**: Start the game by running the executable.
```bash
./csmanager-cli #Linux/Mac
csmanager-cli.exe #Windows
```

## Features

CSManager GO includes the following features to enhance your experience:
- **Dynamic Player Skills**: Players improve after matches and can be trained via bootcamp, but their skills also decay over time based on games played.
- **Realistic League System**: Compete in a league with AI teams, earning 3 points for wins, 1 for draws, and 0 for losses.
- **Tournaments**: Separate from the league, tournaments offer a chance to win trophies and prize money without affecting league standings.
- **Transfer Market**: Buy and sell players to manage your team's strength and finances, with AI teams actively participating.
- **AI Teams**: AI teams compete in matches, make transfers, and improve their players, creating a dynamic and competitive environment.

## Contributing

Contributions are welcome! If you have ideas for new features or find any bugs, please open an issue or submit a pull request. Ensure your code follows the existing style, which is none, and includes tests for new functionality, asking your favorite LLM should work.

## License

This project is licensed under the MIT License. See the LICENSE file for details.
