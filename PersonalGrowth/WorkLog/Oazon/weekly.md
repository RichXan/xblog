# 2025.05.12-2025.05.16
1. 项目环境搭建、熟悉项目代码
2. 分析竞品skylight calendar、拆分理解family calendar产品需求开发
3. tasks、taskInstance需求分析设计，api接口文档设计编写

# 2025.05.19-2025.05.24
1. tasks、taskInstance、task Filter功能接口文档设计及相关接口开发
2. task更新后的socket通知
3. task相关接口联调的修复及优化
4. task重复事件实例过期逻辑讨论设计

# 2025.06.02-2025.06.06
1. 修复family calendar中highest priority的（日历、task、event）bug，以及配合测试修复相关问题
2. 添加Google和MS365的token管理器，支持自动刷新和过期检查功能
3. 添加MS365 API客户端和限流管理器，集成重试机制
4. 修复微软获取refresh token接口报错问题
5. 总结familyCalendar待开发事项

# 2025.06.09-2025.06.14
1. 修复odoo中family calendar 所有缺陷bug，以及配合测试修复相关问题。（剩两个medium缺陷及medium优化项和建议）
2. 过滤webhook导致的重复websocket通知
3. 行事历与伴侣 App 的 Task filter/meals 下的 profile 区分
4. outlook, google事件双向同步增删改查功能，目前已做预留参数待进一步实现及测试。
5. 添加meals、meals categories和recipes的websocket通知功能


# 2025.06.16-2025.06.20
1. 修复calendar loop v1.1.0缺陷。配合测试修复验证相关问题
2. 新增获取设备自身绑定资源信息接口
3. 系统健康检测接口编写及开发
4. 相框模块同步功能
    - 新增相册上传并添加接口
    - 实现压缩图及缩略图存储功能
    - 实现删除同步清理s3文件的处理
    - 实现设备与相框同步功能
    - 兼容适配新增图片与旧图片访问的问题
    - 设备解绑reset时相册同步功能实现
5. google事件双向同步增删改查功能实现

# 2025.06.23-2025.06.27
1. 搭建部署calendarloop测试服务器
    - [x] 搭建自动维护域名证书服务器
    - [x] 搭建calendarloop测试服务器
    - [] 测试服务器拉取docker镜像速度慢
2. 实现google/outlook calendar事件双向同步功能
    - [x] 实现google calendar事件双向同步功能
    - [x] 实现outlook calendar事件双向同步功能
3. 实现日历检测功能
    - [x] 获取日历列表添加日历可用性状态及其对应状态消息。
    - [] 实现日历检测功能接口
4. 修复calendar loop v1.1.0缺陷. 配合测试修复验证相关问题
    - [x] 优化图片压缩时保留元数据，保留原图方向，禁用自动旋转。
    - [x] 修复修改task开启重复事件的时间后，详情页显示重复周期为时间点的问题
    - [x] 修复缩略图上传s3失败问题
    - [x] 修复删除outlook单次事件报错问题
    - [x] 修复无法上传缩略图问题
    - [x] 修复无法创建outlook每月第四个星期三重复任务问题
    - [x] 修复Outlook修改重复事件的所有事件的标题和时间后，成功保存但时间不生效问题
    - [x] 修复Outlook年重复事件返回参数与Google不一致问题
    - [x] 修复创建Outlook每日重复事件时返回wkst字段
    - [x] 双向同步event传参与入参定义
    - [x] 支持outlook修改当前及未来事件
    - [x] 修复无法修改Outlook重复事件当前及以后问题
    - [x] 修复平板上无法修改重复事件的单个事件的时间
    - [x] 修复Outlook修改重复事件的所有事件的标题和时间后，成功保存但时间不生效问题
    - [x] 修复Outlook年重复事件返回参数与Google不一致问题
    - [x] 修复创建Outlook每日重复事件时返回wkst字段
    - [x] 修复Google&Outlook 修改 All 事件，之前事件被清除问题 
    - [x] 修复修改outlook/google重复事件当天及以后，当天事件的原事件存在问题
    - [x] 修复事件卡片错误展示Profile 数量问题

