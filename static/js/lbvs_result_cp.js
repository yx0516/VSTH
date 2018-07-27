/*
**	2016-12-11
**	LBVS  结果展示
**	作者：RCDD杜杰文
*/

/* 	@js global variable
------------------------------------------------------------------------------------------------------------------------*/
//	结果JSON;
var resultJson;

//  供应商信息
var vendorObj = new Array();

//	判断对象是否为空
function isEmptyObject(e) {  
    var t;  
    for (t in e)  
        return !1;  
    return !0  
}
var jobID;

var viewer = null;
/* 	@js DOM Ready
------------------------------------------------------------------------------------------------------------------------*/
$(function() {
	loading();
	//	折叠菜单初始化
	$('#collapseOne').collapse('hide');
	$('#collapseTwo').collapse('hide');
	$('#collapseThree').collapse('hide');
	$('#collapseFour').collapse('hide');
	$('#collapseFive').collapse('hide');
	$('#collapseSix').collapse('hide');
	
	//	 导航栏userName更新
	// var userName = window.localStorage? localStorage.getItem("userName"): Cookie.read("userName");
	// $(".user-info span").html(userName);
	//	 向后台请求结果数据
	jobID = window.localStorage? localStorage.getItem("currentJobID"): Cookie.read("currentJobID");
	
	$.ajax({
  		url : "jobResult",
  		type : "POST",
  		data: {
  			jobId: jobID
  		},
  		dataType: "html"
  	}).done(resultInitial); 
  	 	
});
//	@js 结果数据处理 & 初始化展示
function resultInitial(msg){
	resultJson = JSON.parse(msg);
	console.log(resultJson);
	if (resultJson.State == -1) {
		alert(resultJson.Data);
		console.warn("结果数据返回错误");
		return false;
	} else if (resultJson.State == 0){
		// 默认显示第一个小分子的结构
		var molArray = resultJson.Data.mols;
		var defaultMolStr = molArray[0].MolData +"$$$$\n";
		// 新建3DMol对象
		var element = $('#TDMol-overlay');

		var config = { backgroundColor: 'white' };
		var viewer = $3Dmol.createViewer( element, config );		

		var queryMolStr;
		// type:1,是选择的库分子
		if (resultJson.Data.type==1){
			// viewer ：显示queryMol
		 //  var v1 = viewer;
			// queryMolStr = resultJson.Data.query.Data + "$$$$\n";
			// v1.addAsOneMolecule(queryMolStr, "sdf");
			// v1.setStyle({elem:'H', invert: true},{stick:{colorscheme:"greenCarbon",radius:0.15,hidden:false}});
	  //   v1.zoomTo();
	  //   v1.render();
		  
			// var v2 = viewer;
			// v2.addAsOneMolecule(defaultMolStr, "sdf");
			// v2.setStyle({elem:'H', invert: true},{stick:{radius:0.15}});
	  //   v2.zoomTo();
	  //   v2.render();
	  queryMolStr = resultJson.Data.query.Data + "$$$$\n";
	  var v1 = viewer.addModel(queryMolStr, "sdf");
	  v1.setStyle({elem:'H', invert: true},{stick:{colorscheme:"greenCarbon",radius:0.15,hidden:false}});

	  var v2 = viewer.addModel(defaultMolStr, "sdf");
		v2.setStyle({elem:'H', invert: true},{stick:{radius:0.15}});
                             
		viewer.zoomTo();
		viewer.render();
		    //	1：添加queryMol信息
		    var _molPropEl = $(".queryMolInfo");
		    _molPropEl.append("<a class='list-group-item'><span class='text-info'>Mol ID: </span>" + resultJson.Data.query.Id  +  "</a>");
		    _molPropEl.append("<a class='list-group-item'><span class='text-info'>MolName: </span>" + resultJson.Data.query.Name  +  "</a>");
		    // 判断当前分子是否含有activity信息，没有则显示not available
			var molActivityCnt= resultJson.Data.activities.length;
			if (molActivityCnt==0){
				_molPropEl.append("<a class='list-group-item'><span class='text-info'>Activities:</span>Not available</a>");
				_molPropEl.append("<a class='list-group-item'><span class='text-info'>References:</span>Not available</a>");
			} else {
				// 含有活性信息，则将此信息循环添加进DOM中
				for (var m = 0; m<molActivityCnt; m++) {
					// 构建活性信息字符串
					var activityStr = resultJson.Data.activities[m].Type + resultJson.Data.activities[m].Relation + 
						resultJson.Data.activities[m].Value + resultJson.Data.activities[m].Units;
					// 判断对应文献是否存在，如果不存在则显示为not available
					if (resultJson.Data.references[m].Journal==""){
						_molPropEl.append("<a class='list-group-item'><span class='text-info'>Activities:</span>"+  activityStr +"</a>");
						_molPropEl.append("<a class='list-group-item'><span class='text-info'>References:</span>Not available</a>");
					} else {
						// 构建文献数据字符串
						var refStr = resultJson.Data.references[m].Journal + resultJson.Data.references[m].Year
						+ resultJson.Data.references[m].Issue + "(" + resultJson.Data.references[m].Volume + "): "
						+ resultJson.Data.references[m].First_page + "-" + resultJson.Data.references[m].Last_page +".";					 	
						_molPropEl.append("<a class='list-group-item'><span class='text-info'>Activities:</span>"+  activityStr +"</a>");
						_molPropEl.append("<a class='list-group-item'><span class='text-info'>References:</span>" + refStr + "</a>");
					}
				}
			}
			$('#collapseOne').collapse('show');
			
		} else if (resultJson.Data.type==2){
			// type:2,用户上传的分子
			queryMolStr = resultJson.Data.query.Name + "\n" + resultJson.Data.query.Comment + "\n\n" + resultJson.Data.query.Data + "$$$$\n";
			// viewer ：显示queryMol

		  // 	var v1 = viewer;
				// v1.addAsOneMolecule(queryMolStr, "sdf");
				// v1.setStyle({elem:'H', invert: true},{stick:{colorscheme:"greenCarbon",radius:0.15}});
		  //   v1.zoomTo();
		  //   v1.render();
		  // 	// viewer2 ：显示Library小分子
				// var v2 = viewer;
				// v2.addAsOneMolecule(defaultMolStr, "sdf");
				// v2.setStyle({elem:'H', invert: true},{stick:{radius:0.15}});
		  //   v2.zoomTo();
		  //   v2.render();
		  var v1 = viewer.addModel(queryMolStr, "sdf");
	  	v1.setStyle({elem:'H', invert: true},{stick:{colorscheme:"greenCarbon",radius:0.15,hidden:false}});

	  	var v2 = viewer.addModel(defaultMolStr, "sdf");
			v2.setStyle({elem:'H', invert: true},{stick:{radius:0.15}});
			viewer.zoomTo();
			viewer.render();
				//	1：添加queryMol信息
		    var _molPropEl = $(".queryMolInfo");
		    var attrCnt = resultJson.Data.query.Attrs.length;
		    if (attrCnt != 0) {
		    	for (var n = 0; n < attrCnt; n++) {
		    		_molPropEl.append("<a class='list-group-item'><span class='text-info'>" + resultJson.Data.query.Attrs[n].Name + ":&nbsp;</span>" + resultJson.Data.query.Attrs[n].Value + "</a>");
		    	}
		    } else {
		    	_molPropEl.append("Query molecule property not available");
		    }
		    
		    $('#collapseOne').collapse('show');
		    
		}

		//	2.1：添加小分子的ID及对接打分信息
		var _molIDEl = $(".molIDList");
		var _molScoreEl = $(".score");
		
		for (var i = 0 , molCnt = molArray.length; i < molCnt; i++) {
			if (i==0) {
				_molIDEl.append("<a class='list-group-item active text-info' data-id='"+ resultJson.Data.mols[i].MolName +"'>" + resultJson.Data.mols[i].MolName + "</a>");
				_molScoreEl.append("<a class='list-group-item active text-info' data-id='"+ resultJson.Data.mols[i].MolName +"'>" + resultJson.Data.mols[i].Score + "</a>");
			} else {
				_molIDEl.append("<a class='list-group-item text-info' data-id='"+ resultJson.Data.mols[i].MolName +"'>" + resultJson.Data.mols[i].MolName + "</a>");
				//_molScoreEl.append("<a class='list-group-item text-info' data-id='"+ molArray[i].MolName +"'>" + molArray[i].Score + "</a>");
			}
		}
		
		$('#collapseTwo').collapse('show');
		$('#collapseFour').collapse('show');

		//	2.2：给ID列表和Score列表绑定点击事件
		$(".molIDList a").on({
			click: function(){
				$(".molIDList a").removeClass("active");
				$(this).addClass("active");
				var molID = $(this).attr("data-id");
				var molStr = "";
				
				for (var i = 0 , len = resultJson.Data.mols.length; i < len; i++) {
					if (resultJson.Data.mols[i].MolName == molID ){
						console.log(molID);
						molStr = resultJson.Data.mols[i].MolData + "$$$$\n";
						//molStr = resultJson.Data.mols[i].MolData;
						//$(".score a").removeClass("active");
						$(".score a").remove();
						_molScoreEl.append("<a class='list-group-item text-info' data-id='"+ resultJson.Data.mols[i].MolName +"'>" + resultJson.Data.mols[i].Score + "</a>");
						$(".score a").addClass("active");
					}
				}
				// 更新蛋白质和小分子的显示
				viewer.clear();
				var v1 = viewer.addModel(queryMolStr, "sdf");
		  	v1.setStyle({elem:'H', invert: true},{stick:{colorscheme:"greenCarbon",radius:0.15,hidden:false}});

		  	var v2 = viewer.addModel(molStr, "sdf");
				v2.setStyle({elem:'H', invert: true},{stick:{radius:0.15}});
				viewer.zoomTo();
				viewer.render();
			
				// var v1 = viewer;
				// v1.addAsOneMolecule(queryMolStr, "sdf");
				// v1.setStyle({elem:'H', invert: true},{stick:{colorscheme:"greenCarbon",radius:0.15}});
		  //   v1.zoomTo();
		  //   v1.render();
				
			 //  var v2= viewer; 
				// v2.addAsOneMolecule(molStr, "sdf");
				// v2.setStyle({elem:'H', invert: true},{stick:{radius:0.15}});
				// v2.zoomTo();
				// v2.render();

			  addPropToList();
			}
		});
		addPropToList();
		$(".loading-component").hide();
	}	
} 
/*	@js DOM操作：各种列表的信息添加 
----------------------------------------------------------------------------------------------------------------*/		
function addPropToList () {
	var selectMolName = $(".molIDList .active").text();
	$.ajax({
		url:"getZincProp",
		type: "POST",
		data:{
			molName :　selectMolName
		},
		dataType : "html"
	}).done(loadProperty);
}

