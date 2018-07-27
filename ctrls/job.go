package ctrls

import (
	"encoding/json"
	//	"strconv"
	//	"io"
	"fmt"
	"io/ioutil"
	"os"
	"rcdd/project/VSTH/Admin/pkgs/db"
	"rcdd/project/VSTH/Common/cmd"
	"rcdd/project/VSTH/Common/fdata"
	"strings"
	"time"

	"gopkg.in/mgo.v2/bson"

	"github.com/pborman/uuid"

	plib "gpc/publib"

	"github.com/astaxie/beego"
)

type JobController struct {
	baseController
}

func (this *JobController) ShowHTML() {
	this.TplName = ""
}

/*
** WEGA and Docking job
 */
//func (this *JobController) MyJobs() {

//	queryJob := db.NewQueryJobInfo()
//	queryJob.Match.Add("userid", this.userid)

//	if jobs, err := queryJob.QueryData(); err != nil {
//		this.responseMsg.ErrorMsg("error in jobs query.", nil)
//	} else if numOfJobs, err := queryJob.Count(); err != nil && numOfJobs == 0 {
//		this.responseMsg.SuccessMsg("no jobs available.", nil)
//	} else {
//		idFile := beego.AppConfig.String("idFile")
//		username := beego.AppConfig.String("username")
//		ipAndPort := beego.AppConfig.String("ipAndPort")

//		for i, _ := range jobs { // call script to update jobinfo
//			if jobs[i].Status.IsRunning() {
//				var jobUpdateScript string
//				if jobs[i].Type.IsWEGA() {
//					jobUpdateScript = beego.AppConfig.String("WEGAJobUpdateScript")
//				} else if jobs[i].Type.IsDocking() {
//					jobUpdateScript = beego.AppConfig.String("VinaJobUpdateScript")
//				} else {
//					fmt.Println("job type error, no update script call, type is ", jobs[i].Type)
//				}
//				cmd.RemoteCallJobUpdateScript(idFile, username, ipAndPort, jobUpdateScript, jobs[i].JobId, jobs[i].Library)
//			}
//		}

//		if jobs, err = queryJob.QueryData(); err != nil { // reload jobs
//			this.responseMsg.ErrorMsg("error in reloading jobs.", nil)
//		}

//		data := make(map[string]interface{})
//		data["jobs"] = jobs
//		data["count"] = numOfJobs
//		this.responseMsg.SuccessMsg("success", data)
//	}

//	this.Data["json"] = this.responseMsg
//	this.ServeJSON()

//}

/*
**  WEGA job
 */
func (this *JobController) MyJobs() {
	fmt.Println("Get jobs.")
	queryJob := db.NewQueryJobInfo()
	queryJob.Match.Add("userid", this.userid)

	if jobs, err := queryJob.QueryData(); err != nil {
		this.responseMsg.ErrorMsg("error in jobs query.", nil)
	} else if numOfJobs, err := queryJob.Count(); err != nil && numOfJobs == 0 {
		this.responseMsg.SuccessMsg("no jobs available.", nil)
	} else {
		idFile := beego.AppConfig.String("idFile")
		username := beego.AppConfig.String("username")
		ipAndPort := beego.AppConfig.String("ipAndPort")

		for i, _ := range jobs { // call script to update jobinfo
			if jobs[i].Status.IsRunning() {
				var jobUpdateScript string
				if jobs[i].Type.IsWEGA() {
					jobUpdateScript = beego.AppConfig.String("WEGAJobUpdateScript")
				} else {
					fmt.Println("job type error, no update script call, type is ", jobs[i].Type)
				}
				cmd.RemoteCallJobUpdateScript(idFile, username, ipAndPort, jobUpdateScript, jobs[i].JobId, jobs[i].LibSD)
			}
		}

		if jobs, err = queryJob.QueryData(); err != nil { // reload jobs
			this.responseMsg.ErrorMsg("error in reloading jobs.", nil)
		}

		data := make(map[string]interface{})
		data["jobs"] = jobs
		data["count"] = numOfJobs
		this.responseMsg.SuccessMsg("success", data)
	}

	this.Data["json"] = this.responseMsg
	this.ServeJSON()

}

/*
** WEGA and docking
 */
//func (this *JobController) JobSubmit() {

//	job := db.NewJobInfo()
//	jobInfoJson := strings.TrimSpace(this.GetString("job_json"))
//	err := json.Unmarshal([]byte(jobInfoJson), job)

//	//jobType, err := this.GetInt("type")
//	if err != nil {
//		fmt.Println("get job type failed.")
//		this.responseMsg.ErrorMsg("error in jobs submit.", nil)
//		this.Data["json"] = this.responseMsg
//		this.ServeJSON()
//		return
//	}

//	uuid := uuid.New()
//	job.JobId = strings.Replace(uuid, "-", "_", -1)
//	job.UserId = this.userid
//	job.Status.SetRunning()

//	fmt.Println("job library: ", job.Library)
//	//fmt.Println("ligand Name: ", job.QueryMol)

