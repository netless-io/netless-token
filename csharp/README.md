## C\#

> 需要注意的是，此实现并没有发到 `nuget` 上，需使用方自行复制

### 使用实例

#### SDK Token

```csharp
using Netless;

class Program
{
    static void Main(string[] args)
    {
        string token = NetlessToken.SdkToken(
            "netless ak",
            "netless sk",
            1000 * 60 * 10, // token 有效时间 (10 分钟), 为 0 时, 即永不过期. 单位毫秒
            new SdkContent(
                TokenRole.Admin // 可以选择 Admin/Writter/Reader
            )
        );
    }
}
```

#### Room Token

```csharp
using Netless;

class Program
{
    static void Main(string[] args)
    {
        string token = NetlessToken.RoomToken(
            "netless ak",
            "netless sk",
            1000 * 60 * 10, // token 有效时间 (10 分钟), 为 0 时, 即永不过期. 单位毫秒
            new RoomContent(
                TokenRole.Admin, // 可以选择 Admin/Writter/Reader
                "房间的 UUID"
            )
        );
    }
}
```

#### Task Token

```csharp
using Netless;

class Program
{
    static void Main(string[] args)
    {
        string token = NetlessToken.TaskToken(
            "netless ak",
            "netless sk",
            1000 * 60 * 10, // token 有效时间 (10 分钟), 为 0 时, 即永不过期. 单位毫秒
            new TaskContent(
                TokenRole.Admin, // 可以选择 Admin/Writter/Reader
                "房间的 UUID"
            )
        );
    }
}
```
