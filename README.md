# XC 自动推播插件

## 使用教学
### 前置条件

1. 登录 [PushPlus推送加](http://www.pushplus.plus/push1.html)
2. 获取 `token`
3. 配置 `config.yml`
4. 在 `Robot.exe` 的目录下,打开 `PowerShell`. 运行 `.\Robot.exe config.yml` 即可


### config.yml 详解
```yaml
token: #PushPlus的Token
path: #XC的AutoLog地址,需要绝对路径,【Ex:C:\SXC\Data\AutoLog\C4E3B0D4C0B4BFA9】

ignore: # 填写不想要提醒的副本名称 
  - 副本名称1
  - 副本名称2
```

## 效果展示
- 粗略图<br/>
  ![img.png](README/img.png)
- 详细图<br/>
  ![img.png](README/img1.png)



