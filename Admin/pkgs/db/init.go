package db

import (
	"time"

	"gpc/db/mongodb/types"
	"gpc/db/mongodb/xorm"

	ctype "rcdd/project/VSTH/Admin/pkgs/db/types"
)

const (
	CST_COLL_USER                 = "user"
	CST_COLL_ZINCMOL              = "zincmol"
	CST_COLL_ZINCMOL_SUPPLY       = "zincmol_supply"
	CST_COLL_LIGAND               = "ligand"
	CST_COLL_LIGAND_ACTIVITY      = "ligand_activity"
	CST_COLL_LIGAND_ACTIVITY_REF  = "ligand_activity_ref"
	CST_COLL_TARGET               = "target"
	CST_COLL_TARGET_PDB           = "target_pdb"
	CST_COLL_LIG_ACT_TAR_RELATION = "lig_act_tar_relation"
	CST_COLL_JOBINFO              = "jobinfo"
	CST_COLL_JOBRESULT            = "jobresult"
)

//-----------------------------------------------------------------------------------------------------------//

func init() {
	model := xorm.NewModelInit()
	model.AddDefModels(
		new(User),
		new(ZincMol),
		new(ZincMolProp),
		new(ZincMolSupply),
		new(Ligand),
		new(LigandActivity),
		new(LigandActivityRef),
		new(Target),
		new(TargetPDB),
		new(LigActTarRelation),
		new(JobInfo),
		new(JobResult),
		new(LibInfo),
	)
	model.Init("vs")
}

//-----------------------------------------------------------------------------------------------------------//

// 用户
type User struct {
	xorm.ModelPublic `bson:"-" json:"-"`
	Id               int
	Name             string `xorm:"update" chk:"!null" !null:"!null"`
	Password         string `xorm:"update" chk:"!null" !null:"!null"`
	Phone            string
	Email            string       `xorm:"update" chk:"!null" !null:"!null"`
	Status           types.Status `xorm:"update" chk:"in=1,2" !null:"in=1,2" null:"in!=1,2" def:"1"`
	Logincount       int          `xorm:"update"`
	Lastip           string       `xorm:"update"`
	Lastlogin        time.Time
	CreateTime       time.Time
	UpdateTime       time.Time
}

func NewUser(name string) *User {
	v := &User{
		Name: name,
	}
	v.Init()
	return v
}

func NewUserByEmail(email string) *User {
	v := &User{
		Email: email,
	}
	v.Init()
	return v
}

//-----------------------------------------------------------------------------------------------------------//

// ZINC数据库分子
type ZincMol struct {
	xorm.ModelPublic `bson:"-" json:"-"`
	Id               int
	Name             string `chk:"!null" !null:"!null"`             //ZINC数据库ID号，如：ZINC00000007
	Data             string `chk:"!null" !null:"!null" cfg:"!trim"` //mol格式的结构信息
	//	MW               float64 `xorm:"update" chk:">=0" `              //分子量
	//	LogP             float64 `xorm:"update" `                        //油水分配系数
	//	DesolvApolar     float64 `xorm:"update" `                        //apolar desolvation energies(kcal/mol)非极性去溶剂化能量
	//	DesolvPolar      float64 `xorm:"update" `                        //polar desolvation energies(kcal/mol)极性去溶剂化能量
	//	Hba              int     `xorm:"update" chk:">=0" `              //氢键受体数量
	//	Hbd              int     `xorm:"update" chk:">=0" `              //氢键供体数量
	//	Tpsa             float64 `xorm:"update" `
	//	// Topological Polar Surface Area(拓扑极性表面积)
	//	// * >140 A2 tend to be poor at permeating cell membranes
	//	// * <90 A2 is usually needed to penetrate the blood–brain barrier
	//	Charge   float64 `xorm:"update" `                          //电荷数
	//	Nrb      int     `xorm:"update" chk:">=0" `                //number of rotatable bonds（可旋转键数目）
	//	Smile    string  `xorm:"update" chk:"!null" !null:"!null"` //标准SMILE值
	//	Inchikey string  `xorm:"update" chk:"!null" !null:"!null"` //标准inchikey值
}

func NewZincMol(name string) *ZincMol {
	v := &ZincMol{
		Name: name,
	}
	v.Init()
	return v
}

//-----------------------------------------------------------------------------------------------------------//

