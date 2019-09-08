# 部署私有云过程中遇到的问题详解

> 刚写完一篇 [在Window 10 PC 作为主机（HOST）对外租用虚拟机的一种云桌面的设置](https://blog.csdn.net/WeiXiaoAssassin/article/details/100612612)([也可在Github上查看](https://github.com/wywwwwei/ServiceComputingOnCloud/blob/master/HW1_CloudDesktop/Report.md))，在实验的过程中，或多或少会遇到一些奇怪的问题，在这里会进行一个详细的描述/解决

[TOC]

## 在安装CentOS过程中无法看到鼠标

- 问题详情

  > 在安装CentOS的过程中（安装向导），以及后面安装了GNOME Desktop之后，都看不到鼠标光标，而且点击的位置与在虚拟机显示界面所看到的位置并不是一致的

- 解决方法

  ![1567923430742](https://raw.githubusercontent.com/wywwwwei/ServiceComputingOnCloud/master/HW1_CloudDesktop/pic/1567923430742.png)

## 获取wget `yum install wget`报错

- 问题详情

  ![1567840339963](https://raw.githubusercontent.com/wywwwwei/ServiceComputingOnCloud/master/HW1_CloudDesktop/pic/1567840339963.png)

- 查找原因

  在找原因之前，先来看看CentOS的网络管理

  > 在CentOS7中默认使用NetworkManager守护进程来监控和管理网络设置。nmcli是命令行的NetworkManager工具，会自动把配置写到/etc/sysconfig/network-scripts/目录下面。
  >
  > CentOS7之前的网络管理是通过ifcfg文件配置管理接口(device)，而现在是通过NetworkManager服务管理连接(connection)。一个接口(device)可以有多个连接(connection)，但是同时只允许一个连接(connection)处于激活（active）状态。

  我们可以先通过 `nmcli d` / `nmcli dev status`命令查看当前网络设备连接状态

  ![1567925539481](https://raw.githubusercontent.com/wywwwwei/ServiceComputingOnCloud/master/HW1_CloudDesktop/pic/1567925539481.png)

  可以看到当前所有的网络设备都尚未连接

  在通过 `ip addr` 命令可以看到并没有 inet(ipv4)/inet6(ipv6)地址，意味着当前系统是无法上网的

  ![1567925893349](https://raw.githubusercontent.com/wywwwwei/ServiceComputingOnCloud/master/HW1_CloudDesktop/pic/1567925893349.png)

  其实是因为CentOS默认情况下不启用以太网接口的，我们需要手动去启动或设置开机启动

- 解决方案

  1. 手动启动（未尝试过，建议使用第二种方法）
  
     ```shell
     ifup {interface}  #启动网卡 interface 指网卡ID,如enp0s3/enp0s8
   ifdown {interface} #禁用网卡
     ```
  
     > ifup/ifdown是script，它会直接到 /etc/  sysconfig/network-scripts目录下搜索对应的配置文件，例如ifup  enp0s8，它会找出ifcfg-enp0s8这个文件的内容，然后加以设置。
     >
     > 不过，由于这两个程序主要是搜索设置文件（ifcfg-ethx）来进行启动与关闭的，所以在**使用前请确定ifcfg-xxx是否真的存在于正确的目录内**，否则会启动失败。另外，如果以ifconfig xxx来设置或者是修改了网络接口后，就无法再以ifdown  xxx的方式来关闭了。因为ifdown会分析比较目前的网络参数与ifcfg-xxx是否相符，不符的话，就会放弃这次操作。因此，**使用 ifconfig修改完毕后，应该要以ifconfig xxx down才能够关闭该接口**。

  2. 设置开机启动

     - GUI模式

       1. 命令行输入 `nmtui` ![1567928752026](https://raw.githubusercontent.com/wywwwwei/ServiceComputingOnCloud/master/HW1_CloudDesktop/pic/1567928752026.png)

       2. 选择Edit a connection，可以看到所有的网卡设备

       ![1567929010904](https://raw.githubusercontent.com/wywwwwei/ServiceComputingOnCloud/master/HW1_CloudDesktop/pic/1567929010904.png)

       3. 选中Edit，进入编辑界面![1567929097663](https://raw.githubusercontent.com/wywwwwei/ServiceComputingOnCloud/master/HW1_CloudDesktop/pic/1567929097663.png)

       4. 

          Option|DHCP|StaticIP
      -|-|-
       IPv4 Configuration|Automatic|Manual
	      Other|![1567929560234](https://raw.githubusercontent.com/wywwwwei/ServiceComputingOnCloud/master/HW1_CloudDesktop/pic/1567929560234.png)|![1567929695208](https://raw.githubusercontent.com/wywwwwei/ServiceComputingOnCloud/master/HW1_CloudDesktop/pic/1567929695208.png)
	   
		  5. 重启网络服务
  	
	        ```shell
	        systemctl restart network
	        ```
	   
	        
	
- 命令行模式(参考[Linux下网络配置、查看ip地址、网关信息，DNS信息(以centos7为例)](https://blog.csdn.net/qq_15304853/article/details/78700197))
  
  > **CentOS7网络配置相关文件：**
  >
  > ```shell
  > /etc/resolv.conf             		# DNS配置文件
  > /etc/hosts                      	#主机名到IP地址的映射 ,不改主机名基本不会动他。
  > /etc/sysconfig/network           	#所有的网络接口和路由信息，网关只有最后一个有效。
  > /etc/sysconfig/network-script/ifcfg-<interface-name>      #每一个网络接口的配置信息
  > ```
  > **每一个网卡只能使用一个配置文件**，当有多个配置文件时，后面读取的配置文件信息会覆盖前面的配置信息。所以，一个网卡最好只写一个配置文件。或者之设置一个文件开机自启动，同时/etc/sysconfig/network-script/ifcfg-<interface-name>中不要写网关信息，交给/etc/sysconfig/network来配置。
  
  CentOS 7默认网络接口文件存放于 /etc/sysconfig/network-scripts/ 目录下，一般情况下配置文件默认是：ifcfg-网卡名（由于CentOS的发行及系统升级或许可能会存在网络接口名称与之前版本不一致的情况）
  
  ![1567930177403](https://raw.githubusercontent.com/wywwwwei/ServiceComputingOnCloud/master/HW1_CloudDesktop/pic/1567930177403.png)
  
  上图中配置文件就是	enp0s3对应ifcfg-enp0s3   enp0s8对应ifcfg-enp0s3
  
  1. 用vi打开对应配置文件（按`I`进入插入/编辑模式，按`Esc`回到命令模式，在命令模式下输入`:wq`保存并退出）
  
  2.  将ONBOOT=no改为ONBOOT=yes设置为开机启动
  
     ![1567930395893](https://raw.githubusercontent.com/wywwwwei/ServiceComputingOnCloud/master/HW1_CloudDesktop/pic/1567930395893.png)
     
     3.对两个文件都做该配置后，通过下面的命令重启网络服务
     
     ```shell
     service network restart
     ```
     
     但是执行该命令时可能会出现一些奇怪的错误，可通过打开网卡服务器的DHCP功能（如果不打开第二张网卡无法联通），多次重启网络服务直到成功即可
     
     ![1567930790857](https://raw.githubusercontent.com/wywwwwei/ServiceComputingOnCloud/master/HW1_CloudDesktop/pic/1567930790857.png)
     
     ![1567930828789](https://raw.githubusercontent.com/wywwwwei/ServiceComputingOnCloud/master/HW1_CloudDesktop/pic/1567930828789.png)
     
     ![1567930990300](https://raw.githubusercontent.com/wywwwwei/ServiceComputingOnCloud/master/HW1_CloudDesktop/pic/1567930990300.png)