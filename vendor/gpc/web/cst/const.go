package cst

const (
	// Y | N 字符常量
	Y string = "Y"
	N string = "N"

	// 状态
	STATUS_NOT_SET int = 0         // 未设置
	STATUS_ON      int = 1         // 正常
	STATUS_OFF     int = 2         // 禁止
	STATUS_DEF     int = STATUS_ON // 默认值

	// 激活 状态
	STATE_NOT_SET   int = 0 // 未设置
	STATE_INACTIVE  int = 1 // 未激活
	STATE_ACTIVATED int = 2 // 已激活
	STATE_INVALID   int = 3 // 无效

	// 认证 类型
	AUTH_TYPE_NAME    int = 0 // 用户名
	AUTH_TYPE_EMAIL   int = 1 // 邮箱
	AUTH_TYPE_MOBILE  int = 2 // 手机
	AUTH_TYPE_INVALID int = 3 // 无效
	AUTH_TYPE_ID      int = 4 // 用户 ID

	SCHEME_HTTP  string = "http://"
	SCHEME_HTTPS string = "https://"

	// 字段名
	FIELD_ID            string = "id"
	FIELD_NAME          string = "name"
	FIELD_STATUS        string = "status"
	FIELD_REMARK        string = "remark"
	FIELD_PID           string = "pid"
	FIELD_TAG           string = "tag"
	FIELD_TAG_ID        string = "tag_id"
	FIELD_SEX           string = "sex"
	FIELD_PWD           string = "pwd"
	FIELD_NICK          string = "nick"
	FIELD_EMAIL         string = "email"
	FIELD_EMAIL_STATE   string = "email_state"
	FIELD_MOBILE        string = "mobile"
	FIELD_MOBILE_STATE  string = "mobile_state"
	FIELD_CREATE_TIME   string = "create_time"
	FIELD_UPDATE_TIME   string = "update_time"
	FIELD_CHECK_TIME    string = "check_time"
	FIELD_IP            string = "ip"
	FIELD_TIME          string = "time"
	FIELD_LOGIN_IP      string = "login_ip"
	FIELD_LOGIN_TIME    string = "login_time"
	FIELD_LAST_IP       string = "last_ip"
	FIELD_LAST_TIME     string = "last_time"
	FIELD_REG_IP        string = "reg_ip"
	FIELD_REG_TIME      string = "reg_time"
	FIELD_ALLOW_IP      string = "allow_ip"
	FIELD_SECRET_KEY    string = "secret_key"
	FIELD_API_STATUS    string = "api_status"
	FIELD_MODULE_ID     string = "module_id"
	FIELD_FUNC_ID       string = "func_id"
	FIELD_GROUP_ID      string = "group_id"
	FIELD_PERM_ID       string = "perm_id"
	FIELD_ROLE_ID       string = "role_id"
	FIELD_USER_ID       string = "user_id"
	FIELD_EXPIRE        string = "expire"
	FIELD_BAN_EXPIRE    string = "ban_expire"
	FIELD_PWD_EXPIRE    string = "pwd_expire"
	FIELD_SUCCESS_COUNT string = "success_count"
	FIELD_FAIL_COUNT    string = "fail_count"
	FIELD_FAIL_IP       string = "fail_ip"
	FIELD_FAIL_TIME     string = "fail_time"
	FIELD_QUESTION      string = "question"
	FIELD_ANSWER        string = "answer"
	FIELD_RESET         string = "reset"
	FIELD_DATA          string = "data"
	FIELD_DATA_DEF      string = "data_def"
	FIELD_LIMIT_COUNT   string = "limit_count"
	FIELD_LIMIT_TIME    string = "limit_time"
	FIELD_LIMIT_IP      string = "limit_ip"
	FIELD_ADDRESS       string = "address"
	FIELD_PORT          string = "port"
	FIELD_SCHEME        string = "scheme"
	FIELD_APP_ID        string = "app_id"
	FIELD_URI_LOGIN     string = "uri_login"
	FIELD_URI_QUIT      string = "uri_quit"
	FIELD_URI_SYNC      string = "uri_sync"
	FIELD_LOGIN_KEY     string = "login_key"
	FIELD_USER_KEY      string = "user_key"
	FIELD_TICKET_KEY    string = "ticket_key"
	FIELD_ITEMS         string = "items"
	FIELD_INFO_ID       string = "info_id"
	FIELD_MD5           string = "md5"
	FIELD_DATA_MD5      string = "data_md5"
	FIELD_ORGANIZE_ID   string = "organize_id"
	FIELD_LEVEL         string = "level"
	FIELD_OFFICE_ID     string = "office_id"
	FIELD_MODE          string = "mode" // 模式
	FIELD_SENDER        string = "sender"
	FIELD_SUBJECT       string = "subject"
	FIELD_BODY          string = "body"
	FIELD_MAIL_ID       string = "mail_id"
	FIELD_CHECK_ID      string = "check_id"
	FIELD_TOTAL         string = "total"
	FIELD_HAS_READ      string = "has_read"
	FIELD_DEVICE_ID     string = "device_id"
	FIELD_FILE_PATH     string = "file_path"
	FIELD_JSON_METE     string = "json_mete"
	FIELD_TYPE          string = "type"
	FIELD_TYPE_ID       string = "type_id"
	FIELD_GRAPH_ID      string = "graph_id"
	FIELD_FILTER_ID     string = "filter_id"
	FIELD_PATH          string = "path"

	// SQL 关键字: from | match

	// tag 相关
	TAG_IGNORE    string = "-"    // 忽略tag内容
	TAG_BEEGO_ORM string = "orm"  // beego orm
	TAG_MGO_BSON  string = "bson" // mongodb

	// beego orm 相关
	BEEGO_ORM_TABLE_NAME  string = "TableName"                    // orm 表名 函数名
	BEEGO_ORM_COND_JOIN   string = "__"                           // orm 查询属性的连接符
	BEEGO_ORM_COND_ISNULL string = BEEGO_ORM_COND_JOIN + "isnull" // orm 查询 isnull

	IP_UNLIMITED string = "*" // IP无限制
	ENCRYP_TAG   string = "*" // 加密标记

	NO_PERM = "No permission."
)
