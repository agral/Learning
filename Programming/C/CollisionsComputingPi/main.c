#include <stdio.h>
#include <raylib.h>

int main()
{
    printf("Hello, world!\n");
    InitWindow(1024, 768, "PiBlocks");

    SetTargetFPS(60);
    while (!WindowShouldClose()) {
        BeginDrawing();
        ClearBackground(BLACK);
        DrawRectangle(50, 50, 200, 200, WHITE);
        EndDrawing();
    }

    CloseWindow();
    return 0;
}
