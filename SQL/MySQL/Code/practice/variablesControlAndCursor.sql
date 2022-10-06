#0.准备工作
CREATE DATABASE test16_var_cur;
USE test16_var_cur;
CREATE TABLE employees
AS
SELECT * FROM atguigudb.`employees`;
CREATE TABLE departments
AS
SELECT * FROM atguigudb.`departments`;

USE test16_var_cur;
#无参有返回
#1. 创建函数get_count(),返回公司的员工个数

DELIMITER //

CREATE FUNCTION get_count()
RETURNS INT
DETERMINISTIC
CONTAINS SQL
BEGIN
	RETURN (SELECT COUNT(*) FROM employees);
END //

DELIMITER ;

SELECT get_count();

DELIMITER //
CREATE FUNCTION get_count1() RETURNS INT
DETERMINISTIC
CONTAINS SQL
BEGIN
	DECLARE c INT DEFAULT 0;#定义局部变量
	SELECT COUNT(*) INTO c#赋值
	FROM employees;
	RETURN c;
END //
DELIMITER ;

SELECT get_count1();

#有参有返回
#2. 创建函数ename_salary(),根据员工姓名，返回它的工资

DESC employees;

DELIMITER //

CREATE FUNCTION ename_salary(ename VARCHAR(25))
RETURNS DOUBLE
BEGIN
	DECLARE s DOUBLE;
	SELECT salary INTO s
	FROM employees
	WHERE last_name = ename;
	RETURN s;
END //

DELIMITER ;

SET @en = 'Abel';
SELECT ename_salary(@en);

DELIMITER //
CREATE FUNCTION ename_salary1(emp_name VARCHAR(15))
RETURNS DOUBLE
BEGIN
	SET @sal=0;#定义用户变量
	SELECT salary INTO @sal #赋值
	FROM employees
	WHERE last_name = emp_name;
	RETURN @sal;
END //
DELIMITER ;
#调用
SELECT ename_salary1('Abel');
#3. 创建函数dept_sal() ,根据部门名，返回该部门的平均工资

DESC departments;

DELIMITER //
CREATE FUNCTION dept_sal(dname VARCHAR(30))
RETURNS DOUBLE
BEGIN
	SET @sal=0;
	SELECT AVG(salary) INTO @sal 
	FROM employees e
	JOIN departments d
	ON e.department_id = d.department_id
	WHERE d.department_name = dname;
	RETURN @sal;
END //
DELIMITER ;

SELECT dept_sal('IT');

#4. 创建函数add_float()，实现传入两个float，返回二者之和

DELIMITER //

CREATE FUNCTION add_float(a FLOAT, b FLOAT)
RETURNS FLOAT
BEGIN
	DECLARE s FLOAT;
	SET s = a + b;
	RETURN s;
END //

DELIMITER ;

SET @a := 1.9999, @b := 0.00123;
SELECT add_float(@a, @b);

#1. 创建函数test_if_case()，实现传入成绩，如果成绩>90,返回A，如果成绩>80,返回B，如果成绩>60,返回
#C，否则返回D
#要求：分别使用if结构和case结构实现

DELIMITER //

CREATE FUNCTION test_if_case(score INT)
RETURNS CHAR
BEGIN
	DECLARE res CHAR;
	IF score > 90 
		THEN SET res = 'A';
	ELSEIF score > 80
		THEN SET res = 'B';
	ELSEIF score > 60
		THEN SET res = 'C';
	ELSE
		SET res = 'D';
	END IF;
	RETURN res;
END //

DELIMITER ;

DELIMITER //

CREATE FUNCTION test_if_case1(score INT)
RETURNS CHAR
BEGIN
	DECLARE res CHAR;
	CASE 
	WHEN score > 90 
		THEN SET res = 'A';
	WHEN score > 80
		THEN SET res = 'B';
	WHEN score > 60
		THEN SET res = 'C';
	ELSE
		SET res = 'D';
	END CASE;
	RETURN res;
END //

DELIMITER ;

SELECT test_if_case(66);
SELECT test_if_case1(88);

#2. 创建存储过程test_if_pro()，传入工资值，如果工资值<3000,则删除工资为此值的员工，如果3000 <= 工
#资值 <= 5000,则修改此工资值的员工薪资涨1000，否则涨工资500

DELIMITER //

