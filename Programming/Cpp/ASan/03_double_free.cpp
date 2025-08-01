#include <algorithm>
#include <random>

// A simple program that does vector addition,
// but there's a double free at the end
// (a is `delete[]`ed twice, and b is never `delete[]`ed)
int main() {
    const int SIZE = 1024;
    float *a = new float[SIZE];
    float *b = new float[SIZE];

    std::random_device rd;
    std::mt19937 mt(rd());
    std::uniform_real_distribution dist(0.0f, 1.0f);

    std::generate(a, a+SIZE, [&]{ return dist(mt); });
    std::generate(b, b+SIZE, [&]{ return dist(mt); });

    for (int i = 0; i < SIZE; i++) {
        a[i] += b[i];
    }

    delete[] a;
    delete[] a;
}
