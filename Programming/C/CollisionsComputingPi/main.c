#include <stdio.h>
#include <raylib.h>

const int WIDTH = 1024;
const int HEIGHT = 768;
const double FLOOR_LEVEL = 0.7 * HEIGHT;
const double WALL_X = 50;

typedef struct {
    int size;
    double posX;
    double velX;
    double mass;
} Block;

void draw_scene() {
    // Draw the floor:
    DrawRectangle(WALL_X, FLOOR_LEVEL, WIDTH - WALL_X, 1, YELLOW);

    // Draw the walls:
     DrawRectangle(WALL_X, 0, 1, FLOOR_LEVEL, YELLOW);
}

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

        draw_scene();

        // Draw both blocks
        DrawRectangle(small.posX, FLOOR_LEVEL - small.size, small.size, small.size, WHITE);
        DrawRectangle(big.posX, FLOOR_LEVEL - big.size, big.size, big.size, WHITE);

        EndDrawing();
    }

    CloseWindow();
    return 0;
}
