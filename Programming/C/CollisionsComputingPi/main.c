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

Block small = (Block){50, 0.3 * WIDTH, 0, 1};
Block big = (Block){200, 0.7 * WIDTH, -100, 10000};


void draw_scene() {
    // Draw the floor:
    DrawRectangle(WALL_X, FLOOR_LEVEL, WIDTH - WALL_X, 1, YELLOW);

    // Draw the walls:
     DrawRectangle(WALL_X, 0, 1, FLOOR_LEVEL, YELLOW);

    // Draw both blocks
    DrawRectangle(small.posX, FLOOR_LEVEL - small.size, small.size, small.size, WHITE);
    DrawRectangle(big.posX, FLOOR_LEVEL - big.size, big.size, big.size, WHITE);
}

void update_scene(float deltaTime) {
    // the big block:
    big.posX += big.velX * deltaTime;

    // the small block:
    small.posX += small.velX * deltaTime;

    // Handle block-wall collisions:
    if (big.posX <= WALL_X) {
        // the big block won't ever collide with the wall, but let's have the wall bouncing both blocks.
        big.velX = -big.velX;
    }
    if (small.posX <= WALL_X) {
        // realistically, only the small block collides with the wall. Make it precise:
        // the block does not remain within the wall's bounds, it is moved just outside of it.
        small.posX = WALL_X;
        small.velX = -small.velX;
    }

    // Handle block-block collision:
    // (is the right end of the small block colliding or to the right
    // of the left side of the big block?)
    if (small.posX + small.size >= big.posX) {
        const double smallVx = small.velX;
        const double bigVx = big.velX;
        small.velX = (smallVx * (small.mass - big.mass) / (small.mass + big.mass)) +
                     (bigVx * (2 * big.mass / (small.mass + big.mass))); 
        big.velX = (smallVx * (2 * small.mass / (small.mass + big.mass))) +
                   (bigVx * (big.mass - small.mass) / (small.mass + big.mass));
    }
}

int main()
{
    printf("Hello, world!\n");
    InitWindow(WIDTH, HEIGHT, "PiBlocks");

    SetTargetFPS(60);
    while (!WindowShouldClose()) {
        // Move the scene's objects:

        BeginDrawing();
        ClearBackground(BLACK);

        float deltaTime = GetFrameTime();
        for (int i = 0; i < 10; i++) {
            update_scene(deltaTime/10.0);
        }
        draw_scene();

        EndDrawing();
    }

    CloseWindow();
    return 0;
}
