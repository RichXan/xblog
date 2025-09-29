> 前提：
1. Amazon 开发者账户 & AWS 账户： 创建和管理 Alexa Skill 需要 Amazon 开发者账户，而托管云服务（推荐使用 AWS）需要 AWS 账户。

# 1. 将物联网设备接入aleax

## 1. 设备发现与账户关联流程 (Discovery & Account Linking)
用户初次将您的设备添加到 Alexa 时触发。此流程说明了用户如何首次让 Alexa 识别并添加您的智能设备。

流程说明：
1. 用户操作： 用户在 Alexa App 中搜索并启用您的智能家居技能。
2. 账户关联： Alexa App 会引导用户进行账户关联，即在您的身份验证服务上登录。成功登录后，您的服务会给 Alexa 返回一个 OAuth token。
3. 设备发现： 用户在 Alexa App 中发起 "发现设备" (Discover Devices) 操作。
4. Alexa 调用 Lambda： Alexa 语音服务 (AVS) 将调用您的 Lambda 函数，并发送一个 Discovery Request 请求。
5. Lambda 返回设备列表： 您的 Lambda 函数查询您的后端数据库（例如 AWS IoT Core 的 Thing 注册），获取所有已注册的设备及其支持的 Alexa 功能（例如开/关、亮度控制等），然后将这些信息封装成 Discovery Response 返回给 Alexa。
6. Alexa 添加设备： Alexa 收到响应后，会将这些设备添加到用户的智能家居设备列表中，用户就可以通过语音或 App 控制它们了。

```mermaid
graph TD
    subgraph 用户端交互
        A[开始: 用户打开 Alexa App] --> B[查找并启用您的智能家居技能]
    end

    subgraph Alexa 服务
        C[Alexa App]
    end

    subgraph 您的后端云服务 (AWS 最佳实践)
        D[您的认证/OAuth 2.0 服务]
        E[您的 Lambda 函数 (Alexa Skill 后端)]
        F[您的设备注册数据库/AWS IoT Core Thing 注册]
    end

    subgraph 核心流程
        B -- 提示进行账户关联 --> C
        C -- 打开您的 OAuth 登录页面 --> D
        D -- 用户登录并授权 --> C
        C -- 收到 OAuth access_token --> C
        C -- 用户点击 "发现设备" --> G[Alexa 语音服务 (AVS)]
        G -- 发送 `Discovery Request` 请求 --> E
        E -- 查询已注册设备列表 --> F
        F -- 返回设备列表及能力信息 --> E
        E -- 返回 `Discovery Response` 响应 --> G
        G -- 添加发现的设备到用户账户 --> H[设备出现在 Alexa App 中]
    end

    A --> B
    B --> C
    C --> D
    D --> C
    C --> G
    G --> E
    E --> F
    F --> E
    E --> G
    G --> H
```

## 2. 设备控制流程
此流程说明了用户如何通过语音命令控制您的智能家居设备。


流程说明：

1. 语音指令： 用户向 Alexa 设备（如 Echo Dot）发出语音指令。
2. Alexa 解析： Alexa 语音服务 (AVS) 识别并解析语音指令，确定要调用哪个智能家居技能以及具体的设备操作（例如 "打开客厅灯" 对应 SetPowerState 命令）。
3. Alexa 调用 Lambda： AVS 调用您的 Lambda 函数，发送 Control Request（例如 SetPowerState，包含设备 ID 和目标状态）。
4. Lambda 控制设备： 您的 Lambda 函数接收请求后，通过 AWS IoT Core (MQTT 代理或设备影子) 向您的设备发送控制命令。
    - MQTT 方式： Lambda 向设备订阅的特定 MQTT 主题发布消息。
    - 设备影子方式： Lambda 更新设备影子的 desired 状态，设备监听 desired 状态的变化。
5. 设备执行： 您的智能设备接收到命令后，执行相应的操作（例如打开灯）。
6. 设备报告状态（重要）： 设备执行完操作后，应向 AWS IoT Core 报告其当前状态（例如灯已打开），更新设备影子的 reported 状态或发布状态 MQTT 消息。这确保 Alexa App 或其他查询能获取到设备的最真实状态。
7. Lambda 返回响应： Lambda 函数将操作结果返回给 Alexa。
8. Alexa 反馈： Alexa 收到响应后，给予用户语音反馈（例如 "好的"）。

