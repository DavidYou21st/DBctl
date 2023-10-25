* 该目录存放表关联配置文件，
* 创建以config.yaml的connection名对应的```connection名.json```文件
* json结构：

```json
{
  "main_table": [
    {
      "table": "relation_tableA",
      "type": "has_one",
      "foreignKey": "",
      "relationship_table": "",
      "join_foreign_key": "",
      "join_target_key": ""
    },
    {
      "table": "relation_tableB",
      "type": "has_many",
      "foreignKey": ""
    }
  ]
}
```

* json 字段说明
```
main_table  // 主表名
relation_tableA //要关联的表名A
type //类型支持has_one,has_many,belongs_to,many2many
foreignKey //外键
reference_key //当many2many时，连接表名
join_foreign_key //当many2many时，本表在连接表的外键
join_target_key //当many2many时，关联表在连接表的外键
```

* 类型说明用法
```
has_one:
    一个用户只能报名参加一次某个活动
has_many:
    一个用户拥有多张银行卡
belongs_to:
    银行卡属于用户
many2many:
    多对多
```

* 为什么要使用关系模型
```
优化代码，减少left join查询，可预加载sql
```