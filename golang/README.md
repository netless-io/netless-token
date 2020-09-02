## GoLang

> 需要主要的是，此实现并没有打 `release/tag`，需使用方自行复制

### 使用实例

#### SDK Token

```go
import "your path"

c := token.SDKContent{
    role: token.AdminRole,  // 可以选择 token.ReaderRole / token.WriterRole
}

netlessSDKToken := token.SDKToken(
    "netless ak",
    "netless sk",
    1000 * 60 * 10, // token 有效时间 (10分钟)，为 0 时，即永不过期。单位毫秒
    &c
)
```

#### Room Token

```go
import "your path"

c := token.RoomContent{
    role: token.ReaderRole, // 可以选择 token.AdminRole / token.WriterRole
    uuid: "房间 UUID",
}

netlessRoomToken := token.RoomToken(
    "netless ak",
    "netless sk",
    1000 * 60 * 10, // token 有效时间 (10分钟)，为 0 时，即永不过期。单位毫秒
    &c
)
```

#### Task Token

```go
import "your path"

c := token.TaskContent{
    role: token.WriterRole, // 可以选择 token.AdminRole / token.ReaderRole
    uuid: "房间 UUID",
}

netlessRoomToken := token.TaskToken(
    "netless ak",
    "netless sk",
    1000 * 60 * 10, // token 有效时间 (10分钟)，为 0 时，即永不过期。单位毫秒
    &c
)
```
