#include "Database.hpp"

#include <format>
#include <stdexcept>

namespace records {

Employee& Database::addEmployee(const std::string& firstName, const std::string& lastName) {
    Employee employee {firstName, lastName};
    employee.setEmployeeNumber(m_nextEmployeeNumber++);
    employee.hire();
    m_employees.push_back(employee);
    return m_employees.back();
}

Employee& Database::getEmployee(int employeeNumber) {
    for (auto& employee: m_employees) {
        if (employee.getEmployeeNumber() == employeeNumber) {
            return employee;
        }
    }
    throw std::logic_error {std::format("Employee number {} not found.", employeeNumber)};
}

Employee& Database::getEmployee(const std::string& firstName, const std::string& lastName) {
    for (auto& employee: m_employees) {
        if (employee.getFirstName() == firstName && employee.getLastName() == lastName) {
            return employee;
        }
    }
    throw std::logic_error {std::format("Employee with name {} {} not found.", firstName, lastName)};
}

void Database::displayAll() const {
    for (const auto& employee: m_employees) {
        employee.display();
    }
}

void Database::displayCurrent() const {
    for (const auto& employee: m_employees) {
        if (employee.isHired()) {
            employee.display();
        }
    }
}

void Database::displayFormer() const {
    for (const auto& employee: m_employees) {
        if (!employee.isHired()) {
            employee.display();
        }
    }
}

}
