/*
 **	2016-12-01
 **	myjob页面任务信息更新
 **	作者：RCDD杜杰文
 */

/* 	@js global variable
------------------------------------------------------------------------------------------------------------------------*/
//	刷新每个job时记录此job的相关信息到数组中;
//  此数组的每个对象包含"jobID"和"dataIndex"两个key;
//	此数组依赖于对table不进行翻页，即所有行全部显示;
var jobDictionary = new Array();

/* 	@js DOM Ready
------------------------------------------------------------------------------------------------------------------------*/
$(function() {
    loading();
    $.ajax({
        url: "Myjob",
        type: "POST",
        data: {},
        dataType: "html"
    }).done(loadJobInfo);

    $.ajax({
        url: 'account/loginOrNot',
        type: "POST",
        data: {},
        dataType: "html"
    }).done(isLog);

    function isLog(msg) {
        var jobInfo = JSON.parse(msg);
        console.log(jobInfo);
        var log_ = $('#log');
        if (jobInfo.State == -1) {
            log_.removeClass('logOut');
            log_.addClass('logIn');
            log_.html('Log in');
            log_.attr('href', '/login.html');
            $(".user-info span").html();
            // alert(jobInfo.Msg);
            return false;
        } else {
            if (jobInfo.Data.userid < 0) {
                log_.removeClass('logOut');
                log_.addClass('logIn');
                log_.html('Log in');
                log_.attr('href', '/login.html');
                $(".user-info span").html();
                //$(".loading-component").hide();
                // alert(jobInfo.Msg);
            } else {
                log_.removeClass('logIn');
                log_.addClass('logOut');
                log_.html('Log out');
                log_.attr('href', '/account/logout');
            }
        }
    }
    // 所有job信息刷新
    $(".js-reloadTable").on({
        click: function() {
            window.location.reload();
            $.ajax({
                url: "Myjob",
                type: "POST",
                data: {},
                dataType: "html"
            }).done(loadJobInfo);
        }
    });
});

