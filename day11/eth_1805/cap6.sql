-- 索引创建
create index deptno_index on emp(deptno);
-- 显示索引信息
show index from emp;
-- 删除索引
drop index city_idx on emp;

-- 向emp中插入大量数据
delimiter //

create PROCEDURE proc_emp() 
BEGIN
set @x = 0;
label: LOOP
	set @x = @x +1;
	IF @x=100000 THEN
		LEAVE label; 
	END IF; 
	INSERT INTO emp(ename,hirdate, deptno, sal, age) values(CONCAT('emp',CAST(@x as char)), now(), @x,10000 * RAND(), 40 * RAND());
END LOOP label;

END
// 
delimiter ;
call proc_emp();
-- 在添加索引之前进行查询，获取查询时间
-- 0.066s
select * from emp WHERE deptno = 90000;
-- 添加索引
CREATE INDEX dp_idx ON emp(deptno);
select * from emp;
-- 在添加索引之后进行查询，获取查询时间
-- 0.018s
SHOW INDEX FROM emp; \G

drop index deptno_index on emp;

delete from emp;


