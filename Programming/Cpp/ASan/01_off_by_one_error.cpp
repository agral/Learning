#include <iostream>
#include <vector>

int main() {
    std::vector<int> v = {17, 42, 67, 101};
    for (int i = 0; i <= v.size(); i++) {
        std::cout << v[i] << "\n";
    }
}
