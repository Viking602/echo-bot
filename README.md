# Echo Bot

## 运行方式

```bash

./echo_bot

```

## 连接方式

需要设置 Token 否则会返回 401
```
ws://127.0.0.1:6666/ws
```

## 当前已支持功能

```text
/live add <platform> <room_id> 订阅直播
           bilibili
           douyu

/live del <platform> <room_id> 删除直播
```