function loadJobInfo(msg) {
    var jobInfo = JSON.parse(msg);
    console.log(jobInfo);
    if (jobInfo.State == -1) {
        alert("Failed to load Job Infomation. Please reload this page.");
        return false;
    } else {
        // 没有job，则显示为空
        if (jobInfo.Data.count == 0) {
            $(".js-job-table").bootstrapTable({
                data: ""
            });
            $(".loading-component").hide();
            $("#job-content").show();
            // 有job，处理数据并初始化table
        } else {
            var jobInfArr = new Array();
            // 逐行处理数据
            for (var i = 0, len = jobInfo.Data.jobs.length; i < len; i++) {
                var showJobInfo = new Object();
                showJobInfo.Id = jobInfo.Data.jobs[i].Id;
                //Type==1:WEGA Type==2:Docking
                showJobInfo.Type = jobInfo.Data.jobs[i].Type == 1 ? "WEGA" : "Docking";
                showJobInfo.Status = jobInfo.Data.jobs[i].Status == 1 ? "Running" : "Finish";
                if (jobInfo.Data.jobs[i].Status == -1) {
                    showJobInfo.Status = "Failed";
                }

                showJobInfo.StartTime = jobInfo.Data.jobs[i].StartTime.slice(0, -3);
                if (jobInfo.Data.jobs[i].FinishTime) {
                    var tempTimeArr = jobInfo.Data.jobs[i].FinishTime.replace("T", " ");
                    showJobInfo.FinishTime = tempTimeArr;
                } else {
                    showJobInfo.FinishTime = jobInfo.Data.jobs[i].FinishTime;
                }

//                for (var i2 = 0; i2 < zincArr.length; i2++) {
//                    if (zincArr[i2].end == jobInfo.Data.jobs[i].Library && zincArr[i2].type == jobInfo.Data.jobs[i].Type) {
//                        showJobInfo.Library = zincArr[i2].front;
//                    }
//                }
				showJobInfo.Library = jobInfo.Data.jobs[i].Library;
                showJobInfo.Query = jobInfo.Data.jobs[i].QueryMol;
                //showJobInfo.Target = jobInfo.Data.jobs[i].Target;
                //showJobInfo.Operation = "<a class='btn btn-default '>"
                showJobInfo.Operation = "<a class='js-refresh-job btn' data-jobid='" + jobInfo.Data.jobs[i].JobId + "' <i class='fa fa-refresh'></i></a>"
                    // 1:Run, 2:Finish
                if (jobInfo.Data.jobs[i].Status == 1) {
                    showJobInfo.PerComplete = jobInfo.Data.jobs[i].PerComplete;
                    showJobInfo.Operation = "<a class='btn js-refresh-job' data-jobid='" + jobInfo.Data.jobs[i].JobId + "'><i class='fa fa-refresh'></i><span></span></a>";
                    showJobInfo.Results = "";
                    //showJobInfo.Results = "<a class='js-check-job' href='javascript:void(0)' data-jobType='" + jobInfo.Data.jobs[i].Type + "' data-jobid='" + jobInfo.Data.jobs[i].JobId +"'>Check</a>";
                } else if (jobInfo.Data.jobs[i].Status == 2) {
                    compl = parseFloat(jobInfo.Data.jobs[i].PerComplete.split('%')[0])
                    if (compl >= 99.0) {
                        showJobInfo.PerComplete = '100%';
                    } else {
                        showJobInfo.PerComplete = jobInfo.Data.jobs[i].PerComplete + '&nbsp;(Partial failed)';
                    }
                    showJobInfo.Results = "<a class='js-check-job' href='javascript:void(0)' data-jobType='" + jobInfo.Data.jobs[i].Type + "' data-jobid='" + jobInfo.Data.jobs[i].JobId + "' data-pdb='" + jobInfo.Data.jobs[i].Target + "'>Check</a>";
                    showJobInfo.Operation = "<a class='btn' data-jobid='" + jobInfo.Data.jobs[i].JobId + "'><i class='fa fa-refresh icon-refresh'></i><span></span></a>";
                }
                jobInfArr.push(showJobInfo);
            }

            //	初始化table
            $(".js-job-table").bootstrapTable({
                data: jobInfArr
            });
            $(".loading-component").hide();
            $("#job-content").show();
            //	@js 点击刷新每一行的任务情况。
            $(".js-refresh-job").on({
                click: function() {
                    if (!$(this).hasClass("js-not-refresh")) {
                        $(this).children(".fa").addClass("fa-spin");
                        $(this).children("span").html("&nbsp;&nbsp;updating...");
                        $(this).addClass("js-not-refresh");
                        $(this).attr("disabled", true);
                        //	获取jobID
                        var jobID = $(this).attr("data-jobid");
                        var dataIndex = $(this).parent().parent().attr("data-index");
                        if (jobDictionary.length == 0) {
                            //	创建新对象记录此job信息到数组中					
                            var jobObj = new Object();
                            jobObj.jobID = jobID;
                            jobObj.dataIndex = dataIndex;
                            jobDictionary.push(jobObj);
                            console.log(jobDictionary);
                            $.ajax({
                                url: "jobUpdate",
                                type: "POST",
                                data: {
                                    jobId: jobID
                                },
                                dataType: "html"
                            }).done(updateJob);
                        } else {
                            for (var m = 0, dicLen = jobDictionary.length; m < dicLen; m++) {
                                if (jobDictionary[m].jobID != jobID) {
                                    if (m == jobDictionary.length - 1) {
                                        //	创建新对象记录此job信息到数组中					
                                        var jobObj = new Object();
                                        jobObj.jobID = jobID;
                                        jobObj.dataIndex = dataIndex;
                                        jobDictionary.push(jobObj);
                                    }
                                }
                                console.log(jobDictionary);
                                $.ajax({
                                    url: "jobUpdate",
                                    type: "POST",
                                    data: {
                                        jobId: jobID
                                    },
                                    dataType: "html"
                                }).done(updateJob);
                            }
                        }
                    }
                }
            });
        }
    }
    //	查看当前job运行完成后的结果
    $(".js-check-job").on({
        click: function() {
            var tempJobID = $(this).attr("data-jobid");
            var tempJobType = $(this).attr("data-jobType");
            var tempPDB = $(this).attr("data-pdb");
            //存储，IE6~7 cookie 其他浏览器HTML5本地存储
            if (window.localStorage) {
                localStorage.setItem("currentJobID", tempJobID);
                localStorage.setItem("checkedJobType", tempJobType);
                localStorage.setItem("checkedPDB", tempPDB);
            } else {
                Cookie.write("currentJobID", tempJobID);
                Cookie.write("checkedJobType", tempJobType);
                Cookie.write("checkedPDB", tempPDB);
            }
            if (tempJobType == 1) {
                window.location.href = '/lbvs_result.html';
            } else if (tempJobType == 2) {
                window.location.href = '/sbvs_result.html';
            }

        }
    });
}

