var trgInfo = new Array();
var uniprotInfo = new Array();
var initialPdb = [{
    id: 0,
    text: "1B56"
}];
// 存储zinc的对应信息
$(function() {
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
            $(".loading-component").hide();
            alert(jobInfo.Msg);
            return false;
        } else {
            if (jobInfo.Data.userid < 0) {
                log_.removeClass('logOut');
                log_.addClass('logIn');
                log_.html('Log in');
                log_.attr('href', '/login.html');
                $(".user-info span").html();
                $(".loading-component").hide();
            } else {
                log_.removeClass('logIn');
                log_.addClass('logOut');
                log_.html('Log out');
                log_.attr('href', '/account/logout');
                $(".user-info span").html(jobInfo.Data.username);
                $(".loading-component").hide();
            }
        }
    }

    for (var i = 0, len = Info.length; i < len; i++) {
        trgInfo.push(Info[i].Target);
        uniprotInfo.push(Info[i].Uniprotid);
    }
    //初始化select
    //1.1	Manual selection : Target				 
    $(".js-selLibTarget").select2({
        data: trgInfo,
        theme: "classic"
    });
    //1.2	Manual selection : PDB	
    $(".js-selLibPDB").select2({
        data: initialPdb,
        theme: "classic"
    });
    $("#set-nodes-cnt").select2({
        theme: "classic"
    });
    //设置inputType对应的数据：
    $(".js-selectTarget .select2-selection__rendered").bind('DOMNodeInserted', function(e) {
        var target_uinprot_Name = $(e.target).html();
        var pdbArr = new Array();
        for (var i = 0, len = Info.length; i < len; i++) {
            if (target_uinprot_Name == Info[i].Target.text || target_uinprot_Name == Info[i].Uniprotid.text) {
                pdbArr = Info[i].Pdbs;
                console.log(pdbArr);
                break;
            }
        }
        // 插入数据
        $(".js-selectPDB select").remove();
        $(".js-selectPDB span.select2").remove();
        $(".js-selectPDB").append("<select class='pdb-code' style='width: 77%'></select>");
        $(".pdb-code").select2({
            data: pdbArr,
            theme: "classic"
        });
    });
    //	molecule库初始化
    $(".js-selLibMol").select2({
        theme: "classic"
    });
    var $zincSelect = $(".js-selLibMol");
    $zincSelect.on("change", function(e) {
        var tempName = $(".js-selectLibMol .select2-selection__rendered").html();
        $(".zinc-tips span").html("");
        if (tempName == "ZINC_All-Purchasable") {
            $(".zinc-tips span").html("'All Purchasable' subset of ZINC database, totally 22,724,825 molecules");
        }
        if (tempName == "ZINC_Purchasable_Part_100K") {
            $(".zinc-tips span").html("100,000 molecules selected in order from 'All Purchasable' subset of ZINC database");
        }
        if (tempName == "ZINC_Purchasable_Part_1M") {
            $(".zinc-tips span").html("1,000,000 molecules selected in order from 'All Purchasable' subset of ZINC database");
        }
        if (tempName == "ZINC_Purchasable_Part_10K") {
            $(".zinc-tips span").html("10,000 molecules selected in order from 'All Purchasable' subset of ZINC database");
        }
    });
    /*
     **	Custom checkbox and radio button
     */
    // Smart Wizard	
    $('#wizard').smartWizard({
        transitionEffect: 'slide'
            //labelNext:'Next', 
            //labelPrevious:'Previous', 
            //labelFinish:'Finish'
    });
    $(".buttonFinish").on({
        click: function() {
            //	获取用户选择或输入的靶标和PDB信息
            var pdbCode = $(".js-selectPDB .select2-selection__rendered").html();
            var firstTopRadio = $('.template-first .topSelectSection .radio-btn').hasClass('checkedRadio');
            var firstBottomRadio = $('.template-first .bottomSelectSection .radio-btn').hasClass('checkedRadio');

            //	获取用户选择或输入的小分子库信息
            var molLibrary;
            var tempMolLibrary = $(".js-selectLibMol .select2-selection__rendered").html();

            for (var i = 0; i < zincArr.length; i++) {
                if (tempMolLibrary == zincArr[i].front && zincArr[i].type == 2) {
                    molLibrary = zincArr[i].end;
                }
            }

            var secondTopRadio = $('.template-second .topSelectSection .radio-btn').hasClass('checkedRadio');
            var secondBottomRadio = $('.template-second .bottomSelectSection .radio-btn').hasClass('checkedRadio');

            //	获得用户设置的节点数量
            var nodeCount = $(".nodes-ctrl .select2-selection__rendered").html();
            nodeCount = parseInt(nodeCount);
            var tempName = $(".js-selectLibMol .select2-selection__rendered").html();
            if (tempName == "ZINC_All-Purchasable") {
                if (nodeCount >= 50) {
                    nodeCount = 50
                }
            }

            if (firstTopRadio && secondTopRadio) {
                var jobInfo = new Object();
                jobInfo.Type = 2;
                jobInfo.Target = pdbCode;
                jobInfo.Library = molLibrary;
                jobInfo.Nodes = nodeCount;
                var jobJson = JSON.stringify(jobInfo);
                $(".loading small").text("Submiting");
                loading2();
                $.ajax({
                    url: "/jobSubmit",
                    type: "POST",
                    data: {
                        job_json: jobJson,
                        libUpload: 0
                    },
                    dataType: "html"
                }).done(jobHandle);
            } else if (firstTopRadio && secondBottomRadio) {
                var jobInfo = new Object();
                jobInfo.Type = 2;
                jobInfo.Target = pdbCode;
                jobInfo.Library = "";
                jobInfo.Nodes = nodeCount;
                var jobJson = JSON.stringify(jobInfo);
                var formData = new FormData();
                $(".loading small").text("Submiting");
                loading2();
                // 获取上传文件对象
                var fileObj = document.getElementById("molFile").files[0];
                var fileNameSplit;
                if (!fileObj) {
                    alert("Please upload a sd/mol file!");
                    return false;
                } else {
                    fileNameSplit = fileObj.name.toUpperCase().split(".");
                }
                if (fileObj.size > (1024 * 1024 * 5)) {
                    alert("The file size exceeds the limit(5M), please re-upload a file!");
                    return false;
                }
                //	@js 倒叙遍历数组，来判断"."分割后最后一个字符是否含有"SD/SDF/MOL"
                for (var i = fileNameSplit.length - 1; i >= 0; i--) {
                    if (i = fileNameSplit.length - 1) {
                        if (fileNameSplit[i] == "MOL" || fileNameSplit[i] == "SD" || fileNameSplit[i] == "SDF") {
                            formData.append("job_json", jobJson);
                            formData.append("libUpload", 1);
                            formData.append("library", fileObj);
                            console.log(formData);
                            $.ajax({
                                url: "/jobSubmit",
                                type: "POST",
                                data: formData,
                                dataType: "html",
                                processData: false, // 告诉jQuery不要去处理发送的数据
                                contentType: false // 告诉jQuery不要去设置Content-Type请求头
                            }).done(jobHandle);
                            break;
                        } else {
                            alert("The file format is wrong, please re-upload a sd/mol File!");
                            return false;
                        }

                    }
                }
            }
        }
    });
});



