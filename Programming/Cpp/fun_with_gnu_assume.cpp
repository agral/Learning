#include <cstdio>

int main() {
   puts("Hello, world\n");
    [[gnu::assume(1.0000000000000001 > 1)]];
}
