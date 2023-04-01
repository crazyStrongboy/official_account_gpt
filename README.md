# ChatGPT 微信公众号搭建
本项目旨在通过不到 100 行的 Golang 代码搭建一个微信公众号 ChatGPT，实现基本的智能聊天功能。

## 前置准备

- ChatGpt api token  具体可参考 [地址](https://platform.openai.com/account/api-keys)
- 微信公众号开发者账号  具体参考 [地址](https://mp.weixin.qq.com/)
- 服务器，必须有一个公网服务器部署服务

## 步骤

### 1. 一键部署


### 2. 运行命令

运行代码。使用以下命令运行代码：

./offical_account_gpt  --wx_app_id=xx --wx_ori_id=yy --wx_token=zz --wx_aes_key=kk --token=tt


其中，`wx_app_id`、`wx_ori_id`、`wx_token`、`wx_aes_key` 分别为微信公众号的开发者ID、原始ID、令牌和消息加解密密钥，`token` 为 ChatGpt API token。

### 3. 配置微信公众号

在微信公众平台中，进入“设置与开发/基本配置”页面，启用服务器配置：

- URL：`http://yourserver.com/wechat`
- Token：`your_wechat_token`
- 消息加解密模式：安全模式

将 `yourserver.com` 替换为你的服务器域名或 IP 地址，将 `your_wechat_token` 替换为你自己定义的 Token。

### 4. 测试

现在，你可以在微信公众号中向 ChatGPT 发送消息，ChatGPT 将会自动回复你的消息。

## 总结

通过不到 100 行的 Golang 代码，我们成功地搭建了一个微信公众号 ChatGPT，实现了基本的智能聊天功能。当然，这只是一个简单的示例，你可以根据自己的需求扩展功能，让 ChatGPT 变得更加强大。]()