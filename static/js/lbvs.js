function createCode() {
    var code = "";
    var codeLength = 4;
    var codeChars = new Array(0, 1, 2, 3, 4, 5, 6, 7, 8, 9,
        'a', 'b', 'c', 'd', 'e', 'f', 'g', 'h', 'i', 'j', 'k', 'l', 'm', 'n', 'o', 'p', 'q', 'r', 's', 't', 'u', 'v', 'w', 'x', 'y', 'z',
        'A', 'B', 'C', 'D', 'E', 'F', 'G', 'H', 'I', 'J', 'K', 'L', 'M', 'N', 'O', 'P', 'Q', 'R', 'S', 'T', 'U', 'V', 'W', 'X', 'Y', 'Z');
    for (var i = 0; i < codeLength; i++) {
        var charNum = Math.floor(Math.random() * 52);
        code += codeChars[charNum];
    }
    return code
}

var myLibs = [];
var sysLibs = [];

$(function() {
    if (window.localStorage) {
        localStorage.removeItem("step1");
    } else {
        Cookie.write("step1", '');
    }

    initializeSelect2();
    $.ajax({
        url: 'account/loginOrNot',
        type: "POST",
        data: {},
        dataType: "html"
    }).done(isLog);

    //	监控select的变化情况，动态请求分子结构数据
    var $eventSelect = $(".js-selectTarget select");
    $eventSelect.on("change", function(e) {
        var targetId = "";

        var tempName = $(".js-selectTarget .select2-selection__rendered").html();
        if (tempName.indexOf('Select') >= 0) {
            if (window.localStorage) {
                localStorage.setItem("step1", '');
            } else {
                Cookie.write("step1", '');
            }
        } else {
            if (window.localStorage) {
                localStorage.setItem("step1", 'ok');
            } else {
                Cookie.write("step1", 'ok');
            }
        }
        for (var i = 0, len = lbvsTargetInfo.length; i < len; i++) {
            if (lbvsTargetInfo[i].text == tempName) {
                targetId = lbvsTargetInfo[i].id;
                break
            }
        }
        console.log('selectTarget OK');
        console.log(targetId);
        if (targetId != 0) {
            //	@js 加载loading
            loading();
            $.ajax({
                url: "/GetQuery",
                type: "POST",
                data: {
                    targetId: targetId
                },
                dataType: "html"
            }).done(showMolStruc);
        } else {
            return false;
        }
    });

    function isLog(msg) {
        var jobInfo = JSON.parse(msg);
        console.log(jobInfo);
        //获取 system library
        $.ajax({
            url: 'SysLib',
            type: "POST",
            data: {},
            dataType: "html"
        }).done(loadSysLib);


        var log_ = $('#log');
        if (jobInfo.State == -1) {
            log_.removeClass('logOut');
            log_.addClass('logIn');
            log_.html('Log in');
            log_.attr('href', '/login.html');
            $(".user-info span").html();
            $('.user-btn').removeClass('canCheckedRadio');
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
                $('.user-btn').removeClass('canCheckedRadio');
                $('.user-tips').show();
            } else {
                log_.removeClass('logIn');
                log_.addClass('logOut');
                log_.html('Log out');
                log_.attr('href', '/account/logout');
                $(".user-info span").html(jobInfo.Data.username);
                $(".loading-component").hide();

                $('.user-tips').hide();
                $.ajax({
                    url: 'MyLib',
                    type: "POST",
                    data: {
                        userid: jobInfo.Data.userid
                    },
                    dataType: "html"
                }).done(loadMyLib);
            }
        }
    }

    //初始化select
    function initializeSelect2() {
        //1.1	Manual selection : Target			 
        $(".js-selTarget").select2({
            /*placeholder: 'Select an option',*/
            data: lbvsTargetInfo,
            theme: "classic"
        });
        $("#set-nodes-cnt").select2({
            theme: "classic"
        });

    }


    function loadMyLib(msg) {
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
                $('.js-myLib').select2({
                    theme: "classic"
                });
            } else {
                var libData = [];
                myLibs = jobInfo.Data.libs;
                for (var i = 0; i < myLibs.length; i++) {
                    libData.push({
                        'id': i,
                        'text': myLibs[i].Name
                    });
                }
                console.warn(libData);
                $('.js-myLib').select2({
                    theme: "classic",
                    data: libData
                });
            }
        }
    }

    function loadSysLib(msg) {
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
                $(".js-sysLib").select2({
                    theme: "classic"
                });
            } else {
                var libData = [];
                sysLibs = jobInfo.Data.libs;
                for (var i = 0; i < sysLibs.length; i++) {
                    libData.push({
                        'id': i,
                        'text': sysLibs[i].Name
                    });
                }
                console.warn(libData);
                $('.js-sysLib').select2({
                    theme: "classic",
                    data: libData
                });
            }
        }
    }

    $('#wizard').smartWizard({
        transitionEffect: 'slide'
            //labelNext:'Next', 
            //labelPrevious:'Previous', 
            //labelFinish:'Finish'
    });
    $(".buttonFinish").on({
        click: function() {
            //	获取用户选择或输入的靶标和PDB信息
            var molTemplateStr;
            var firstTopRadio = $('.template-first .topSelectSection .big-radio').hasClass('checkedRadio');
            var firstBottomRadio = $('.template-first .bottomSelectSection .big-radio').hasClass('checkedRadio');

            //	获取用户选择或输入的小分子库信息
            var molLibrary;

            var secondTopRadio = $('#step-2 .topSelectSection .big-radio').hasClass('checkedRadio');
            var secondBottomRadio = $('#step-2 .bottomSelectSection .big-radio').hasClass('checkedRadio');
            //	获得用户设置的节点数量
            var nodeCount = $(".nodes-ctrl .select2-selection__rendered").html();

            if (firstTopRadio && secondTopRadio) {
                var step1 = window.localStorage ? localStorage.getItem("step1") : Cookie.read("step1");
                if (!step1) {
                    alert("Please select one ligand or upload your own ligands in Step 1.");
                    return false;
                }
                //@todo: 这两个元素没有用到
                // var selectMolel = $(".mol-structure .checkedRadio").parent().attr("data-index");
                // var mollength = perPageMolArr.length;

                molTemplateStr = $("#molContainer-left .selectMOlname").html();

                var molLibrary = $(".js-selectLibMol .select2-selection__rendered").html();
                var libInfo = null;
                (function IIFE() {
                    for (var i = 0; i < sysLibs.length; i++) {
                        if (sysLibs[i].Name == molLibrary) {
                            libInfo = sysLibs[i];
                        }
                    }
                })();
                var libJson = JSON.stringify(libInfo);

                var jobInfo = new Object();
                jobInfo.QueryMol = molTemplateStr;
                jobInfo.Nodes = parseInt(nodeCount);
                var jobJson = JSON.stringify(jobInfo);
                loading2();
                $.ajax({
                    url: "/jobSubmit",
                    type: "POST",
                    data: {
                        libJson: libJson,
                        jobJson: jobJson,
                        queryUpload: 0
                    },
                    dataType: "html"
                }).done(jobHandle);
            }
            if (firstTopRadio && secondBottomRadio) {
                var step1 = window.localStorage ? localStorage.getItem("step1") : Cookie.read("step1");
                if (!step1) {
                    alert("Please select one ligand or upload your own ligands in Step 1.");
                    return false;
                }
                // var selectMolel = $(".mol-structure .checkedRadio").parent().attr("data-index");
                // var molLibrary = $(".js-myLib .select2-selection__rendered").html();
                // var mollength = perPageMolArr.length;
                molTemplateStr = $("#molContainer-left .selectMOlname").html();

                var jobInfo = new Object();
                jobInfo.QueryMol = molTemplateStr;
                jobInfo.Nodes = parseInt(nodeCount);
                var libInfo = null;
                (function IIFE() {
                    for (var i = 0; i < myLibs.length; i++) {
                        if (myLibs[i].Name == molLibrary) {
                            libInfo = myLibs[i];
                        }
                    }
                })();
                var libJson = JSON.stringify(libInfo);

                var jobJson = JSON.stringify(jobInfo);
                loading2();
                $.ajax({
                    url: "/jobSubmit",
                    type: "POST",
                    data: {
                        libJson: libJson,
                        jobJson: jobJson,
                        queryUpload: 0
                    },
                    dataType: "html"
                }).done(jobHandle);
            }

            if (firstBottomRadio && secondTopRadio) {
                var molLibrary = $(".js-selectLibMol .select2-selection__rendered").html();
                var libInfo = null;
                (function IIFE() {
                    for (var i = 0; i < sysLibs.length; i++) {
                        if (sysLibs[i].Name == molLibrary) {
                            libInfo = sysLibs[i];
                        }
                    }
                })();
                var libJson = JSON.stringify(libInfo);

                var jobInfo = new Object();
                jobInfo.QueryMol = '';
                jobInfo.Nodes = parseInt(nodeCount);
                var jobJson = JSON.stringify(jobInfo);
                var formData = new FormData();

                // 获取上传文件对象
                var fileObj1 = document.getElementById("queryMol").files[0];
                var fileName1Split;
                if (!fileObj1) {
                    alert("Please upload a sd/mol file!");
                    return false;
                } else {
                    fileName1Split = fileObj1.name.toUpperCase().split(".");
                }
                if (fileObj1.size > (1024 * 1024 * 5)) {
                    alert("The file size exceeds the limit(5M), please re-upload a file!");
                    return false;
                }
                //	@js 倒叙遍历数组，来判断"."分割后最后一个字符是否含有"SD/SDF/MOL" 
                for (var i = fileName1Split.length - 1; i >= 0; i--) {
                    if (i = fileName1Split.length - 1) {
                        if ((fileName1Split[i] == "MOL" || fileName1Split[i] == "SD" || fileName1Split[i] == "SDF")) {
                            formData.append("libJson", libJson);
                            formData.append("jobJson", jobJson);
                            formData.append("queryUpload", 1);
                            formData.append("query", fileObj1);
                            loading2();
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
                            alert("The file format is wrong, please re-upload a SD/SDF file or mol File!");
                            return false;
                        }
                    }
                }
            }

            if (firstBottomRadio && secondBottomRadio) {
                var molLibrary = $(".js-select-myLib .select2-selection__rendered").html();
                var libInfo = null;

                for (var i = 0; i < myLibs.length; i++) {
                    if (myLibs[i].Name == molLibrary) {
                        //alert(myLibs[i])
                        //alert(molLibrary)
                        libInfo = myLibs[i];
                    }
                }

                var libJson = JSON.stringify(libInfo);
                //return

                var jobInfo = new Object();
                jobInfo.QueryMol = '';
                jobInfo.Nodes = parseInt(nodeCount);
                var jobJson = JSON.stringify(jobInfo);
                var formData = new FormData();

                // 获取上传文件对象
                var fileObj1 = document.getElementById("queryMol").files[0];
                var fileName1Split;
                if (!fileObj1) {
                    alert("Please upload a sd/mol file!");
                    return false;
                } else {
                    fileName1Split = fileObj1.name.toUpperCase().split(".");
                }
                if (fileObj1.size > (1024 * 1024 * 5)) {
                    alert("The file size exceeds the limit(5M), please re-upload a file!");
                    return false;
                }
                //	@js 倒叙遍历数组，来判断"."分割后最后一个字符是否含有"SD/SDF/MOL" 
                for (var i = fileName1Split.length - 1; i >= 0; i--) {
                    if (i = fileName1Split.length - 1) {
                        if ((fileName1Split[i] == "MOL" || fileName1Split[i] == "SD" || fileName1Split[i] == "SDF")) {
                            formData.append("libJson", libJson);
                            formData.append("jobJson", jobJson);
                            formData.append("queryUpload", 1);
                            formData.append("query", fileObj1);
                            loading2();
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
                            alert("The file format is wrong, please re-upload a SD/SDF file or mol File!");
                            return false;
                        }
                    }
                }
            }
        }
    });


});

