/*
 * @Author: lisheng
 * @Date: 2022-10-28 17:23:34
 * @LastEditTime: 2022-11-03 00:40:43
 * @LastEditors: lisheng
 * @Description: 错误信息之 数据库相关
 * @FilePath: /gitee.com/liqiyuworks/jf-go-kit/base/statuscode/db_code.go
 */
package statuscode

var (
	ERROR_MYSQL_CONNECT       = &Status{-100001, "Mysql connect error"}            // mysql 连接错误
	ERROR_MYSQL_CLOSE         = &Status{-100002, "Mysql close error"}              // mysql 关闭错误
	ERROR_MYSQL_AUTO_MIGRATE  = &Status{-100003, "Mysql table auto migrate error"} // mysql 自动表迁移错误
	ERROR_MYSQL_INSERT_RECORD = &Status{-100004, "Mysql insert record error"}      // mysql 插入数据错误
	ERROR_MYSQL_QUERY_RECORD  = &Status{-100005, "Mysql query record error"}       // mysql 查询单条数据错误
	ERROR_MYSQL_FIND_RECORDS  = &Status{-100006, "Mysql find records error"}       // mysql 查询多条数据错误
	ERROR_MYSQL_DELETE_RECORD = &Status{-100007, "Mysql delete record error"}      // mysql 删除数据错误
	ERROR_MYSQL_UPDATE_RECORD = &Status{-100008, "Mysql update record error"}      // mysql 更新数据错误

	ERROR_CK_CONNECT       = &Status{-200001, "Clickhouse connect error"}            // clickhouse 连接错误
	ERROR_CK_CLOSE         = &Status{-200002, "Clickhouse close error"}              // clickhouse 关闭错误
	ERROR_CK_AUTO_MIGRATE  = &Status{-200003, "Clickhouse table auto migrate error"} // clickhouse 自动表迁移错误
	ERROR_CK_INSERT_RECORD = &Status{-200004, "Clickhouse insert record error"}      // clickhouse 插入数据错误
	ERROR_CK_QUERY_RECORD  = &Status{-200005, "Clickhouse query record error"}       // clickhouse 查询单条数据错误
	ERROR_CK_FIND_RECORDS  = &Status{-200006, "Clickhouse find records error"}       // clickhouse 查询多条数据错误
	ERROR_CK_DELETE_RECORD = &Status{-200007, "Clickhouse delete record error"}      // clickhouse 删除数据错误
	ERROR_CK_UPDATE_RECORD = &Status{-200008, "Clickhouse update record error"}      // clickhouse 更新数据错误

	ERROR_PG_CONNECT       = &Status{-300001, "pg connect error"}            // pg 连接错误
	ERROR_PG_CLOSE         = &Status{-300002, "pg close error"}              // pg 关闭错误
	ERROR_PG_AUTO_MIGRATE  = &Status{-300003, "pg table auto migrate error"} // pg 自动表迁移错误
	ERROR_PG_INSERT_RECORD = &Status{-300004, "Pg insert record error"}      // Pg 插入数据错误
	ERROR_PG_QUERY_RECORD  = &Status{-300005, "Pg query record error"}       // Pg 查询单条数据错误
	ERROR_PG_FIND_RECORDS  = &Status{-300006, "Pg find records error"}       // Pg 查询多条数据错误
	ERROR_PG_DELETE_RECORD = &Status{-300007, "Pg delete record error"}      // Pg 删除数据错误
	ERROR_PG_UPDATE_RECORD = &Status{-300008, "Pg update record error"}      // Pg 更新数据错误

	ERROR_MONGO_CONNECT       = &Status{-400001, "Mongo connect error"}            // Mongo 连接错误
	ERROR_MONGO_CLOSE         = &Status{-400002, "Mongo close error"}              // Mongo 关闭错误
	ERROR_MONGO_AUTO_MIGRATE  = &Status{-400003, "Mongo table auto migrate error"} // Mongo 自动表迁移错误
	ERROR_MONGO_INSERT_RECORD = &Status{-400004, "Mongo insert record error"}      // Mongo 插入数据错误
	ERROR_MONGO_EXIST_RECORD  = &Status{-400005, "Mongo exist record error"}       // Mongo 是否存在该条数据错误
	ERROR_MONGO_QUERY_RECORD  = &Status{-400006, "Mongo query record error"}       // Mongo 查询单条数据错误
	ERROR_MONGO_FIND_RECORDS  = &Status{-400007, "Mongo find records error"}       // Mongo 查询多条数据错误
	ERROR_MONGO_DELETE_RECORD = &Status{-400008, "Mongo delete record error"}      // Mongo 删除数据错误
	ERROR_MONGO_UPDATE_RECORD = &Status{-400009, "Mongo update record error"}      // Mongo 更新数据错误

	ERROR_TDENGINE_CONNECT       = &Status{-500001, "TDengine connect error"}       // TDengine 连接错误
	ERROR_TDENGINE_CLOSE         = &Status{-500002, "TDengine close error"}         // TDengine 关闭错误
	ERROR_TDENGINE_INSERT_RECORD = &Status{-500003, "TDengine insert record error"} // TDengine 插入数据错误
	ERROR_TDENGINE_EXIST_RECORD  = &Status{-500004, "TDengine exist record error"}  // TDengine 是否存在该条数据错误
	ERROR_TDENGINE_QUERY_RECORD  = &Status{-500005, "TDengine query record error"}  // TDengine 查询单条数据错误
	ERROR_TDENGINE_FIND_RECORDS  = &Status{-500006, "TDengine find records error"}  // TDengine 查询多条数据错误
	ERROR_TDENGINE_DELETE_RECORD = &Status{-500007, "TDengine delete record error"} // TDengine 删除数据错误
	ERROR_TDENGINE_UPDATE_RECORD = &Status{-500008, "TDengine update record error"} // TDengine 更新数据错误

	ERROR_REDIS_CONNECT         = &Status{-600001, "Redis connect error"}         // Redis 连接错误
	ERROR_REDIS_CLOSE           = &Status{-600002, "Redis close error"}           // Redis 关闭错误
	ERROR_REDIS_STRING_GET      = &Status{-600003, "Redis string get error"}      // Redis string get 错误
	ERROR_REDIS_STRING_SET      = &Status{-600004, "Redis string set error"}      // Redis string set  错误
	ERROR_REDIS_SET_ADD         = &Status{-600005, "Redis set add error"}         // Redis set add  错误
	ERROR_REDIS_SET_MEMBERS     = &Status{-600006, "Redis set members error"}     // Redis Set members  错误
	ERROR_REDIS_SET_ZANGE       = &Status{-600007, "Redis set zange error"}       // Redis Set zange  错误
	ERROR_REDIS_STRING_DELETE   = &Status{-600008, "Redis string delete error"}   // Redis string delete  错误
	ERROR_REDIS_HASH_GET        = &Status{-600009, "Redis hash get error"}        // Redis hash get 错误
	ERROR_REDIS_HASH_GET_FIELDS = &Status{-600010, "Redis hash get fields error"} // Redis hash get fields 错误
	ERROR_REDIS_HASH_SET        = &Status{-600011, "Redis hash set error"}        // Redis hash set 错误
	ERROR_REDIS_HASH_SET_FIELDS = &Status{-600012, "Redis hash set fields error"} // Redis hash set fields 错误
	ERROR_REDIS_HASH_DELETE     = &Status{-600013, "Redis hash delete error"}     // Redis hash delete  错误
)
