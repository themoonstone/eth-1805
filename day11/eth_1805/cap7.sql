-- 存储过程说明
delimiter //
create PROCEDURE test2(in inparams int, out outparams VARCHAR(20)) 
BEGIN
declare var char(10);
if inparams = 1 then
set var = 'hello';
else 
set var = 'world';
end if;
insert into emp(ename) values(var);
END //
delimiter ;

-- 调用 存储过程
call test2(1, @OUT);


select * from emp;

show PROCEDURE status like 'test2';

show create table emp;

delimiter //

create PROCEDURE test_loop() 
BEGIN
set @x = 0;
label: LOOP
	set @x = @x +1;
	IF @x=100 THEN
		LEAVE label; 
	END IF; 
	INSERT INTO emp(ename, deptno) values('test', 100);
END LOOP label;

END
// 
delimiter ;
call test_loop();

-- loop 与leave结合使用

SELECT * from emp;