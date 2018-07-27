# etcdsh [![Build Status](https://travis-ci.org/kamilhark/etcdsh.svg?branch=master)](https://travis-ci.org/kamilhark/etcdsh.svg?branch=master)

`etcdsh` is a command line tool for [etcd](https://github.com/coreos/etcd).
etcdsh provides filesystem-like access to etcd structure. 
Although there is an official command line tool [etcdctl](https://github.com/coreos/etcd/tree/master/etcdctl), it is annoying you have to enter the same (often very long) keys again for every command. etcdsh tries to make it simpler and faster by providing history and tab completion.

## Building
`etcdsh` is written using go language, it can be built using standard go build tool:

`go get github.com/kamilhark/etcdsh`

## Downloads binaries
 * [mac](https://github.com/kamilhark/etcdsh/releases/download/0.0.1-ALPHA/etcdsh-mac.zip) 
 * [linux](https://github.com/kamilhark/etcdsh/releases/download/0.0.2-ALPHA/etcdsh_mac.tar.gz)

## Examples
<pre>
<code>./etcdsh --url http://localhost:4001</code>
<code>connected to etcd 2.0.0</code>
<code>/>cd foo/bar</code>
<code>/foo/bar>ls</code>
<code>dir1</code>
<code>dir2</code>
<code>/foo/bar>set key value</code>
<code>/foo/bar>get key</code>
<code>value</code>
<code>/foo/bar>rm key</code>
<code>/foo/bar>exit</code>
</pre>

