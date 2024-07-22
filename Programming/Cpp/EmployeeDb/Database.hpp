#pragma once

#include <string>
#include <vector>

#include "Employee.hpp"

namespace records {
const int FIRST_EMPLOYEE_NUMBER {1000};

class Database {
public:
    Employee& addEmployee(const std::string& firstName, const std::string& lastName);
    Employee& getEmployee(int employeeNumber);
    Employee& getEmployee(const std::string& firstName, const std::string& lastName);

    void displayAll() const;
    void displayCurrent() const;
    void displayFormer() const;

private:
    std::vector<Employee> m_employees;
    int m_nextEmployeeNumber {FIRST_EMPLOYEE_NUMBER};
};
}
