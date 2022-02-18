# Subscription Center

SubCenter —— 集成各种任务并进行实时推送的中间件。

![GitHub go.mod Go version](https://img.shields.io/github/go-mod/go-version/ThreeCatsLoveFish/SubCenter)
[![wakatime](https://wakatime.com/badge/github/ThreeCatsLoveFish/SubCenter.svg)](https://wakatime.com/badge/github/ThreeCatsLoveFish/SubCenter)

## 项目依赖

### 推送服务

由于推送的设备要求与维护成本较高，请自行前往相关网站注册，如需接入可发Issue

- [x] [Server酱-Turbo版](https://sct.ftqq.com/)
- [x] [PushDeer (IOS)](https://github.com/easychen/pushdeer)
- [ ] [PushPlus (微信)](https://www.pushplus.plus/)

### 数据服务

- [x] 天选时刻：[BLTH](https://github.com/andywang425/BLTH)
- [x] A股价格：[东方财富](https://push2.eastmoney.com/)
- [ ] M股价格：[富途](https://www.futunn.com/)
- [ ] 数字货币：[币安](https://www.binance.com/)
- [ ] 钱包变动：[BSC](https://github.com/binance-chain/bsc)
- [ ] ...

## 使用说明

### 环境搭建

```bash
git clone https://github.com/ThreeCatsLoveFish/SubCenter.git
make
./output/subcenter
```

### 推送配置

1. `config/push.toml` 文件中增加需要推送的endpoint和对应token(key)
2. `config/task.toml` 文件中增加需要订阅的内容和发送频率
