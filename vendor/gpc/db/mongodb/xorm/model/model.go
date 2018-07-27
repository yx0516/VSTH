package model

import (
	"reflect"

	"gopkg.in/mgo.v2/bson"

	plib "gpc/publib"
)

//-----------------------------------------------------------------------------------------------------------//

// 模型 接口
type Model interface {
	// 初始化
	Init()

	// 是否存在指定ID
	ExistId() (ok bool, err error)

	// 查询条件 id 或 X , 优先使用 id
	MakeMatch() (bson.M, error)

	// 插入时检查参数
	CheckInsert() error

	// 插入
	Insert() (int, error)

	// 读取数据【fields 指定只读取哪些字段的数据返回回来】【没有数据会返回 err:not found】
	Read(fields ...string) error

	// 通过 Id 来读取【fields 指定只读取哪些字段的数据返回回来】【没有数据会返回 err:not found】
	ReadById(fields ...string) error

	// 更新 匹配 指定字段
	UpdateFields(match bson.M, fields ...string) error

	// 更新 匹配
	UpdateFieldValues(match, fieldValues bson.M) error

	// 更新 指定字段【勾入处理函数】
	UpdateByFields(fields plib.StringArray) error

	// 自动 更新 非空字段
	Update() error

	// 删除
	Delete() error

	// 通过 Id 来删除
	DeleteById() error

	// 通过多个 ID 来读取
	ReadInIds(result interface{}, ids ...int) error

	// 通过多个 ID 来删除
	DeleteInIds(ids ...int) (int, error)

	StringPretty() string
	GetCollName() (name string)
	GetModelName() (name string)
	GetModelFullName() (name string)

	GetFieldIdValue() (id int)
	GetFieldNameValue() (name string)
	SetFieldIdValue(id int) error
	SetFieldNameValue(name string) error

	MakeModelPtr() Model
	MakeModelSlicePtr() interface{}
	GetModelType() reflect.Type

	Copy() (Model, error)
	CopyTo(mdl Model) (err error)
}

//-----------------------------------------------------------------------------------------------------------//

// 查询 接口
type Query interface {
	Query(rows interface{}) error     // 查询数据
	GetQuery() ([]interface{}, error) // 查询数据【保存到接口数组里】
	GetCollName() string              // 获取 集合名称
	Count() (int, error)              // 查询条数
	CountAll() (int, error)           // 查询条数【无过滤条件】
}

//-----------------------------------------------------------------------------------------------------------//
