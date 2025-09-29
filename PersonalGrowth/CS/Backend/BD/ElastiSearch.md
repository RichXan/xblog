# ElasticSearch

## 与MySQL关系型数据库的区别

| MySQL     | ElasticSearch | 说明| 
| ---------------- | ------------- | ------------- |
| 数据库(database)           | 索引(index)     | 索引(index)，就是文档的集合，类似数据库的表(table)| 
| 表(table)               | 类型(type)     | 类型（Type），就是文档的集合，类似数据库中的表（Table）| 
| 表结构(schema)           | 映射(mapping)   | 映射（Mapping），就是文档的结构，类似数据库中的表结构（Schema）| 
| 行(rows)         | 文档(document) | 文档（Document），就是一条条的数据，类似数据库中的行（Row），文档都是JSON格式 | 
| 列、字段(columns) | 字段(field)     | 字段（Field），就是JSON文档中的字段，类似数据库中的列（Column）|
| SQL           | DSL   | SQL是关系型数据库中的概念，DSL是ElasticSearch中的概念|


## 索引库操作
### 1. Mapping映射关系
mapping是对索引库中文档的约束，常见的mapping属性包括：

- type：字段数据类型，常见的简单类型有：

    - 字符串：text（可分词的文本）、keyword（精确值，例如：品牌、国家、ip地址）
        - keyword类型只能整体搜索，不支持搜索部分内容

    - 数值：long、integer、short、byte、double、float、

    - 布尔：boolean

    - 日期：date

    - 对象：object

- index：是否创建索引，默认为true

- analyzer：使用哪种分词器

- properties：该字段的子字段


### 2. 索引库的CRUD
> 索引库的CRUD操作，类似于MySQL数据库的CRUD操作
> - 创建索引库：PUT /索引库名
> - 查询索引库：GET /索引库名
> - 删除索引库：DELETE /索引库名
> - 修改索引库（添加字段）：PUT /索引库名/_mapping
> - 修改索引库（删除字段）：DELETE /索引库名/_mapping/字段名


参考文档
> [ElasticSearch(ES从入门到精通一篇就够了)](https://www.cnblogs.com/buchizicai/p/17093719.html)