function updateJob(msg) {
    var currentJobInfo = JSON.parse(msg);
    console.log(currentJobInfo);
    if (currentJobInfo.State != 0) {
        alert("Error :" + currentJobInfo.Data + "\\n Please reload the page.");
        return false;
    }
    //	查询任务字典，找到对应任务的信息
    var trIndex;
    for (var i = 0, len = jobDictionary.length; i < len; i++) {
        if (currentJobInfo.Data.job.JobId == jobDictionary[i].jobID) {
            trIndex = jobDictionary[i].dataIndex;
        }
    }
    //	更新当前Job的信息
    $(".fixed-table-body tr").each(function() {
        // 进入当前job的DOM
        var _this = $(this);
        if ($(this).attr("data-index") == trIndex) {
            //	1.当前任务完成时，删除按钮的类名，从而失去点击事件
            if (currentJobInfo.Data.job.Status == 2) {
                $(".js-refresh-job").removeClass("js-refresh-job");
            }
            //	2.更新finishTime和Result信息
            _this.children("td").each(function(index) {
                if (index == 2) {
                    _status = currentJobInfo.Data.job.Status == 1 ? "Running" : "Finish";
                    if (currentJobInfo.Data.job.Status == -1) {
                        _status = "Failed";
                    }
                    $(this).html(_status);
                }
                if (index == 4) {
                    $(this).html(currentJobInfo.Data.job.FinishTime);
                } else if (index == 5) {
                    perComplete = '';
                    if (currentJobInfo.Data.job.Status == 1) {
                        perComplete = currentJobInfo.Data.job.PerComplete;
                    } else if (currentJobInfo.Data.job.Status == 2) {
                        compl = parseFloat(currentJobInfo.Data.job.PerComplete.split('%')[0]);
                        if (compl >= 99.0) {
                            perComplete = '100%';
                        } else {
                            perComplete = currentJobInfo.Data.job.PerComplete + '&nbsp;(Partial failed)';
                        }
                    }
                    $(this).html(perComplete);
                } else if (index == 9) {
                    //index==9
                    if (currentJobInfo.Data.job.Status == 1) {
                        $(this).html("");
                    } else {
                        $(this).html("<a class='js-check-job' href='javascript:void(0)' data-jobid='" + currentJobInfo.Data.job.JobId + "'>Check</a>");
                    }
                } else if (index == 10) {
                    //	去掉动画及disable状态
                    $(this).find("span").html("");
                    $(this).find(".fa").removeClass("fa-spin");

                    if (currentJobInfo.Data.job.Status == 2) {
                        $(this).children("a").removeClass("js-refresh-job");
                        $(this).children("a").attr("disabled", true);
                    } else {
                        $(this).children("a").removeClass("js-not-refresh");
                        $(this).children("a").attr("disabled", false);
                    }
                }
            });
        }
    });
}

//	@js loading加载
function loading() {
    //得到页面高度 
    var yScroll = (document.documentElement.scrollHeight > document.documentElement.clientHeight) ? document.documentElement.scrollHeight : document.documentElement.clientHeight;
    //得到container的宽度
    var loadingCmpWidth = $(".job-info").width();
    //设置loading-component的宽度和高度
    $(".loading-component").width(loadingCmpWidth);
    $(".loading-component").height(yScroll);

    //
    var loadingWidth = $(".loading").width();
    var loadingHeight = $(".loading").height();
    var left1 = loadingCmpWidth / 2 - loadingWidth / 2;
    var top1 = yScroll / 2 - loadingHeight / 2;
    $(".loading").css("margin-left", left1 + "px");
    $(".loading").css("margin-top", top1 + "px");
    $(".loading-component").css("display", "block");
}