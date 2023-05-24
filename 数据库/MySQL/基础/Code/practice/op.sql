# 1.选择工资不在5000到12000的员工的姓名和工资
SELECT last_name, salary
FROM employees
WHERE salary NOT BETWEEN 5000 AND 12000

# 2.选择在20或50号部门工作的员工姓名和部门号

SELECT	last_name, department_id
FROM employees
WHERE department_id IN (20,50)

# 3.选择公司中没有管理者的员工姓名及job_id

SELECT	last_name, job_id, manager_id
FROM employees
WHERE manager_id IS NULL

# 4.选择公司中有奖金的员工姓名，工资和奖金级别

SELECT	last_name, salary, commission_pct
FROM employees
WHERE commission_pct IS NOT NULL

# 5.选择员工姓名的第三个字母是a的员工姓名

SELECT	last_name
FROM employees
WHERE last_name LIKE '__a%'

# 6.选择姓名中有字母a和k的员工姓名

SELECT	last_name
FROM employees
WHERE last_name LIKE '%a%k%' OR last_name LIKE '%k%a%'

# 7.显示出表 employees 表中 first_name 以 'e'结尾的员工信息

SELECT	*
FROM employees
WHERE first_name LIKE '%e'

# 8.显示出表 employees 部门编号在 80-100 之间的姓名、工种

SELECT	department_id, last_name, job_id
FROM employees
WHERE department_id BETWEEN 80 AND 100

# 9.显示出表 employees 的 manager_id 是 100,101,110 的员工姓名、工资、管理者id

SELECT	last_name, salary, manager_id
FROM employees
WHERE manager_id IN (100, 101, 110)