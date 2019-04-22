select * from emp where deptno = 11;
set autocommit=0;
-- 开启tx
start TRANSACTION;

-- 插入一条数据
insert into emp values('tx', '20190422',11, 3000,19,'BJ','M');
-- 手动提交
commit


-- commit and chain
insert into emp values('tx', '20190422',13, 3000,19,'BJ','M');
select * from emp where deptno = 13;

-- 启动
start TRANSACTION;

insert into emp values('tx', '20190422',14, 3000,19,'BJ','M');
select * from emp where deptno = 14;
-- 提交当前事务之后开启一个新的事务
commit and chain;
insert into emp values('tx', '20190422',14, 3000,19,'BJ','M');
commit;

-- 事务回滚
select * from emp where deptno = 15;

start TRANSACTION;
insert into emp values('sp', '20190422',14, 3000,19,'BJ','M');

insert into emp values('sp', '20190422',14, 3000,19,'BJ','M');
SAVEPOINT test;
insert into emp values('sp2', '20190422',14, 3000,19,'BJ','M');

ROLLBACK to savepoint test;

commit 

select * from emp where ename='sp2';