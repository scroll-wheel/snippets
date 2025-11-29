import java.util.Scanner;

public class Print {
    // Clears the current screen
    public static void clearScreen() {  
        System.out.print("\033[H\033[2J");  
        System.out.flush();  
    }

    // Reads and returns the string read from the user
    public static String readLine(String prompt) {
        System.out.print(prompt + " ");
        Scanner scanner = new Scanner(System.in);
        return scanner.nextLine().trim();
    }
}