# 2025.06.30-2025.07.04
1. 修复calendarloop v1.1.0缺陷，配合测试修复验证相关问题
2. 实现相同google/outlook日历事件去重合并profiles信息功能
3. 设计并实现日历可用性检测接口
    - [x] 实现日历可用性检测接口
    - [x] 实现统一日历状态错误消息和错误代码定义
    - [x] 实现获取日历列表中添加日历可用性状态及其对应状态消息,以支持前端多语言功能
    - [x] 总结归纳不同日历错误对应的引导操作,以及错误代码的定义同步
4. 讨论设备分享功能设计方案, 设备备案以及产测流程功能设计方案
    - [x] 讨论设备分享功能, 设备备案以及产测流程设计方案
    - [x] 编写并确认设备分享功能设计文档, 设备备案以及产测流程功能设计文档
5. 设备资源分享功能设计及开发
    - [x] 实现生成资源邀请接口 (支持设置最大邀请人数, 有效期, 邀请者备注)
    - [x] 实现获取资源邀请接口 (获取资源名, 资源所属者信息)
    - [x] 实现支持受邀者退出资源功能
    - [x] 实现使用邀请token加入资源接口
    - [x] 实现资源成员管理列表接口
    - [x] 实现资源成员移除接口
6. 实现平板设备备案及产测功能(进行中)

# 2025.07.07-2025.07.11
- 对接联调日历可用性检测, 设备资源分享功能
- 实现平板设备备案及产测功能相关接口
    - [x] 产测计划脚本生成
    - [x] 产测计划管理
    - [x] license管理
    - [x] 产测报告结果管理, 设备激活记录
    - [x] 设备根据license交换设备登录凭证
- 部署calendarloop邮件服务器, 将品牌统一为calendarloop 
- 迁移calendarloop到新生产服务器
- 更换calendarloop的邮件服务器为aws ses中继发送邮件

# 2025.07.14-2025.07.18
- [1] 配合前端对接,联调并完善平板设备产测相关功能
- [1] 讨论及编写搭建 caldav 日历服务器相关文档及定义
    - 搭建配置 caldav 日历服务器
    - 持久化事件信息到 caldav 日历服务器
    - 实现 caldav 日历服务器事件同步功能
- [2] 部署搭建calendarloop的mongodb服务
    - 搭建mongodb服务，并配置副本集
    - 配置mongodb相关配置以支持事务回滚功能
    - 迁移原mongodb数据到新mongodb服务

# 2025.07.21-2025.07.25
- [x] 实现设备版本管理功能及相关接口
- [x] 实现第三方授权信息管理功能及相关接口
- [x] 实现设备自身信息管理功能及相关接口
- [x] 设备产测相关功能联调，数据提供及功能支持
- [x] 处理日历可用性检查相关缺陷
- 搭建配置 Caldav 日历服务器
    - [] 部署生产环境caldav服务器
- 集成第三方 Caldav 日历
    - [x] 集成icloud/cozi/yahoo日历源, icalendar格式内容解析及集成至日历服务器中
    - [x] calendarProvider / calendar 相关数据模型的集成及适配
    - [] event相关接口适配
    - [] 实现日历事件异步缓存功能
    
# 2025.07.28-2025.08.01
- [x] 生产环境搭建/部署及验证caldav日历服务器，postgres数据库
- [x] 集成第三方 caldav 日历（icloud/cozi/yahoo）
    - [x] 联调日历源集成接口
    - [x] 实现event相关接口适配
    - [x] 实现icloud/cozi/yahoo日历源事件定时异步缓存功能
- [x] 联调 产测 相关接口
    - [x] 产测备案注册时sn和mac的校验
    - [x] 支持通过sn和mac地址获取设备licenses信息
    - [x] 测试联调，提供产测测试的设备licenses及相关数据
- [x] 联调 设备版本 管理相关接口，新增升级链接信息
- [] 联调其他日历源集成功能，测试并修复相关缺陷

# 2025.08.04-2025.08.08
- [x] 生产环境搭建/部署及验证caldav日历服务器，postgres数据库
- [x] 集成第三方 caldav 日历（icloud/cozi/yahoo）
    - [x] 联调日历源集成接口
    - [x] 实现event相关接口适配
    - [x] 实现icloud/cozi/yahoo日历源事件定时异步缓存功能
