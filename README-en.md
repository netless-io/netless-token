# netless-token

[中文版](README.md)

This project is used to check out the Token that can be recognized by the [Agora Interactive Whiteboard](https://docs.agora.io/en/whiteboard/product_whiteboard?platform=Android) service. For details, please refer to [Interactive Whiteboard Token Overview](https://docs.agora.io/en/whiteboard/whiteboard_token_overview?platform=RESTful).

## How to use

You need to obtain ``AK`` and ``SK`` (please refer to [Get access keys](https://docs.agora.io/en/whiteboard/whiteboard_token_overview?platform=Android#get-access-keys) for obtaining methods. After that, select the sample codes in the repo according to your own language and migrate them to your own project. Finally, when needed, Pass in ``AK`` and ``SK`` and call the function to generate Token.

- [JavaScript](/Node/JavaScript)
- [TypeScript](/Node/TypeScript)
- [C#](/csharp)
- [Go](/golang)
- [PHP](/php)
- [Ruby](/ruby)
- [Python](/python)

If the repo does not provide sample codes in a language that meets your needs, you can do the following.

- Try to imitate the equivalent codes of the language you need based on the existing language.
- To apply for a token by initiating an HTTP request to the Agora Interactive Whiteboard service, refer to [Generate a Token Using RESTful API](https://docs.agora.io/en/whiteboard/generate_whiteboard_token?platform=RESTful). But we do **NOT** recommend this approach.

## Note

- ``AK`` and ``SK`` are important assets of your company or team. Do not transmit it to the client, or directly code it to the client. Getting ``AK`` and ``SK`` means getting everything, allowing malicious people to steal and seriously damage your asset security.
- Tokens that never expire may bring security risks to your business. Imagine that if someone acquires a highly authorized token, he can use the token to harm your system, and the only way you can invalidate the token is to disable the access key pair of the token—this is a side effect Great operation.
- Don't leak the sdkToken to the client (or front-end), and don't store the sdkToken in the database or write it into the configuration file. It should be checked out temporarily while in use, and the expiration time should be set as short as possible. The permission level of sdkToken is very high, and the leakage will endanger business security.