//	记录当前选择的靶标所对应的分子数量
var ligndCount;
//	每页加载的3dmol窗口数量:默认为8
var molCntPerPage = 8;
//	显示完所有分子需要的页数：默认为1
var maxPages = 1;
//	当前页码:默认为1
var currentPage = 1;

var perPageMolArr = new Array();

var selectMolStr = "";

var allMol;

function showMolStruc(msg) {
    perPageMolArr = [];
    var molJson = JSON.parse(msg);
    if (molJson.State != 0) {
        $("#step-1 .bottomSelectSection").css("margin-top", "0");
        alert("No data returned from the serve.");

        $(".JQ_pagination").hide();
        clear3DMol(perPageMolArr);
        $(".mol-structure").hide();
        $(".loading-component").css("display", "none");
        return false;
    }
    //clear3DMol();
    console.log(molJson);
    ligndCount = molJson.Data.ligands.length;

    allMol = molJson.Data.ligands;

    //  向perPageMolArr添加数据
    if (molCntPerPage <= ligndCount) {
        for (var i = 0; i < molCntPerPage; i++) {
            perPageMolArr.push(molJson.Data.ligands[i]);
        }
    } else {
        for (var i = 0; i < ligndCount; i++) {
            perPageMolArr.push(molJson.Data.ligands[i]);
        }
    }
    clear3DMol(perPageMolArr);
    console.log(perPageMolArr);

    $(".loading-component").css("display", "none");
    load3DMol(perPageMolArr);
    // 初始化3Dmol&分子显示
    if (perPageMolArr.length > 4) {
        $("#wizard").children("ul").css("height", "726");
        $("#step-1").css("height", "720");
        $(".stepContainer").css("height", "720");
        $("#step-1 .bottomSelectSection").css("margin-top", "25px");
    } else {
        $("#wizard").children("ul").css("height", "526");
        $("#step-1").css("height", "520");
        $(".stepContainer").css("height", "520");
        $("#step-1 .bottomSelectSection").css("margin-top", "25px");
    }

    if (ligndCount % molCntPerPage === 0) {
        maxPages = ligndCount / molCntPerPage
    } else {
        maxPages = Math.ceil(ligndCount / molCntPerPage);
    }
    currentPage = 1;
    //初始化翻页组件
    var tempPaginationClass = "JQ_pagination_" + ligndCount;
    $("#paginate").remove();
    $("#step-1 .topSelectSection").append("<div class='JQ_pagination pull-right " + tempPaginationClass + "' id='paginate'>" +
        "<a href='#' class='first' data-action='first'>&laquo;</a><a href='#' class='previous' data-action='previous'>&lsaquo;</a>" +
        "<input type='text' readonly='readonly' data-max-page='40' /><a href='#' class='next' data-action='next'>&rsaquo;</a><a href='#' class='last' data-action='last'>&raquo;</a></div>");
    $("." + tempPaginationClass).jqPagination({
        current_page: currentPage,
        max_page: maxPages,
        paged: function(page) {
            //点击页码要做的操作

            page = parseInt(page, 10);
            currentPage = page;
            console.log("currentPage==" + currentPage);
            var startMol = (page - 1) * molCntPerPage;
            var endMol = page * molCntPerPage - 1;
            var selectMolID = window.localStorage ? localStorage.getItem("selectMolID") : Cookie.read("selectMolID");
            if (endMol <= ligndCount) {
                perPageMolArr = [];
                perPageMolArr = molJson.Data.ligands.slice(startMol, endMol + 1);
                clear3DMol(perPageMolArr);
                load3DMol(perPageMolArr);

                // 重新加载选中分子
                $("#molContainer-left canvas").remove();
                loadSelectMol(selectMolStr);
                //	@js 监测当前页是否有先前选中的目标
                $(".small-radio").removeClass("checkedRadio");
                var localCurrentPage = window.localStorage ? localStorage.getItem("currentPage") : Cookie.read("currentPage");

                if (currentPage == localCurrentPage) {
                    $(".js-molInfoPopup a").each(function() {
                        if ($(this).attr("data-id") == selectMolID) {
                            $(this).parent().parent().find(".small-radio").addClass("checkedRadio");
                        }
                    });
                }
            } else {
                perPageMolArr = [];
                perPageMolArr = molJson.Data.ligands.slice(startMol);
                clear3DMol(perPageMolArr);
                load3DMol(perPageMolArr);
                // 重新加载选中分子
                $("#molContainer-left canvas").remove();
                loadSelectMol(selectMolStr);
                //	@js 监测当前页是否有先前选中的目标
                $(".small-radio").removeClass("checkedRadio");
                var localCurrentPage = window.localStorage ? localStorage.getItem("currentPage") : Cookie.read("currentPage");

                if (currentPage == localCurrentPage) {
                    $(".js-molInfoPopup a").each(function() {
                        if ($(this).attr("data-id") == selectMolID) {
                            $(this).parent().parent().find(".small-radio").addClass("checkedRadio");
                        }
                    });
                }

            }
        }
    });
    $(".JQ_pagination").show();
    $(".mol-structure").show();

    /*	
     **	@bug	$("#wizard").children("ul").css("height","506");存在时，底部按钮条的宽度会变一下。
     */
    /*$("#step-1").css("height","500");
    $(".stepContainer").css("height","500");
    $(".set-position").hide();*/
    //$("#wizard").children("ul").css("height","506");



}

