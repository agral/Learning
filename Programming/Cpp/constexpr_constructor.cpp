#include <cstdint>

class Rectangle {
public:
    constexpr Rectangle(int width, int height) : m_width{width}, m_height{height} {}
    constexpr int64_t getArea() const { return static_cast<int64_t>(m_width) * m_height; }

private:
    int m_width{0};
    int m_height{0};
};

int main() {
    constexpr Rectangle r{3, 14};

    // With a constexpr constructor and a constexpr getArea method, r.getArea() can be used
    // from within a constant expression:
    int someArray[r.getArea()]; // now, this is allowed.
}
