### Token
这里的token是JWT（Json Web Token）.

token是由authorization server签名发布的.

authorization server就是使用oauth和openid connect等协议, 并且用户可以使用用户名密码等方式来换取token的服务, 这些token让我们拥有了访问一些资源的权限.

private key 是用来签发token用的, 它只存在于authorization server, 永远不要让别人知道private key.

public key 是用来验证token的, authorization server 会让所有人都知道这个public key的.

api或者各种消费者应该一直检验token, 以确保它们没有被篡改. 它们可以向authorization server申请public key, 并使用它来验证token. 也可以把token传递给authorization server, 并让authorization server来进行验证.


### 参考文档
- [OAuth 2.0 RFC 6749](https://datatracker.ietf.org/doc/html/rfc6749#section-1.1)
- [IdentityServer8](https://identity-server.readthedocs.io/en/latest/index.html)