//	清空前一页的所有3dMol数据
function clear3DMol() {
    //	原生3DMol拥有一个clear函数，可以清空，但需由对象调用。
    //	故直接删除DOM元素
    $(".mol-canvas").children("canvas").remove();

}

function load3DMol(molData) {
    var currentMolCnt = molData.length;
    show(currentMolCnt);
    hide(currentMolCnt);
    //	加载mol数据
    for (var i = 0; i < currentMolCnt; i++) {
        var cid = "container-" + i;
        var element = $('.' + cid);
        element.attr("data-index", i);
        element.children(".funcBars1").children(".viewMolInfo").children("a").attr("data-id", molData[i].Id);
        var molInfoEl = element.children(".funcBars2").children(".simpleInfo");
        console.log(molInfoEl);
        molInfoEl.children(".molName").text(molData[i].Name);
        // molInfoEl.children(".molID").text(molData[i].Id);
        // molInfoEl.children(".molID").css('display', 'none');

        for (var m = 0, len = molData[i].Activities.length; m < len; m++) {
            if (m == 0) {
                var activityStr = molData[i].Activities[m].Type + molData[i].Activities[m].Relation +
                    molData[i].Activities[m].Value + molData[i].Activities[m].Units;
                molInfoEl.children(".activity").text(activityStr);
                break;
            }
        }
        var config = { backgroundColor: 'white' };
        var tDmolel = $('#' + cid);
        var viewer = $3Dmol.createViewer(tDmolel, config);
        var v = viewer;
        v.addAsOneMolecule(molData[i].Data, "sdf");
        v.zoomTo();
        v.render();
    }
}