//	//	t := time.Now().UnixNano()
//	//day := time.Now().Format("20060102")
//	if job.Type.IsWEGA() { // WEGA
//		queryUpload, err := this.GetInt("queryUpload")
//		if err != nil {
//			fmt.Println("error: get queryUpload failed.")
//			this.responseMsg.ErrorMsg("error in jobs submit.", nil)
//			this.Data["json"] = this.responseMsg
//			this.ServeJSON()
//			return
//		}
//		var query string      // for "query" option of job running script.
//		if queryUpload == 0 { // no query upload, job.QueryMol = ligand.Name
//			// job exists or not
//			jobQuery := db.NewQueryJobInfo()
//			jobQuery.Match.Add("querymol", job.QueryMol).Add("library", job.Library)
//			jobs, err := jobQuery.QueryData()
//			if err == nil && len(jobs) > 0 && jobs[0].Status != -1 { // exist
//				fmt.Println("WEGA job already exist.")
//				this.responseMsg.ErrorMsg("Job already exists", job)
//				this.Data["json"] = this.responseMsg
//				this.ServeJSON()
//				return
//			}

//			LigandSavePath := beego.AppConfig.String("LigandSavePath")
//			LigandSavePathRemote := beego.AppConfig.String("LigandSavePathRemote")
//			ligandName := job.QueryMol
//			fmt.Println("ligand id:  ", ligandName)
//			if err != nil {
//				fmt.Println("get query error, ligand id error: ", job.QueryMol)
//				this.responseMsg.ErrorMsg("error in jobs submit.", nil)
//				this.Data["json"] = this.responseMsg
//				this.ServeJSON()
//				return
//			}
//			//ligand := db.NewLigand()
//			//ligand.Name = ligandName
//			ligandQuery := db.NewQueryLigand()
//			ligandQuery.Match.Add("name", ligandName)
//			ligands, err := ligandQuery.QueryData()
//			if err != nil || len(ligands) < 1 {
//				fmt.Println("get ligand error: ", ligandName)
//				this.responseMsg.ErrorMsg("error in jobs submit.", nil)
//				this.Data["json"] = this.responseMsg
//				this.ServeJSON()
//				return
//			}

//			ligandFilePath := fmt.Sprintf("%s/%s.sdf", LigandSavePath, ligandName)
//			remoteLigandFilePath := fmt.Sprintf("%s/%s.sdf", LigandSavePathRemote, ligandName)
//			if !plib.FileIsExist(ligandFilePath) { // ligand file not exist
//				fi, err1 := os.Create(ligandFilePath)
//				defer fi.Close()
//				if err1 != nil {
//					fmt.Println("create ligand file error: ", ligandFilePath)
//					this.responseMsg.ErrorMsg("error in jobs submit.", nil)
//					this.Data["json"] = this.responseMsg
//					this.ServeJSON()
//					return
//				}
//				fi.WriteString(ligands[0].Data)
//			}
//			// call scp to save ligand file to remote computing node
//			idFile := beego.AppConfig.String("idFile")
//			username := beego.AppConfig.String("username")
//			ipAndPort := beego.AppConfig.String("ipAndPort")
//			err = cmd.CallScp(idFile, username, ipAndPort, ligandFilePath, remoteLigandFilePath)
//			if err != nil {
//				fmt.Println("scp ligand file error: ", ligandFilePath)
//				this.responseMsg.ErrorMsg("error in jobs submit.", nil)
//				this.Data["json"] = this.responseMsg
//				this.ServeJSON()
//				return
//			}
//			//job.QueryMol = fmt.Sprintf("%s.sdf", ligandName)
//			query = fmt.Sprintf("%s.sdf", ligandName)

//			// call

//		} else if queryUpload == 1 { // upload query molecule
//			//queryMol := this.GetString("query")
//			queryUploadSavePath := beego.AppConfig.String("queryUploadSavePath")
//			savepath := fmt.Sprintf("%s/%d", queryUploadSavePath, this.userid)
//			os.MkdirAll(savepath, os.ModePerm)
//			filename := fmt.Sprintf("%s/%s.sdf", savepath, job.JobId)
//			if err = this.SaveToFile("query", filename); err != nil {
//				fmt.Println("save upload query file error: ", filename)
//				this.responseMsg.ErrorMsg("", nil)
//				this.Data["json"] = this.responseMsg
//				this.ServeJSON()
//				return
//			}

//			// call scp command to transfer file
//			LigandSavePathRemote := beego.AppConfig.String("LigandSavePathRemote")
//			remoteFilename := fmt.Sprintf("%s/%s.sdf", LigandSavePathRemote, job.JobId)

//			idFile := beego.AppConfig.String("idFile")
//			username := beego.AppConfig.String("username")
//			ipAndPort := beego.AppConfig.String("ipAndPort")
//			err := cmd.CallScp(idFile, username, ipAndPort, filename, remoteFilename)
//			if err != nil {
//				fmt.Println("scp upload query file error: ", filename)
//				this.responseMsg.ErrorMsg("error in jobs submit.", nil)
//				this.Data["json"] = this.responseMsg
//				this.ServeJSON()
//				return
//			}

//			job.QueryMol = fmt.Sprintf("%s.sdf", job.JobId)
//			query = fmt.Sprintf("%s.sdf", job.JobId)
//		}

//		libUpload, err := this.GetInt("libUpload")
//		if err != nil {
//			fmt.Println("error: get libUpload failed.")
//			this.responseMsg.ErrorMsg("error in jobs submit.", nil)
//			this.Data["json"] = this.responseMsg
//			this.ServeJSON()
//			return
//		}
//		if libUpload == 0 { // no lib upload
//			//job.Library = "zinc_conf_test"
//			if job.Library == "zinc_conf_test" {
//				job.Nodes = 8
//			}

//		} else if libUpload == 1 { // upload user's library
//			//_, library, err := this.GetFile("uploadLib")
//			libUploadSavePath := fmt.Sprintf("%s/%d", beego.AppConfig.String("libUploadSavePath"), this.userid)
//			os.MkdirAll(libUploadSavePath, os.ModePerm)
//			filename := fmt.Sprintf("%s/%s.sdf", libUploadSavePath, job.JobId)

