package types

const (
	// 分子-结构类型
	MOL_STRU_TYPE_NOT_SET int = 0                // 未设置
	MOL_STRU_TYPE_2D      int = 1                // 二维拓扑结构
	MOL_STRU_TYPE_3DMIN   int = 2                // 三维最低能量结构
	MOL_STRU_TYPE_3DPDB   int = 3                // 三维X-ray结构
	MOL_STRU_TYPE_DEF     int = MOL_STRU_TYPE_2D // 默认值

	// job type
	JOB_TYPE_NOT_SET int = 0 // not set
	JOB_TYPE_WEGA    int = 1 // WEGA
	JOB_TYPE_DOCKING int = 2 // DOCKING

	// job status
	JOB_STATUS_NOT_SET  int = 0
	JOB_STATUS_ERROR    int = -1
	JOB_STATUS_RUNNING  int = 1
	JOB_STATUS_FINISHED int = 2
)

//-----------------------------------------------------------------------------------------------------------//

// 分子-结构类型
type MolStruType int

func (self *MolStruType) Set2D() {
	*self = MolStruType(MOL_STRU_TYPE_2D)
}

func (self *MolStruType) Set3DMin() {
	*self = MolStruType(MOL_STRU_TYPE_3DMIN)
}

func (self *MolStruType) Set3DPDB() {
	*self = MolStruType(MOL_STRU_TYPE_3DPDB)
}

func (self MolStruType) Is2D() bool {
	return int(self) == MOL_STRU_TYPE_2D
}

func (self MolStruType) Is3DMin() bool {
	return int(self) == MOL_STRU_TYPE_3DMIN
}

func (self MolStruType) Is3DPDB() bool {
	return int(self) == MOL_STRU_TYPE_3DPDB
}
func (self MolStruType) IsNotSet() bool {
	return int(self) == MOL_STRU_TYPE_NOT_SET
}

//-----------------------------------------------------------------------------------------------------------//
type JobType int

func (self *JobType) SetWEGA() {
	*self = JobType(JOB_TYPE_WEGA)
}

func (self *JobType) SetDocking() {
	*self = JobType(JOB_TYPE_DOCKING)
}

func (self JobType) IsWEGA() bool {
	return int(self) == JOB_TYPE_WEGA
}

func (self JobType) IsDocking() bool {
	return int(self) == JOB_TYPE_DOCKING
}

//-----------------------------------------------------------------------------------------------------------//
type JobStatus int

func (self *JobStatus) SetERROR() {
	*self = JobStatus(JOB_STATUS_ERROR)
}

func (self *JobStatus) SetRunning() {
	*self = JobStatus(JOB_STATUS_RUNNING)
}

func (self *JobStatus) SetFinished() {
	*self = JobStatus(JOB_STATUS_FINISHED)
}

func (self JobStatus) IsERROR() bool {
	return int(self) == JOB_STATUS_ERROR
}

func (self JobStatus) IsRunning() bool {
	return int(self) == JOB_STATUS_RUNNING
}

func (self JobStatus) IsFinished() bool {
	return int(self) == JOB_STATUS_FINISHED
}