- [x] 联调 产测 相关接口
    - [x] 产测备案注册时sn和mac的校验
    - [x] 支持通过sn和mac地址获取设备licenses信息
    - [x] 测试联调，提供产测测试的设备licenses及相关数据
- [x] 联调 设备版本 管理相关接口，新增升级链接信息
- [] 联调其他日历源集成功能，测试并修复相关缺陷

# 2025.08.04-2025.08.08
- [x] 实现icloud / yahoo / cozi 日历权限诊断功能
- [x] 修复处理icloud/yahoo/cozi日历源集成功能，测试并修复相关缺陷
    - [x] 添加从Yahoo获取日历名称的功能
    - [x] 修复日历源集成时，日历事件写入失败问题
    - [x] 优化外部日历同步逻辑，处理重复事件中特殊事件
    - [x] 处理重复事件，支持RECURRENCE-ID事件的UID修改和覆盖实例的应用
    - [x] 修复Cozi/Yahoo/icloud重复事件时区问题
    - [x] 修复Cozi/Yahoo/icloud长标题事件无法同步标题信息问题
    - [x] 修复Cozi/Yahoo/icloud仅修改或仅删除当前重复事件未同步问题
    - [x] 修复第三方事件同步时重复socket通知问题
- [x] 修改设备产测流程及逻辑相关接口，提供产测相关数据及license信息
    - [x] 支持通过MAC和SN查询许可证信息
    - [x] 修复更新设备版本时的唯一索引冲突检查逻辑
    - [x] 修改产测服务逻辑相关接口将mac调整为model字段
    - [x] 提供产测相关数据license信息(测试、周工)
- [x] 配置搭建calendarloop生产环境及gitlab环境配置
- [] 测试环境部署calendarloop

# 2025.08.11-2025.08.15
Calendarloop
- [x] calendarloop测试环境部署及gitlab环境配置
- [x] 产测流程调整及文档整理，兼容烧录、不烧录、回收license的情况，并且支持删除回收license功能
- [x] 更新默认菜单列表逻辑，确保未激活设备显示默认菜单项
- [x] 添加错误信息多语言支持，新增德语、西班牙语、法语、意大利语和俄语的错误信息和通知文本
- [x] 协助排查客户反馈calendarloop相关问题

Mocreo
- [x] 对接Alexa Skill Kit, 完善Oauth、lambda配置信息

# 2025.08.18-2025.08.22
- [x] 对接Alexa Skill相关接口（Alexa、DiscoveryTemperatureSensor、StateReport）
- [x] 实现并跑通基础的Alexa Skill使用流程
- [x] Aleax接入Mecreo，及相关接口编写，lambda函数编写，alexa实现目标与计划总结
- [x] 熟悉mocreo v3代码

# 2025.08.25-2025.08.29
- MOCREO
    - 搭建部署MOCREO v3调试环境
    - Alexa Skill 的基础配置与测试，Alexa Skill发布到线上
    - 支持alexa向服务器读取设备状态
    - 实现 App-to-App Account Linking功能
        - 相关跳转接口及参数配置(iOS Universal Link / Android App Link)
        - 验证 iOS 与 Android 平台的跳转流程  >>> 进行中，需要前端配合联调
    - 实现Alexa与mocreo的用户账号关联信息记录功能 >>> 进行中

- Calendarloop
    - 调整产测备案、认证流程流程，文档更新
        - 支持根据加密后的工厂名匹配对应licnese
        - 支持产测动态传入排除的菜单项
        - 支持传入加密后的菜单项及加密后的工厂名信息

# 八月份
- 对于Calendarloop项目的开发，第三方日历源的集成，产测流程的调整，对calendarloop做阶段性收尾的工作。后续

在做集成兼容icloud, yahoo, cozi日历源中，锻炼了我在接口抽象、通用模型设计上的能力。后续在针对其他不同协议日历源的差异，需要多考虑边界情况，代码逻辑中实现兼容性适配。对于网络维护这一块的能力需要多加锻炼，多思考服务器维护的思路，以及更新部署服务器的效率。

