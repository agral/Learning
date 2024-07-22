#pragma once

#include <string>

namespace records {
class Employee {
public:
    static const int c_defaultStartingSalary = 50000;
    static const int c_defaultRaiseAndDemeritAmount = 2000;

    Employee(const std::string& firstName, const std::string& lastName);

    void promote(int raiseAmount=c_defaultRaiseAndDemeritAmount);
    void demote(int demeritAmount=c_defaultRaiseAndDemeritAmount);
    void hire();
    void fire();
    void display() const;

    void setFirstName(const std::string& newFirstName);
    const std::string& getFirstName() const;
    void setLastName(const std::string& newLastName);
    const std::string& getLastName() const;
    void setEmployeeNumber(int newEmployeeNumber);
    int getEmployeeNumber() const;
    void setSalary(int newSalary);
    int getSalary() const;
    bool isHired() const;

private:
    std::string m_firstName;
    std::string m_lastName;
    int m_employeeNumber {-1};
    int m_salary {c_defaultStartingSalary};
    bool m_hired {false};
};
}
