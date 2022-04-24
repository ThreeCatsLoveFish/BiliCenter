# SubCenter

中文 | [English](README-en.md)

推送中心 —— 集成各种任务并进行实时推送的中间件。

![GitHub go.mod Go version](https://img.shields.io/github/go-mod/go-version/ThreeCatsLoveFish/SubCenter)
[![wakatime](https://wakatime.com/badge/github/ThreeCatsLoveFish/SubCenter.svg)](https://wakatime.com/badge/github/ThreeCatsLoveFish/SubCenter)

## 项目依赖

### 推送服务

由于推送的设备要求与维护成本较高，请自行前往相关网站注册，如需接入可发Issue

- [Server酱-Turbo版](https://sct.ftqq.com/)
- [PushDeer (IOS)](https://github.com/easychen/pushdeer)
- [PushPlus (微信)](https://www.pushplus.plus/)

### 数据服务

- 天选时刻：[awpush](https://github.com/andywang425/BLTH-server)

## 使用说明

### 环境搭建

```bash
git clone https://github.com/ThreeCatsLoveFish/SubCenter.git
cd SubCenter
make build
```

### 配置文件

```
config
 ├─bili.toml
 ├─push.toml
 └─task.toml
```

### 运行程序

```bash
make run
```
