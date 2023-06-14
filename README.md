# go-executor
对 jar 包进行管理，(相当于 ```java -jar xxx.jar```)

### 使用
启动时扫描所有同级的 jar 包并加入启动。

### 日志
所有 jar 包的控制台输出 (stdout) 都会保存到 logs/ 下，命名为 jar包名-yyyy-MM-dd.log。

### 命令
- ```help``` 查看帮助
- ```pid``` 列出所有管理的 jar 包名和对应进程的 pid
- ```kill pid``` 结束一个进程 (仅限管理的jar包)
- ```run xxx.jar``` 加入管理并运行一个 jar 包