//			if err = this.SaveToFile("library", filename); err != nil {
//				this.responseMsg.ErrorMsg("", nil)
//				this.Data["json"] = this.responseMsg
//				this.ServeJSON()
//				return
//			}

//			// call scp command to transfer library file

//			job.Library = fmt.Sprintf("%s.sdf", job.JobId)
//		}

//		//target := strings.TrimSpace(this.GetString("target"))

//		if _, err := job.Insert(); err != nil {
//			fmt.Println("job insert error")
//			this.responseMsg.ErrorMsg("error in WEGA job submit.", nil)
//			this.Data["json"] = this.responseMsg
//			this.ServeJSON()
//			return
//		}
//		fmt.Println(job.Status)

//		// call script to submit job
//		idFile := beego.AppConfig.String("idFile")
//		username := beego.AppConfig.String("username")
//		ipAndPort := beego.AppConfig.String("ipAndPort")
//		script := beego.AppConfig.String("WEGAScript")
//		err = cmd.RemoteCallWEGAScript(idFile, username, ipAndPort, script, job.JobId, query, job.Library, job.Nodes)
//		if err != nil {
//			fmt.Println("error to call script to submit WEGA job.")
//			this.responseMsg.ErrorMsg("error to call script to submit job.", nil)
//			this.Data["json"] = this.responseMsg
//			this.ServeJSON()
//			return
//		}

//	} else if job.Type.IsDocking() { // Docking
//		//pdbCode := strings.TrimSpace(this.GetString("PDBCode"))
//		//		targetUpload, err := this.GetInt("targetUpload")
//		//		if err != nil {
//		//			fmt.Println("WEGA parameter error: get queryUpload failed.")
//		//			this.responseMsg.ErrorMsg("error in jobs submit.", nil)
//		//			this.ServeJSON()
//		//			return
//		//		}
//		libUpload, err := this.GetInt("libUpload")
//		if err != nil {
//			fmt.Println("error: get libUpload failed.")
//			this.responseMsg.ErrorMsg("error in jobs submit.", nil)
//			this.Data["json"] = this.responseMsg
//			this.ServeJSON()
//			return
//		}
//		if libUpload == 0 { // no lib upload
//			//job.Library = "zinc_test"

//		} else if libUpload == 1 { // upload user's library
//			//_, library, err := this.GetFile("uploadLib")
//			libUploadSavePath := fmt.Sprintf("%s/%d", beego.AppConfig.String("libUploadSavePath"), this.userid)
//			if err := os.MkdirAll(libUploadSavePath, os.ModeDir); err != nil {
//				fmt.Println("error in creating user directory")
//				this.responseMsg.ErrorMsg("", nil)
//				this.Data["json"] = this.responseMsg
//				this.ServeJSON()
//				return
//			}

//			filename := fmt.Sprintf("%s/%s.sdf", libUploadSavePath, job.JobId)

//			if err = this.SaveToFile("library", filename); err != nil {
//				fmt.Println("error in saving library")
//				this.responseMsg.ErrorMsg("", nil)
//				this.Data["json"] = this.responseMsg
//				this.ServeJSON()
//				return
//			}

//			// call scp command to transfer file: to do

//			job.Library = fmt.Sprintf("%s.sdf", job.JobId)
//		}

//		// job exists or not
//		jobQuery := db.NewQueryJobInfo()
//		jobQuery.Match.Add("target", job.Target).Add("library", job.Library)
//		jobs, err := jobQuery.QueryData()
//		if err == nil && len(jobs) > 0 && jobs[0].Status != -1 { // exist
//			fmt.Println("WEGA job already exist.")
//			this.responseMsg.ErrorMsg("Job already exists", job)
//			this.Data["json"] = this.responseMsg
//			this.ServeJSON()
//			return
//		}

//		if _, err := job.Insert(); err != nil {
//			fmt.Println("job insert error")
//			this.responseMsg.ErrorMsg("error in Vina job submit.", nil)
//			this.Data["json"] = this.responseMsg
//			this.ServeJSON()
//			return
//		}

//		//return
//		// remote call script to submit job
//		idFile := beego.AppConfig.String("idFile")
//		username := beego.AppConfig.String("username")
//		ipAndPort := beego.AppConfig.String("ipAndPort")
//		script := beego.AppConfig.String("VinaScript")

//		err = cmd.RemoteCallVinaScript(idFile, username, ipAndPort, script, job.JobId, strings.ToLower(job.Target), job.Library, job.Nodes)
//		if err != nil {
//			fmt.Println("error to call script to submit Vina job.")
//			fmt.Println(err)
//			this.responseMsg.ErrorMsg("error to call script to submit job.", nil)
//			this.Data["json"] = this.responseMsg
//			this.ServeJSON()
//			return
//		}

//	} else {
//		fmt.Println("job type error.")
//		this.responseMsg.ErrorMsg("error in jobs submit.", nil)
//		this.Data["json"] = this.responseMsg
//		this.ServeJSON()
//		return
//	}

//	this.responseMsg.SuccessMsg("", nil)
//	this.Data["json"] = this.responseMsg
//	this.ServeJSON()
//}

/*
** WEGA
 */
