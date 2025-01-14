# JWT 中不同的签名算法

## 1. ECDSA（椭圆曲线数字签名算法）
基于椭圆曲线密码学
比 RSA 使用更短的密钥就能提供同等安全性
签名速度快，密钥长度短
提供三种变体：ES256、ES384、ES512
- 可以使用多种椭圆曲线（如 P-256、P-384、P-521）
- 需要高质量的随机数生成器
- 实现复杂，容易出现安全漏洞
- 某些操作不是常量时间，可能受到时序攻击

## 2. ED25519
现代的椭圆曲线签名算法
安全性高，性能优异
密钥长度短（公钥32字节，私钥64字节）
没有子变体，只有一种实现
- 基于 Edwards25519 曲线，专门为数字签名优化
- 内置防护措施，抵抗多种密码分析攻击
- 没有可配置参数，降低了错误配置的风险
- 所有操作都是常量时间的，更好地防止侧信道攻击


## 3. HMAC（哈希消息认证码）
使用同一个密钥进行签名和验证（对称加密）
计算速度快，实现简单
提供三种变体：HS256、HS384、HS512
适合单服务器场景

## 4. None
不进行任何签名
极不安全，不建议在生产环境使用
仅用于测试目的

## 5. RSA
使用公钥/私钥对（非对称加密）
广泛使用，成熟可靠
密钥较长，计算相对较慢
提供三种变体：RS256、RS384、RS512

## 6. RSAPSS（RSA 概率签名方案）
RSA 的改进版本
增加了随机性，理论上更安全
提供三种变体：PS256、PS384、PS512
计算开销比标准 RSA 略大

# 选择建议

单服务器场景：
- 推荐使用 HMAC（HS256）
- 实现简单，性能好

分布式系统：
- 推荐使用 ES256（ECDSA）或 RS256（RSA）
- ECDSA 更现代，性能更好
- RSA 更传统，兼容性更好

追求最新最好：
- 推荐使用 ED25519
- 安全性好，性能优异

最常见的选择：
- RS256 - 使用最广泛，生态支持最好
- ES256 - 更现代的选择，逐渐增多

性能对比（大致排序，从快到慢）：
- ED25519
- HMAC
- ECDSA
- RSA/RSAPSS

## 安全性对比：
- ED25519、ECDSA、RSA-PSS 被认为是最安全的
- 标准 RSA 和 HMAC 也很安全（如果正确使用）
- None 完全不安全

## 选择建议：
- 如果不确定选什么，用 RS256
- 如果想要更现代的方案，用 ES256 或 ED25519
- 如果是单服务器且想要简单，用 HS256

## 示例代码 - 不同算法的使用：

```go
import "github.com/golang-jwt/jwt/v5"

func generateTokens() {
    // HMAC
    hmacToken := jwt.New(jwt.SigningMethodHS256)
    
    // RSA
    rsaToken := jwt.New(jwt.SigningMethodRS256)
    
    // ECDSA
    ecdsaToken := jwt.New(jwt.SigningMethodES256)
    
    // ED25519
    ed25519Token := jwt.New(jwt.SigningMethodEdDSA)
    
    // RSAPSS
    rsapssToken := jwt.New(jwt.SigningMethodPS256)
}
```

# ED25519 与 ECDSA 的对比
## 1. 安全性设计
- ED25519
  - 基于 Edwards25519 曲线，专门为数字签名优化
  - 内置防护措施，抵抗多种密码分析攻击
  - 没有可配置参数，降低了错误配置的风险
  - 所有操作都是常量时间的，更好地防止侧信道攻击
- ECDSA
  - 可以使用多种椭圆曲线（如 P-256、P-384、P-521）
  - 需要高质量的随机数生成器
  - 实现复杂，容易出现安全漏洞
  - 某些操作不是常量时间，可能受到时序攻击

## 2. 性能特点
- ED25519

```go
import "crypto/ed25519"

// 生成密钥对
publicKey, privateKey, _ := ed25519.GenerateKey(nil)

// 签名 - 简单直接
signature := ed25519.Sign(privateKey, message)

// 验证 - 同样简单
valid := ed25519.Verify(publicKey, message, signature)
```

- ECDSA
```go
import (
    "crypto/ecdsa"
    "crypto/elliptic"
    "crypto/rand"
)

// 生成密钥对 - 需要指定曲线
privateKey, _ := ecdsa.GenerateKey(elliptic.P256(), rand.Reader)

// 签名 - 需要额外的随机数
r, s, _ := ecdsa.Sign(rand.Reader, privateKey, hash)

// 验证
valid := ecdsa.Verify(&privateKey.PublicKey, hash, r, s)
```

## 3. 主要优势对比
| 特性 | ED25519 | ECDSA |
|------|----------|--------|
| 密钥长度 | 固定32字节公钥/64字节私钥 | 变长，取决于曲线 |
| 签名长度 | 固定64字节 | 变长，通常较大 |
| 速度 | 非常快 | 相对较慢 |
| 随机数需求 | 仅在密钥生成时需要 | 每次签名都需要 |
| 实现复杂度 | 简单 | 复杂 |
| 配置选项 | 无需配置 | 需要选择曲线和参数 |


## 4. 具体优势
1. 确定性签名：

```go
// ED25519 - 相同消息总是产生相同的签名
signature1 := ed25519.Sign(privateKey, message)
signature2 := ed25519.Sign(privateKey, message)
// signature1 == signature2 总是为真

// ECDSA - 相同消息可能产生不同的签名
sig1, _ := ecdsa.Sign(rand.Reader, privateKey, hash)
sig2, _ := ecdsa.Sign(rand.Reader, privateKey, hash)
// sig1 != sig2 很可能不同
```
2. 更简单的 API：
```go
// ED25519 - 简单的 API
type Ed25519PrivateKey []byte
type Ed25519PublicKey []byte

// ECDSA - 更复杂的结构
type ECDSAPrivateKey struct {
    D *big.Int      // 私钥
    PublicKey       // 嵌入的公钥
}

type ECDSAPublicKey struct {
    elliptic.Curve  // 曲线参数
    X, Y *big.Int   // 公钥点坐标
}
```
## 5. 总结：
- ED25519 是更现代的选择，提供了更好的安全性和性能
- 设计更简单，不容易出错
- 不需要配置，开箱即用
- 如果没有特殊的兼容性要求，建议使用 ED25519