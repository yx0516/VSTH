在源码 github.com/mattn/go-sqlite3/sqlite3.go 里添加编译选项，开启外键级联删除
#cgo CFLAGS: -DSQLITE_DEFAULT_FOREIGN_KEYS=1

默认关闭的


### 时区不对，少 8 小时
### github.com/mattn/go-sqlite3
# 修改插入数据时
sqlite3.go  的 807 行  .UTC() 替换成 .Local()
807             b := []byte(v.Local().Format(SQLiteTimestampFormats[0]))

# 读取数据时
# 912 和 957 行  time.UTC 替换成 time.Local

912  epoch := time.Date(1970, 1, 1, 0, 0, 0, 0, time.Local)

957                 if timeVal, err = time.ParseInLocation(format, s, time.Local); err == nil {