func (this *JobController) JobSubmit() {

	job := db.NewJobInfo()
	jobInfoJson := strings.TrimSpace(this.GetString("jobJson"))
	fmt.Println("job: ", jobInfoJson)
	err := json.Unmarshal([]byte(jobInfoJson), job)
	fmt.Println("job: ", jobInfoJson)
	//jobType, err := this.GetInt("type")
	if err != nil {
		fmt.Println(err)
		fmt.Println("get job information failed.")
		this.responseMsg.ErrorMsg("error in jobs submit.", nil)
		this.Data["json"] = this.responseMsg
		this.ServeJSON()
		return
	}

	lib := db.NewLibInfo()
	libInfoJson := strings.TrimSpace(this.GetString("libJson"))
	fmt.Println("lib: ", libInfoJson)
	err = json.Unmarshal([]byte(libInfoJson), lib)
	fmt.Println("lib: ", libInfoJson)
	if err != nil {
		fmt.Println(err)
		fmt.Println("get library information failed.")
		this.responseMsg.ErrorMsg("error in jobs submit.", nil)
		this.Data["json"] = this.responseMsg
		this.ServeJSON()
		return
	}

	uuid := uuid.New()

	job.JobId = strings.Replace(uuid, "-", "_", -1)
	job.UserId = this.userid
	job.Status.SetRunning()

	//fmt.Println("job library: ", job.Library)
	//fmt.Println("ligand Name: ", job.QueryMol)

	//	t := time.Now().UnixNano()
	//day := time.Now().Format("20060102")
	queryUpload, err := this.GetInt("queryUpload")
	if err != nil {
		fmt.Println("error: get queryUpload failed.")
		this.responseMsg.ErrorMsg("error in jobs submit.", nil)
		this.Data["json"] = this.responseMsg
		this.ServeJSON()
		return
	}
	var query string      // for "query" option of job running script.
	if queryUpload == 0 { // no query upload, job.QueryMol = ligand.Name

		// check job exists or not
		jobQuery := db.NewQueryJobInfo()
		jobQuery.Match.Add("querymol", job.QueryMol).Add("library", lib.Name).Add("userid", this.userid)
		jobs, err := jobQuery.QueryData()
		if err == nil && len(jobs) > 0 && jobs[0].Status != -1 { // exist
			fmt.Println("WEGA job already exist.")
			this.responseMsg.ErrorMsg("This job already exists, please check the job list.", job)
			this.Data["json"] = this.responseMsg
			this.ServeJSON()
			return
		}

		// get mol of ligand from databse and save ligand
		LigandSavePath := beego.AppConfig.String("LigandSavePath")
		LigandSavePathRemote := beego.AppConfig.String("LigandSavePathRemote")
		ligandName := job.QueryMol
		fmt.Println("ligand id:  ", ligandName)
		if err != nil {
			fmt.Println("get query error, ligand id error: ", job.QueryMol)
			this.responseMsg.ErrorMsg("error in jobs submit.", nil)
			this.Data["json"] = this.responseMsg
			this.ServeJSON()
			return
		}
		//ligand := db.NewLigand()
		//ligand.Name = ligandName
		ligandQuery := db.NewQueryLigand()
		ligandQuery.Match.Add("name", ligandName)
		ligands, err := ligandQuery.QueryData()
		if err != nil || len(ligands) < 1 {
			fmt.Println("get ligand error: ", ligandName)
			this.responseMsg.ErrorMsg("error in jobs submit.", nil)
			this.Data["json"] = this.responseMsg
			this.ServeJSON()
			return
		}

		ligandFilePath := fmt.Sprintf("%s/%s.sdf", LigandSavePath, ligandName)
		remoteLigandFilePath := fmt.Sprintf("%s/%s.sdf", LigandSavePathRemote, ligandName)
		if !plib.FileIsExist(ligandFilePath) { // ligand file not exist
			fi, err1 := os.Create(ligandFilePath)
			defer fi.Close()
			if err1 != nil {
				fmt.Println("create ligand file error: ", ligandFilePath)
				this.responseMsg.ErrorMsg("error in jobs submit.", nil)
				this.Data["json"] = this.responseMsg
				this.ServeJSON()
				return
			}
			fi.WriteString(ligands[0].Data)
		}
		// call scp to save ligand file to remote computing node
		idFile := beego.AppConfig.String("idFile")
		username := beego.AppConfig.String("username")
		ipAndPort := beego.AppConfig.String("ipAndPort")
		err = cmd.CallScp(idFile, username, ipAndPort, ligandFilePath, remoteLigandFilePath)
		if err != nil {
			fmt.Println("scp ligand file error: ", ligandFilePath)
			this.responseMsg.ErrorMsg("error in jobs submit.", nil)
			this.Data["json"] = this.responseMsg
			this.ServeJSON()
			return
		}
		//job.QueryMol = fmt.Sprintf("%s.sdf", ligandName)
		query = fmt.Sprintf("%s.sdf", ligandName)

	} else if queryUpload == 1 { // upload query molecule
		//queryMol := this.GetString("query")
		queryUploadSavePath := beego.AppConfig.String("queryUploadSavePath")
		savepath := fmt.Sprintf("%s/%d", queryUploadSavePath, this.userid)
		os.MkdirAll(savepath, os.ModePerm)
		filename := fmt.Sprintf("%s/%s.sdf", savepath, job.JobId)
		if err = this.SaveToFile("query", filename); err != nil {
			fmt.Println("save upload query file error: ", filename)
			this.responseMsg.ErrorMsg("", nil)
			this.Data["json"] = this.responseMsg
			this.ServeJSON()
			return
		}

		// call scp command to transfer file
		LigandSavePathRemote := beego.AppConfig.String("LigandSavePathRemote")
		remoteFilename := fmt.Sprintf("%s/%s.sdf", LigandSavePathRemote, job.JobId)

		idFile := beego.AppConfig.String("idFile")
		username := beego.AppConfig.String("username")
		ipAndPort := beego.AppConfig.String("ipAndPort")
		err := cmd.CallScp(idFile, username, ipAndPort, filename, remoteFilename)
		if err != nil {
			fmt.Println("scp upload query file error: ", filename)
			this.responseMsg.ErrorMsg("error in jobs submit.", nil)
			this.Data["json"] = this.responseMsg
			this.ServeJSON()
			return
		}

		job.QueryMol = fmt.Sprintf("%s.sdf", job.JobId)
		query = fmt.Sprintf("%s.sdf", job.JobId)
	}

	var script string
	if lib.Type == 1 { // system library
		script = beego.AppConfig.String("WEGAScript")
		job.LibSD = lib.Name
	} else if lib.Type == 2 { // user's library
		script = beego.AppConfig.String("WEGAScript_userlib")
		job.LibSD = lib.LibId // UUID for the SD file Name of user library

		job.Nodes = 1 // user's library only use 1 note
	} else {
		fmt.Println("library type error, type is ", lib.Type)
	}

	job.Library = lib.Name
	job.Type = 1 // WEGA

	//target := strings.TrimSpace(this.GetString("target"))
	//fmt.Println(job.Id)
	//fmt.Println(job)
	if _, err := job.Insert(); err != nil {
		fmt.Println("job insert error")
		fmt.Print(err)
		this.responseMsg.ErrorMsg("Error in WEGA job submit.", nil)
		this.Data["json"] = this.responseMsg
		this.ServeJSON()
		return
	}
	fmt.Println(job.Status)

	// call script to submit job
	idFile := beego.AppConfig.String("idFile")
	username := beego.AppConfig.String("username")
	ipAndPort := beego.AppConfig.String("ipAndPort")

	err = cmd.RemoteCallWEGAScript(idFile, username, ipAndPort, script, job.JobId, query, job.LibSD, job.Nodes)
	if err != nil {
		fmt.Println("error to call script to submit WEGA job.")
		this.responseMsg.ErrorMsg("Error in WEGA job submit.", nil)
		this.Data["json"] = this.responseMsg
		this.ServeJSON()
		return
	}

	this.responseMsg.SuccessMsg("", nil)
	this.Data["json"] = this.responseMsg
	this.ServeJSON()
}

