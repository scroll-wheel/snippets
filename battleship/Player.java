public class Player {
    
    // These are the lengths of all of the ships.
    private static final int[] SHIP_LENGTHS = {2, 3, 3, 4, 5};

    // Instance variables
    private Ship[] ships;
    private Grid playersGrid;
    private Grid opponentsGrid;
    private boolean isComputer;

    // Player constructor
    public Player(boolean isComputer) {
        ships = new Ship[5];
        for(int i = 0; i < SHIP_LENGTHS.length; i++) {
            ships[i] = new Ship(SHIP_LENGTHS[i]);
        }
        playersGrid = new Grid();
        opponentsGrid = new Grid();
        this.isComputer = isComputer;
    }

    // Sets a ship’s row, column and direction
    // and adds it to the current player’s grid
    public boolean chooseShipLocation(Ship s, int row, int col, int direction) {
        s.setLocation(row, col);
        s.setDirection(direction);
        return playersGrid.addShip(s);
    }

    // Takes in an opponent guess for a row, col location,
    // and records the guess, and returns a boolean indicating
    // whether the guess was a hit
    public boolean recordOpponentGuess(int row, int col) {
        if(playersGrid.alreadyGuessed(row, col)) {
            return false;
        }
        else if(playersGrid.hasShip(row, col)) {
            playersGrid.markHit(row, col);
            Ship ship = playersGrid.get(row, col).getShip();
            ship.decrementHealth();
            if(ship.getHealth() == 0) {
                System.out.println("SHIP SUNK!");
                ship.decrementHealth();
            }
            return true;
        } else {
            playersGrid.markMiss(row, col);
            return false;
        }
    }

    // Returns the ship indexed at x
    public Ship getShip(int x) {
        return ships[x];
    }

    // Returns the player's grid
    public Grid getPlayersGrid() {
        return playersGrid;
    }

    // Returns the opponent's grid
    public Grid getOpponentsGrid() {
        return opponentsGrid;
    }

    public boolean isComputer() {
        return isComputer;
    }

}