package tag

const (
	CST_STATUS_OK   uint8 = iota // 0 正常
	CST_STATUS_FAIL              // 1 不通过
	CST_STATUS_ERR               // 2 错误

	CST_TAG_CHK      string = "chk"   // 检查 字段 值 是否合法
	CST_TAG_NULL     string = "null"  // 检查 字段 值 是否为空
	CST_TAG_NOT_NULL string = "!null" // 检查 字段 值 是否为非空
	CST_TAG_DEF      string = "def"   // 对于非空字段可以指定设置默认值
	CST_TAG_VAL      string = "val"   // 结构体 可指定 获取值的字段名
	CST_LOGIC_AND    string = "and:"  // and 逻辑
	CST_LOGIC_OR     string = "or:"   // or 逻辑
	CST_EXP_IGNORE   string = "-"     // 忽略检查

	CST_TAG_CFG      string = "cfg"   // 配置
	CST_CFG_NOT_TRIM string = "!trim" // 对值不执行 trim 操作

	CST_TAG_XORM      string = "xorm"   // xorm 注解
	CST_XORM_UPDATE   string = "update" // 自定义可 更新 字段
	CST_XORM_QUERY    string = "query"  // 自定义可 查询 字段
	CST_XORM_ZERO_NIL string = "ztnil"  // 如果 Update 时的字段值为0,则设置为nil【在更新键值fvs阶段】
	CST_XORM_NIL_DEL  string = "nildel" // 如果 Update 时的字段值为 nil 则 删除掉
	CST_XORM_ADD_FK   string = "addfk"  // 添加 外键 查询字段

	// 结构体字段值可以设置当不能获取到值时(默认会为 nil)，可以设置指定类型的 默认值
	CST_VAL_TYPE_NUM string = "num" // 默认值  0
	CST_VAL_TYPE_STR string = "str" // 默认值 ""
)