```mermaid
graph TD
    subgraph 用户与 Alexa 交互
        A[开始: 用户发出语音指令: "Alexa, 打开客厅灯"]
    end

    subgraph Alexa 服务
        B[Alexa 语音服务 (AVS)]
    end

    subgraph 您的后端云服务 (AWS 最佳实践)
        C[您的 Lambda 函数 (Alexa Skill 后端)]
        D[AWS IoT Core (MQTT 代理/设备影子)]
    end

    subgraph 您的 IoT 设备
        E[您的智能物联网设备]
    end

    A -- 语音识别与意图解析 --> B
    B -- 调用识别到的智能家居技能 --> C
    C -- 发送 `Control Request` (例如: `SetPowerState`) --> C
    C -- 将指令转换为 MQTT 消息 或 更新设备影子 --> D
    D -- MQTT 消息发布 或 设备影子更新 --> E
    E -- 设备收到命令并执行操作 (例如: 开灯) --> F[设备操作完成]
    E -- (可选) 向 AWS IoT Core 报告新状态 --> D
    D -- (可选) 设备影子更新 或 接收状态 --> D
    C -- 返回 `Control Response` (成功/失败) --> B
    B -- 语音反馈 --> G[Alexa 回复: "好的"]

    A --> B
    B --> C
    C --> D
    D --> E
    E --> F
    F --> D
    D --> C
    C --> B
    B --> G
```

## 3. 设备状态报告流程 (Proactive Events)
此流程说明了当设备状态在没有 Alexa 命令的情况下发生变化时，如何主动通知 Alexa，以保持状态同步。这对于手动操作设备（如按物理按钮）或传感器数据变化非常有用。

流程说明：

1. 设备状态变化： 您的智能设备上的状态发生变化，例如用户通过物理按钮打开了灯，或者传感器检测到温度变化。
2. 设备报告到云服务： 设备将新的状态通过 MQTT 消息发布到 AWS IoT Core，或者更新其设备影子的 reported 状态。
3. 触发 Lambda： AWS IoT Core 的规则引擎监测到这些状态变化，并触发您的 Lambda 函数。
4. Lambda 发送主动事件： Lambda 函数将设备的新状态封装成一个 Alexa ChangeReport 事件（一种 Proactive Event），并通过 Alexa 事件网关发送给 Alexa。
5. Alexa 更新状态： Alexa 收到 ChangeReport 后，会更新其内部的设备状态，并在 Alexa App 中同步显示。这样，即使是非语音控制导致的状态变化，也能在 Alexa App 中准确反映。

```mermaid
graph TD
    subgraph 您的 IoT 设备
        A[开始: 智能设备状态发生变化 (例如: 手动开关, 传感器读数更新)]
    end

    subgraph 您的后端云服务 (AWS 最佳实践)
        B[AWS IoT Core]
        C[AWS IoT Core 规则引擎]
        D[您的 Lambda 函数 (处理状态更新)]
    end

    subgraph Alexa 服务
        E[Alexa 事件网关 (Event Gateway)]
    end

    subgraph Alexa 用户端
        F[Alexa App 设备状态更新]
    end

    A -- 发布新状态到 MQTT topic 或更新设备影子 --> B
    B -- 触发 AWS IoT Core 规则 --> C
    C -- 调用您的 Lambda 函数 --> D
    D -- 格式化 `ChangeReport` (Proactive Event) --> E
    E -- 发送状态更新到 Alexa --> F

    A --> B
    B --> C
    C --> D
    D --> E
    E --> F
```


# 2. Aleax支持的语言交互模型
Alexa 支持两种类型的语音交互模型：

