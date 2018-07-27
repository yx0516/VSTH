/*
 **	2018-04-09
 **	Library页面任务信息更新
 **	作者：RCDD杜杰文
 */

/* 	@js global variable
------------------------------------------------------------------------------------------------------------------------*/
//	刷新每个job时记录此job的相关信息到数组中;
//  此数组的每个对象包含"jobID"和"dataIndex"两个key;
//	此数组依赖于对table不进行翻页，即所有行全部显示;
/* 	@js DOM Ready
------------------------------------------------------------------------------------------------------------------------*/
$(function() {
    loading();

    // 所有Library信息刷新
    urls = {
            sysLibrary: 'SysLib',
            myLibrary: 'MyLib'
        }
        // 获取所有Library信息
    url = window.location.href;
    url = url.toUpperCase();
    console.log(url.toUpperCase());

    $.ajax({
        url: 'account/loginOrNot',
        type: "POST",
        data: {},
        dataType: "html"
    }).done(isLog);


    function isLog(msg) {
        var jobInfo = JSON.parse(msg);
        console.log(jobInfo);
        if (url.indexOf('SYS') >= 0) {
            $.ajax({
                url: urls.sysLibrary,
                type: "POST",
                data: {},
                dataType: "html"
            }).done(loadJobInfo);
        } else {
            var log_ = $('#log');
            if (jobInfo.State == -1) {
                log_.removeClass('logOut');
                log_.addClass('logIn');
                log_.html('Log in');
                log_.attr('href', '/login.html');
                $(".user-info span").html();
                alert('No library available. Please login first');
                return false;
            } else {
                if (jobInfo.Data.userid < 0) {
                    log_.removeClass('logOut');
                    log_.addClass('logIn');
                    log_.html('Log in');
                    log_.attr('href', '/login.html');
                    $(".user-info span").html();
                    $(".loading-component").hide();
                    alert('No library available. Please login first');
                } else {
                    log_.removeClass('logIn');
                    log_.addClass('logOut');
                    log_.html('Log out');
                    log_.attr('href', '/account/logout');
                    $(".loading-component").hide();
                    $(".user-info span").html(jobInfo.Data.username);
                    $(".btn-loadfile").removeClass('disabled');
                    $.ajax({
                        url: urls.myLibrary,
                        type: "POST",
                        data: {
                            userid: jobInfo.Data.userid
                        },
                        dataType: "html"
                    }).done(loadJobInfo);
                }
            }
        }
    }
})

//开启上传文件的模态框
$(".btn-loadfile").on({
    click: function(e) {
        if ($(this).hasClass('disabled')) {
            e.preventDefault();
        } else {
            $(".tips").css('display', '');
            $(".tips").css('visibility', 'hidden');
            var feedback = $(".feedback");
            feedback.css("display", "none");
            feedback.html("");
            $('#myModal').modal('show');
            $(".tips").css('visibility', 'hidden');
            $(".feedback").html();
            $(".feedback").empty();
            $(".feedback").css('display', 'none');
            $(".js-loadFile").removeClass('disabled');
        }
    }
});

$(".js-loadFile").on({
    click: function(e) {
        if ($(this).hasClass('disabled')) {
            e.preventDefault();
        } else {
            var LibName = $('#myLibName').val();
            if (LibName) {
                $(this).addClass('disabled');
                $.ajax({
                    url: "/CheckLibNameExist",
                    type: "POST",
                    data: {
                        name: LibName
                    },
                    dataType: "html"
                }).done(fileIsExit);
            } else {
                alert('Please enter a library name.')
            }
        }
    }
});

$(".js-reload-myLibrary").on({
    click: function() {
        window.location.reload();
        $.ajax({
            url: urls.myLibrary,
            type: "POST",
            data: { userid: userid },
            dataType: "html"
        }).done(loadJobInfo);
    }
});

$(".js-reload-sysLibrary").on({
    click: function() {
        window.location.reload();
        $.ajax({
            url: urls.sysLibrary,
            type: "POST",
            data: {},
            dataType: "html"
        }).done(loadJobInfo);
    }
});

function loadJobInfo(msg) {
    var jobInfo = JSON.parse(msg);
    console.log(jobInfo);
    if (jobInfo.State == -1) {
        $(".loading-component").hide();
        alert(jobInfo.Msg);
        // alert("Failed to load Library Infomation. Please reload this page.");
        return false;
    } else {
        // 没有job，则显示为空
        if (!jobInfo.Data.libs) {
            $(".js-job-table").bootstrapTable({
                data: ""
            });
            $(".loading-component").hide();
            $("#job-content").show();
        } else {
            $(".js-job-table").bootstrapTable({
                data: jobInfo.Data.libs
            });
            $(".loading-component").hide();
            $("#job-content").show();
        }
    }
}

function fileIsExit(msg) {
    var msgJson = JSON.parse(msg);
    console.log(msgJson);
    if (msgJson.State == -1) {
        alert(msgJson.Msg);
        $(".js-loadFile").removeClass('disabled');
        return false;
    } else {
        uploadfile();
    }
}

function uploadfile() {
    var jobInfo = new Object();
    var formData = new FormData();
    // $(".loading small").text("Submiting");
    // loading2();
    // 获取上传文件对象
    var fileObj = document.getElementById("libFile").files[0];

    var fileNameSplit;
    if (!fileObj) {
        alert("Please upload a sd/sdf file!");
        $(".js-loadFile").removeClass('disabled');
        return false;
    } else {
        fileNameSplit = fileObj.name.toUpperCase().split(".");
    }
    if (fileObj.size > (1024 * 1024 * 100)) {
        $(".tips").css('visibility', 'hidden');
        $(".feedback").empty();
        $(".feedback").html();
        $(".feedback").html("The file size exceeds the limit(100M), please re-upload a file!");
        $(".feedback").css('display', 'inline-block');
        $(".js-loadFile").removeClass('disabled');
        return false;
    }
    var LibName = $('#myLibName').val();
    //	@js 倒叙遍历数组，来判断"."分割后最后一个字符是否含有"SD/SDF/MOL"
    for (var i = fileNameSplit.length - 1; i >= 0; i--) {
        if (i = fileNameSplit.length - 1) {
            if (fileNameSplit[i] == "SD" || fileNameSplit[i] == "SDF") {
                formData.append("name", LibName);
                formData.append("library", fileObj);
                console.log(formData);
                $.ajax({
                    url: "/UploadLib",
                    type: "POST",
                    data: formData,
                    dataType: "html",
                    processData: false, // 告诉jQuery不要去处理发送的数据
                    contentType: false // 告诉jQuery不要去设置Content-Type请求头
                }).done(jobHandle);
                $(".tips").css('visibility', 'visible');
                $(".feedback").css('display', 'none');
                break;
            } else {
                $(".tips").css('visibility', 'hidden');
                $(".feedback").empty();
                $(".feedback").html();
                $(".feedback").append("<p>The file format is wrong, please re-upload a <span style='font-weight:bold;color:red;'>sd/sdf</span> File!</p>");
                $(".feedback").css('display', 'inline-block');
                $(".js-loadFile").removeClass('disabled');
                return false;
            }
        }
    }
}

function jobHandle(msg) {
    var jobInfo = JSON.parse(msg);
    console.log(jobInfo);
    $(".loading-component").hide();
    $(".tips").css('visibility', 'hidden');
    $(".feedback").empty();
    $(".feedback").html();
    $(".feedback").html(jobInfo.Msg);
    $(".feedback").css('display', 'inline-block');
    $(".js-loadFile").removeClass('disabled');
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