# go-zero 学习笔记



## goctl 使用

### 基础命令

- goctl --help // 查看goctl命令帮助

### 生成api模块
```bash
goctl api new modelname
```

### 生成model命令

```bash
goctl model mysql datasource -url="root:123456@tcp(127.0.0.1:3306)/zebblog" -table="sys_user" -dir ./model -c
```




