/*******************************************************************************************
*
*   raylib [core] example - Basic window
*
*   Welcome to raylib!
*
*   To test examples, just press F6 and execute raylib_compile_execute script
*   Note that compiled executable is placed in the same folder as .c file
*
*   You can find all basic examples on C:\raylib\raylib\examples folder or
*   raylib official webpage: www.raylib.com
*
*   Enjoy using raylib. :)
*
*   This example has been created using raylib 1.0 (www.raylib.com)
*   raylib is licensed under an unmodified zlib/libpng license (View raylib.h for details)
*
*   Copyright (c) 2013-2016 Ramon Santamaria (@raysan5)
*
********************************************************************************************/

#include <cstdlib>
#include <ctime>
#include <iostream>
#include <vector>
#include <string>
#include "raylib.h"

//------------------------------------------------------------------------------------
// Program main entry point
//------------------------------------------------------------------------------------
int main(void)
{
    // Initialization
    //--------------------------------------------------------------------------------------
    const int screenWidth = 800;
    const int screenHeight = 450;

    InitWindow(screenWidth, screenHeight, "raylib [core] example - basic window");

    SetTargetFPS(60);               // Set our game to run at 60 frames-per-second
    //--------------------------------------------------------------------------------------

    const int squareSize = 50;
    const int n = 8;
    std::vector<std::vector<bool>> grid;

    std::srand(std::time(nullptr));

    for(int i = 0; i < n; i++) {
        std::vector<bool> row;
        for(int j = 0; j < n; j++) {
            row.push_back(0);
            // row.push_back(std::rand() % 2);
        }
        grid.push_back(row);
    }

    for(int m = 0; m < n * n; m++) {
        int i = std::rand() % n;
        int j = std::rand() % n;
        for(int k = -1; k <= 1; k++) {
            for(int l = -1; l <= 1; l++) {
                if(0 <= i + k && i + k < n && 0 <= j + l && j + l < n)
                    grid[i + k][j + l] = !grid[i + k][j + l];
            }
        }
    }

    // Main game loop
    while (!WindowShouldClose())    // Detect window close button or ESC key
    {
        // Update
        //----------------------------------------------------------------------------------
        // TODO: Update your variables here
        //----------------------------------------------------------------------------------
        for(int i = 0; i < n; i++) {
            for(int j = 0; j < n; j++) {
                int x = 0 + squareSize * j * 1.1;
                int y = 0 + squareSize * i * 1.1;

                int mouseX = GetMouseX();
                int mouseY = GetMouseY();

                if(IsMouseButtonPressed(0) && x <= mouseX && mouseX <= x + squareSize && y <= mouseY && mouseY <= y + squareSize) {
                    for(int k = -1; k <= 1; k++) {
                        for(int l = -1; l <= 1; l++) {
                            if(0 <= i + k && i + k < n && 0 <= j + l && j + l < n)
                                grid[i + k][j + l] = !grid[i + k][j + l];
                        }
                    }

                    bool notSolved = false;
                    for(int k = 0; k < n; k++) {
                        for(int l = 0; l < n; l++) {
                            notSolved |= grid[k][l];
                        }
                    }
                    if(!notSolved) {
                        for(int m = 0; m < n * n; m++) {
                            int i = std::rand() % n;
                            int j = std::rand() % n;
                            for(int k = -1; k <= 1; k++) {
                                for(int l = -1; l <= 1; l++) {
                                    if(0 <= i + k && i + k < n && 0 <= j + l && j + l < n)
                                        grid[i + k][j + l] = !grid[i + k][j + l];
                                }
                            }
                        }
                    }
                }
            }
        }

        // Draw
        //----------------------------------------------------------------------------------
        BeginDrawing();

            ClearBackground(BLACK);

            // Vector2 v = GetMouseDelta();

            // int dx = v.x;
            // int dy = v.y;
            // std::string str = std::to_string(dx) + " " + std::to_string(dx) + " " + std::to_string(IsKeyDown('1'));

            // DrawText(str.c_str(), 190, 200, 20, RAYWHITE);

            for(int i = 0; i < n; i++) {
                for(int j = 0; j < n; j++) {
                    Rectangle rect;
                    rect.x = 0 + squareSize * j * 1.1;
                    rect.y = 0 + squareSize * i * 1.1;
                    rect.width = squareSize;
                    rect.height = squareSize;
                    
                    if(!grid[i][j])
                        DrawRectangleLinesEx(rect, 3, RAYWHITE);
                    else
                        DrawRectangle(0 + squareSize * j * 1.1, 0 + squareSize * i * 1.1, squareSize, squareSize, RAYWHITE);
                }
            }

        EndDrawing();
        //----------------------------------------------------------------------------------
    }

    // De-Initialization
    //--------------------------------------------------------------------------------------
    CloseWindow();        // Close window and OpenGL context
    //--------------------------------------------------------------------------------------

    return 0;
}