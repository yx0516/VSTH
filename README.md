# VSTH
 VSTH is a web server that allows the end user to virtually screen the big compound library by employing high-performance computing power available on Tianhe-2.The VSTH is free and open to all users at http://114.67.37.143:9081.


# 1. Code structure
VSTH is implemented with WEGA engine. 3Dmol.js is used as the chemical structure viewer. Beego framework, Bootstrap are used to serve the web pages. HHVSF framework is used to accelerate large scale VS on supercomputers. 

```
$ tree VSTH
├─Admin        # ORM models defination.
├─Common
├─conf         # Configure files.
├─ctrls        # Controller objects.
├─doc
├─logs
├─routers
├─static       # Js files, css files, images etc.
├─tools
├─vendor       # Third part packages used in VSTH.
└─views        # HTML files.
```

VSTH implementation details are described in the following table.

| Tools	|Version	|Purpose	|Linkage|
| --------   | -----   | ---- | ----- |
|WEGA	       |1.0	      |Molecule structure similarity calculation	 |- |
|OpenBabel	  |2.4.1	    |Chemical file format conversion	           |http://openbabel.org/wiki/Main_Page |
|3Dmol.js	   |-	        |3D structure display	                      |http://3dmol.csb.pitt.edu/ | 
|HHVSF	      |1.0	      |A framework to accelerate large scale VS on supercomputer	 | https://github.com/pincher-chen/HHVSF |
|Beego	      |1.8.3	    |A web server framework based on golang 	| https://beego.me/ |
|Bootstrap	  |3.3.7	    |Front-end component library	| http://getbootstrap.com/ |
|MongoDB	    |3.2.9	    |Storage database	| https://www.mongodb.com/ |


# 2. How to compile
The back-end of VSTH is written by golang. It is easy to launch web services by using go command.
```
$ go build main.go
```


