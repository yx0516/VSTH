package main

import (
	"bufio"
	"encoding/json"
	//"bufio"
	"fmt"
	//"io"
	"io/ioutil"
	"log"
	"os"
	//"strconv"
	"strings"

	"github.com/davecgh/go-spew/spew"

	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"

	plib "gpc/publib"

	rs "gpc/util/randstr"

	"gpc/db/mongodb/conn"
	"gpc/db/mongodb/conn/gsession"
	"gpc/db/mongodb/xorm"

	"rcdd/project/VSTH/Admin/pkgs/db"
	//"rcdd/project/VSTH/Common/fdata"
	//"rcdd/project/VSTH/ctrls"
)

var (
	pp  = fmt.Println
	ppp = func() { pp("######################################") }
	_   = plib.Debug
	_   = spew.Config
	_   = log.Ldate
	_   = make(bson.M)
	_   = xorm.Log
	_   = os.ErrNotExist

	session  *conn.Session
	database *mgo.Database

	randStrNum func(n int) string
)

func init() {
	//	mgo.SetDebug(true)
	//	mgo.SetLogger(log.New(os.Stdout, "[mgo] ", log.Ldate|log.Ltime))

	randStrNum = rs.RandAlphaDigit

	c := conn.NewConn(
		"127.0.0.1",
		27017,
	)
	c.DB = "vs"
	c.UserName = "appUser"
	c.UserPwd = "appPwd"
	c.Timeout = 3

	var err error
	if session, err = c.Login(); err != nil {
		plib.OsExitPrint(1, err.Error())
	}
	database = session.DB(c.DB)

	// 初始化 orm session
	gsession.InitSession(session.Session, c.DB)
}

type Pdb struct {
	Id   int
	Text string
}

type Target struct {
	Id   int
	Text string
}

type UniprotId struct {
	Id   int
	Text string
}
type DataFormat struct {
	Target    Target
	Uniprotid UniprotId
	Pdbs      []*Pdb
}

//
// 在Docking计算时，用于杰文进行前台展示 靶标-PDB 数据关联信息
//
func pdbqtsToTargetPDBJson() {

	var jsonData []*DataFormat

	// 获取指定目录下的所有文件，不进入下一级目录搜索，可以匹配后缀过滤。
	dirs, err := ioutil.ReadDir("E:/tmp/pdbqt/20161017_PBD_CL/pdb_CL") // fold to save pdbqts
	if err != nil {
		panic(err)
	}
	for _, dir := range dirs {
		if dir.IsDir() {
			pdbCode := strings.ToLower(strings.TrimSpace(dir.Name()))
			queryPDB := db.NewQueryTargetPDB()
			queryPDB.Match.Add("pdbcode", pdbCode) //通过PDBcode查找targetPDB
			if targetPDBs, err := queryPDB.QueryData(); err != nil {
				panic(err)
			} else if len(targetPDBs) == 0 {
				pp("targetPDB not exist: " + pdbCode)
			} else {
				if len(targetPDBs) > 1 {
					//					pp("next")
				}

				for i, _ := range targetPDBs {
					if len(targetPDBs) > 1 {
						//						pp(targetPDBs[i].UniprotId)
					}
					queryTarget := db.NewQueryTarget()
					queryTarget.Match.Add("uniprotid", targetPDBs[i].UniprotId) //通过uniprotId查找target

					if targets, err := queryTarget.QueryData(); err != nil {
						pp("target query error: " + targetPDBs[i].UniprotId)
						panic(err)
					} else if len(targets) == 0 {
						pp("target not exist:" + targetPDBs[i].UniprotId + ", " + pdbCode)
					} else if len(targets) == 1 {
						//						plib.Dump(targets[0])
						flag := 0
						if len(jsonData) != 0 {
							for i, _ := range jsonData {
								if jsonData[i].Target.Text == targets[0].Name || jsonData[i].Uniprotid.Text == targets[0].UniprotId {
									var tmppdb Pdb
									tmppdb.Id = len(jsonData[i].Pdbs)
									tmppdb.Text = strings.ToUpper(pdbCode)
									jsonData[i].Pdbs = append(jsonData[i].Pdbs, &tmppdb)
									flag = 1
									break
								}
							}
						}
						if flag == 0 {
							var targetPDBJson DataFormat
							targetPDBJson.Target.Id = len(jsonData)
							targetPDBJson.Target.Text = targets[0].Name
							targetPDBJson.Uniprotid.Id = len(jsonData)
							targetPDBJson.Uniprotid.Text = targets[0].UniprotId
							//							plib.Dump(targetPDBJson)

							var tmppdb Pdb
							tmppdb.Id = 0
							tmppdb.Text = strings.ToUpper(pdbCode)

							targetPDBJson.Pdbs = append(targetPDBJson.Pdbs, &tmppdb)

							jsonData = append(jsonData, &targetPDBJson)
						}
					} else {
						pp("multiple targets: " + targetPDBs[i].UniprotId)
					}
				}

			}
		}
	}

	data, _ := json.MarshalIndent(jsonData, "", "    ")
	plib.WriteByteToFile("E:/Dev/Go/src/rcdd/project/VSTH/tools/GenStaticData/staticData/target_pdb_list_JsonData", data, 1)
	//	plib.Dump(jsonData)
	//	pp(string(data))

}

func InArrayString(strArray []string, str string) bool {
	for _, v := range strArray {
		if v == str {
			return true
		}
	}
	return false
}

//
// 在WEGA计算时，用于杰文展示可选取活性小分子的靶标的信息
//
func getAllTarget() {

	fi, _ := os.Open("E:/百度云/科研项目/丁鹏/PTS/数据/statistics data/target_activity_molecule_relation/relation.txt")
	defer fi.Close()

	br := bufio.NewReader(fi)
	// skip first line: column name
	br.ReadLine()

	var targets []string
	var targetsJson []*Target
	var count = 0
	for {
		bytes, _, err := br.ReadLine()
		str := string(bytes)
		if err != nil {
			fmt.Println(err)
			break
		}
		//fmt.Println(line)
		tmp := strings.Split(str, "\t")

		if len(tmp) != 7 {
			pp(tmp)
			continue
		}

		targetUniprot := strings.TrimSpace(tmp[5])
		if InArrayString(targets, targetUniprot) {

		} else {
			targets = append(targets, targetUniprot)
			target := db.NewTarget(targetUniprot)
			if err := target.Read(); err != nil {
				pp("target read error: " + targetUniprot)
			} else {
				var trgJson Target
				trgJson.Id = target.Id
				trgJson.Text = target.Name
				targetsJson = append(targetsJson, &trgJson)
				count++
			}

		}
	}
	pp(count)
	//	data, _ := json.MarshalIndent(targetsJson, "", "    ")
	data, _ := json.Marshal(targetsJson)
	plib.Dump(string(data))
	_, err := plib.WriteByteToFile("E:/Dev/Go/src/rcdd/project/VSTH/tools/GenStaticData/staticData/target_JsonData", data, 1)
	pp(err)

}

func main() {
	defer session.SafeClose()
	//	pdbqtsToTargetPDBJson()
	getAllTarget()
}
