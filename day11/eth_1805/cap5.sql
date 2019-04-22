SELECT
	abs( - 0.8 );
SELECT
	RAND( ) -- 字符串内置函数
SELECT
	CHAR_LENGTH( "全栈" );
SELECT
	length( "全栈" );
-- 日期和时间
SELECT
	CURDATE( );
SELECT MONTH
	( "20190422" );
-- 条件判断函数
-- if
SELECT
IF
	( sal > 2000, 'high', 'low' ) 
FROM
	emp;
	
-- case
SELECT
CASE
	
WHEN
	sal <= 2000 THEN
	'low' ELSE 'high' 
END 
FROM
emp;

-- 系统信息
select version();

-- md5
select MD5("themoonstone")