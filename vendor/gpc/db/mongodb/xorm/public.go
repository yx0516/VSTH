package xorm

import (
	"errors"
	"fmt"
	"strings"

	"gpc/app/log/xmgo"
	"gpc/util/ref"
	"gpc/util/tag"

	mdl "gpc/db/mongodb/xorm/model"

	plib "gpc/publib"
)

const (
	CST_FILTER_OPT_WHITE uint8 = 0                    // 白名单过滤 字段
	CST_FILTER_OPT_BLACK uint8 = 1                    // 黑名单过滤 字段
	CST_FILTER_OPT_NONE  uint8 = 2                    // 不过滤 字段
	CST_FILTER_OPT_DEF   uint8 = CST_FILTER_OPT_WHITE // 默认过滤规则

	CST_COLL_SYS_ID_COUNTER = "SYS_ID_COUNTER" // 存储自增Id的集合名称
)

var (
	// 名称 转换 函数：驼峰大小写 字符串转下划线连接格式【 XxYy to xx_yy 】
	funcNameConvertSnake func(string) string = ref.FieldNameConvertFunc(ref.CST_FNAME_CONVERT_TO_SNAKE)

	// 日志对象【公开，如果有需要，在 init 阶段可覆盖】
	Log = xmgo.Log

	obj = struct { // 内部 对象
		pack PackString // 包装 字符串
		chk  ChkField   // 检查 字段
		err  *ErrType   // 错误 对象
	}{
		err: NewErrType(Log),
	}
)

// Hook 更新函数【vSelf 是 Copy 读取后的对象】
type HookUpdate func(fieldValues MapSI, vSelf interface{}) error

//-----------------------------------------------------------------------------------------------------------//

// 获取 集合 名称
func GetCollName(objetc mdl.Model) (name string, err error) {
	// 优先查询 方法指定 的名称
	if name, err = tag.GetStructName(objetc, "CollName", nil); err == nil {
		return
	}
	return tag.GetStructName(objetc, "", funcNameConvertSnake)
}

//-----------------------------------------------------------------------------------------------------------//

// 功能 未实现
func NewFuncNotImplemented() error {
	return errors.New("this feature is not yet implemented!")
}

// 生成 过滤匹配 id|name 错误
func NewErrFilterMatchByIdName(fields ...string) error {
	arr := []string{"id", "name"}
	return NewErrFilterMatch(append(arr, fields...)...)
}

// 生成 过滤匹配 错误
func NewErrFilterMatch(fields ...string) error {
	return fmt.Errorf("filter match error.(%s)", strings.Join(fields, "|"))
}

//-----------------------------------------------------------------------------------------------------------//

// 字符串-接口 的 Map
type MapSI map[string]interface{}

func (self MapSI) Contains(key string) (ok bool) {
	_, ok = self[key]
	return
}

func (self MapSI) Len() int {
	return len(self)
}

func (self MapSI) Delete(keys ...string) {
	for _, key := range keys {
		delete(self, key)
	}
}

// 只保留指定键
func (self MapSI) Retain(keys ...string) {
	delKeys := self.Keys()
	delKeys.Deletes(true, keys...)
	self.Delete(delKeys...)
}

// 获取所有的键
func (self MapSI) Keys() (keys plib.StringArray) {
	for key := range self {
		keys.Add(key)
	}
	return
}

// 获取指定键内容重新生成一个新的 MapSI
func (self MapSI) GetMap(keys ...string) (m MapSI) {
	m = make(MapSI)
	for _, key := range keys {
		if val, ok := self[key]; ok {
			m[key] = val
		}
	}
	return
}

func (self MapSI) Add(key string, val interface{}) {
	self[key] = val
}

// 追加一个 Map
func (self MapSI) Append(m map[string]interface{}) {
	for key, val := range m {
		self[key] = val
	}
}

// 清空
func (self MapSI) Clear() {
	self.Delete(self.Keys()...)
}

//-----------------------------------------------------------------------------------------------------------//
