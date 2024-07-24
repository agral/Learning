#include <iostream>
#include <stdexcept>
#include <format>
#include <string>

#include "Database.hpp"
#include "Employee.hpp"

using namespace records;

int displayMenu();
void hire(Database& db);
void fire(Database& db);
void promote(Database& db);

int main() {
    Database db;
    bool isDone {false};
    while (!isDone) {
        int selection = displayMenu();
        if (selection == 0) {
            isDone = true;
        }
        else if (selection == 1) {
            hire(db);
        }
        else if (selection == 2) {
            fire(db);
        }
        else if (selection == 3) {
            promote(db);
        }
        else if (selection == 4) {
            db.displayAll();
        }
        else if (selection == 5) {
            db.displayCurrent();
        }
        else if (selection == 6) {
            db.displayFormer();
        }
        else {
            std::cerr << "Unknown command. Ignoring.\n";
        }
    }
}

int displayMenu() {
    std::cout   << "Employee Database\n"
                << "-----------------\n"
                << "1) Hire a new employee\n"
                << "2) Fire an employee\n"
                << "3) Promote an employee\n"
                << "4) List all employees\n"
                << "5) List all current employees\n"
                << "6) List all former employees\n"
                << "0) Quit\n\n"
                << "--> ";
    int selection;
    std::cin >> selection;
    return selection;
}

void hire(Database& db) {
    std::string firstName, lastName;
    std::cout << "First name: ";
    std::cin >> firstName;
    std::cout << "Last name: ";
    std::cin >> lastName;
    auto& employee = db.addEmployee(firstName, lastName);
    std::cout << std::format("Hired employee: {} {}, assigned employee number {}.\n",
        firstName, lastName, employee.getEmployeeNumber());
}

void fire(Database& db) {
    int employeeNumber;
    std::cout << "Employee number: ";
    std::cin >> employeeNumber;
    try {
        auto& employee = db.getEmployee(employeeNumber);
        employee.fire();
        std::cout << std::format("Employee {} {} ({}) terminated.\n",
            employee.getFirstName(), employee.getLastName(), employee.getEmployeeNumber());
    } catch (const std::logic_error& e) {
        std::cerr << std::format("Unable to terminate employee {}. {}\n", employeeNumber, e.what());
    }
}

void promote(Database& db) {
    int employeeNumber;
    std::cout << "Employee number: ";
    std::cin >> employeeNumber;
    int raiseAmount;
    std::cout << "How much of a raise?: ";
    std::cin >> raiseAmount;
    try {
        auto& employee = db.getEmployee(employeeNumber);
        employee.promote(raiseAmount);
        std::cout << std::format("Employee {} {} ({}) promoted.\n",
            employee.getFirstName(), employee.getLastName(), employee.getEmployeeNumber());
    } catch (const std::logic_error& e) {
        std::cerr << std::format("Unable to promote employee {}. {}\n", employeeNumber, e.what());
    }
}

int previousMain() {
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
    return 0;
}
