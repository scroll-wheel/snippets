public class Grid {

    // Constants for number of rows and columns.
    public static final int NUM_ROWS = 10;
    public static final int NUM_COLS = 10;

    // Constants for grid print type
    public static final int PRINT_STATUS = 0;
    public static final int PRINT_SHIPS = 1;

    // Instance variable
    private Location[][] grid;

    // Create a new Grid. Initialize each Location in the grid
    // to be a new Location object.
    public Grid() {
        grid = new Location[NUM_ROWS][NUM_COLS];
        for(int i = 0; i < grid.length; i++) {
            for(int j = 0; j < grid[i].length; j++) {
                grid[i][j] = new Location();
            }
        }
    }

    // Mark a hit in this location by calling the markHit method
    // on the Location object.  
    public void markHit(int row, int col) {
        grid[row][col].markHit();
    }

    // Mark a miss on this location.    
    public void markMiss(int row, int col) {
        grid[row][col].markMiss();
    }

    // Set the status of this location object.
    public void setStatus(int row, int col, int status) {
        grid[row][col].setStatus(status);
    }

    // Get the status of this location in the grid  
    public int getStatus(int row, int col) {
        return grid[row][col].getStatus();
    }

    // Return whether or not this Location has already been guessed.
    public boolean alreadyGuessed(int row, int col) {
        return !grid[row][col].isUnguessed();
    }

    // Set whether or not there is a ship at this location to the val
    // Returns false if the location is out of bounds and true otherwise   
    public boolean setShip(int row, int col, Ship ship) {
        if(row < 0 || row >= numRows() || col < 0 || col >= numCols() || hasShip(row, col)) {
            return false;
        } else {
            grid[row][col].setShip(ship);
            return true;
        }
    }

    // Return whether or not there is a ship here   
    public boolean hasShip(int row, int col) {
        return grid[row][col].hasShip();
    }


    // Get the Location object at this row and column position
    public Location get(int row, int col) {
        return grid[row][col];
    }

    // Return the number of rows in the Grid
    public int numRows() {
        return grid.length;
    }

    // Return the number of columns in the grid
    public int numCols() {
        return grid[0].length;
    }


    // Print the Grid status including a header at the top
    // that shows the columns 1-10 as well as letters across
    // the side for A-J
    // If there is no guess print a -
    // If it was a miss print a O
    // If it was a hit, print an X
    // A sample print out would look something like this:
    // 
    //   1 2 3 4 5 6 7 8 9 10 
    // A - - - - - - - - - - 
    // B - - - - - - - - - - 
    // C - - - O - - - - - - 
    // D - O - - - - - - - - 
    // E - X - - - - - - - - 
    // F - X - - - - - - - - 
    // G - X - - - - - - - - 
    // H - O - - - - - - - - 
    // I - - - - - - - - - - 
    // J - - - - - - - - - - 
    public void printStatus() {
        System.out.println("  1 2 3 4 5 6 7 8 9 10");
        char start = 'A';
        for(int i = 0; i < grid.length; i++) {
            char cur = (char)(start + i);
            String result = cur + " ";
            for(int j = 0; j < grid[i].length; j++) {
                switch(this.getStatus(i, j)) {
                    case Location.UNGUESSED:
                        result += "- ";
                        break;
                    case Location.HIT:
                        result += "X ";
                        break;
                    case Location.MISSED:
                        result += "O ";
                        break;
                    default: result += "";
                }
            }
            System.out.println(result);
        }
    }

    // Print the grid and whether there is a ship at each location.
    // If there is no ship, you will print a - and if there is a
    // ship you will print a X. You can find out if there was a ship
    // by calling the hasShip method.
    //
    //   1 2 3 4 5 6 7 8 9 10 
    // A - - - - - - - - - - 
    // B - X - - - - - - - - 
    // C - X - - - - - - - - 
    // D - - - - - - - - - - 
    // E X X X - - - - - - - 
    // F - - - - - - - - - - 
    // G - - - - - - - - - - 
    // H - - - X X X X - X - 
    // I - - - - - - - - X - 
    // J - - - - - - - - X - 
    public void printShips() {
        System.out.println("  1 2 3 4 5 6 7 8 9 10");
        char start = 'A';
        for(int i = 0; i < grid.length; i++) {
            char cur = (char)(start + i);
            String result = cur + " ";
            for(int j = 0; j < grid[i].length; j++) {
                if(this.hasShip(i, j)) {
                    result += "X ";
                } else {
                    result += "- ";
                }
            }
            System.out.println(result);
        }
    }

    /**
     * This method can be called on your own grid. To add a ship
     * we will go to the ships location and mark a true value
     * in every location that the ship takes up.
     */
    public boolean addShip(Ship s) {
        if(s.getDirection() == Ship.HORIZONTAL) {
            for(int i = 0; i < s.getLength(); i++) {
                if(!setShip(s.getRow(), s.getCol() + i, null)) {
                    return false;
                }
            }
            for(int i = 0; i < s.getLength(); i++) {
                setShip(s.getRow(), s.getCol() + i, s);
            }
        }
        else if(s.getDirection() == Ship.VERTICAL) {
            for(int i = 0; i < s.getLength(); i++) {
                if(!setShip(s.getRow() + i, s.getCol(), null)) {
                    return false;
                }
            }
            for(int i = 0; i < s.getLength(); i++) {
                setShip(s.getRow() + i, s.getCol(), s);
            }
        }
        return true;
    }

}