function loadProperty(msg){
	var propertyJson = JSON.parse(msg);
	console.log(propertyJson);
	$(".molPropList a").remove();
	if (propertyJson.State == -1) {
		//alert(propertyJson.Data);
		var _molProEl = $(".molPropList");
		_molProEl.append("<a class='list-group-item'><span class='text-info'>MolName:</span> Not available</a>");
		_molProEl.append("<a class='list-group-item'><span class='text-info'>MolWeight:</span> Not available</a>");
		_molProEl.append("<a class='list-group-item'><span class='text-info'>LogP:</span> Not available</a>");
		_molProEl.append("<a class='list-group-item'><span class='text-info'>Apolar desolvation energies:</span> Not available</a>");
		_molProEl.append("<a class='list-group-item'><span class='text-info'>Polar desolvation energies:</span> Not available</a>");
		_molProEl.append("<a class='list-group-item'><span class='text-info'>Hba:</span> Not available</a>");
		_molProEl.append("<a class='list-group-item'><span class='text-info'>Hbd:</span> Not available</a>");
		_molProEl.append("<a class='list-group-item'><span class='text-info'>Tpsa:</span> Not available</a>");
		_molProEl.append("<a class='list-group-item'><span class='text-info'>Charge:</span> Not available</a>");
		_molProEl.append("<a class='list-group-item'><span class='text-info'>Rotatable bonds:</span> Not available</a>");
		_molProEl.append("<a class='list-group-item'><span class='text-info'>Smile:</span> Not available</a>");
		_molProEl.append("<a class='list-group-item'><span class='text-info'>Inchikey:</span> Not available</a>");
		return false;
	} else {

		var _molProEl = $(".molPropList");
		var molObj = propertyJson.Data.zincMolProp;
		/*------------------------------- @js 2：小分子属性列表信息添加 ----------------------------------*/
		_molProEl.append("<a class='list-group-item'><span class='text-info'>MolName:</span> " + molObj.ZincMolName + "</a>");
		_molProEl.append("<a class='list-group-item'><span class='text-info'>MolWeight:</span> " + molObj.MW + "</a>");
		_molProEl.append("<a class='list-group-item'><span class='text-info'>LogP:</span> " + molObj.LogP + "</a>");
		_molProEl.append("<a class='list-group-item'><span class='text-info'>Apolar desolvation energies:</span> " + molObj.DesolvApolar + "</a>");
		_molProEl.append("<a class='list-group-item'><span class='text-info'>Polar desolvation energies:</span> " + molObj.DesolvPolar + "</a>");
		_molProEl.append("<a class='list-group-item'><span class='text-info'>Hba:</span> " + molObj.Hba + "</a>");
		_molProEl.append("<a class='list-group-item'><span class='text-info'>Hbd:</span> " + molObj.Hbd + "</a>");
		_molProEl.append("<a class='list-group-item'><span class='text-info'>Tpsa:</span> " + molObj.Tpsa + "</a>");
		_molProEl.append("<a class='list-group-item'><span class='text-info'>Charge:</span> " + molObj.Charge + "</a>");
		_molProEl.append("<a class='list-group-item'><span class='text-info'>Rotatable bonds:</span> " + molObj.Nrb + "</a>");
		_molProEl.append("<a class='list-group-item'><span class='text-info'>Smile:</span> " + molObj.Smile + "</a>");
		_molProEl.append("<a class='list-group-item'><span class='text-info'>Inchikey:</span> " + molObj.Inchikey + "</a>");

		$('#collapseThree').collapse('show');
		var selectMolName = $(".molIDList .active").text();
		$.ajax({
			url:"getZincSupply",
			type: "POST",
			data:{
				molName :　selectMolName
			},
			dataType : "html"
		}).done(loadSupply);
	}		
}