// ZINC数据库分子-属性
type ZincMolProp struct {
	xorm.ModelPublic `bson:"-" json:"-"`
	Id               int
	ZincMolName      string  `chk:"!null" !null:"!null"` //ZINC数据库ID号，如：ZINC00000007
	MW               float64 `xorm:"update" chk:">=0" `  //分子量
	LogP             float64 `xorm:"update" `            //油水分配系数
	DesolvApolar     float64 `xorm:"update" `            //apolar desolvation energies(kcal/mol)非极性去溶剂化能量
	DesolvPolar      float64 `xorm:"update" `            //polar desolvation energies(kcal/mol)极性去溶剂化能量
	Hba              int     `xorm:"update" chk:">=0" `  //氢键受体数量
	Hbd              int     `xorm:"update" chk:">=0" `  //氢键供体数量
	Tpsa             float64 `xorm:"update" `
	// Topological Polar Surface Area(拓扑极性表面积)
	// * >140 A2 tend to be poor at permeating cell membranes
	// * <90 A2 is usually needed to penetrate the blood–brain barrier
	Charge   float64 `xorm:"update" `                          //电荷数
	Nrb      int     `xorm:"update" chk:">=0" `                //number of rotatable bonds（可旋转键数目）
	Smile    string  `xorm:"update" chk:"!null" !null:"!null"` //标准SMILE值
	Inchikey string  `xorm:"update" chk:"!null" !null:"!null"` //标准inchikey值
}

func NewZincMolProp(zincMolName string) *ZincMolProp {
	v := &ZincMolProp{
		ZincMolName: zincMolName,
	}
	v.Init()
	return v
}

//-----------------------------------------------------------------------------------------------------------//

// ZINC数据库分子-供应商
type ZincMolSupply struct {
	xorm.ModelPublic `bson:"-" json:"-"`
	Id               int
	ZincMolName      string `chk:"!null" !null:"!null"` //zincMol collection中的Name字段)，如：ZINC00000007
	Supplier         string `chk:"!null" !null:"!null"` //供应商名称
	SupplierCode     string `chk:"!null" !null:"!null"` //供应商自身代码
	Website          string `chk:"!null" !null:"!null"` //如果没有，用NA表示
	Email            string `chk:"!null" !null:"!null"` //如果没有，用NA表示
	Phone            string `chk:"!null" !null:"!null"` //如果没有，用NA表示
	Fax              string `chk:"!null" !null:"!null"` //如果没有，用NA表示
}

func NewZincMolSupply() *ZincMolSupply {
	v := &ZincMolSupply{}
	v.Init()
	return v
}

//-----------------------------------------------------------------------------------------------------------//

//
// ChEMBL数据库+PDB-bind数据库 Ligand
// 存在一个ligand同时出现在两个数据库中（Inchikey相同），但构象不一样
//
type Ligand struct {
	xorm.ModelPublic `bson:"-" json:"-"`
	Id               int
	Name             string            `xorm:"update" chk:"!null" !null:"!null"`                  //uniqueId，小分子的唯一标识，原ChEMBL数据库ID或原PDB-bind数据库ID
	Data             string            `chk:"!null" !null:"!null" cfg:"!trim"`                    //mol格式的结构信息
	Smile            string            `xorm:"update"`                                            //标准SMILE值
	Inchikey         string            `xorm:"update"`                                            //标准inchikey值
	Type             ctype.MolStruType `chk:"in=1,2,3" !null:"in=1,2,3" null:"in!=1,2,3" def:"1"` // 1:2D二维结构, 2:3D能量最低, 3:3D-PDB x-ray结构
	//PDBCode          string
}

func NewLigand() *Ligand {
	v := &Ligand{}
	v.Init()
	return v
}

//-----------------------------------------------------------------------------------------------------------//

// Ligand-活性(一对多)
type LigandActivity struct {
	xorm.ModelPublic `bson:"-" json:"-"`
	Id               int
	Name             string  `chk:"!null" !null:"!null"`    //原ChEMBL数据库的活性记录编号，PDB: pdbCode
	Type             string  `chk:"!null" !null:"!null"`    // Ki Kd IC50 EC50
	Relation         string  `chk:"!null" !null:"in=>,<,="` //大于，小于，等于；>/</=
	Value            float64 //
	Units            string  `chk:"!null" !null:"!null"` //活性数值的单位：nm，um等
}

func NewLigandActivity() *LigandActivity {
	v := &LigandActivity{}
	v.Init()
	return v
}

//-----------------------------------------------------------------------------------------------------------//

// Ligand-活性-文献（一对一）
type LigandActivityRef struct {
	xorm.ModelPublic `bson:"-" json:"-"`
	Id               int
	LigandActivityId int    `xorm:"update" chk:">0" !null:">0" null:"<=0"` //LigandActivity collection中的Id字段
	Name             string // chembl doc_id
	Journal          string `chk:"!null" !null:"!null"`
	Year             int    `chk:">1900" !null:">1900" null:"<=1900"` //发表年份
	Volume           int    //卷号
	Issue            int    //期号
	First_page       int    `chk:">0" !null:">0" null:"<=0"` //起始页码
	Last_page        int    `chk:">0" !null:">0" null:"<=0"` //终止页码
	//Doi              string //文献的DOI编号
	//Pubmed_id        string //文献的Pubmed编号
	//Title            string `chk:"!null"` //标题
	//Authors          string `chk:"!null"` //作者
	//Type             string //类别
	//Abstract         string //摘要
}

