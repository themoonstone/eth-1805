-- dept 
-- emp
-- 使用inner join内连接
select * from emp inner join dept on emp.deptno=dept.deptno;

-- 使用left join查询
select * from emp left JOIN dept on emp.deptno=dept.deptno;

-- 使用right join查询
select * from emp right JOIN dept on emp.deptno=dept.deptno;

-- 子查询
-- in
select * from emp WHERE deptno in(select deptno FROM dept);

select * from emp where  sal > 1000;
-- exists
select * from emp WHERE exists (select deptno from dept where deptno=1);

-- 合并查询结果(默认去重)
select deptno from emp union select deptno from dept;

-- 合并查询结果
select deptno from emp union all select deptno from dept;


-- 别名(as 可以省略)
select ename as e from emp as e right JOIN dept as d on e.deptno=d.deptno;

-- 正则
select * from emp where ename REGEXP 'n$';

-- update 
update emp set ename = 'Kobe' where ename='Tom';

select * from emp;

-- 删除
DELETE from emp where ename='Kobe';

DELETE emp, dept from emp, dept where emp.deptno = 3 and emp.deptno = dept.deptno;

-- 查询删除之后的结果
select * from dept;