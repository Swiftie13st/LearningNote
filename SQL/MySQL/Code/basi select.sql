#5. 最基本的SELECT语句： SELECT 字段1,字段2,... FROM 表名 
SELECT 1 + 1,3 * 2;

SELECT 1 + 1,3 * 2
FROM DUAL; #dual：伪表

# *:表中的所有的字段（或列）
SELECT * FROM employees;

SELECT employee_id,last_name,salary
FROM employees;

#6. 列的别名
# as:全称：alias(别名),可以省略
# 列的别名可以使用一对""引起来，不要使用''。
SELECT employee_id emp_id,last_name AS lname,department_id "部门id",salary * 12 AS "annual sal"
FROM employees;

# 7. 去除重复行
#查询员工表中一共有哪些部门id呢？
#错误的:没有去重的情况
SELECT department_id
FROM employees;
#正确的：去重的情况
SELECT DISTINCT department_id
FROM employees;


SELECT salary,DISTINCT department_id
FROM employees;

#仅仅是没有报错，但是没有实际意义。
SELECT DISTINCT department_id,salary
FROM employees;


#8. 空值参与运算
# 1. 空值：null
# 2. null不等同于0，''，'null'
SELECT * FROM employees;

#3. 空值参与运算：结果一定也为空。
SELECT employee_id,salary "月工资",salary * (1 + commission_pct) * 12 "年工资",commission_pct
FROM employees;
#实际问题的解决方案：引入IFNULL
SELECT employee_id,salary "月工资",salary * (1 + IFNULL(commission_pct,0)) * 12 "年工资",commission_pct
FROM `employees`;

SELECT * FROM ORDER
SELECT * FROM `order`

# 查询常数
SELECT '尚硅谷',123,employee_id,last_name
FROM employees;

#12.过滤数据

# 查询90号部门的员工信息
SELECT * 
FROM employees
# 过滤条件,声明在FROM结构的后面
WHERE department_id = 90;

#练习：查询last_name为'King'的员工信息
SELECT * 
FROM EMPLOYEES
WHERE LAST_NAME = 'King'; 
