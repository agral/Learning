#include <stdio.h>
#include <raylib.h>

int main()
{
    printf("Hello, world!\n");
    InitWindow(800, 600, "PiBlocks");

    SetTargetFPS(60);
    while (!WindowShouldClose()) {
        BeginDrawing();
        EndDrawing();
    }

    CloseWindow();
    return 0;
}