# 2025.09.01-2025.09.05
- MOCREO
    - 实现邮件html发送功能
        - [x] 支持过滤支持dismiss功能设备的内容展示
        - [x] 实现并验证邮件html发送功能
        - [x] 在下发dismiss功能下添加设备固件版本判断
        - [x] 在告警通知html模板中兼容不支持dismiss功能设备的内容展示
    - 实现Alexa与mocreo的用户账号关联信息记录功能 
    - 公开API文档边界条件补充

- Calendarloop
    - 协助排查socket通知发送及接收问题
    - 服务器磁盘inode占用问题排查及处理

# 2025.09.08-2025.09.12
- MOCREO
    - Stoplight Elements部署搭建，openapi中文、英文文档整理
    - 公开API文档边界条件补充及odoo相关bug修复
    - 支持 H3 继电器控制
    - 补充LoRa设备信号强度转换百分比方法
    - NS1 / NS2 蜂鸣器控制
    - App和邮件dismiss功能token有效期调整

# 2025.09.15-2025.09.19
- MOCREO
    - 用户绑定手机号功能开发，添加手机号码绑定和验证码发送相关接口
    - 修改邮件功能开发
    - LW1首次setup进行按键的通知推送内容修改
    - 修改通知规则结构
    - 设备重复告警事件修复，用户版本已知晓确认功能开发
    - Dismiss指令时修改为向sensor所有的代理网关和指定beep的网关下发dismiss指令
    - 设备切换资产后没有通知设备更新规则问题修复
    - 处理dismiss重复弹窗问题

- [1] MOCREO
    - [1] aws 短信通知集成，手机号绑定
    - [1] 修改用户邮箱功能开发
    - [2] LW1首次setup进行按键的通知推送内容修改
    - [2] 修改规则结构、支持ios重要推送通知
    - [2] Amazon 用户邀评逻辑处理
    - [2] V3 公开API文档调整

# 2025.09.22-2025.09.26
- MOCREO
    - aws 短信通知集成
    - 支持ios重要推送通知
    - 导出lw1设备数据时water_level字段处理
    - Dismiss 成功后通知推送，消息内容确认
    - Alexa Uninversal Link Nginx转发配置

开发1.3.0版本相关功能，实现LW1、SW3、带蜂鸣器NS1/NS2集成，实现Alexa功能集成，完成短信通知功能，html通知邮件改进与Dismiss功能支持，V3公开API文档边界条件处理及环境部署，手机号/邮件修改功能开发

# 2025.10.9-2025.10.11
- MOCREO
    - 开发集成aws sms 短信通知、手机号绑定、更换、获取功能开发联调
    - Dismiss功能方案讨论及方案确认
    - 邮件内容新增iPhone 17用户的重要通知，更新FAQ链接格式
    - LD1的设备描述文件审核及更新
    - 更新生产环境部署方式

# 2025.10.13-2025.10.17
- MOCREO
    - dismiss功能逻辑调整（需要兼容旧固件 在分级规则触发之后还会重新上报trigger）
    - aws 短信额度低通知功能
    - LD1 门锁状态支持
    - 配合测试联调验证alexa、aws sms相关功能


# 本周工作
- Mocreo
    - 邮件html模板内容调整及测试联调验证
    - Stoplight Elements部署搭建，openapi文档整理
    - dismiss功能完善及验证
    - app iot server对接联调及问题排查处理
    - Alexa联调对接
        - 联调对接App-to-App Account Linking功能, 并验证iOS与Android平台的跳转流程
        - Alexa Skill功能迭代发布 >>> 新功能需要等MOCREO v1.3.0版本部署上线后发布

    - [] odoo相关bug修复
    - [x] 补充LoRa设备信号强度转换百分比方法 >>> Loki
    - [x] 公开api 文档补充，上线生产环境 >>> Loki
    - [] 短信通知
    - [] 修改规则结构、支持ios重要推送通知 >>> Loki
    - [x] 修改邮箱
    - [] Dismiss 成功后推送通知 >>> Loki 
    - [x] NS1/NS2 蜂鸣器控制 >>> Loki 
    - [x] H3 继电器控制