CREATE PROCEDURE test_if_pro(IN sal DOUBLE)
BEGIN
	IF sal < 3000
		THEN DELETE FROM employees WHERE salary = sal;
	ELSEIF sal <= 5000
		THEN UPDATE employees SET salary = salary + 1000 WHERE salary = sal;
	ELSE 
		UPDATE employees SET salary = salary + 500 WHERE salary = sal;
	END IF;
END //

DELIMITER ;

SELECT * FROM employees;

CALL test_if_pro(2900);
CALL test_if_pro(24000);
CALL test_if_pro(4200);

#3. 创建存储过程insert_data(),传入参数为 IN 的 INT 类型变量 insert_count,实现向admin表中批量插
#入insert_count条记录
CREATE TABLE admin(
id INT PRIMARY KEY AUTO_INCREMENT,
user_name VARCHAR(25) NOT NULL,
user_pwd VARCHAR(35) NOT NULL
);
SELECT * FROM admin;

DELIMITER //

CREATE PROCEDURE insert_data(IN insert_count INT)
BEGIN
	DECLARE i INT DEFAULT 0;
	WHILE i < insert_count DO
		SET i = i + 1;
		INSERT INTO admin(user_name, user_pwd)
		VALUES('name', 'passwd');
	END WHILE;
END // 

DELIMITER ;

CALL insert_data(4);

#创建存储过程update_salary()，参数1为 IN 的INT型变量dept_id，表示部门id；参数2为 IN的INT型变量
#change_sal_count，表示要调整薪资的员工个数。查询指定id部门的员工信息，按照salary升序排列，根
#据hire_date的情况，调整前change_sal_count个员工的薪资，详情如下。

DELIMITER //

CREATE PROCEDURE update_salary(IN dept_id INT, IN change_sal_count INT)
BEGIN
	DECLARE i INT DEFAULT 0;
	DECLARE emp_id INT;
	DECLARE hire_year YEAR;
	
	DECLARE emp_cursor CURSOR FOR 
		SELECT employee_id, YEAR(hire_date) 
		FROM employees
		WHERE department_id = dept_id 
		ORDER BY salary;
	OPEN emp_cursor;
	
	WHILE i < change_sal_count DO
		SET i = i + 1;
		FETCH emp_cursor INTO emp_id, hire_year;
		IF hire_year < 1995
			THEN UPDATE employees SET salary = salary * 1.2 WHERE emp_id = employee_id;
		ELSEIF hire_year <= 1998
			THEN UPDATE employees SET salary = salary * 1.15 WHERE emp_id = employee_id;
		ELSEIF hire_year <= 2001
			THEN UPDATE employees SET salary = salary * 1.1 WHERE emp_id = employee_id;
		ELSE 
			UPDATE employees SET salary = salary * 1.05 WHERE emp_id = employee_id;
		END IF;
	END WHILE;
	
	CLOSE emp_cursor;
END // 

DELIMITER ;



SELECT * FROM employees
WHERE department_id = 90;
CALL update_salary(90, 3)



DELIMITER //
CREATE PROCEDURE update_salary1(IN dept_id INT,IN change_sal_count INT)
BEGIN
	#声明变量
	DECLARE int_count INT DEFAULT 0;
	DECLARE salary_rate DOUBLE DEFAULT 0.0;
	DECLARE emp_id INT;
	DECLARE emp_hire_date DATE;
	#声明游标
	DECLARE emp_cursor CURSOR FOR SELECT employee_id,hire_date FROM employees
	WHERE department_id = dept_id ORDER BY salary ;
	#打开游标
	OPEN emp_cursor;
	WHILE int_count < change_sal_count DO
		#使用游标
		FETCH emp_cursor INTO emp_id,emp_hire_date;
		IF(YEAR(emp_hire_date) < 1995)
		THEN SET salary_rate = 1.2;
		ELSEIF(YEAR(emp_hire_date) <= 1998)
		THEN SET salary_rate = 1.15;
		ELSEIF(YEAR(emp_hire_date) <= 2001)
		THEN SET salary_rate = 1.10;
		ELSE SET salary_rate = 1.05;
		END IF;
		#更新工资
		UPDATE employees SET salary = salary * salary_rate
		WHERE employee_id = emp_id;
		#迭代条件
		SET int_count = int_count + 1;
	END WHILE;
	#关闭游标
	CLOSE emp_cursor;
END //
DELIMITER ;

# 调用
CALL update_salary(50,2);