/*
 **Custom checkbox and radio button
 */
$(".canCheckedRadio").on('click', function() {
    var _this = $(this),
        block = _this.parent().parent().parent();
    //console.log(block);

    block.find('input:radio').attr('checked', false);
    block.find(".canCheckedRadio").removeClass('checkedRadio');
    _this.addClass('checkedRadio');
    _this.find('input:radio').attr('checked', true);
    var firstTopRadio = $('.template-first .topSelectSection .canCheckedRadio').hasClass('checkedRadio');
    var firstBottomRadio = $('.template-first .bottomSelectSection .canCheckedRadio').hasClass('checkedRadio');
    //var secondTopRadio = $('.template-second .topSelectSection .radio-btn').hasClass('checkedRadio');
    //var secondBottomRadio = $('.template-second .bottomSelectSection .radio-btn').hasClass('checkedRadio');
    if (firstBottomRadio) {
        /*$("#wizard").children("ul").css("height","606");
        $("#step-1").css("height","600");
        $(".stepContainer").css("height","600");

        $(".set-position").show();*/
    };
    if (firstTopRadio) {

        /*	
         **	@bug	$("#wizard").children("ul").css("height","506");存在时，底部按钮条的宽度会变一下。
         */
        /*$("#step-1").css("height","500");
        $(".stepContainer").css("height","500");
        $(".set-position").hide();*/
        //$("#wizard").children("ul").css("height","506");
    }
});

//
function jobHandle(msg) {
    var jobInfo = JSON.parse(msg);
    console.log(jobInfo);
    if (jobInfo.State == -1) {
        if (jobInfo.Msg.indexOf('exist') >= 0) {
            alert("This job already exists.");
            window.location.href = "/Myjob.html";
            return false;
        }
        alert("Failed to load Job Infomation. Please reload this page.");
        location.reload();
        return false;
    } else {
        window.location.href = "/Myjob.html";
    }
}
//	@js loading加载
function loading() {
    var loadingWidth = $("#wizard").width() + 20;
    var loadingHeight = $("#wizard").height() + 20;
    $(".loading-component").width(loadingWidth);
    $(".loading-component").height(loadingHeight);
    $(".loading-component").css("display", "block");
}
//	@js submiting加载
function loading2() {
    var loadingWidth = $("#wizard").width() + 20;
    var loadingHeight = $("#wizard").height() + 20;
    $(".loading-component2").width(loadingWidth);
    $(".loading-component2").height(loadingHeight);
    $(".loading-component2").css("display", "block");
}