# 待办事项
- [x] 整理归纳通知需要下发到哪个hub上？（代理hub、关联hub、所有hub）
- [x] LW1首次setup进行按键的通知推送内容修改
    - setup/scan 扫码绑定接口添加 redis key三小时缓存，在上报 {"type": "EVENT", "deviceId": "0030AE1118301F00", "data": {"eventType": "refresh", "params": {}}}  的信息中判断是否存在首次上报的数据，如果是首次绑定的话就发送deviceInit首次绑定的消息，否则发送test notification消息
- [x] 在切换资产后，需要通知Hub重新获取一次规则

- 添加已知晓接口，用户更新到1.3.0之后，将用户所有资产下所有设备的告警规则rule的condition.all中的params.refreshTime若为600修改为0。

- 短信通知功能开发
- LW1告警短信发送

- dismiss链接点击后下发给该node的所有代理hub和选择beep的hub


- 水位对应映射关系，过滤水位传入的max,min聚合参数，只返回水位对应的映射关系
\-1(顶部漏水)
0 (无水)
1\~n(底部漏水，漏水严重性)【1，2，3，4】
2，是低水位漏水
3，是中水位漏水
4，高水位漏水


- 触发告警时创建会话信息，在通知中返回会话id
- 在recover的时候删除会话id

- 抽象会话id，每一次rule的触发（一次active/inactive）的切换就是一次会话

- 设备获取command添加超时删除command不在下发处理(待测试什么情况下会出现)

- createCommandFromApp时如果为dismiss，创建指令的时候, 如果传入的params.repeat为true将rule.attributes.alertDismissTime设置为当前时间+一个小时的时间戳。如果传入的params.repeat为false将rule.attributes.alertDismissTime设置为forever。在收到设备上报的告警恢复的消息时，将rule.attributes.alertDismissTime设置为空。

- 用户刚注册的这封邮件需要调整:
1. 如果用户填写了个人信息，就不要用 Hello friend 了。
2. 介绍了 Youtube 设置视频，但是跳转链接呢?
3. FAQ 引导页面最好是去往一个目录页面，而不是 H5-Lite

- 告知测试1.3.0功能的测试点，要在周五前上，除了短信功能。



新Dismiss逻辑：
分级规则激活后，没点击dismiss，不管多久
1. 触发值相同，不重复触发告警
2. 触发值不同，不重复触发告警

分级规则激活后，点击临时dismiss一个小时内
1. 触发值相同，不重复触发告警
2. 触发值不同，不重复触发告警

分级规则激活后，点击临时dismiss一个小时后
1. 触发值相同，重复触发告警
2. 触发值不同，重复触发告警

分级规则激活后，点击永久dismiss，不管多久
1. 触发值相同，不重复触发告警
2. 触发值不同，不重复触发告警

服务器改动项：
1. 重构规则引擎，需要动态维护规则dismiss状态、是否为永久dismiss字段。根据dismiss的状态来判断是否需要下发指令给设备。
2. dismiss多久根据设备描述文件.attr.alarmRepeatInterval决定。
3. 所有设备需要测试验证

===========================
现有Dismiss逻辑：

服务器：
普通规则激活后：
1. refreshTime > 0：按刷新时间间隔触发
2. refreshTime = 0：触发后1小时内不再触发
分级规则激活后：
1. 触发值相同时，按 refreshTime 判断（过滤重复trigger）
2. 触发值变化时，立即触发（绕过时间限制）

设备侧：
1. 触发值相同时，不重复上报trigger
2. 触发值变化时，上报trigger
3. refreshTime默认为600s，现在已弃用。临时dismiss后，设备根据alarmRepeatInterval进行重复上报trigger

===========================
待确认：
1. 在dismiss的情况下，触发水浸等级变化，Hub需不需要重新上报trigger（需不需要发送邮件）
2. 指定多hub beep时，按其中一个hub停止beep时，另外的hub是否也会停止beep？
3. 在按物理按键hub停止beep时，设备测后续是如何上报trigger？




