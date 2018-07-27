var mymol
var marvinSketcherInstance;

$(document).ready(function handleDocumentReady (e) {
	var p = MarvinJSUtil.getEditor("#sketch");
	p.then(function (sketcherInstance) {
		marvinSketcherInstance = sketcherInstance;
		//initControl();
	}, function (error) {
		alert("Cannot retrieve sketcher instance from iframe:"+error);
	});
});
function submit(){
	marvinSketcherInstance.exportStructure("mol").then(function(source) {
		ajax(source);
	})
}

function ajax(source){
	Mol = source;
	Method = 1
	Simi = 0.8
	console.log("Mol:%s",Mol)
	$.ajax({
			url: "submit",
			type: "POST",
			data: {mol:Mol,
			       simi:Simi,
			       method:Method
				   
			},
			dataType: "html"
		}).done(showResponse)
}

/*
function initControl () {

	// get mol button
	$("#btn-getmol").on("click", function (e) {
		marvinSketcherInstance.exportStructure("mol").then(function(source) {
			//$("#molsource").text(source);
			//console.log("source = %s",source)
			mymol = source;
			console.log("mymol1 = %s",mymol)
		}, function(error) {
			alert("Molecule export failed:"+error);	
		});
	});
}*/

/*
function submit(){
	marvinSketcherInstance.exportStructure("mol").then(function(source) 					    {
			//$("#molsource").text(source);
			//console.log("source = %s",source)
			mymol = source;
			console.log("mymol1 = %s",mymol)
	}, function(error) {
			alert("Molecule export failed:"+error);	
	});
		console.log(mymol)
}
*/


/*

	$.ajax({
			url: "submit",
			type: "POST",
			data: {molecule:mol
			},
			dataType: "html"
		}).done(showResponse)
*/