function hide(cnt) {
    for (var i = cnt; i < molCntPerPage; i++) {
        var cid = "container-" + i;
        $("." + cid).css("display", "none");
    }
}

function show(cnt) {
    for (var i = 0; i < molCntPerPage; i++) {
        var cid = "container-" + i;
        $("." + cid).show();
        $(".list ." + cid).css("display", "flex");
    }
}
$(".js-molInfoPopup a").on({
    click: function() {
        var elID = $(this).attr("data-id");
        for (var i = 0, len = perPageMolArr.length; i < len; i++) {
            if (elID == perPageMolArr[i].Id) {
                var molActivityCnt = perPageMolArr[i].Activities.length;
                var refCnt = perPageMolArr[i].References.length;
                $(".molDetail tbody").append("<tr><th>MolID</th><td>" + perPageMolArr[i].Id +
                    "</td><th>Mol Name</th><td>" + perPageMolArr[i].Name + "</td></tr>");
                // 判断当前分子是否含有activity信息，没有则显示not available
                if (molActivityCnt == 0) {
                    $(".molDetail tbody").append("<tr><th>Activities</th><td>Not available</td>" +
                        "<th>References</th><td>Not available</td></tr>");
                } else {
                    // 含有活性信息，则将此信息循环添加进DOM中
                    for (var m = 0; m < molActivityCnt; m++) {
                        // 构建活性信息字符串
                        var activityStr = perPageMolArr[i].Activities[m].Type + perPageMolArr[i].Activities[m].Relation +
                            perPageMolArr[i].Activities[m].Value + perPageMolArr[i].Activities[m].Units;
                        // 判断对应文献是否存在，如果不存在则显示为not available
                        if (perPageMolArr[i].References[m].Journal == "") {
                            $(".molDetail tbody").append("<tr><th>Activities</th><td>" + activityStr + "</td>" +
                                "<th>References</th><td>Not available</td></tr>");
                        } else {
                            // 构建文献数据字符串
                            var refStr = perPageMolArr[i].References[m].Journal + perPageMolArr[i].References[m].Year +
                                perPageMolArr[i].References[m].Issue + "(" + perPageMolArr[i].References[m].Volume + "): " +
                                perPageMolArr[i].References[m].First_page + "-" + perPageMolArr[i].References[m].Last_page + ".";
                            $(".molDetail tbody").append("<tr><th>Activities</th><td>" + activityStr + "</td>" +
                                "<th>References</th><td>" + refStr + "</td></tr>");
                        }
                    }
                }
                // 显示分子的活性信息
                $(".molDetail").show();
                break
            }
        }
    }
});
$(".close").on({
    click: function() {
        $(this).parent(".panel-body").find("tr").remove();
        $(".molDetail").css("display", "none");
    }
});
/*
 **Custom checkbox and radio button
 */
