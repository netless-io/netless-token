## PHP

> 需要注意的是，此实现并没有发到 `packagist` 上，需使用方自行复制

### 使用实例

#### SDK Token

```php
use Netless\Token\Generate;

$netlessToken = new Generate;
$sdkToken = $netlessToken->sdkToken(
    "netless ak",
    "netless sk",
    1000 * 60 * 10, // token 有效时间 (10分钟)，为 0 时，即永不过期。单位毫秒
    array(
        "role" => Generate::AdminRole // 可以选择 Generate::ReaderRole / TokenRole.WriterRole
    )
);
```

#### Room Token

```php
use Netless\Token\Generate;

$netlessToken = new Generate;
$roomToken = $netlessToken->roomToken(
    "netless ak",
    "netless sk",
    1000 * 60 * 10, // token 有效时间 (10分钟)，为 0 时，即永不过期。单位毫秒
    array(
        "role" => Generate::ReaderRole, // 可以选择 Generate::AdminRole / TokenRole.WriterRole
        "uuid" => "房间的 UUID"
    )
);
```

#### Task Token

```php
use Netless\Token\Generate;

$netlessToken = new Generate;
$taskToken = $netlessToken->taskToken(
    "netless ak",
    "netless sk",
    1000 * 60 * 10, // token 有效时间 (10分钟)，为 0 时，即永不过期。单位毫秒
    array(
        "role" => Generate::WriterRole // 可以选择 Generate::AdminRole / TokenRole.ReaderRole
        "uuid" => "任务的 UUID"
    )
);
```
