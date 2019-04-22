SELECT * from country;

SELECT * from country where number > 15000000;

-- 求每个地区的总人口、平均人口
select sum(number), AVG(number), area from country GROUP BY area;

-- 筛选出平均人口大于15000000的地区
select sum(number), AVG(number), area from country GROUP BY area HAVING AVG(number) > 15000000;

-- 求每个地区的总人口、平均人口，将结果进行排序
select sum(number), AVG(number), area from country GROUP BY area ORDER BY AVG(number);

-- 结果去重
select distinct(area) from country;

-- limit
-- limit(m, n):从第m条开始，取n条数据
select * from country limit 2,2;

-- 聚合函数
-- count
select count(number) from country;
-- max
select max(number) from country;

