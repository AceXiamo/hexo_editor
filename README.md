# hexo_editor
<img src="https://xiamo.oss-accelerate.aliyuncs.com/xiamo/WordPress/2021/12/20211220154336504.png" />
因为编辑Hexo博客的方式实在麻烦，并且后台插件不是很好看，所以决定自己写，于是也就开始了这个项目，顺便在其中学习一下Golang 

### 预览
测试博客地址： 🔗 [Click here](http://demo.hexo.xiamoqwq.com)

博客管理端地址： 🔗 [Click here](http://admin.hexo.xiamoqwq.com)

账号：xiamo
密码：xiamo

### 大概会有的功能
- 博客编辑 (一些基础功能，暂时已经完成了)
- 表情 (这里有markdown-emoji和表情图)
- 内容管理 (存放一些静态资源，以及在博客中使用这些静态资源)
- 发布说说
- css/js配置 (因为我的博客中用到了大量的css和一部分js，所以可以配置一个全局css/js，或者单独导入部分css/js到某个博客中)

### 关于启动方式
首次运行项目会产生一个配置文件，补充完整后重新启动，再运行Vue即可

### 配置文件中的内容
`port`  后端运行端口

`hexo_root`  hexo的根目录（此处用于限制可访问路径，比如说填写：'/data/hexo'，那么就只能访问/编辑该目录下的文件

`username`  登陆用户名

`password`  登陆密码

### 前端传送门
🔗 [Click here](https://github.com/xm17906193/hexo_editor_vue)

### 我的博客
🔗 [Click here](https://qwq.link)

### 关于说说
说说的数据来自我的另一个项目 "树" ,各位如果想要加上这么一个小模块可以到博客编辑器中自取

说说在这个目录下： `/data/hexo/source/shuoshuo/`

###### 数据来源
如果你有已经完成的说说可以直接修改js代码中的接口地址

如果没有的话则可以到 `https://tree.qwq.link` 中发布一条说说，然后从接口返回中取到你的userId

然后copy目录下index.md再将js代码中的userId替换掉即可💦