func (this *JobController) GetLigandOfTarget() {
	targetId, err := this.GetInt("targetId")
	if err != nil {
		fmt.Println("get targetId error")
		this.responseMsg.ErrorMsg("", nil)
		this.Data["json"] = this.responseMsg
		this.ServeJSON()
		return
	}
	relationQuery := db.NewQueryLigActTarRelation()
	relationQuery.Match.Add("targetid", targetId)
	relations, err := relationQuery.QueryData()
	if err != nil {
		fmt.Println("get relations of target error")
		this.responseMsg.ErrorMsg("", nil)
		this.Data["json"] = this.responseMsg
		this.ServeJSON()
		return
	}

	if len(relations) > 0 {
		var ligandsOfTarget []*db.LigandForOneTarget
		for i, _ := range relations {
			ligand := db.NewLigand()
			ligand.Id = relations[i].LigandId
			err := ligand.ReadById()
			if err != nil {
				fmt.Println("get ligand error: ", ligand.Id)
				continue
			}

			ligandActivity := db.NewLigandActivity()
			ligandActivity.Id = relations[i].LigandActivityId
			ligandActivity.ReadById()

			referece := db.NewLigandActivityRef()
			referece.LigandActivityId = relations[i].LigandActivityId
			referece.Read()

			flag := -1 // if this ligand has already existed in ligandsOfTargets, -1: not exist, j: the existed index
			for j, _ := range ligandsOfTarget {
				if ligand.Id == ligandsOfTarget[j].Id {
					flag = j
					break
				}
			}

			if flag == -1 {
				var ligandData db.LigandForOneTarget
				ligandData.Id = ligand.Id
				ligandData.Name = ligand.Name
				ligandData.Data = ligand.Data
				ligandData.Activities = append(ligandData.Activities, ligandActivity)
				ligandData.References = append(ligandData.References, referece)
				ligandsOfTarget = append(ligandsOfTarget, &ligandData)

			} else {
				ligandsOfTarget[flag].Activities = append(ligandsOfTarget[flag].Activities, ligandActivity)
				ligandsOfTarget[flag].References = append(ligandsOfTarget[flag].References, referece)
			}

		}

		data := make(map[string]interface{})
		data["ligands"] = ligandsOfTarget
		this.responseMsg.SuccessMsg("", data)
		this.Data["json"] = this.responseMsg
		this.ServeJSON()

	} else {
		fmt.Println("number of relations error: ", len(relations))
		this.responseMsg.ErrorMsg("", nil)
		this.Data["json"] = this.responseMsg
		this.ServeJSON()
		return
	}
}

func (this *JobController) GetPDBData() {
	pdbCode := strings.ToLower(strings.TrimSpace(this.GetString("PDBCode")))
	targetPDB := db.NewTargetPDBByCode(pdbCode)
	if err := targetPDB.Read(); err != nil {
		fmt.Println("read PDB error: ", pdbCode)
		this.responseMsg.ErrorMsg("", nil)
		this.Data["json"] = this.responseMsg
		this.ServeJSON()
	} else {
		pdbFileData, err := ioutil.ReadFile(targetPDB.DataFile)
		if err != nil {
			fmt.Println("read PDBFile error: ", targetPDB.DataFile)
			this.responseMsg.ErrorMsg("", nil)
			this.Data["json"] = this.responseMsg
			this.ServeJSON()
		} else {
			data := make(map[string]interface{})
			data["pdb"] = string(pdbFileData)
			this.responseMsg.SuccessMsg("", data)
			this.Data["json"] = this.responseMsg
			this.ServeJSON()
		}
	}

}

