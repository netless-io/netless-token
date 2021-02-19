# netless-token

[English](README-en.md)

该项目用于签出 [Netless](https://netless.link) 服务可识别的 Token，具体请参考[《签出 Token｜项目与权限》](https://developer.netless.link/document-zh/home/project-and-authority#签出-token)。

## 如何使用

你需要获取 ``AK``、和 ``SK``（获取方式请参考[《项目与权限》](https://developer.netless.link/document-zh/home/project-and-authority#签出-token)。之后，根据自己所掌握的语言，选择该 repo 中的 sample codes，将其迁移到自己的项目中。最后，在需要时，传入 ``AK`` 和 ``SK`` 并调用函数生成 Token。

- [JavaScript](/Node/JavaScript)
- [TypeScript](/Node/TypeScript)
- [C#](/csharp)
- [Go](/golang)
- [PHP](/php)
- [Ruby](/ruby)
- [Python][/python]

如果该 repo 没有提供满足你需求的语言的 sample codes，你可以。

- 尝试根据已有的语言仿写你需要语言的等效 codes。
- 通过向 Netless 服务发起 HTTP 请求来申请 Token，参考[《生成 Token》](https://developer.netless.link/server-zh/home/server-token)。但我们**不推荐**这种做法。

## 注意事项

- ``AK``、``SK`` 是你的公司或团队的重要资产。切勿将其传输给客户端，或直接用代码写死在客户端。拿到了 ``AK``、``SK`` 就是拿到了一切，让恶意人士窃取严重破坏你的资产安全。
- 永不过期的 Token 可能为你的业务带来安全隐患。想象一下，如果某人获取了一个权限很高的 Token，他就可以用该 Token 危害你的系统，而你将该 Token 失效的唯一手段只有禁用该 Token 的访问密钥对——这是一个副作用极大的操作。
- 不要将 sdkToken 泄漏到客户端（或前端），也不要将 sdkToken 存入数据库或写入配置文件。应该在使用时临时签出，过期时间尽可能设短。sdkToken 的权限级别很高，泄漏后会危害业务安全。
