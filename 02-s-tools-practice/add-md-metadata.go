package main

import (
	"bufio"
	"fmt"
	"io"
	"io/ioutil"
	"os"
	"path"
	"path/filepath"
	"strings"
	"syscall"
	"time"
)

// 递归获取指定目录下的所有文件名
func GetAllFile(pathname string) ([]string, error) {
	result := []string{}

	fis, err := ioutil.ReadDir(pathname)
	if err != nil {
		fmt.Printf("读取文件目录失败，pathname=%v, err=%v \n", pathname, err)
		return result, err
	}

	// 所有文件/文件夹
	for _, fi := range fis {
		fullname := pathname + "/" + fi.Name()
		//获取文件后缀
		fileSuffix := path.Ext(fi.Name())

		// 是文件夹则递归进入获取;是文件，则压入数组
		if fi.IsDir() {
			temp, err := GetAllFile(fullname)
			if err != nil {
				fmt.Printf("读取文件目录失败,fullname=%v, err=%v", fullname, err)
				return result, err
			}
			result = append(result, temp...)
		} else if fileSuffix == ".md" {
			result = append(result, fullname)
		}

	}

	return result, nil
}

// 把秒级的时间戳转为time格式
func SecondToTime(sec int64) time.Time {
	return time.Unix(sec, 0)
}

func fileInsertInfo(path_dir string, filename string, instert_data []string) {
	// 打开要操作的文件 os.O_RDWR: 可读可写
	p := path_dir + "/" + filename
	temp_path := path_dir + "/" + "temp.md"
	file, err := os.OpenFile(p, os.O_RDWR, 0544)
	if err != nil {
		fmt.Printf("File open failed! err: %v\n", err)
		return
	}
	reader := bufio.NewReader(file)
	line, err := reader.ReadString('\n') // 依次读一行
	if err != nil {
		fmt.Printf("Read file failed! err: %v\n", err)
		return
	}
	for line == "" {
		line, err = reader.ReadString('\n')
	}
	if strings.Contains(line, "---") {
		file.Close()
		return
	}
	// 新建临时文件
	tempFile, err := os.OpenFile(temp_path, os.O_CREATE|os.O_TRUNC|os.O_WRONLY, 0644)
	if err != nil {
		fmt.Printf("Temp create failed! err: %v\n", err)
		return
	}
	writer := bufio.NewWriter(tempFile)
	_ = writer.Flush()
	// 在第二行插入新内容
	// 将原文件写入临时文件
	for _, s := range instert_data {
		_, _ = writer.WriteString(s)
	}
	_ = writer.Flush()

	// 写入临时文件
	_, _ = writer.WriteString(line)
	_ = writer.Flush()
	// 移动光标
	//_, _ = file.Seek(int64(len(line)), 0)
	// 写入要插入的内容
	// var tempInfo string
	// tempInfo = "hi world!\n"
	// _, _ = tempFile.WriteString(tempInfo)
	// 把源文件的后续内容写入临时文件
	for {
		line, err := reader.ReadString('\n') // 依次读一行
		_, _ = writer.WriteString(line)
		if err == io.EOF {
			break
		}

		if err != nil {
			fmt.Printf("File raed failed! err: %v\n", err)
			return
		}

	}
	_ = writer.Flush()

	file.Close()
	tempFile.Close()
	err = os.Rename(temp_path, p)
	if err != nil {
		fmt.Printf("Rename file raed failed! err: %v\n", err)
		return
	}
	fmt.Println(p, "add info success.")
}

func get_tags(file_path string) (res []string) {
	tags := []string{"python", "js", "git", "java", "go", "html", "css", "svn", "c++", "c#", "vba", "github"}
	content, err := ioutil.ReadFile(file_path)
	if err != nil {
		panic(err)
	}
	content_s := strings.ToLower(string(content))
	// fmt.Println(content_s)
	for _, v := range tags {
		if strings.Contains(content_s, strings.ToLower(v)) {
			res = append(res, v)
		}
	}
	return
}

func main() {
	directory, err := os.Getwd() //get the current directory using the built-in function
	if err != nil {
		fmt.Println(err) //print the error if obtained
	}
	args := os.Args
	if len(args) < 2 {
		fmt.Println("please input file path ...")
		return
	}
	// 递归获取目录下的所有文件
	p := directory + "/" + args[1]
	fmt.Println(p)
	var files []string
	files, _ = GetAllFile(p)

	fmt.Println("目录下的所有文件如下")
	for i := 0; i < len(files); i++ {
		fmt.Println("文件名：", files[i])

		// 获取文件原来的访问时间，修改时间
		finfo, _ := os.Stat(files[i])

		// linux环境下代码如下
		//linuxFileAttr := finfo.Sys().(*syscall.Stat_t)
		//fmt.Println("文件创建时间", SecondToTime(linuxFileAttr.Ctim.Sec))
		//fmt.Println("最后访问时间", SecondToTime(linuxFileAttr.Atim.Sec))
		//fmt.Println("最后修改时间", SecondToTime(linuxFileAttr.Mtim.Sec))

		// windows下代码如下
		winFileAttr := finfo.Sys().(*syscall.Win32FileAttributeData)
		fmt.Println("文件创建时间：", SecondToTime(winFileAttr.CreationTime.Nanoseconds()/1e9).Format("2006-01-02 15:04:05"))
		fmt.Println("最后访问时间：", SecondToTime(winFileAttr.LastAccessTime.Nanoseconds()/1e9))
		fmt.Println("最后修改时间：", SecondToTime(winFileAttr.LastWriteTime.Nanoseconds()/1e9))

		date := SecondToTime(winFileAttr.CreationTime.Nanoseconds() / 1e9).Format("2006-01-02 15:04:05")
		title := strings.TrimSuffix(path.Base(files[i]), ".md")
		folders := strings.Split(files[i], "/")
		fmt.Println(folders)
		category := make([]string, 0, 1)
		flag_is_start := false
		last_folder_name := strings.Split(args[1], "/")[len(strings.Split(args[1], "/"))-1]
		for _, v := range folders {
			if flag_is_start && !strings.Contains(v, ".md") {
				category = append(category, v)
			}
			if v == last_folder_name {
				flag_is_start = true
			}
		}
		tags := get_tags(files[i])
		fmt.Println(tags)
		fmt.Println(category)
		fmt.Println(date, title)

		info := make([]string, 0, 1)
		info = append(info, "---\n")
		info = append(info, "title: "+title+"\n")
		info = append(info, "date: "+date+"\n")
		for i, v := range category {
			if i == 0 {
				info = append(info, "category: \n")
			}
			info = append(info, "  - "+v+"\n")
		}
		for i, v := range tags {
			if i == 0 {
				info = append(info, "tag: \n")
			}
			info = append(info, "  - "+v+"\n")
		}
		info = append(info, "---\n")
		fmt.Println(info)
		fileInsertInfo(filepath.Dir(files[i]), path.Base(files[i]), info)
		fmt.Println("done")
	}
}
