# Subscription Center

SubCenter是一个集成各种任务并进行实时推送的中间件，本身不提供数据与推送服务。

> NOTE: 项目开发初期，欢迎提出新意见


### 项目依赖

#### 推送服务

由于推送的设备要求与维护成本较高，请自行前往相关网站注册，如需接入可发Issue

- [x] [Server酱-Turbo版 (微信)](https://sct.ftqq.com/)
- [ ] [PushDeer (IOS)](https://github.com/easychen/pushdeer)
- [ ] ...

#### 数据服务

- [ ] 数字货币：[币安](https://www.binance.com/)
- [ ] 钱包变动：[BSC](https://github.com/binance-chain/bsc)
- [ ] A股价格：[新浪](https://finance.sina.com.cn/stock/)
- [ ] M股价格：[富途](https://www.futunn.com/)
- [ ] ...

### 通知种类

- [ ] 定时通知
- [ ] 事件通知
- [ ] 价格变动通知
- [ ] ...

### 使用说明

#### 环境搭建

```bash
git clone https://github.com/ThreeCatsLoveFish/SubCenter.git
make
./output/subcenter
```

#### 推送配置

1. `config/push.toml` 文件中增加需要推送的endpoint和对应token(key)
2. `config/task.toml` 文件中增加需要订阅的内容和发送频率
