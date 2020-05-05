package handle

import (
	"archive/zip"
	"io"
	"io/ioutil"
	"log"
	"moriaty.com/cia/cia-common/base/constant"
	"os"
	"os/exec"
	"path/filepath"
)

// 执行测试
func Exec(file, testFile, platformName, platformVersion, deviceName, apkName, imgLocation string) error {

	_, err := os.Stat(imgLocation)
	if err != nil {
		if os.IsNotExist(err) {
			err = os.MkdirAll(imgLocation, 0755)
			if err != nil {
				log.Printf("create dir [%s] failed, err: %v", imgLocation, err)
				return err
			}
		}
	}

	dest, _ := filepath.Split(file)
	err = unzipFile(file, dest)
	if err != nil {
		log.Printf("unzip file [%s] failed, err: %v", constant.FILE_LOCATION+"test.zip", err)
		return err
	}
	defer removeUnzipFile(testFile)

	cmd := exec.Command("python", testFile, platformName, platformVersion, deviceName, apkName, imgLocation)
	stdout, err := cmd.StdoutPipe()
	if err != nil {
		log.Printf("stdout failed, err: %v", err)
		return err
	}
	defer stdout.Close()

	err = cmd.Run()
	if err != nil {
		log.Printf("run cmd failed, err: %v", err)
		return err
	}

	_, err = ioutil.ReadAll(stdout)
	if err != nil {
		if err.Error() == "read |0: file already closed" {
			log.Println("exec success")
			return nil
		}
		log.Printf("read out failed, err: %v", err)
		return err
	}
	log.Println("exec success")

	return nil
}

// 删除文件
func removeUnzipFile(file string) {
	dest, _ := filepath.Split(file)
	err := os.RemoveAll(dest)
	if err != nil {
		log.Printf("remove dir [%s] failed, err: %v", dest, err)
		return
	}
	log.Printf("remove dir [%s] success", dest)
	return
}

// 解压文件
func unzipFile(zipFile, dest string) (err error) {
	//目标文件夹不存在则创建
	_, err = os.Stat(dest)
	if err != nil {
		if os.IsNotExist(err) {
			err = os.MkdirAll(dest, 0755)
			if err != nil {
				log.Printf("create dir [%s] failed, err: %v", dest, err)
				return err
			}
		}
	}

	reader, err := zip.OpenReader(zipFile)
	if err != nil {
		log.Printf("open file [%s] failed, err: %v", zipFile, err)
		return err
	}

	defer reader.Close()

	for _, file := range reader.File {
		if file.FileInfo().IsDir() {
			err = os.MkdirAll(dest+string(os.PathSeparator)+file.Name, 0755)
			if err != nil {
				log.Printf("create dir [%s] failed, err: %v", dest+string(os.PathSeparator)+file.Name, err)
				return err
			}
			continue
		} else {
			rc, err := file.Open()
			if err != nil {
				log.Printf("open file [%s] failed, err: %v", file.Name, err)
				return err
			}

			filename := dest + "/" + file.Name

			w, err := os.Create(filename)
			if err != nil {
				log.Printf("create file [%s] failed, err: %v", filename, err)
				rc.Close()
				return err
			}

			_, err = io.Copy(w, rc)
			if err != nil {
				log.Printf("copy file [%s] failed, err: %v", filename, err)
				w.Close()
				rc.Close()
				return err
			}
			w.Close()
			rc.Close()
		}
	}
	return
}
