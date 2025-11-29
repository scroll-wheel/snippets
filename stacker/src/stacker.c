#include <signal.h>
#include <stdbool.h>
#include <stdio.h>
#include <stdlib.h>
#include <termios.h>
#include <time.h>
#include <unistd.h>

#define WIDTH 9
#define HEIGHT 15

#define EMPTY 0
#define BLOCK 1
#define WALL 2

#define LEFT 0
#define RIGHT 1

int grid[HEIGHT][WIDTH];
int direction = LEFT;

void draw() {
	printf("\e[1;1H\e[2J");
	for(int i = 0; i < HEIGHT; i++) {
		for(int j = 0; j < WIDTH; j++) {
			if(grid[i][j] & WALL)
				printf("██");
			else if(grid[i][j] & BLOCK)
				printf("\e[0;94m██\e[0m");
			else
				printf("  ");
		}
		printf("\n");
	}
}

void moveBlocks(int level) {
	int i = HEIGHT - level;
	int leftmost, rightmost;
	if(grid[i][0] & BLOCK) direction = RIGHT;
	if(grid[i][WIDTH - 1] & BLOCK) direction = LEFT;
	for(int j = 0; j < WIDTH; j++) {
		if(grid[i][j] & BLOCK) {
			leftmost = j;
			break;
		}
	}
	for(int j = WIDTH - 1; j >= 0; j--) {
		if(grid[i][j] & BLOCK) {
			rightmost = j;
			break;
		}
	}
	if(direction == LEFT) {
		grid[i][leftmost-1] += BLOCK;
		grid[i][rightmost] -= BLOCK;
	} else {
		grid[i][leftmost] -= BLOCK;
		grid[i][rightmost+1] += BLOCK;
	}
}

int placeBlocks(int level) {
	int i = HEIGHT - level;
	int count = 0;
	for(int j = 0; j < WIDTH; j++) {
		if(grid[i][j] == BLOCK && (level == 1 || grid[i+1][j] == BLOCK)) {
			count++;
		}
	}
	return count;
}

void animateBlink(int x, int y, int z, bool gameOver) {
	int a = -1, b = -1, c = -1;
	for(int i = 0; i < HEIGHT; i++) {
		if(a == -1 && x != -1 && grid[i][x] == BLOCK) a = i;
		if(b == -1 && y != -1 && grid[i][y] == BLOCK) b = i;
		if(c == -1 && z != -1 && grid[i][z] == BLOCK) c = i;
	}
	for(int n = 0; n < 3; n++) {
		if(a != -1 && x != -1) grid[a][x] = BLOCK;
		if(b != -1 && y != -1) grid[b][y] = BLOCK;
		if(c != -1 && z != -1) grid[c][z] = BLOCK;
		draw();
		usleep(250000);
		
		if(a != -1 && x != -1) grid[a][x] = EMPTY;
		if(b != -1 && y != -1) grid[b][y] = EMPTY;
		if(c != -1 && z != -1) grid[c][z] = EMPTY;
		draw();
		usleep(250000);
	}
	if(gameOver) {
		if(a != -1 && x != -1) grid[a][x] = BLOCK;
		if(b != -1 && y != -1) grid[b][y] = BLOCK;
		if(c != -1 && z != -1) grid[c][z] = BLOCK;
		draw();
		usleep(250000);
	}
}

void animateFall(int level, int speed) {
	int i = HEIGHT - level;
	int x = -1, y = -1;
	while(i < HEIGHT - 1) {
		int count = 0;
		for(int j = 0; j < WIDTH; j++) {
			if(grid[i][j] == BLOCK && grid[i+1][j] != BLOCK) {
				if(x == -1) x = j;
				else if(y == -1) y = j;
				grid[i][j] -= BLOCK;
				grid[i+1][j] += BLOCK;
				count++;
			}
		}
		if(count == 0)
			break;
		usleep(speed * 1000);
		draw();
		i++;
	}
	if(x != -1) animateBlink(x, y, -1, false);
}

void initNewLevel(int level, int nBlocks) {
	int i = HEIGHT - level;
	int j = (rand() % (WIDTH - 2)) + 1;
	grid[i][j] += BLOCK;
	if(nBlocks == 2) {
		int k = rand() % 2 ? j - 1 : j + 1;
		grid[i][k] += BLOCK;
	} else if(nBlocks == 3) {
		grid[i][j - 1] += BLOCK;
		grid[i][j + 1] += BLOCK;
	}
	direction = rand() % 2;
}

// Set terminal input to (non)canonical processing mode
void setCanonical(bool enable) {
	static bool enabled = true;
	static struct termios old;
	struct termios new;

	if (enable && !enabled) {
		printf("\033[?25h\033[m");
		tcsetattr(STDIN_FILENO, TCSANOW, &old);
		enabled = true;
	} else if (!enable && enabled) {
		printf("\033[?25l\033[2J");
		tcgetattr(STDIN_FILENO, &new);
		old = new;
		new.c_lflag &= (~ICANON & ~ECHO);
		tcsetattr(STDIN_FILENO, TCSANOW, &new);
		enabled = false;
	}
}

int kbhit() {
	struct timeval tv;
	fd_set fds;
	tv.tv_sec = 0;
	tv.tv_usec = 0;
	FD_ZERO(&fds);
	FD_SET(STDIN_FILENO, &fds);
	select(STDIN_FILENO+1, &fds, NULL, NULL, &tv);
	return FD_ISSET(STDIN_FILENO, &fds);
}

void sigintCallback(int signum) {
	setCanonical(true);
	exit(signum);
}

int min(int a, int b) {
	return a > b ? b : a;
}

int main() {
	signal(SIGINT, sigintCallback);
	setCanonical(false);
	srand(time(0));

	// Initialize grid
	for(int i = 0; i < HEIGHT; i++) {
		for(int j = 0; j < WIDTH; j++) {
			grid[i][j] = EMPTY;
		}
	}
	for(int i = 0; i < HEIGHT; i++) {
		grid[i][0] += WALL;
		grid[i][WIDTH - 1] += WALL;
	}

	int level = 1;
	int speedCount = 0;
	int speeds[HEIGHT + 1] = {-1, 100, 98, 96, 94, 92, 90, 88, 86, 84, 82, 80, 78, 76, 74, 72};
	int maxBlocks[HEIGHT + 1] = {-1, 3, 3, 3, 2, 2, 2, 2, 2, 2, 1, 1, 1, 1, 1, 1};
	int speed = speeds[level];
	// bool spaceHeld = false;
	char c = 0;

	initNewLevel(level, maxBlocks[level]);
	while(level <= HEIGHT) {
		usleep(1000);
		speedCount++;
		
		// TODO: Space bar latch
		if(kbhit()) {
			c = fgetc(stdin);
			if(c == 'q')
				break;
			if(c == ' ') {
				speedCount = 0;
				int n = placeBlocks(level);
				if(n != 0) {
					animateFall(level, speed);
				} else {
					int i = HEIGHT - level;
					int x = -1, y = -1, z = -1;
					for(int j = 0; j < WIDTH; j++) {
						if(grid[i][j] == BLOCK) {
							if(x == -1) x = j;
							else if(y == -1) y = j;
							else z = j;
						}
					}
					animateBlink(x, y, z, true);
					break;
				}
				level++;
				speed = speeds[level];
				initNewLevel(level, min(maxBlocks[level], n));
				usleep(500000);
				draw();
			}
		}

		if(speedCount == speed) {
			speedCount = 0;
			moveBlocks(level);
			draw();
		}
	}
	setCanonical(true);
	return 0;
}
