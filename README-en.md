# BiliCenter

[中文](README.md) | English

BiliCenter is a middleware that integrates task subscriptions and real-time push

![GitHub go.mod Go version](https://img.shields.io/github/go-mod/go-version/ThreeCatsLoveFish/SubCenter)
[![wakatime](https://wakatime.com/badge/github/ThreeCatsLoveFish/SubCenter.svg)](https://wakatime.com/badge/github/ThreeCatsLoveFish/SubCenter)

## Dependency

### Push Service

- [Server酱-Turbo版](https://sct.ftqq.com/)
- [PushDeer (IOS)](https://github.com/easychen/pushdeer)
- [PushPlus (WeChat)](https://www.pushplus.plus/)

### Data Service

- Bilibili live award [awpush](https://github.com/andywang425/BLTH-server)
- Bilibili live medal [MedalHelper](https://github.com/ThreeCatsLoveFish/MedalHelper)

## Usage

1. Clone source code
   ```bash
   git clone https://github.com/ThreeCatsLoveFish/BiliCenter.git
   cd BiliCenter
   make build
   ```
1. Add config files, fill in personal info and remove *.sample* suffix
   ```
   config
    ├─bili.toml.sample
    ├─bili.toml
    ├─push.toml.sample
    ├─push.toml
    ├─task.toml.sample
    └─task.toml
   ```
1. Launch Sub-center
   ```bash
   make run
   ```
