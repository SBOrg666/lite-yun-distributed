# Lite Yun 使用说明

## server端启动参数
- -u：登录用户名（必需）
- -p：登录用户密码（必需）
- -port：端口号（默认8000）
- -l：服务器日志文件路径（默认/var/log/pacman.log）

---
## API接口说明
- /login POST  
    **请求参数：**  
    username=XXX  
    password=XXX  
    **返回结果：**  
    成功返回“ok”，并且response的header中有set-cookie字段，里面有个session字段，后面每一次请求都需要带有这个cookie  
    失败返回“failed”

- /systemInfo POST  
    **请求参数：**  
    无  
    **返回结果：**  
    ```
    {"cpu_info":[11.392486124946615,11.617307213709251,11.289336782556795,11.474873603595276],
    "disk_info":[["/","11804180480","15786254336"],["/home","12733849600","95640199168"],["/boot","74871808","100663296"]],"log_info":"",
    "mem_info":["3821301760","8267571200"],
    "network_info":[["12","1","2","3","4","5"],["0","0","0","0","0","12078919"],["0","0","0","0","0","80095365"],["0","0","0","0","0","92174284"]],
    "swap_info":["0","8128102400"],"sys_info":["Dell-XPS-13-9343","arch","4.16.6-1-ARCH","Intel(R) Core(TM) i7-5500U CPU @ 2.40GHz","3000.0 Mhz","8267571200","2018-05-06 13:25:15 +0800 CST"]}
    ```

- /systemInfo websocket  
    同上

- /getProcessInfo POST  
    **请求参数：**  
    无  
    **返回结果：**  
    ```
    {"ProcessInfo":[{"cpu_percent":"0.004972614070283306","create_time":"1525584315000","exe":"","memory_percent":"0.107260525","name":"systemd","pid":"1","status":"Sleep","username":"root"},
    {"cpu_percent":"0","create_time":"1525584315000","exe":"","memory_percent":"0","name":"kthreadd","pid":"2","status":"Sleep","username":"root"},
    {"cpu_percent":"0","create_time":"1525584315000","exe":"","memory_percent":"0","name":"kworker/0:0H","pid":"4","status":"Idle","username":"root"}}
    ```

- /processInfo websocket  
    同上

- /manageProcess POST  
    **请求参数：**  
    pid=pid  
    operation=resume|suspend|kill|terminate  
    createTime=createTime  
    **返回结果：**  
    成功返回“pid succeed”，失败返回错误原因

- path GET  
    **请求参数：**  
    path=path  
    **返回结果：**  
    ```
    {"dirs":[{"Url":"/path?path=/home/zhang/Desktop/New Folder","DirName":"New Folder","Permission":"drwxr-xr-x","Size":"4096","Owner":"zhang","Group":"zhang","Mtime":"2018-05-06 16:45:31.185591279 +0800 CST","Access":true}],
    "files":[{"FileName":".directory","Permission":"-rw-r--r--","Size":"65","Owner":"zhang","Group":"zhang","Mtime":"2018-01-12 19:20:44.765814395 +0800 CST","Access":true},{"FileName":"Home.desktop","Permission":"-rw-r--r--","Size":"2401","Owner":"zhang","Group":"zhang","Mtime":"2018-01-12 19:20:44.765814395 +0800 CST","Access":true},{"FileName":"README.md","Permission":"-rw-r--r--","Size":"2","Owner":"zhang","Group":"zhang","Mtime":"2018-05-06 15:56:05.730774909 +0800 CST","Access":true},{"FileName":"Text File","Permission":"-rw-r--r--","Size":"2","Owner":"zhang","Group":"zhang","Mtime":"2018-05-06 14:39:53.023221685 +0800 CST","Access":true},{"FileName":"trash.desktop","Permission":"-rw-r--r--","Size":"2820","Owner":"zhang","Group":"zhang","Mtime":"2018-01-12 19:20:44.765814395 +0800 CST","Access":true}],
    "path":"/home/zhang/Desktop",
    "writable":true}
    ```

- /download POST  
    **请求参数：**  
    files=[XXX,XXX,XXX]  
    **返回结果：**  
    成功返回“ok zip文件路径”，失败返回错误原因

- /download GET  
    **请求参数：**  
    name=zip文件路径  
    **返回结果：**  
    成功返回文件，失败返回错误原因

- /upload POST  
    **请求参数：**  
    query参数：  
    path=path（上传路径）  
    post参数：  
    file=file  
    **返回参数：**  
    成功返回{"name":filename}，失败返回错误原因

- /delete DELETE  
    **请求参数：**  
    files=[XXX,XXX,XXX]  
    **返回结果：**  
    成功返回“ok”，失败返回错误原因
