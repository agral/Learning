#include <stdio.h>
#include <raylib.h>

const int WIDTH = 1024;
const int HEIGHT = 768;
const double FLOOR_LEVEL = 0.7 * HEIGHT;

typedef struct {
    int size;
    double posX;
    double velX;
    double mass;
} Block;

int main()
{
    printf("Hello, world!\n");
    InitWindow(WIDTH, HEIGHT, "PiBlocks");

    Block small = (Block){50, 0.3 * WIDTH, 0, 1};
    Block big = (Block){200, 0.7 * WIDTH, -1, 10000};

    SetTargetFPS(60);
    while (!WindowShouldClose()) {
        // Move the scene's objects:

        BeginDrawing();
        ClearBackground(BLACK);

        // Draw both blocks
        DrawRectangle(small.posX, FLOOR_LEVEL - small.size, small.size, small.size, WHITE);
        DrawRectangle(big.posX, FLOOR_LEVEL - big.size, big.size, big.size, WHITE);

        EndDrawing();
    }

    CloseWindow();
    return 0;
}
