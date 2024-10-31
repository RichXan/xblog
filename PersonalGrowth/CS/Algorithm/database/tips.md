# Mysql
## 函数
常见函数
> SUM(x)： 返回x列的和
> 
> SUM(DISTINCT x)： 返回x列的不重复值的和
> 
> AVG(x)： 返回x列的平均值，通常与 GROUP BY 搭配使用
> 
> COUNT(x)： 返回x列的行数
> 
> MAX(x)： 返回x列的最大值
> 
> MIN(x)： 返回x列的最小值
> 
> ROUND(x, y)： 返回x列的值四舍五入到小数点后y位的结果
> 
> MOD(x, y)： 返回x列的值除以y的余数
> 
> CONCAT(x, y)： 返回x列和y列的拼接结果
> 


[字符串函数](https://dev.mysql.com/doc/refman/8.4/en/string-functions.html)
> CHAR_LENGTH(x)： 返回x列的**字符串**长度
>   
> LENGTH(x)： 返回x列的字节长度
> 
> SUBSTRING(x, y, z)： 返回x列的从y位置开始的z个字符
> 
> 

[data&time函数](https://dev.mysql.com/doc/refman/8.4/en/date-and-time-functions.html)
> NOW()： 返回当前日期和时间
> 
> CURDATE()： 返回当前日期
> 
> CURTIME()： 返回当前时间
> 
> DATEDIFF(x, y)： 返回x列和y列的日期差
> 
> DATE_FORMAT(x, y)： 返回x列的日期按照y格式化
> 
> DAYOFWEEK(x)： 返回x列的日期是星期几
> 
MySQL 使用三值逻辑 —— TRUE, FALSE 和 UNKNOWN。任何与 NULL 值进行的比较都会与第三种值 UNKNOWN 做比较。这个“任何值”包括 NULL 本身！这就是为什么 MySQL 提供 IS NULL 和 IS NOT NULL 两种操作来对 NULL 特殊判断。对null使用比较运算符，最终结果都是unknown。



## 联表查询
- left join 左连接，返回左表的全部记录和右表的全部记录，如果左表的记录在右表中没有匹配的记录，则右表的对应字段值为null。

- right join 右连接，返回右表的全部记录和左表的全部记录，如果右表的记录在左表中没有匹配的记录，则左表的对应字段值为null。

- inner join 内连接，返回两表中都存在的记录，如果两表的记录在对方中没有匹配的记录，则不返回。

- full join 全连接，返回两表中的全部记录，如果两表的记录在对方中没有匹配的记录，则对应的字段值为null。

- cross join 交叉连接，返回两表的笛卡尔积，即两表的记录两两组合。

## 查询技巧
1. HAVING子句通常与GROUP BY子句一起使用，以根据指定的条件过滤分组。如果省略GROUP BY子句，则HAVING子句的行为与WHERE子句类似。
```sql
-- 找出至少有五个直接下属的经理。
select Manager.Name as Name
from
Employee as Manager join Employee as Report
on Manager.Id = Report.ManagerId
group by Manager.Id
having count(Report.Id) >= 5
```

2. 使用 DISTINCT 关键字来从表 Views 中检索唯一元素。
```sql
-- 找出所有作者的唯一id，这些作者的id与他们自己文章的id相同。
SELECT
    DISTINCT author_id as id
FROM 
    Views 
WHERE 
    author_id = viewer_id;
```
3. 使用 LEFT JOIN 和 IFNULL 函数来计算每个用户的确认率。
> [MySQL中IF()、IFNULL()、NULLIF()、ISNULL()函数的使用](https://blog.csdn.net/pan_junbiao/article/details/85928004)
>
> MySQL 使用三值逻辑 —— TRUE, FALSE 和 UNKNOWN。任何与 NULL 值进行的比较都会与第三种值 UNKNOWN 做比较。这个“任何值”包括 NULL 本身！这就是为什么 MySQL 提供 IS NULL 和 IS NOT NULL 两种操作来对 NULL 特殊判断。
>
> 对null使用比较运算符，最终结果都是unknown。
```sql
-- 计算每个用户的确认率。
SELECT
    s.user_id,
    ROUND(IFNULL(AVG(c.action='confirmed'), 0), 2) AS confirmation_rate
FROM
    Signups AS s
LEFT JOIN
    Confirmations AS c
ON
    s.user_id = c.user_id
GROUP BY
    s.user_id
```