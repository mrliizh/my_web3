

# Linux网络

重启网络

sudo systemctl restart NetworkManager



```shell
ip addr 
//新命令，更强大，优于ifconfig

mrliizh@ubuntu:~$ ip addr
1: lo: <LOOPBACK,UP,LOWER_UP> mtu 65536 qdisc noqueue state UNKNOWN group default qlen 1000
    link/loopback 00:00:00:00:00:00 brd 00:00:00:00:00:00
    inet 127.0.0.1/8 scope host lo
       valid_lft forever preferred_lft forever
    inet6 ::1/128 scope host 
       valid_lft forever preferred_lft forever
2: ens33: <BROADCAST,MULTICAST,UP,LOWER_UP> mtu 1500 qdisc fq_codel state UP group default qlen 1000
    link/ether 00:0c:29:72:a2:7a brd ff:ff:ff:ff:ff:ff
    altname enp2s1
    inet 192.168.89.195/24 brd 192.168.89.255 scope global dynamic noprefixroute ens33
       valid_lft 114sec preferred_lft 114sec
    inet6 2409:895a:34a9:b586:4ba1:2b62:bb25:8717/64 scope global temporary dynamic 
       valid_lft 7138sec preferred_lft 7138sec
    inet6 2409:895a:34a9:b586:40f7:165a:ae54:ff08/64 scope global dynamic mngtmpaddr noprefixroute 
       valid_lft 7138sec preferred_lft 7138sec
    inet6 fe80::3372:73c9:5519:961d/64 scope link noprefixroute 
       valid_lft forever preferred_lft forever
3: br-0a864fda9cf4: <NO-CARRIER,BROADCAST,MULTICAST,UP> mtu 1500 qdisc noqueue state DOWN group default 
    link/ether 02:42:ea:17:6a:bc brd ff:ff:ff:ff:ff:ff
    inet 172.18.0.1/16 brd 172.18.255.255 scope global br-0a864fda9cf4
       valid_lft forever preferred_lft forever
4: docker0: <NO-CARRIER,BROADCAST,MULTICAST,UP> mtu 1500 qdisc noqueue state DOWN group default 
    link/ether 02:42:e0:af:74:68 brd ff:ff:ff:ff:ff:ff
    inet 172.17.0.1/16 brd 172.17.255.255 scope global docker0
       valid_lft forever preferred_lft forever



```



```shell
ssh-copy-id


//ssh-copy-id 会将本地的公钥（~/.ssh/id_rsa.pub）复制到远程服务器的 //~/.ssh/authorized_keys 文件中。
```





```shell
netstat  -tuln | grep 80  //查看80端口情况
lsof -i 80
ntlp
```



# shell

```shell
ls -l

ls -a  显示包含隐形文件

ll //等于ls-l
```



```shell
pwd    // print work directory 打印当前目录

cat  //

head  文件名

head --lines=2  文件名

tail  --lines=3  文件名

less 文件名

q退出


file 文件名

echo 文件名

a=123
echo $a    //打印变量a
echo ${a}  //一样的{}方便区分
```

```shell
while :; do ps axj | head -1&&ps axj | grep mytest|grep -v grep; sleep 1;echo "--------"; done
```



```shell
ps -elLf | (head -1 && grep nginx) | grep -v grep | awk '{print NR, $0}'
```

pushd

popd



# powershell

win + x

```cmd
start test.ps1     ##默认打开方式打开文件
&"D:\Program Files (x86)\Notepad++\notepad++.exe" test.ps1  ##以某个可执行程序打开文件,需要加&
cd ~
./test.ps1 #运行
```



# cmd

```cmd
start test.ps1
"D:\Program Files (x86)\Notepad++\notepad++.exe" test.ps1  ##以某个可执行程序打开文件，不需要&

echo %homepath%

test.ps1 #运行,不需要./
```





### windows文件操作

windows中家目录是%userprofile%，相当于Linux中的~

cmd操作

cd  c:

cd  %userprofile%



文件资源管理器操作

双击我的电脑    地址栏中输入%userprofile%，再回车

测试SSH连接：在终端中运行 ssh -T git@github.com。如果配置正确，会看到一条欢迎信息；如果有错误，终端会显示相关信息



windows指令  Linux

dir  =  ls

type 指令 = cat



notepad 加文件名 ，以notepad应用打开文件



刷新dns服务器

ipconfig /flushdns











mrliizh@ubuntu:~$ ifconfig ens33
ens33: flags=4163<UP,BROADCAST,RUNNING,MULTICAST>  mtu 1500
        inet 192.168.89.195  netmask 255.255.255.0  broadcast 192.168.89.255
        inet6 2409:895a:34a9:b586:40f7:165a:ae54:ff08  prefixlen 64  scopeid 0x0<global>
        inet6 2409:895a:34a9:b586:4ba1:2b62:bb25:8717  prefixlen 64  scopeid 0x0<global>
        inet6 fe80::3372:73c9:5519:961d  prefixlen 64  scopeid 0x20<link>
        ether 00:0c:29:72:a2:7a  txqueuelen 1000  (Ethernet)
        RX packets 52250  bytes 55503555 (55.5 MB)
        RX errors 0  dropped 0  overruns 0  frame 0
        TX packets 9878  bytes 1130893 (1.1 MB)
        TX errors 0  dropped 0 overruns 0  carrier 0  collisions 0



192.168.89.195

mrliizh@ubuntu:~/.ssh$ ps aux |grep sshd
root        1029  0.0  0.1  12196  7408 ?        Ss   Nov07   0:00 sshd: /usr/sbin/sshd -D [listener] 0 of 10-100 startups
root        4612  0.0  0.2  13744  8868 ?        Ss   01:09   0:00 sshd: mrliizh [priv]
mrliizh     4665  0.0  0.1  13752  5300 ?        S    01:09   0:00 sshd: mrliizh@pts/1
mrliizh     4702  0.0  0.0   9040   720 pts/1    S+   01:31   0:00 grep --color=auto sshd



git pull 冲突了后