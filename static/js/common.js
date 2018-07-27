/*
 **	2016-11-01
 **	所有通用函数都包含于此文件内
 **  作者：杜杰文
 */

// 存储zinc的对应信息
var zincArr = [{
    type: 1,
    front: "ZINC_All-Purchasable",
    end: "zinc_conformer_all"
}, {
    type: 1,
    front: "ZINC_Purchasable_Part_1M",
    end: "zinc_conformer_100w"
}, {
    type: 2,
    front: "ZINC_All-Purchasable",
    end: "zinc_ligand_1000w_sort"
}, {
    type: 2,
    front: "ZINC_Purchasable_Part_100K",
    end: "zinc_ligand_10w_sort"
}, {
    type: 2,
    front: "ZINC_Purchasable_Part_1M",
    end: "zinc_ligand_100w_sort"
}, {
    type: 2,
    front: "ZINC_Purchasable_Part_10K",
    end: "zinc_ligand_1w_sort"
}];


/* @js	获取地址栏的jobid,用正则
------------------------------------------------------------------------------------------------------------*/
function GetQueryString(name) {
    var reg = new RegExp("(^|&)" + name + "=([^&]*)(&|$)");
    var r = window.location.search.substr(1).match(reg);
    if (r != null) return unescape(r[2]);
    return null;
}

/* @js  侧边栏下拉菜单 & 鼠标滑过字体变大 & 各种链接
------------------------------------------------------------------------------------------------------------*/
$(".gn-icon-creat").on({
    click: function() {
        var _gn_submenu = $(".gn-submenu").attr("data-display");
        if (_gn_submenu.indexOf("none") >= 0) {
            $(".gn-submenu").slideDown(400);
            $(".gn-submenu").attr("data-display", "block");
        } else {
            $(".gn-submenu").slideUp(400);
            $(".gn-submenu").attr("data-display", "none");
        }
    }
});
$(".gn-icon-library").on({
    click: function() {
        var _gn_submenu = $(".gn-submenu2").attr("data-display");
        if (_gn_submenu.indexOf("none") >= 0) {
            $(".gn-submenu2").slideDown(400);
            $(".gn-submenu2").attr("data-display", "block");
        } else {
            $(".gn-submenu2").slideUp(400);
            $(".gn-submenu2").attr("data-display", "none");
        }
    }
});
$(".gn-icon-sys-library").on({
    click: function() {
        window.location.href = "/sysLibrary.html";
    }
});
$(".gn-icon-my-library").on({
    click: function() {
        window.location.href = "/myLibrary.html";
    }
});

$(".gn-menu li").on({
    mouseenter: function() {
        $(this).children("a").css("font-weight", "700");
    },
    mouseleave: function() {
        $(this).children("a").css("font-weight", "300");
    }
});

$(".gn-icon-jobInfo").on({
    click: function() {
        // if (!$(".gn-submenu").hasClass("index-ul")) {
        //     window.location.href = "/Myjob.html";
        // } else {
        //     return false;
        // }
        window.location.href = "/Myjob.html";
    }
});

$(".gn-icon-LBVS").on({
    click: function() {
        window.location.href = "/lbvs.html";
    }
});
$(".gn-icon-SBVS").on({
    click: function() {
        window.location.href = "/sbvs.html";
    }
});
/*$(".gn-icon-MD").on({
    click: function(){
        window.location.href= "/md.html";
    }
});*/
/*$(".gn-icon-help").on({
    click: function(){
        window.location.href= "/help.html";
    }
});
$(".gn-icon-info").on({
    click: function(){
        window.location.href= "/info.html";
    }
});*/
// 退出操作
$(".logOut").on({
    click: function() {
        window.location.href = "index.html";
    }
});
$(function() {
    $(".underDev").css("color", "#c1c9d1");
    //  @js 清除文件，使用自带的clear方法
    $(".clearSD").on({
        click: function() {
            console.warn($(this).parent().find('input[type="file"]').val());
            $(this).parent().find('input[type="file"]').filestyle('clear');
        }
    });
});