$(".canCheckedRadio").on('click', function() {
    var _this = $(this),
        block = _this.parent().parent().parent();
    //console.log(block);
    block.find('.canCheckedRadio').children('input:radio').attr('checked', false);
    block.find(".canCheckedRadio").removeClass('checkedRadio');
    _this.addClass('checkedRadio');
    _this.find('.canCheckedRadio').children('input:radio').attr('checked', true);
});

$(".small-radio").on({
    click: function() {
        var _this = $(this),
            block = _this.parent().parent();
        //此点击只为了在左侧窗口查看分子
        block.parent().find('.small-radio input').attr('checked', false);
        block.parent().find(".small-radio").removeClass('checkedRadio');
        _this.addClass('checkedRadio');
        _this.find('input:radio').attr('checked', true);

        // 1.获取数据
        var selectMolID = block.find('.viewMolInfo').children("a").attr("data-id");
        var selectMolName = block.find('.simpleInfo').children(".molName").text();
        var selectMolAct = block.find('.simpleInfo').children(".activity").text();

        for (var i = 0, len = allMol.length; i < len; i++) {
            if (allMol[i].Id == selectMolID) {
                selectMolStr = allMol[i].Data;
                break
            }
        }

        // 2.存储数据
        //存储，IE6~7 cookie 其他浏览器HTML5本地存储
        console.log("currentPage====" + currentPage);
        if (window.localStorage) {
            localStorage.setItem("selectMolID", selectMolID);
            localStorage.setItem("currentPage", currentPage);
        } else {
            Cookie.write("selectMolID", selectMolID);
            Cookie.write("currentPage", currentPage);
        }
        // 3.显示数据
        $(".selMolID-left span").html(selectMolID);
        $("#molContainer-left .activity").html(selectMolAct);
        $("#molContainer-left .selectMOlname").html(selectMolName);
        $("#molContainer-left").children("canvas").remove();
        console.log(selectMolStr);
        loadSelectMol(selectMolStr);
        $(".selectMolContent").show();
    }
});


//
function jobHandle(msg) {
    var jobInfo = JSON.parse(msg);
    console.log(jobInfo);
    if (jobInfo.State == -1) {
        if (jobInfo.Msg.indexOf("exist") >= 0) {
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

function loadSelectMol(selectMolStr) {
    var element = $("#molContainer-left");
    var config = { backgroundColor: 'white' };
    var viewer = $3Dmol.createViewer(element, config);
    var v = viewer;
    v.addAsOneMolecule(selectMolStr, "sdf");
    v.zoomTo();
    v.render();
};