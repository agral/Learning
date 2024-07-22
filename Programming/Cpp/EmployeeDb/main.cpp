#include "Employee.hpp"

int main() {
    records::Employee johnny{"John", "Doe"};
    johnny.display();

    johnny.hire();
    johnny.setSalary(100000);
    johnny.setEmployeeNumber(1000);
    johnny.promote(8000);
    johnny.display();
}