/*
** remote call script to update status and results of job
 */
func (this *JobController) JobUpdate() {
	jobId := strings.TrimSpace(this.GetString("jobId"))
	jobInfo := db.NewJobInfo()
	jobInfo.JobId = jobId
	fmt.Println(jobId)
	if err := jobInfo.Read(); err != nil {
		fmt.Println("jobinfo does not exist: ", jobId)
		this.responseMsg.ErrorMsg("", nil)
		this.Data["json"] = this.responseMsg
		this.ServeJSON()
		return
	}

	// remote call script to update job
	idFile := beego.AppConfig.String("idFile")
	username := beego.AppConfig.String("username")
	ipAndPort := beego.AppConfig.String("ipAndPort")

	var jobUpdateScript string
	if jobInfo.Type.IsWEGA() {
		jobUpdateScript = beego.AppConfig.String("WEGAJobUpdateScript")
	} else if jobInfo.Type.IsDocking() {
		jobUpdateScript = beego.AppConfig.String("VinaJobUpdateScript")
	} else {
		fmt.Println("job type error, no update script call, type is ", jobInfo.Type)
	}
	err := cmd.RemoteCallJobUpdateScript(idFile, username, ipAndPort, jobUpdateScript, jobId, jobInfo.LibSD)

	if err != nil {
		fmt.Println("remote call JobResInsert script error: ", err)
		this.responseMsg.ErrorMsg("", nil)
		this.Data["json"] = this.responseMsg
		this.ServeJSON()
		return
	}

	if err := jobInfo.Read(); err != nil {
		fmt.Println("reload job error, jobinfo does not exist: ", jobId)
		this.responseMsg.ErrorMsg("", nil)
		this.Data["json"] = this.responseMsg
		this.ServeJSON()
		return
	}
	data := make(map[string]interface{})
	data["job"] = jobInfo
	this.responseMsg.SuccessMsg("", data)
	this.Data["json"] = this.responseMsg
	this.ServeJSON()

}

func (this *JobController) JobResults() {
	jobId := strings.TrimSpace(this.GetString("jobId"))
	jobInfo := db.NewJobInfo()
	jobInfo.JobId = jobId
	if err := jobInfo.Read(); err != nil {
		fmt.Println("jobinfo does not exist: ", jobId)
		this.responseMsg.ErrorMsg("", nil)
		this.Data["json"] = this.responseMsg
		this.ServeJSON()
		return
	}

	data := make(map[string]interface{})
	jobResQuery := db.NewQueryJobResult()
	jobResQuery.Match.Add("jobid", jobId)
	if jobRes, err := jobResQuery.QueryData(); err != nil || len(jobRes) == 0 {
		fmt.Println("no results for jobId: ", jobId)
		this.responseMsg.ErrorMsg("", nil)
		this.Data["json"] = this.responseMsg
		this.ServeJSON()
	} else {
		data["mols"] = jobRes
		if jobInfo.Type.IsWEGA() { // get query molecule
			LigandSavePath := beego.AppConfig.String("LigandSavePath")
			ligandFile := fmt.Sprintf("%s/%s.sdf", LigandSavePath, jobInfo.QueryMol)

			queryUploadSavePath := beego.AppConfig.String("queryUploadSavePath")
			queryUploadFile := fmt.Sprintf("%s/%d/%s", queryUploadSavePath, this.userid, jobInfo.QueryMol)

			if plib.FileIsExist(queryUploadFile) { // uploaded query filePath
				mol, err := fdata.ReadByFile(queryUploadFile, 1)
				if err != nil {
					fmt.Println("get upload query failed, filePath: ", jobInfo.QueryMol)
					this.responseMsg.ErrorMsg("", nil)
					this.Data["json"] = this.responseMsg
					this.ServeJSON()
					return
				} else {
					data["query"] = mol[0]
					data["type"] = 2 // 2 for upload
				}

			} else if plib.FileIsExist(ligandFile) { // ligand
				//ligand := db.NewLigand()
				//ligandName := jobInfo.QueryMol
				ligandQuery := db.NewQueryLigand()
				ligandQuery.Match.Add("name", jobInfo.QueryMol)
				ligands, err := ligandQuery.QueryData()
				if err != nil || len(ligands) == 0 {
					fmt.Println("get ligand failed, ligandId: ", jobInfo.QueryMol)
					this.responseMsg.ErrorMsg("", nil)
					this.Data["json"] = this.responseMsg
					this.ServeJSON()
					return
				} else {
					data["query"] = ligands[0]
					data["type"] = 1 // 1 for upload

					relationQuery := db.NewQueryLigActTarRelation()
					relationQuery.Match.Add("ligandid", ligands[0].Id)
					relations, err := relationQuery.QueryData()
					var activities []*db.LigandActivity
					var references []*db.LigandActivityRef
					if err != nil {
						fmt.Println("result data error: no activity data for ligand.")
					} else if len(relations) > 0 {
						for i, _ := range relations {
							ligandActivity := db.NewLigandActivity()
							ligandActivity.Id = relations[i].LigandActivityId
							ligandActivity.ReadById()
							activities = append(activities, ligandActivity)
							referece := db.NewLigandActivityRef()
							referece.LigandActivityId = relations[i].LigandActivityId
							referece.Read()
							references = append(references, referece)
						}

					}
					data["activities"] = activities
					data["references"] = references

				}

			}

		} else if jobInfo.Type.IsDocking() { // get PDBfile of target
			pdbCode := strings.ToLower(jobInfo.Target)
			targetPDB := db.NewTargetPDBByCode(pdbCode)
			if err := targetPDB.Read(); err != nil {
				fmt.Println("read PDB error: ", pdbCode)
				this.responseMsg.ErrorMsg("", nil)
				this.Data["json"] = this.responseMsg
				this.ServeJSON()
				return
			} else {
				filePath := fmt.Sprintf("%s/%s/%s", beego.AppConfig.String("pdbFilePath"), pdbCode, targetPDB.DataFile)

				if pdbFileData, err := ioutil.ReadFile(filePath); err != nil {
					fmt.Println("read PDBFile error: ", targetPDB.DataFile)
					this.responseMsg.ErrorMsg("", nil)
					this.Data["json"] = this.responseMsg
					this.ServeJSON()
					return
				} else {
					data["PDB"] = string(pdbFileData)
				}
			}

			target := db.NewTarget(targetPDB.UniprotId)
			if err := target.Read(); err != nil {
				data["target"] = nil
			} else {
				data["target"] = target
			}
		}
		//tmpdata, _ := json.Marshal(data)
		//fmt.Println(string(tmpdata))
		this.responseMsg.SuccessMsg("", data)
		this.Data["json"] = this.responseMsg
		this.ServeJSON()

		fmt.Println("save resultFile if necessary")

		if plib.FileIsExist(jobInfo.ResultFile) {
			fmt.Println("result file exist")
		} else { // save resultFile if not exist
			resFilePath := fmt.Sprintf("%s/%s.sdf", beego.AppConfig.String("jobResultFile"), jobInfo.JobId)
			fi, err1 := os.Create(resFilePath)
			defer fi.Close()
			if err1 != nil {
				fmt.Println("create result file error: ", err1)
			}
			for _, mol := range jobRes {

				fi.WriteString(mol.MolData) // save mol data

				fi.WriteString("> <score>\n")
				fi.WriteString(plib.FloatToStr(mol.Score) + "\n\n")
				fi.WriteString("$$$$\n")
				// get supply: to do
				// get props: to do
			}

			jobInfo.ResultFile = resFilePath // update the resultfile of jobinfo
			if err := jobInfo.UpdateFields(bson.M{"jobid": jobId}, "resultfile"); err != nil {
				fmt.Println("jobinfo update error")
			}
		}
	}

}

