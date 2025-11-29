public class Location {

    // Status constants
    public static final int UNGUESSED = 0;
    public static final int HIT = 1;
    public static final int MISSED = 2;

    // Instance variables 
    private int status;
    private Ship ship;

    // Location constructor. 
    public Location() {
        this.status = UNGUESSED;
        this.ship = null;
    }

    // Was this Location a hit?
    public boolean checkHit() {
        return this.status == HIT;
    }

    // Was this location a miss?
    public boolean checkMiss() {
        return this.status == MISSED;
    }

    // Was this location unguessed?
    public boolean isUnguessed() {
        return this.status == UNGUESSED;
    }

    // Mark this location a hit.
    public void markHit() {
        this.status = HIT;
    }

    // Mark this location a miss.
    public void markMiss() {
        this.status = MISSED;
    }

    // Return whether or not this location has a ship.
    public boolean hasShip() {
        return this.ship != null;
    }

    // Set the value of whether this location has a ship.
    public void setShip(Ship ship) {
        this.ship = ship;
    }

    public Ship getShip() {
        return this.ship;
    }

    // Set the status of this Location.
    public void setStatus(int status) {
        this.status = status;
    }

    // Get the status of this Location.
    public int getStatus() {
        return this.status;
    }

}