function loadSupply(msg){
	var supplyJson = JSON.parse(msg);
	if (supplyJson.State == -1) {
		alert(supplyJson.Data);
		return false;
	}
	console.log(supplyJson);
	$(".vendorList a").remove();
	$(".vendorDetail a").remove();
	var _vendorEl = $(".vendorList");
	var _vendorDetailEl = $(".vendorDetail");
	if (!supplyJson.Data.zincMolSupply){
		_vendorEl.append("<a class='list-group-item><span class='text-info'>Supplier:</span> Not available</a>");
		_vendorDetailEl.append("<a class='list-group-item'><span class='text-info'>SupplierCode:</span> Not available</a>"
								+ "<a class='list-group-item'><span class='text-info'>Phone:</span> Not available</a>"
								+ "<a class='list-group-item'><span class='text-info'>Fax:</span> Not available</a>"
								+ "<a class='list-group-item'><span class='text-info'>Email:</span> Not available</a>"
								+ "<a class='list-group-item'><span class='text-info'>WebSite:</span> Not available</a>");
		return
	} else {
		vendorObj = supplyJson.Data.zincMolSupply;
			/*------------------------------- @js 3: 小分子供应商信息添加 ------------------------------------*/
		for (var i = 0 , len = vendorObj.length; i < len; i++) {
			if (i==0) {
				_vendorEl.append("<a class='list-group-item active' data-id='" + vendorObj[i].Id + "'><span class='text-info'>Supplier:</span> " + vendorObj[i].Supplier + "</a>");
				_vendorDetailEl.append("<a class='list-group-item'><span class='text-info'>SupplierCode:</span> " + vendorObj[i].SupplierCode + "</a>"
					+ "<a class='list-group-item'><span class='text-info'>Phone:</span> " + vendorObj[i].Phone + "</a>"
					+ "<a class='list-group-item'><span class='text-info'>Fax:</span> " + vendorObj[i].Fax + "</a>"
					+ "<a class='list-group-item'><span class='text-info'>Email:</span> " + vendorObj[i].Email + "</a>"
					+ "<a class='list-group-item'><span class='text-info'>WebSite:</span> " + vendorObj[i].WebSite + "</a>");
			} else {
				_vendorEl.append("<a class='list-group-item' data-id='" + vendorObj[i].Id + "'><span class='text-info'>Supplier:</span> " + vendorObj[i].Supplier + "</a>");
			}
		}
		$('#collapseFive').collapse('show');
		$('#collapseSix').collapse('show');
		$(".vendorList a").on({
				click: function(){
					$(".vendorList a").removeClass("active");
					$(this).addClass("active");
					var vendorID = $(this).attr("data-id");
					console.log(vendorObj);
					for (var m = 0 , len = vendorObj.length; m < len; m++) {
						if (vendorObj[m].Id == vendorID ){
							$(".vendorDetail a").remove();
							_vendorDetailEl.append("<a class='list-group-item'><span class='text-info'>SupplierCode:</span> " + vendorObj[m].SupplierCode + "</a>"
								+ "<a class='list-group-item'><span class='text-info'>Phone:</span> " + vendorObj[m].Phone + "</a>"
								+ "<a class='list-group-item'><span class='text-info'>Fax:</span> " + vendorObj[m].Fax + "</a>"
								+ "<a class='list-group-item'><span class='text-info'>Email:</span> " + vendorObj[m].Email + "</a>"
								+ "<a class='list-group-item'><span class='text-info'>WebSite:</span> " + vendorObj[m].WebSite + "</a>");
							
						}
					}				
				}
		});	
	}	
}

//	@js loading加载
function loading(){
	
	var loadingCmpWidth = $(".sbvs-result").width()+30; 
	var loadingCmpHeight = $(".sbvs-result").height()+120;
	//设置loading-component的宽度和高度
	$(".loading-component").width(loadingCmpWidth);
	$(".loading-component").height(loadingCmpHeight);

	var loadingWidth = $(".loading").width();
	var loadingHeight = $(".loading").height();
	var left1 = loadingCmpWidth/2-loadingWidth/2;
    var top1 = loadingCmpHeight/2-loadingHeight/2;
    $(".loading").css("margin-left",left1 + "px");
    $(".loading").css("margin-top",top1 + "px");
	$(".loading-component").css("display","block");
}

$(".js-downloadMol").on({
	click: function() {
		var jobID = window.localStorage? localStorage.getItem("currentJobID"): Cookie.read("currentJobID");
		url = 'downloadJobResult?jobId=' + jobID;
		window.location.href = url;
		
	}
});