func (this *JobController) DownloadJobResults() {
	jobId := strings.TrimSpace(this.GetString("jobId"))
	fmt.Println("Download job: ", jobId)
	jobInfo := db.NewJobInfo()
	jobInfo.JobId = jobId
	if err := jobInfo.Read(); err != nil {
		fmt.Println("jobinfo does not exist: ", jobId)
		this.responseMsg.ErrorMsg("", nil)
		this.Data["json"] = this.responseMsg
		this.ServeJSON()
		return
	}

	if plib.FileIsExist(jobInfo.ResultFile) {
		fmt.Println("jobResult file exists, ready to download.")

	} else { // save resultFile if not exist
		jobResQuery := db.NewQueryJobResult()
		jobResQuery.Match.Add("jobid", jobId)
		jobRes, err := jobResQuery.QueryData()
		if err != nil || len(jobRes) == 0 {
			fmt.Println("no results for jobId: ", jobId)
			this.responseMsg.ErrorMsg("", nil)
			this.Data["json"] = this.responseMsg
			this.ServeJSON()
			return
		}

		resFilePath := fmt.Sprintf("%s/%s.sdf", beego.AppConfig.String("jobResultFile"), jobInfo.JobId)
		fi, err1 := os.Create(resFilePath)
		defer fi.Close()
		if err1 != nil {
			fmt.Println("create result file error: ", err1)
		}
		for _, mol := range jobRes {

			fi.WriteString(mol.MolData) // save mol data

			fi.WriteString("> <score>\n")
			fi.WriteString(plib.FloatToStr(mol.Score) + "\n\n")
			fi.WriteString("$$$$\n")
			// get supply: to do
			// get props: to do
		}

		jobInfo.ResultFile = resFilePath // update the resultfile of jobinfo
		if err := jobInfo.Update(); err != nil {
			fmt.Println("jobinfo update error")
		}
	}
	this.Ctx.Output.Download(jobInfo.ResultFile)
	//this.Ctx.Output.Download()
}

func (this *JobController) GetPropsOfZincMol() {
	zincMolName := strings.TrimSpace(this.GetString("molName"))
	zincMolProp := db.NewZincMolProp(zincMolName)
	if err := zincMolProp.Read(); err != nil {
		fmt.Println("read property of zincMol error: ", zincMolName)
		this.responseMsg.ErrorMsg("", nil)
		this.Data["json"] = this.responseMsg
		this.ServeJSON()
	} else {
		data := make(map[string]interface{})
		data["zincMolProp"] = zincMolProp
		this.responseMsg.SuccessMsg("", data)
		this.Data["json"] = this.responseMsg
		this.ServeJSON()
	}
}

