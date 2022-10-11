# juejin-auto
> 稀土掘金自动化任务，例如签到等功能。

# features

- [x] 掘金自动化签到
- [x] 钉钉消息推送（目前只支持内置单模板）
- [ ] 邮箱消息推送

# use
1. 首先需要将项目 `fork` 到自己的 `github` 仓库中
2. 进入项目主页点击 `Settings` > `Security` > `Secrets` > `Actions`  配置信息
  - 配置掘金 `cookie`: key 配置为 `JUEJIN_COOKIE`，value 填入从掘金网站获取的cookie
  - 配置钉钉机器人 `token`： key 配置为 `DINGTALK_BOT_TOKEN`，value 填入从钉钉机器人 token
3. 定时任务配置了每天9点运行，但是由于 github actions 并不准时，大概在12点左右才能执行

    
     
