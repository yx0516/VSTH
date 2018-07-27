package cst

const (
	// Y | N 字符常量
	Y string = "Y"
	N string = "N"

	// 状态
	Status_NotSet int = 0         // 未设置
	Status_On     int = 1         // 正常
	Status_Off    int = 2         // 禁止
	Status_Def    int = Status_On // 默认值

	// 激活 状态
	State_NotSet    int = 0 // 未设置
	State_Inactive  int = 1 // 未激活
	State_Activated int = 2 // 已激活
	State_Invalid   int = 3 // 无效

	// 认证 类型
	AuthType_Name    int = 0 // 用户名
	AuthType_Email   int = 1 // 邮箱
	AuthType_Mobile  int = 2 // 手机
	AuthType_Invalid int = 3 // 无效
	AuthType_Id      int = 4 // 用户 id

	Scheme_HTTP  string = "http://"
	Scheme_HTTPS string = "https://"

	// 字段名
	Field_Id           string = "id"
	Field_Name         string = "name"
	Field_Status       string = "status"
	Field_Remark       string = "remark"
	Field_Pid          string = "pid"
	Field_Tag          string = "tag"
	Field_TagId        string = "tagid"
	Field_Sex          string = "sex"
	Field_Pwd          string = "pwd"
	Field_Nick         string = "nick"
	Field_Email        string = "email"
	Field_EmailState   string = "emailstate"
	Field_Mobile       string = "mobile"
	Field_MobileState  string = "mobilestate"
	Field_CreateTime   string = "createtime"
	Field_UpdateTime   string = "updatetime"
	Field_CheckTime    string = "checktime"
	Field_IP           string = "ip"
	Field_Time         string = "time"
	Field_LoginIP      string = "loginip"
	Field_LoginTime    string = "logintime"
	Field_LastIP       string = "lastip"
	Field_LastTime     string = "lasttime"
	Field_RegIP        string = "regip"
	Field_RegTime      string = "regtime"
	Field_AllowIP      string = "allowip"
	Field_SecretKey    string = "secretkey"
	Field_ApiStatus    string = "apistatus"
	Field_ModuleId     string = "moduleid"
	Field_FuncId       string = "funcid"
	Field_GroupId      string = "groupid"
	Field_PermId       string = "permid"
	Field_RoleId       string = "roleid"
	Field_UserId       string = "userid"
	Field_Expire       string = "expire"
	Field_BanExpire    string = "banexpire"
	Field_PwdExpire    string = "pwdexpire"
	Field_SuccessCount string = "successcount"
	Field_FailCount    string = "failcount"
	Field_FailIP       string = "failip"
	Field_FailTime     string = "failtime"
	Field_Question     string = "question"
	Field_Answer       string = "answer"
	Field_Reset        string = "reset"
	Field_Data         string = "data"
	Field_DataDef      string = "datadef"
	Field_LimitCount   string = "limitcount"
	Field_LimitTime    string = "limittime"
	Field_LimitIP      string = "limitip"
	Field_Address      string = "address"
	Field_Port         string = "port"
	Field_Scheme       string = "scheme"
	Field_AppId        string = "appid"
	Field_UriLogin     string = "urilogin"
	Field_UriQuit      string = "uriquit"
	Field_UriSync      string = "urisync"
	Field_LoginKey     string = "loginkey"
	Field_UserKey      string = "userkey"
	Field_TicketKey    string = "ticketkey"
	Field_Items        string = "items"
	Field_InfoId       string = "infoid"
	Field_Md5          string = "md5"
	Field_DataMd5      string = "datamd5"
	Field_OrganizeId   string = "organizeid"
	Field_Level        string = "level"
	Field_OfficeId     string = "officeid"
	Field_Mode         string = "mode"
	Field_Sender       string = "sender"
	Field_Subject      string = "subject"
	Field_Body         string = "body"
	Field_MailId       string = "mailid"
	Field_CheckId      string = "checkid"
	Field_Total        string = "total"
	Field_HasRead      string = "hasread"
	Field_DeviceId     string = "deviceid"
	Field_File         string = "file"
	Field_Path         string = "path"
	Field_FilePath     string = "filepath"
	Field_JsonMete     string = "jsonmete"
	Field_Type         string = "type"
	Field_TypeId       string = "typeid"
	Field_GraphId      string = "graphid"
	Field_FilterId     string = "filterid"
	Field_ChemId       string = "chemid"
	Field_InChiKey     string = "inchikey"
	Field_MolId        string = "molid"
	Field_Value        string = "value"
)
