## Ruby

> 需要注意的是，此实现并没有发到 `gem` 上，需使用方自行复制

### 使用实例

#### SDK Token

```ruby
require './lib/token.rb'

NetlessToken.sdk_token(
    "netless ak",
    "netless sk",
    1000 * 60 * 10, # token 有效时间 (10分钟)，为 0 时，即永不过期。单位毫秒
    {
        :role => NetlessToken::ROLE::ADMIN  # 可以选择 NetlessToken::ROLE::WRITER / NetlessToken::ROLE::READER 
    }
)
```

#### Room Token

```ruby
require './lib/token.rb'

NetlessToken.room_token(
    "netless ak",
    "netless sk",
    1000 * 60 * 10, # token 有效时间 (10分钟)，为 0 时，即永不过期。单位毫秒
    {
        :role => NetlessToken::ROLE::READER  # 可以选择 NetlessToken::ROLE::WRITER / NetlessToken::ROLE::ADMIN 
        :uuid => "房间的 UUID"
    }
)
```

#### Task Token

```ruby
require './lib/token.rb'

NetlessToken.task_token(
    "netless ak",
    "netless sk",
    1000 * 60 * 10, # token 有效时间 (10分钟)，为 0 时，即永不过期。单位毫秒
    {
        :role => NetlessToken::ROLE::READER  # 可以选择 NetlessToken::ROLE::WRITER / NetlessToken::ROLE::ADMIN 
        :uuid => "任务的 UUID"
    }
)
```