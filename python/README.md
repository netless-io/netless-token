## Python

> 需要注意的是，此实现并没有发到 `pip` 上，需使用方自行复制

### 使用实例

#### SDK Token

```py
import netless_token

netless_token.sdk_token(
    "netless ak",
    "netless sk",
    1000 * 60 * 10, # token 有效时间 (10分钟)，为 0 时，即永不过期。单位毫秒
    role=netless_token.ADMIN  # 可以选择 netless_token.WRITER / netless_token.READER 
)
```

#### Room Token

```py
import netless_token

netless_token.room_token(
    "netless ak",
    "netless sk",
    1000 * 60 * 10, # token 有效时间 (10分钟)，为 0 时，即永不过期。单位毫秒
    role=netless_token.ADMIN,  # 可以选择 netless_token.WRITER / netless_token.READER 
    uuid="房间的 UUID"
)
```

#### Task Token

```py
import netless_token

netless_token.task_token(
    "netless ak",
    "netless sk",
    1000 * 60 * 10, # token 有效时间 (10分钟)，为 0 时，即永不过期。单位毫秒
    role=netless_token.READER,  # 可以选择 netless_token.WRITER / netless_token.ADMIN 
    uuid="任务的 UUID"
)
```
