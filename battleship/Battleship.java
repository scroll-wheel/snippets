public class Battleship {
    public static void main(String[] args) {
        Player player = new Player(false);
        Player computer = new Player(true);

        printIntroduction();
        placeShips(player);
        Print.clearScreen();

        displayGrid("Your current grid of ships.", player.getPlayersGrid(), Grid.PRINT_SHIPS);
        placeShips(computer);
        Print.readLine("Hit enter to start guessing.");

        int turnCount = 1;
        int playerHits = 0;
        int computerHits = 0;
        while(playerHits < 17 && computerHits < 17) {
            if(turnCount % 2 == 0) {
                if(askForGuess(computer, player)) {
                    computerHits++;
                }
                player.getPlayersGrid().printStatus();
                System.out.println("Total Hits = " + computerHits + " out of " + 17);
            } else {
                if(askForGuess(player, computer)) {
                    playerHits++;
                }
                computer.getPlayersGrid().printStatus();
                System.out.println("Total Hits = " + playerHits + " out of " + 17);
            }
            turnCount++;
        }
        if(playerHits == 17) {
            System.out.println("You win!");
        } else {
            System.out.println("The computer wins!");
        }
    }

    private static void printIntroduction() {
        Print.clearScreen();
        System.out.println("=======================");
        System.out.println("Welcome to Battle Ship");
        System.out.println("=======================");
        System.out.println("Time to place your ships.");
    }

    private static void displayGrid(String message, Grid grid, int printType) {
        System.out.println(message);
        if(printType == Grid.PRINT_STATUS) {
            grid.printStatus();
        } else {
            grid.printShips();
        }
    }

    private static int getRow() {
        while(true) {
            String rowString = Print.readLine("Which row? (A-J)").toUpperCase();
            if(rowString.equals("")) {
                continue;
            }
            int row = rowString.charAt(0) - 'A';
            if(row < 0 || row > 9) {
                System.out.println("Invalid row, please try again.");
            } else {
                return row;
            }
        }
    }

    private static int getCol() {
        while(true) {
            String colString = Print.readLine("Which column? (1-10)");
            if(colString.equals("")) {
                continue;
            }
            try {
                int col = Integer.parseInt(colString) - 1;
                if(col < 0 || col > 9) {
                    System.out.println("Invalid column, please try again.");
                } else {
                    return col;
                }
            } catch (NumberFormatException e) {
                System.out.println("Invalid column, please try again.");
            }
        }
    }

    private static int getDirection() {
        while(true) {
            String directionString = Print.readLine("Horizontal or vertical?");
            if(directionString.equals("")) {
                continue;
            } else if(directionString.toUpperCase().charAt(0) == 'H') {
                return Ship.HORIZONTAL;
            } else if (directionString.toUpperCase().charAt(0) == 'V') {
                return Ship.VERTICAL;
            } else {
                System.out.println("Invalid direction, please try again.");
            }
        }
    }

    private static void placeShips(Player player) {
        if(player.isComputer()) {
            Print.readLine("Hit enter for the enemy to choose their ship locations.");
            for(int i = 0; i < 5; i++) {
                Ship cur = player.getShip(i);
                while(true) {
                    int row = (int)(Math.random() * 10);
                    int col = (int)(Math.random() * 10);
                    int direction = (int)(Math.random() * 2);
                    if(player.chooseShipLocation(cur, row, col, direction)) {
                        break;
                    }
                }
            }
            System.out.println("The enemy has placed their ships.");
        } else {
            for(int i = 0; i < 5; i++) {
                Print.readLine("Hit enter to place the next ship.");
                Print.clearScreen();
                displayGrid("Your current grid of ships.", player.getPlayersGrid(), Grid.PRINT_SHIPS);
                Ship cur = player.getShip(i);
                System.out.println("Now you need to place a ship of length " + cur.getLength());
                while(true) {

                    int row = getRow();
                    int col = getCol();

                    int direction = getDirection();

                    if(player.chooseShipLocation(cur, row, col, direction)) {
                        break;
                    } else {
                        System.out.println("Invalid ship placement. Please try again.");
                    }
                }
            }
        }
    }

    public static boolean askForGuess(Player player, Player enemy) {
        if(player.isComputer()) {
            Print.readLine("Hit enter for the computer's turn.");
            Print.clearScreen();
            char row;
            char col;
            while(true) {
                row = (char)(Math.random() * 10);
                col = (char)(Math.random() * 10);
                if(!player.getOpponentsGrid().alreadyGuessed(row, col)) {
                    break;
                }
            }
            System.out.println("Computer player guesses row " + ((char)(row + 'A')) + " and column " + (col + 1));
            if(enemy.recordOpponentGuess(row, col)) {
                System.out.println("Computer hit!");
                return true;
            } else {
                System.out.println("Computer missed.");
                return false;
            }
        } else {
            Print.readLine("Hit enter for your turn.");
            Print.clearScreen();
            displayGrid("Enemy grid", enemy.getPlayersGrid(), Grid.PRINT_STATUS);
            System.out.println("It's your turn to guess.");
            int row = getRow();
            int col = getCol();
            Print.clearScreen();
            if(enemy.recordOpponentGuess(row, col)) {
                System.out.println("You got a hit!");
                return true;
            } else {
                System.out.println("Nope, that was a miss.");
                return false;
            }
        }
        
    }
}