func (this *JobController) GetSupplysOfZincMol() {
	zincMolName := strings.TrimSpace(this.GetString("molName"))
	q := db.NewQueryZincMolSupply()
	q.Match.Add("zincmolname", zincMolName)

	if zincMolSupplys, err := q.QueryData(); err != nil {
		fmt.Println("read supplys of zincMol error: ", zincMolName)
		this.responseMsg.ErrorMsg("", nil)
		this.Data["json"] = this.responseMsg
		this.ServeJSON()
	} else {
		data := make(map[string]interface{})
		data["zincMolSupply"] = zincMolSupplys
		this.responseMsg.SuccessMsg("", data)
		this.Data["json"] = this.responseMsg
		this.ServeJSON()
	}
}

/***** library functions ******/
func (this *JobController) SysLibs() {
	fmt.Println("System library")
	queryLibs := db.NewQueryLibInfo()
	queryLibs.Match.Add("type", 1)
	if libs, err := queryLibs.QueryData(); err != nil {
		fmt.Println("error in lib query")
		this.responseMsg.ErrorMsg("No library available.", nil)
		//this.ServeJSON()
	} else if numOfLibs, err := queryLibs.Count(); err != nil && numOfLibs == 0 {
		fmt.Println("No library available")
		this.responseMsg.SuccessMsg("No library available.", nil)
		//this.ServeJSON()
	} else {
		//fmt.Println(libs)
		data := make(map[string]interface{})
		data["libs"] = libs
		//data["count"] = numOfLibs
		this.responseMsg.SuccessMsg("success", data)
		//this.ServeJSON()
	}
	this.Data["json"] = this.responseMsg
	this.ServeJSON()
}
func (this *JobController) MyLibs() {
	fmt.Println("My library")
	if this.userid == -1 {
		this.responseMsg.ErrorMsg("No library available. Please login first.", nil)
		this.Data["json"] = this.responseMsg
		this.ServeJSON()
		return
	}

	queryLibs := db.NewQueryLibInfo()
	queryLibs.Match.Add("userid", this.userid)
	if libs, err := queryLibs.QueryData(); err != nil {
		this.responseMsg.ErrorMsg("No library available.", nil)

	} else if numOfLibs, err := queryLibs.Count(); err != nil && numOfLibs == 0 {
		this.responseMsg.ErrorMsg("No library available. You can upload your own library for virtual screening.", nil)

	} else {
		data := make(map[string]interface{})
		data["libs"] = libs
		//data["count"] = numOfLibs
		this.responseMsg.SuccessMsg("success", data)

	}
	this.Data["json"] = this.responseMsg
	this.ServeJSON()

}

func (this *JobController) CheckLibNameExist() {
	fmt.Println("check lib name exist.")
	libname := strings.TrimSpace(this.GetString("name"))
	fmt.Println("lib name: ", libname)

	// check if the library name exists in this user's library
	queryLib := db.NewQueryLibInfo()
	queryLib.Match.Add("userid", this.userid).Add("name", libname)
	if libs, err := queryLib.QueryData(); err != nil {
		this.responseMsg.ErrorMsg("Error in library query.", nil)
		this.Data["json"] = this.responseMsg
		this.ServeJSON()
		return
	} else if numOfLibs := len(libs); numOfLibs > 0 {
		this.responseMsg.ErrorMsg("The library name has already existed, please try another one", nil)
		this.Data["json"] = this.responseMsg
		this.ServeJSON()
		return
	} else {
		this.responseMsg.SuccessMsg("", nil)
		this.Data["json"] = this.responseMsg
		this.ServeJSON()
		return
	}
}

func (this *JobController) UploadLib() {
	fmt.Println("upload user library")
	lib := db.NewLibInfo()
	lib.UserId = this.userid
	lib.Name = this.GetString("name")

	// upload library
	uuid := uuid.New()
	lib.LibId = strings.Replace(uuid, "-", "_", -1)

	libUploadSavePath := fmt.Sprintf("%s/%d", beego.AppConfig.String("libUploadSavePath"), lib.UserId)
	os.MkdirAll(libUploadSavePath, os.ModePerm)
	filename := fmt.Sprintf("%s/%s.sdf", libUploadSavePath, lib.LibId)

	if err := this.SaveToFile("library", filename); err != nil {
		this.responseMsg.ErrorMsg("", nil)
		this.Data["json"] = this.responseMsg
		this.ServeJSON()
		return
	}

	// call scp command to transfer file
	libUploadSavePathRemote := beego.AppConfig.String("libUploadSavePathRemote")
	remoteFilename := fmt.Sprintf("%s/%s.sdf", libUploadSavePathRemote, lib.LibId)

	idFile := beego.AppConfig.String("idFile")
	username := beego.AppConfig.String("username")
	ipAndPort := beego.AppConfig.String("ipAndPort")
	err := cmd.CallScp(idFile, username, ipAndPort, filename, remoteFilename)
	if err != nil {
		fmt.Println("scp upload library file error: ", filename)
		this.responseMsg.ErrorMsg("Error in library upload. Please try again.", nil)
		this.Data["json"] = this.responseMsg
		this.ServeJSON()
		return
	}

	lib.SDfile = remoteFilename
	lib.Type = 2 // user library
	lib.UploadTime = time.Now().Format("2006-01-02 15:04:05")
	if lib.Confs, err = fdata.MolCountsByFile(filename); err != nil {
		fmt.Println("error in confs of library")
	}
	//flib, _, _ := this.GetFile("library")
	//buf_len, _ := flib.Seek(0, os.SEEK_END)
	if _, err = lib.Insert(); err != nil {
		fmt.Println("error in library insert")
	}

	this.responseMsg.SuccessMsg("Library upload successfully.", nil)
	this.Data["json"] = this.responseMsg
	this.ServeJSON()

}
