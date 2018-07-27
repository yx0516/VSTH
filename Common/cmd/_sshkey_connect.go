package main

import (
	"bytes"
	"fmt"

	"github.com/golang/crypto/ssh"
)

const (
	KEY = `-----BEGIN RSA PRIVATE KEY-----
MIIEowIBAAKCAQEAn4y2JVQmyvTn7CenVVHAD7WsODTtcumB2VJCLE/hcvlIdb3V
5GdJmghtEsUXtk3I/JAio2bQn13KdXxDWVYM0eOUmHPnNMnpgHDGjR1bu/I+tgh3
DZmCQjDfXO5Ij7Lg9Xg1MEyTZBb4K1J6hUBOxLcMOqnD50Z+wD+VUsnGgVFAQ+Wn
iwAwUynPalx+k2Be2GxWx5TsEb6fvvZGxkIK0tsouYfKmUnzskl0G1ssdUqxO+F9
HJ2+ztyEmT/8NUWpgeh5kUVhJVy77ctid7LezYBZiCRqiwh72wYJ1KIhsaIh0Qbl
0E+MiPVyswqCw8v5XgpVWAVt3Ly+tQln3UEMmQIBIwKCAQEAmv24QYTyfAPZ+1m4
fsRxbli17XU+b6EBy8xASE2ZLdw3wtWrNaYqPdxMsyXchTWerDQ+5+A4YEx7iBJQ
DaQMdB7oHxGBg7zUM6DA4Nqw4nZLjEK8y2HsQEy7uVyBAJ8jrKCospOH3ZKn75Hz
XN9idfOletDM70vLli8jV7yyNHlH1MiW/lsaY7Lm5YcmxiFeeQ4EelUEfywdiVmK
RHUJqZBga6B7ZwnbiN7ERBqkDsN3lIA6C+oZnXSuFcXIwoKTZMZSElLZcbPiw2LI
D6RLi+dCzid0C2itrolU0sVMhhi6QBGovkCC8DlMdeYYxRMlZjns+c1Jkl0uTwmn
okl66wKBgQDRJhec2uEzLmtx+7tPhaPMF/9Z7bFmMWjY1Ez3zuyNUH3GfB0xnAWm
1xDd3zbVs0p8ppG8talQug52yfnZbqbUEB1Z3JP4hPOaWErsEh6wlFoaNLcNw//7
fuxcjXOyZbvC2S8AtW4TIfMUgl7QHWt2KcgMwoUAZ7Kjme3I3jZnpQKBgQDDSkrY
qki8VLNgr4b/nCHAjxrLEKofXWwg+2O5GkZsCIT++8zsx+dURsu9IzYXSWQqbWJX
eM69H+UAQGPyOp/ya+/pVolMjT8DcgQv+x/CVk7F2ZAK94Yn7DH0E9vNsszqiLao
Vwh6JG0H84ogrlRVLOiJWB4LIiKnq6C9ud0e5QKBgQChV+ZUblX0SGGD0NJTSdYh
GdOdJRPSfeMq7OrrDVdlsx8yt4Q8NogEXMPeW7yWOdpgKLmRk3P8ckWsGCckE4gC
rVh4hZa4ZpAJWg3pT86IN+28cc8KnoOkwP96mQF6/gXfdFd1k0ZJRhNKVfFet5wK
soRhnV9JdJEfHvlDLxQGzwKBgEiJTwAEu+4uFr2DBkGvBTjksi4qwthzKCmB8dcJ
wmKkFCQxo+LrKglcH8nogc4ikusv8VOwh02PgPF3AI0rt8BxO9pTV5F2SprhNMFd
Rk9/JJKhRCFUn4sr2AozQwNCaV5tSyiVWuuJ31q5ixNz83AX/plp30X2vGopzf1T
qebnAoGBALV0vBFnf5lbyVRApnZK6Cx+xKeoUN2B5fXstpGEaGRSMar3ffOBqFbR
wGc1Ku8XK5YC5YLTFgecqZfIWXZcqKrLMPWKNyJeEGu29027G0RAL2iay9aoNI56
EP+VJfJESNBTMChUdNchR3/uOFZtxbj9QT8BWiuLkUx9dYimeqUa
-----END RSA PRIVATE KEY-----`

	USER_NAME = "root"
	ADDR      = "192.168.200.186:22"
)

func rumCmd(strCmd string) (string, error) {
	signer, err := ssh.ParsePrivateKey([]byte(KEY))
	if err != nil {
		return "", err
	}

	config := &ssh.ClientConfig{
		User: USER_NAME,
		Auth: []ssh.AuthMethod{
			ssh.PublicKeys(signer),
		},
	}

	sshClient, err := ssh.Dial("tcp", ADDR, config)
	if err != nil {
		return "", err
	}
	defer sshClient.Close()

	sshSession, err := sshClient.NewSession()
	if err != nil {
		return "", err
	}
	defer sshSession.Close()

	buf := new(bytes.Buffer)
	sshSession.Stdout = buf
	if err = sshSession.Run(strCmd); err != nil {
		return "", err
	} else {
		return buf.String(), nil
	}
}

func main() {
	if out, err := rumCmd("df -h"); err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(out)
	}
}
