package chaos

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"syscall"
)

func ReadDirNoLink(path string) (s []string) {
	_, err := os.Stat(path)
	if err != nil {
		err = os.MkdirAll(path, os.ModePerm)
		if err != nil {
			return
		}
	}
	list, err := ioutil.ReadDir(path)
	if err != nil {
		fmt.Printf("UserActRecConsumer |read dir err |message=%v\n", err)
		return
	}

	for _, v := range list {
		if IsLinkFile(path + "/" + v.Name()) {
			continue
		}
		log.Println("UserActRecConsumer |fils |message=", v.Name())
		s = append(s, v.Name())
	}

	return
}

func IsLinkFile(filename string) bool {
	fi, err := os.Lstat(filename)
	if err != nil {
		return false
	}
	s, ok := fi.Sys().(*syscall.Stat_t)
	if !ok {
		return false
	}
	inode := uint64(s.Ino)
	nlink := uint32(s.Nlink)
	if nlink > 1 {
		log.Printf("Inode %v has %v other hardlinks besides %v \n.", inode, nlink, filename)
	}

	if fi.Mode()&os.ModeSymlink != 0 {
		_, err := os.Readlink(filename)
		if err != nil {
			return false
		}
		return true
	} else {
		return false
	}

	return true
}

func IsExist(path string) bool {
	if _, err := os.Stat(path); err == nil {
		return true
	} else {
		if os.IsNotExist(err) {
			return false
		}
	}
	return false
}

func RemoveFile(dis, src string) {

}
