#include "Employee.hpp"

#include <iostream>
#include <format>

namespace records {
Employee::Employee(const std::string& firstName, const std::string& lastName)
: m_firstName(firstName), m_lastName(lastName) {
}

void Employee::promote(int raiseAmount) {
    setSalary(getSalary() + raiseAmount);
}

void Employee::demote(int demeritAmount) {
    setSalary(getSalary() - demeritAmount);
}

void Employee::hire() { m_hired = true; }
void Employee::fire() { m_hired = false; }

void Employee::display() const {
    std::cout << std::format("Employee: {}, {}", getLastName(), getFirstName()) << "\n"
        << "------------------------------\n"
        << (isHired() ? "Current" : "Former") << " Employee\n"
        << std::format("Employee Number: {}\n", getEmployeeNumber())
        << std::format("Salary: ${}\n\n", getSalary());
}

void Employee::setFirstName(const std::string& newFirstName) {
    m_firstName = newFirstName;
}
const std::string& Employee::getFirstName() const {
    return m_firstName;
}

void Employee::setLastName(const std::string& newLastName) {
    m_lastName = newLastName;
}
const std::string& Employee::getLastName() const {
    return m_lastName;
}

void Employee::setEmployeeNumber(int newEmployeeNumber) {
    m_employeeNumber = newEmployeeNumber;
}
int Employee::getEmployeeNumber() const {
    return m_employeeNumber;
}

void Employee::setSalary(int newSalary) {
    m_salary = newSalary;
}
int Employee::getSalary() const {
    return m_salary;
}

bool Employee::isHired() const {
    return m_hired;
}

} // namespace records
