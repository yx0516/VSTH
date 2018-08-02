# VSTH
 VSTH is a web server that allows the end user to virtually screen the big compound library by employing high-performance computing power available on Tianhe-2.


# 1. Code structure
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


# 2. How to compile
The back-end of VSTH is written by golang. It is easy to launch web services by using go commnandm.
```
$ go build main.go
```


