go 代理配置 安装工具会用到

1. 右键 我的电脑 -> 属性 -> 高级系统设置 -> 环境变量
2. 在 “[你的用户名]的用户变量” 中点击 ”新建“ 按钮
3. 在 “变量名” 输入框并新增 “GOPROXY”
4. 在对应的 “变量值” 输入框中新增 “https://proxy.golang.com.cn,direct”
5. 最后点击 “确定” 按钮保存设置

package 名字和文件夹名不一致时 引用需要注意


![图片 alt](https://github.com/yungphilo/unknow_name/blob/office/demo1/import-package.png)


git config --global user.name "你的GitHub用户名"
git config --global user.email "你的GitHub注册邮箱"
然后生成ssh密钥文件

ssh-keygen -t rsa -C "你的GitHub注册邮箱"
然后三个回车就可以了。

(adsbygoogle = window.adsbygoogle || []).push({});

然后在C:\Users\admin.ssh文件夹中找到id_rsa.pub,将里面所有的内容全部复制出来。
在这个网址https://github.com/settings/keys点New SSH key

Title随便输入，Key粘贴刚才复制的key,然后点击Add SSH key

验证一下是否连接成功

在Git Bash中输入

ssh git@github.com
如果显示下面这样就是成功了

![图片alt](https://pic2.zhimg.com/80/v2-87ec0cf43e93d2a01aac117d4f0ad2dd_720w.webp "图片title")

***markdown 本地图片不支持绝对路径***

![图片alt](https://github.com/yungphilo/unknow_name/blob/office/demo1/gitclone.png)

***-b 是选择分支 默认是master分支，github.com后面是 ":" 不是 "/"***