- **预构建的语音交互模型** ：其交互模式是声明式而非对话式。在此模型中，ASK 定义了用户为调用技能而说的一组单词。例如，用户可以说“Alexa，打开灯” 或 “Alexa，关掉电视”。 您只需定义您的技能即可接受这些预定义的请求。
- **自定义语音交互模型** ： 这种模型允许创建高度灵活和对话式的体验。自定义模型提供最大的灵活性，但最复杂。您设计整个语音交互。使用自定义模型时，通常必须定义用户向技能传达相同请求的每种方式。例如，“Alexa，计划从西雅图到丹佛的旅行”、“Alexa，我想从西雅图去丹佛旅行” 和 “Alexa，计划去丹佛旅行”。

# 3. 设备发现与账户关联流程 (Discovery & Account Linking)


# 4. 设备控制流程

# 5. 设备状态报告流程 (Proactive Events)


# 6. node.js lambda
```nodejs
// lambda_function.js
const https = require('https'); // For HTTPS requests to your custom service
const util = require('util'); // For util.inspect (optional, for deep logging)

// 配置日志级别
const LOG_LEVEL = process.env.LOG_LEVEL || 'INFO';
const logger = {
    info: (msg) => LOG_LEVEL === 'INFO' || LOG_LEVEL === 'DEBUG' ? console.log(`INFO: ${msg}`) : null,
    warn: (msg) => LOG_LEVEL === 'WARN' || LOG_LEVEL === 'INFO' || LOG_LEVEL === 'DEBUG' ? console.warn(`WARN: ${msg}`) : null,
    error: (msg, error) => {
        if (LOG_LEVEL === 'ERROR' || LOG_LEVEL === 'WARN' || LOG_LEVEL === 'INFO' || LOG_LEVEL === 'DEBUG') {
            console.error(`ERROR: ${msg}`);
            if (error) console.error(error);
        }
    },
    debug: (msg) => LOG_LEVEL === 'DEBUG' ? console.log(`DEBUG: ${msg}`) : null,
};

// 从环境变量获取您的自定义服务URL和API Key
const YOUR_SERVICE_BASE_URL = process.env.YOUR_SERVICE_BASE_URL || 'https://your.custom.service.com/api/v1';
const YOUR_SERVICE_API_KEY = process.env.YOUR_SERVICE_API_KEY || 'YOUR_SUPER_SECRET_API_KEY'; // 生产环境请使用更安全的管理方式

/**
 * 辅助函数: 调用您的自定义后端服务获取设备的最新状态。
 * @param {string} deviceId 从 Alexa 请求中获取的设备唯一ID。
 * @returns {Promise<object|null>} 设备的温度和湿度数据，或在失败时返回null。
 */
async function getDeviceInfoFromYourService(deviceId) {
    return new Promise((resolve, reject) => {
        // 假设您的服务查询API是 GET /api/v1/devices/{device_id}/status
        const url = new URL(`${YOUR_SERVICE_BASE_URL}/devices/${deviceId}/status`);
        
        const options = {
            hostname: url.hostname,
            port: url.port || 443, // Default HTTPS port
            path: url.pathname + url.search,
            method: 'GET',
            headers: {
                'Authorization': `Bearer ${YOUR_SERVICE_API_KEY}`, // 根据您的认证方式调整
                'Content-Type': 'application/json',
            },
            timeout: 5000, // 5秒超时
        };

        logger.info(`Calling your service for device ${deviceId} at ${url.toString()}`);

        const req = https.request(options, (res) => {
            let data = '';
            res.on('data', (chunk) => {
                data += chunk;
            });
            res.on('end', () => {
                if (res.statusCode >= 200 && res.statusCode < 300) {
                    try {
                        const deviceData = JSON.parse(data);
                        logger.info(`Received data from your service: ${util.inspect(deviceData, {depth: null})}`);
                        resolve(deviceData);
                    } catch (e) {
                        logger.error(`Error parsing JSON from your service: ${e}`);
                        reject(null); // 解析失败
                    }
                } else {
                    logger.error(`Your service responded with status ${res.statusCode}: ${data}`);
                    reject(null); // 服务端错误
                }
            });
        });

        req.on('error', (e) => {
            logger.error(`Error during HTTP request to your service: ${e.message}`, e);
            reject(null); // 网络错误
        });

        req.on('timeout', () => {
            req.destroy();
            logger.error(`HTTP request to your service timed out.`);
            reject(null); // 超时
        });

        req.end();
    });
}

/**
 * 辅助函数: 构造 Alexa 响应的通用结构。
 * @param {object} header 响应头。
 * @param {object} payload 响应体。
 * @param {string} [endpointId] 设备的唯一ID (可选)。
 * @returns {object} Alexa 响应对象。
 */
function buildAlexaResponse(header, payload, endpointId = null) {
    const response = {
        event: {
            header: header,
            payload: payload
        }
    };
    if (endpointId) {
        response.event.endpoint = { endpointId: endpointId };
    }
    
    logger.info(`Alexa Response: ${JSON.stringify(response, null, 2)}`);
    return response;
}

/**
 * 处理 Alexa 的 Discovery 请求，返回设备列表。
 * @param {object} event Lambda 事件对象。
 * @returns {object} Alexa Discover.Response。
 */
async function handleDiscoveryRequest(event) {
    logger.info("Handling Discovery Request");
    
    // 这里硬编码一个设备作为示例。
    // 在真实场景中，你会从您的后端数据库加载用户拥有的设备列表。
    const sampleDevice = {
        endpointId: "myCustomTempHumiditySensor001", // 在 Alexa 系统中唯一的ID
        manufacturerName: "MyCustomIoT",
        description: "Custom Temperature and Humidity Sensor",
        friendlyName: "客厅温湿度传感器", // 用户在 Alexa App 中看到的名称
        displayCategories: ["TEMPERATURE_SENSOR", "OTHER"],
        capabilities: [
            {
                type: "AlexaInterface",
                interface: "Alexa.TemperatureSensor",
                version: "3",
                properties: {
                    supported: [{ name: "temperature" }],
                    retrievable: true,
                    proactivelyReported: false // 如果设备能主动推送状态，设置为true
                }
            },
            {
                type: "AlexaInterface",
                interface: "Alexa.RangeController", // 用于表示湿度
                instance: "humidity", // 实例名称，对应属性
                version: "3",
                properties: {
                    supported: [{ name: "rangeValue" }],
                    retrievable: true,
                    proactivelyReported: false
                },
                capabilityResources: {
                    friendlyNames: [
                        { "@type": "text", "value": { "text": "湿度", "locale": "zh-CN" } }
                    ]
                },
                configuration: {
                    supportedRange: { minimumValue: 0, maximumValue: 100 },
                    unitOfMeasure: "PERCENT",
                    presets: []
                }
            },
            {
                type: "AlexaInterface",
                interface: "Alexa.EndpointHealth",
                version: "3",
                properties: {
                    supported: [{ name: "connectivity" }],
                    retrievable: true,
                    proactivelyReported: false
                }
            }
        ],
        cookie: { // 任意自定义数据，会在后续请求中原样返回
            internalDeviceId: "ABC-123-DEVICE" // 这是您后端服务中识别设备的ID
        }
    };

    const header = {
        namespace: "Alexa.Discovery",
        name: "Discover.Response",
        payloadVersion: "3",
        messageId: event.directive.header.messageId
    };
    const payload = {
        endpoints: [sampleDevice]
    };
    
    return buildAlexaResponse(header, payload);
}

/**
 * 处理 Alexa 的 ReportState 请求，用于查询设备当前状态。
 * @param {object} event Lambda 事件对象。
 * @returns {object} Alexa StateReport 响应。
 */
async function handleReportStateRequest(event) {
    logger.info("Handling ReportState Request");
    
    const endpointId = event.directive.endpoint.endpointId;
    const internalDeviceId = event.directive.endpoint.cookie.internalDeviceId;
    const bearerToken = event.directive.payload.scope.token; // 用户 access token

    if (!internalDeviceId) {
        logger.error(`Internal Device ID not found in cookie for endpoint: ${endpointId}`);
        const header = {
            namespace: "Alexa",
            name: "ErrorResponse",
            payloadVersion: "3",
            messageId: event.directive.header.messageId
        };
        const payload = {
            type: "NO_SUCH_ENDPOINT",
            message: "Device configuration error."
        };
        return buildAlexaResponse(header, payload, endpointId);
    }

    // 调用您的自定义服务获取设备数据
    const deviceData = await getDeviceInfoFromYourService(internalDeviceId);
    if (!deviceData) {
        const header = {
            namespace: "Alexa",
            name: "ErrorResponse",
            payloadVersion: "3",
            messageId: event.directive.header.messageId
        };
        const payload = {
            type: "ENDPOINT_UNREACHABLE",
            message: "Could not retrieve device state from your service."
        };
        return buildAlexaResponse(header, payload, endpointId);
    }

    const properties = [];
    const currentTime = new Date().toISOString(); // 实际应是您设备数据的时间戳

    // 处理 Alexa.TemperatureSensor 接口
    const temperatureValue = deviceData.temperature;
    const temperatureUnit = deviceData.unit || "CELSIUS"; // 默认为摄氏度
    if (temperatureValue !== undefined && temperatureValue !== null) {
        properties.push({
            namespace: "Alexa.TemperatureSensor",
            name: "temperature",
            value: {
                value: temperatureValue,
                unit: temperatureUnit
            },
            timeOfSample: currentTime, 
            uncertaintyInMilliseconds: 0
        });
    }

    // 处理 Alexa.RangeController 用于湿度
    const humidityValue = deviceData.humidity;
    if (humidityValue !== undefined && humidityValue !== null) {
        properties.push({
            namespace: "Alexa.RangeController",
            name: "rangeValue",
            instance: "humidity", // 对应 Discovery 中设置的 'instance'
            value: humidityValue,
            timeOfSample: currentTime,
            uncertaintyInMilliseconds: 0
        });
    }

    // 处理 Alexa.EndpointHealth 接口
    const connectivityStatus = deviceData.connectivity || "OK";
    properties.push({
        namespace: "Alexa.EndpointHealth",
        name: "connectivity",
        value: {
            value: connectivityStatus
        },
        timeOfSample: currentTime,
        uncertaintyInMilliseconds: 0
    });

    const header = {
        namespace: "Alexa",
        name: "StateReport",
        payloadVersion: "3",
        messageId: event.directive.header.messageId
    };
    const payload = {
        context: {
            properties: properties
        },
        scope: {
            type: "BearerToken",
            token: bearerToken // 将用户的 access token 返回以供 Alexa 使用
        }
    };
    
    return buildAlexaResponse(header, payload, endpointId);
}

/**
 * Lambda 的入口点。
 * 根据 Alexa 请求的命名空间和名称将其分派给相应的处理函数。
 * @param {object} event Lambda 事件对象。
 * @param {object} context Lambda 上下文对象。
 * @returns {Promise<object>} Alexa 响应对象。
 */
exports.handler = async (event, context) => {
    logger.info(`Received Alexa Request: ${JSON.stringify(event, null, 2)}`);

    try {
        const namespace = event.directive.header.namespace;
        const name = event.directive.header.name;

        if (namespace === "Alexa.Discovery" && name === "Discover") {
            return await handleDiscoveryRequest(event);
        } else if (namespace === "Alexa" && name === "ReportState") {
            return await handleReportStateRequest(event);
        } else {
            logger.warn(`Unsupported directive: Namespace=${namespace}, Name=${name}`);
            const header = {
                namespace: "Alexa",
                name: "ErrorResponse",
                payloadVersion: "3",
                messageId: event.directive.header.messageId
            };
            const payload = {
                type: "INVALID_DIRECTIVE",
                message: `Unsupported directive: ${namespace}.${name}`
            };
            return buildAlexaResponse(header, payload);
        }

    } catch (e) {
        logger.error(`Error processing request: ${e.message}`, e);
        const header = {
            namespace: "Alexa",
            name: "ErrorResponse",
            payloadVersion: "3",
            messageId: event.directive.header.messageId
        };
        const payload = {
            type: "INTERNAL_ERROR",
            message: `An unexpected error occurred: ${e.message}`
        };
        return buildAlexaResponse(header, payload);
    }
};
```
