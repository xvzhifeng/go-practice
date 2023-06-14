---
title: git代码统计
date: 2023-06-12 01:11:07
catagery: 
  - git
tag: 
  - python
  - git
  - github
---

## 命令行

### 查看git上的个人代码量：

```shell
git log --author="username" --pretty=tformat: --numstat | awk '{ add += $1; subs += $2; loc += $1 - $2 } END { printf "added lines: %s, removed lines: %s, total lines: %s\n", add, subs, loc }' -
```

结果示例：(记得修改 username)

```apache
added lines: 120745, removed lines: 71738, total lines: 49007
```

### 统计每个人增删行数

```shell
git log --format='%aN' | sort -u | while read name; do echo -en "$name\t"; git log --author="$name" --pretty=tformat: --numstat | awk '{ add += $1; subs += $2; loc += $1 - $2 } END { printf "added lines: %s, removed lines: %s, total lines: %s\n", add, subs, loc }' -; done
```

结果示例

```sh
xuzhif	added lines: 4323, removed lines: 164, total lines: 4159
苏木	added lines: 274, removed lines: 0, total lines: 274
徐志锋	added lines: 75592, removed lines: 104, total lines: 75488
```

### 查看仓库提交者排名前 5

```bash
git log --pretty='%aN' | sort | uniq -c | sort -k1 -n -r | head -n 5
```

### 贡献值统计

```bash
git log --pretty='%aN' | sort -u | wc -l
```

### 提交数统计

```1c
git log --oneline | wc -l
```

### 添加或修改的代码行数：

```nginx
git log --stat|perl -ne 'END { print $c } $c += $1 if /(\d+) insertions/'
```

## 使用gitstats

[GitStats项目](https://link.segmentfault.com/?enc=d222BALJB10GykwUzDRSjw%3D%3D.QdabMZdVXgkMJO1EraIGV90nZKF6pmbb%2Fn1dy7um7U3pdR0lwWDfQUTOd4y4nTdF)，用Python开发的一个工具，通过封装Git命令来实现统计出来代码情况并且生成可浏览的网页。官方文档可以参考这里。

### 使用方法

```awk
git clone git://github.com/hoxu/gitstats.git
cd gitstats
./gitstats 你的项目的位置 生成统计的文件夹位置
```

可能会提示没有安装gnuplot画图程序，那么需要安装再执行：

```awk
//mac osx
brew install gnuplot
//centos linux
yum install gnuplot
```

## 使用cloc

```cmake
npm install -g cloc
// 运行一下命令
cloc .
```

