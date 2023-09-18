# 加密

1. 用secret加请求体进行加密

[来源](https://www.jianshu.com/p/ad410836587a)

### 防止篡改

**参数签名**

1. 按照请求参数名的字母升序排列非空请求参数（包含AccessKey），使用URL键值对的格式（即key1=value1&key2=value2…）拼接成字符串stringA；
2. 在stringA最后拼接上Secretkey得到字符串stringSignTemp；
3. 对stringSignTemp进行MD5运算，并将得到的字符串所有字符转换为大写，得到sign值。

请求携带参数**AccessKey**和**Sign**，只有拥有合法的身份AccessKey和正确的签名Sign才能放行。这样就解决了身份验证和参数篡改问题，即使请求参数被劫持，由于获取不到SecretKey（**仅作本地加密使用，不参与网络传输**），无法伪造合法的请求。

### 实现

`请求接口：http://api.test.com/test?name=hello&home=world&work=java`

- 客户端
    1. 生成当前时间戳timestamp=now和唯一随机字符串nonce=random
    2. 按照请求参数名的字母升序排列非空请求参数（包含AccessKey)`stringA="AccessKey=access&home=world&name=hello&work=java&timestamp=now&nonce=random";`
    3. 拼接密钥SecretKey`stringSignTemp="AccessKey=access&home=world&name=hello&work=java&timestamp=now&nonce=random&SecretKey=secret";`
    4. MD5并转换为大写`sign=MD5(stringSignTemp).toUpperCase();`
    5. 最终请求`http://api.test.com/test?name=hello&home=world&work=java&timestamp=now&nonce=nonce&sign=sign;`

