package cmd

import (
	"fmt"
	"os/exec"
	"strconv"
	//	"strings"
)

func RemoteCallScript(idFile string, username string, ip string, script string) error {

	userAndIp := username + "@" + ip
	cmd := exec.Command("ssh", "-i", idFile, userAndIp, "sh", script)

	if buf, err := cmd.Output(); err != nil {
		fmt.Println(cmd)
		return err
	} else {
		fmt.Println(string(buf))
		return nil
	}

}

func CallScp(idFile string, username string, ip string, fromFile string, toFilePath string) error {

	userAndIpToFile := username + "@" + ip + ":" + toFilePath
	cmd := exec.Command("scp", "-i", idFile, fromFile, userAndIpToFile)
	fmt.Println(cmd)
	if buf, err := cmd.Output(); err != nil {
		fmt.Println(cmd)
		return err
	} else {
		fmt.Println(string(buf))
		return nil
	}

}

func RemoteCallJobResInsertScript(idFile string, username string, ip string, script string, jobId string, dbName string) error {

	userAndIp := username + "@" + ip
	cmd := exec.Command("ssh", "-i", idFile, userAndIp, "sh", script, jobId, dbName)
	fmt.Println(cmd)
	if buf, err := cmd.Output(); err != nil {
		fmt.Println(cmd)
		return err
	} else {
		fmt.Println(string(buf))
		return nil
	}

}

func RemoteCallJobUpdateScript(idFile string, username string, ip string, script string, jobId string, dbName string) error {

	userAndIp := username + "@" + ip
	cmd := exec.Command("ssh", "-i", idFile, userAndIp, "sh", script, jobId, dbName)
	fmt.Println(cmd)
	if buf, err := cmd.Output(); err != nil {
		fmt.Println(cmd)
		return err
	} else {
		fmt.Println(string(buf))
		return nil
	}

}

func RemoteCallVinaScript(idFile string, username string, ip string, script string, jobId string, pdbCode string, library string, nodes int) error {

	userAndIp := username + "@" + ip
	cmd := exec.Command("ssh", "-i", idFile, userAndIp, "sh", script, jobId, pdbCode, library, strconv.FormatInt(int64(nodes), 10))
	fmt.Println(cmd)
	if buf, err := cmd.Output(); err != nil {
		fmt.Println(cmd)
		return err
	} else {
		fmt.Println(string(buf))
		return nil
	}
}

func RemoteCallWEGAScript(idFile string, username string, ip string, script string, jobId string, query string, library string, nodes int) error {

	userAndIp := username + "@" + ip
	cmd := exec.Command("ssh", "-i", idFile, userAndIp, "sh", script, jobId, query, library, strconv.FormatInt(int64(nodes), 10))
	fmt.Println(cmd)
	if buf, err := cmd.Output(); err != nil {
		fmt.Println(cmd)
		return err
	} else {
		fmt.Println(string(buf))
		return nil
	}
}
