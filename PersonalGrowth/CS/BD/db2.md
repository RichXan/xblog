# 优化sql

可以通过**存储结构**来创建几十万几百万条数据

所有查询语句前加上 **EXPLAIN** 可以分析这条sql执行的效率

## 1.优化带排序的分页查询

```sql
-- 浅分页 # 1.200s
select *user_no, user_name, score from student order by score desc limit 5,20;

-- 深分页 # 2.029s
select *user_no, user_name, score from student order by score desc limit 80000,20;
```

方法1：增加order by排序字段索引脚本
```sql
-- 增加order by排序字段索引脚本
ALTER TABLE student ADD index idx_score(score)

-- 浅分页 # 0.038s
select *user_no, user_name, score from student order by score desc limit 5,20;

-- 深分页 # 1.911s
select *user_no, user_name, score from student order by score desc limit 80000,20;

-- 3.165s
select user_no, user_name, score from student FORCE INDEX(idx_score) order by score desc limit 80000, 20; 
```

方法2：增加联合索引

缺点：select多增加一个字段的时候，sql语句无法走联合索引查询
```sql
-- order by 和 select字段加上联合索引
ALTER TABLE student ADD index idx_score_name_no (score,user_name, user_no);

-- 深分页 # 0.048s
select *user_no, user_name, score from student order by score desc limit 80000,20; 

```

### 总结

带排序的分页查询优化

1. 浅分页可以给order by字段添加索引
2. 深分页可以给order by和select字段添加联合索引
3. 可以通过手动回表，强制去走索引
4. 从业务方着手，去限制他的分页查询或者修改前后端交互（将煤业最后一条数据的id和分数传递过来）


## datetime和timestamp的区别
| 类型        | 占用字节 | 存储格式                     |
| ----------- | -------- | ---------------------------- |
| datetime    | 8        | YYYY-MM-DD HH:mm:ss           |
| timestamp   | 4        | YYYY-MM-DD HH:mm:ss           |


timestamp 只占 4 个字节，而且是以utc的格式储存， 它会自动检索当前时区并进行转换。 datetime以 8 个字节储存，不会进行时区的检索. 也就是说，对于timestamp来说，如果储存时的时区和检索时的时区不一样，那么拿出来的数据也不一样。 对于datetime来说，存什么拿到的就是什么。 还有一个区别就是如果存进去的是NULL，timestamp会自动储存当前时间，而 datetime会储存 NULL。
