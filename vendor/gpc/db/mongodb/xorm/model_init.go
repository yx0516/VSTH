package xorm

import (
	"gpc/db/mongodb/xorm/model"
	"gpc/util/ref"
	"gpc/util/tag"
	"gpc/web/cst"

	plib "gpc/publib"
)

var (
	// 结构体 注解 缓存
	globalStructCache *tag.StructCache = tag.NewStructCache(
		ref.NewFiledNameMgoConvert(ref.CST_FNAME_CONVERT_TO_LOWER),
		cst.TAG_MGO_BSON, tag.CST_TAG_XORM, tag.CST_TAG_CFG,
	)

	// 结构体 字段 缓存
	globalStructField *tag.AdminStructField = tag.NewAdminStructField()
)

// Hook 模块注册的函数
type HookModelInit func(cellName string, st *tag.StructTag, sf *tag.StructField, sfMap map[string]*tag.StructField)

//-----------------------------------------------------------------------------------------------------------//

type Model struct {
	Object                 model.Model
	IsFilterQueryFields    bool
	IsQueryAddStrudtFields bool
	IsWhiteAddStrudtFields bool
	IsBlackAddStrudtFields bool
	HookFunc               HookModelInit
}

func NewModel(object model.Model) *Model {
	object.Init()
	return &Model{
		Object:                 object,
		IsFilterQueryFields:    false,
		IsQueryAddStrudtFields: false,
		IsWhiteAddStrudtFields: false,
		IsBlackAddStrudtFields: false,
	}
}

// 模型 注册类
type ModelInit struct {
	models   []*Model
	modelMap map[string]*Model
}

func NewModelInit() *ModelInit {
	return &ModelInit{
		modelMap: make(map[string]*Model),
	}
}

func (self *ModelInit) getModelFullName(intf model.Model) string {
	structFullName, err := ref.RefStructTypeFullName(intf)
	if err != nil {
		errMsg := "Get struct full name error:" + err.Error()
		Log.Emergency(errMsg)
		plib.TimeSleep(1 * 1000)
		plib.OsExit(1, errMsg)
	}
	return structFullName
}

func (self *ModelInit) checkModel(model *Model) {
	if model == nil {
		errMsg := "model is nil."
		Log.Emergency(errMsg)
		plib.TimeSleep(1 * 1000)
		plib.OsExit(1, errMsg)
	}
}

// 添加 默认配置的 模型实例
func (self *ModelInit) AddDefModel(intf model.Model) {
	self.AddModel(NewModel(intf))
}

func (self *ModelInit) AddDefModels(intfs ...model.Model) {
	for _, intf := range intfs {
		self.AddModel(NewModel(intf))
	}
}

func (self *ModelInit) AddModel(m *Model) {
	self.checkModel(m)
	self.models = append(self.models, m)
	if self.modelMap == nil {
		self.modelMap = make(map[string]*Model)
	}
	self.modelMap[self.getModelFullName(m.Object)] = m
}

func (self *ModelInit) AddModels(models ...*Model) {
	for _, model := range models {
		self.AddModel(model)
	}
}

func (self *ModelInit) addCache(m *Model, sfMap map[string]*tag.StructField) {
	self.checkModel(m)
	structFullName := self.getModelFullName(m.Object)
	if globalStructCache.HasExist(structFullName) {
		return
	}
	m.Object.Init()

	if st, err := globalStructCache.Add(m.Object); err != nil {
		errMsg := "Add struct tag cache error:" + err.Error()
		Log.Emergency(errMsg)
		plib.TimeSleep(1 * 1000)
		plib.OsExit(1, errMsg)
	} else {
		// 获取 结构体 字段 信息
		sf := tag.NewStructField(st, m.IsFilterQueryFields)
		globalStructField.Add(st.Struct.FullName, sf)

		if m.IsQueryAddStrudtFields {
			for _, field := range sf.Query {
				if m, ok := st.Fields[field]; ok {
					sf.Query.Add(m.Name)
				}
			}
			sf.Query.RemoveDuplicate()
		}

		if m.IsWhiteAddStrudtFields {
			for _, field := range sf.White {
				if m, ok := st.Fields[field]; ok {
					sf.White.Add(m.Name)
				}
			}
			sf.White.RemoveDuplicate()
		}

		if m.IsBlackAddStrudtFields {
			for _, field := range sf.Black {
				if m, ok := st.Fields[field]; ok {
					sf.Black.Add(m.Name)
				}
			}
			sf.Black.RemoveDuplicate()
		}

		if m.HookFunc != nil {
			cellName := m.Object.GetCollName()
			sfMap[cellName] = sf
			m.HookFunc(cellName, st, sf, sfMap)
		}
	}
	return
}

// 异常 将直接终止执行
func (self *ModelInit) Init(modelName ...string) {
	sfMap := make(map[string]*tag.StructField)
	for _, model := range self.models {
		self.addCache(model, sfMap)
	}
}

//-----------------------------------------------------------------------------------------------------------//