func NewLigandActivityRef() *LigandActivityRef {
	v := &LigandActivityRef{}
	v.Init()
	return v
}

//-----------------------------------------------------------------------------------------------------------//

// 靶标
type Target struct {
	xorm.ModelPublic `bson:"-" json:"-"`
	Id               int
	UniprotId        string //uniprot数据库编号
	TtdId            string //TTD数据库编号
	ChemblId         string //
	Name             string //靶标名字
	Type             string //靶标类型
	Gene             string //靶标所属基因名称
	Organism         string //靶标所属种属
	Function         string //靶标的功能描述

}

func NewTarget(uniprotId string) *Target {
	v := &Target{
		UniprotId: uniprotId,
	}
	v.Init()
	return v
}

//-----------------------------------------------------------------------------------------------------------//

// 靶标
type TargetPDB struct {
	xorm.ModelPublic `bson:"-" json:"-"`
	Id               int
	TargetId         int     //Target Collection Id
	UniprotId        string  //uniprot数据库编号
	PDBCode          string  //PDB数据库编号
	Resolution       float64 //分辨率
	Year             int     //发布年份
	DataFile         string  //结构文件名

}

func NewTargetPDB() *TargetPDB {
	v := &TargetPDB{}
	v.Init()
	return v
}

func NewTargetPDBByCode(pdbCode string) *TargetPDB {
	v := &TargetPDB{
		PDBCode: pdbCode,
	}
	v.Init()
	return v
}

func (self *TargetPDB) CollName() string {
	return "target_pdb"
}

//-----------------------------------------------------------------------------------------------------------//

// Ligand-活性-靶标 相互关系
type LigActTarRelation struct {
	xorm.ModelPublic `bson:"-" json:"-"`
	Id               int
	LigandId         int `chk:">0" !null:">0" null:"<=0"` //Ligand collection中的Id字段
	TargetId         int `chk:">0" !null:">0" null:"<=0"` //Target collection中的Id字段
	LigandActivityId int `chk:">0" !null:">0" null:"<=0"` //LigandActivity collection中的Id字段
}

func NewLigActTarRelation() *LigActTarRelation {
	v := &LigActTarRelation{}
	v.Init()
	return v
}

//-----------------------------------------------------------------------------------------------------------//

type LigandForOneTarget struct {
	Id         int               //Ligand collection中的Id字段
	Name       string            //uniqueId，小分子的唯一标识，原ChEMBL数据库ID或原PDB-bind数据库ID
	Data       string            //mol格式的结构信息
	Type       ctype.MolStruType // 1:2D二维结构, 2:3D能量最低构象, 3:3D-PDB x-ray构象
	Activities []*LigandActivity
	References []*LigandActivityRef
}

//-----------------------------------------------------------------------------------------------------------//

// 任务信息
type JobInfo struct {
	xorm.ModelPublic `bson:"-" json:"-"`
	Id               int
	JobId            string `chk:"!null" !null:"!null"`
	UserId           int
	Type             ctype.JobType `chk:"in=1,2" !null:"in=1,2" null:"in!=1,2"` //1:WEGA, 2:Docking
	Target           string        // for Docking, using PDBCode: 1ABC
	QueryMol         string        // for WEGA, 如用户选择，为Ligand collection中的Id字段 ，如用户上传，为上传后保存的文件路径
	Library          string        // library name
	LibSD            string        // library SD file name, for user library
	Nodes            int
	Status           ctype.JobStatus `xorm:"update" chk:"in=-1,1,2" !null:"in=-1,1,2" null:"in!=-1,1,2"` //-1:error, 1:Run, 2:Finish
	StartTime        string          `xorm:"update" `
	FinishTime       string          `xorm:"update" `
	PerComplete      string          `xorm:"update" ` // percent of complete
	ResultFile       string          `xorm:"update" ` // result file of mols and props for download
}

func NewJobInfo() *JobInfo {
	v := &JobInfo{}
	v.Init()
	return v
}

//-----------------------------------------------------------------------------------------------------------//

// 任务结果信息
type JobResult struct {
	xorm.ModelPublic `bson:"-" json:"-"`
	Id               int
	JobId            string
	MolName          string // ZincMolName or molname of user's library molecule
	MolData          string
	Score            float64
}

func NewJobResult() *JobResult {
	v := &JobResult{}
	v.Init()
	return v
}

//-----------------------------------------------------------------------------------------------------------//

// 筛选库信息，筛选库需要是已经产生三维构象的结构
type LibInfo struct {
	xorm.ModelPublic `bson:"-" json:"-"`
	Id               int
	LibId            string `chk:"!null" !null:"!null"`
	UserId           int
	Name             string
	Type             int //1:sys, 2:user
	UploadTime       string
	SDfile           string // absolute path
	Confs            int    // number of 3D conformations
	//Size             int    //storage: MB
}

func NewLibInfo() *LibInfo {
	v := &LibInfo{}
	v.Init()
	return v
}

//-----------------------------------------------------------------------------------------------------------//
