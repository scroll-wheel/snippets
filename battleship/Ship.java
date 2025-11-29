public class Ship {

    // Direction constants
    public static final int UNSET = -1;
    public static final int HORIZONTAL = 0;
    public static final int VERTICAL = 1;

    // Instance variables
    private int row;
    private int col;
    private int length;
    private int health;
    private int direction;

    // Constructor. Create a ship and set the length.
    public Ship(int length) {
        this.row = UNSET;
        this.col = UNSET;
        this.length = length;
        this.health = length;
        this.direction = UNSET;
    }

    // Has the location been initialized
    public boolean isLocationSet() {
        return this.row != UNSET && this.col != UNSET; 
    }

    // Has the direction been initialized
    public boolean isDirectionSet() {
        return this.direction != UNSET;
    }

    // Set the location of the ship
    public void setLocation(int row, int col) {
        this.row = row;
        this.col = col;
    }

    // Set the direction of the ship
    public void setDirection(int direction) {
        this.direction = direction;
    }

    // Decrement the health of the ship
    public void decrementHealth() {
        this.health--;
    }

    // Getter for the row value
    public int getRow() {
        return this.row;
    }

    // Getter for the column value
    public int getCol() {
        return this.col;
    }

    // Getter for the length of the ship
    public int getLength() {
        return this.length;
    }

    // Getter for the direction
    public int getDirection() {
        return this.direction;
    }

    // Getter for the health
    public int getHealth() {
        return this.health;
    }

    // Helper method to get a string value from the direction
    private String directionToString() {
        switch(this.getDirection()) {
            case UNSET: return "UNSET";
            case HORIZONTAL: return "HORIZONTAL";
            case VERTICAL: return "VERTICAL";
            default: return "";
        }
    }

    // Helper method to get a (row, col) string value from the location
    private String locationToString() {
        return "(" + this.getRow() + ", " + this.getCol() + ")";
    }

    // toString value for this Ship
    public String toString() {
        return "Location: " + this.locationToString() + "\n"
             + "Length: " + this.getLength() + "\n"
             + "Direction: " + this.directionToString() + "\n";
    }

}