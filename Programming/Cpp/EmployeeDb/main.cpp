#include <iostream>

#include "Database.hpp"
#include "Employee.hpp"

int main() {
    records::Employee johnny{"John", "Doe"};
    johnny.display();

    johnny.hire();
    johnny.setSalary(100000);
    johnny.setEmployeeNumber(1000);
    johnny.promote(8000);
    johnny.display();

    records::Database myDatabase{};
    records::Employee& grim = myDatabase.addEmployee("Grim", "Reaper");
    grim.setSalary(200'000);

    records::Employee& capt = myDatabase.addEmployee("Captain", "Nemo");
    capt.setSalary(100'000);
    capt.promote();

    records::Employee& bill = myDatabase.addEmployee("Buffalo", "Bill");
    bill.fire();

    std::cout << "All employees:\n";
    myDatabase.displayAll();
}
