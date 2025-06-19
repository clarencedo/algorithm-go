CREATE TABLE departments (
    dept_id INT PRIMARY KEY AUTO_INCREMENT,
    dept_name VARCHAR(100) NOT NULL,
    location VARCHAR(100),
    budget DECIMAL(12,2)
);
CREATE TABLE employees (
    emp_id INT PRIMARY KEY AUTO_INCREMENT,
    emp_name VARCHAR(100) NOT NULL,
    salary DECIMAL(10,2) NOT NULL,
    hire_date DATE NOT NULL,
    dept_id INT NOT NULL,
    manager_id INT,
    job_title VARCHAR(100),
    email VARCHAR(100) UNIQUE,
    FOREIGN KEY (dept_id) REFERENCES departments(dept_id),
    FOREIGN KEY (manager_id) REFERENCES employees(emp_id)
);
CREATE TABLE employee_history (
    history_id INT PRIMARY KEY AUTO_INCREMENT,
    emp_id INT NOT NULL,
    dept_id INT NOT NULL,
    position VARCHAR(100),
    start_date DATE NOT NULL,
    end_date DATE,
    salary DECIMAL(10,2) NOT NULL,
    FOREIGN KEY (emp_id) REFERENCES employees(emp_id),
    FOREIGN KEY (dept_id) REFERENCES departments(dept_id)
);

-- 1. 查询工资高于本部门平均工资的员工
SELECT e.emp_name, e.salary, d.dept_name
FROM employees e
JOIN departments d ON e.dept_id = d.dept_id
WHERE e.salary > (SELECT AVG(salary) FROM employees WHERE dept_id = e.dept_id);

-- 2. 查询每个部门工资排名前三的员工
SELECT e.emp_name, e.salary, d.dept_name
FROM employees e
JOIN departments d ON e.dept_id = d.dept_id
WHERE (
    SELECT COUNT(*)
    FROM employees e2
    WHERE e2.dept_id = e.dept_id AND e2.salary > e.salary
) < 3;

-- 3. 查询某个员工的所有下属（包括间接下属）
WITH RECURSIVE subordinates AS (
    SELECT emp_id, emp_name
    FROM employees
    WHERE emp_id = ? -- 替换为你要查询的员工ID
    UNION ALL
    SELECT e.emp_id, e.emp_name
    FROM employees e
    JOIN subordinates s ON e.manager_id = s.emp_id
)
SELECT * FROM subordinates;