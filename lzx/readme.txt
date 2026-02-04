学习golang语言

wafw00f https://example.org
    wafw00f 是一个WAF 指纹识别工具：它会向目标站点发一些探测请求，观察响应头/状态码/重定向/特征页面等，来判断站点前面是否有 Web Application Firewall，以及是哪一家。
    ASCII 画 + WAFW00F : v2.4.2
        只是工具的 banner 和版本信息。
    [*] Checking https://example.org
        表示开始对该 URL 做探测。
    [+] The site https://example.org is behind Cloudflare (Cloudflare Inc.) WAF.
        结论：这个站点前面有 Cloudflare 的 WAF/防护。
        注意它说的是 “behind Cloudflare WAF”——也就是你的请求先到 Cloudflare，再转发到源站。
    [~] Number of requests: 2
        wafw00f 为了得到这个结论，总共只发了 2 次请求（说明这个目标的特征很明显，所以很快识别出来）。

cd tests 进入tests文件夹，然后输入pytest 即可运行测试

WAF (Web Application Firewall) Web 应用防火墙   
     WAF 是 Web 应用防火墙，用于保护 Web 应用免受恶意攻击。
     WAF 的作用是：
     1. 拦截非法请求
     2. 防止 SQL 注入攻击
     3. 防止 XSS 攻击
     4. 防止拒绝服务攻击
     5. 防止信息泄露攻击
     6. 防止跨站脚本攻击
     7. 防止跨站请求伪造攻击
     8. 防止拒绝服务攻击
     9. 防止信息泄露攻击

WAF 指纹（WAF Fingerprinting）
    WAF产品运行时留下的特征信息，用于识别网站使用的是哪种WAF

wafw00f 这个项目的作用：识别目标网站使用的是哪种 WAF，帮助安全测试人员了解防护措施，或帮助攻击者绕过特定 WAF（这也是为什么需要隐藏 WAF 指纹）。
wafw00f 识别 WAF 的原理：
    1. 发送一些请求，观察响应头/状态码/重定向/特征页面等，来判断站点前面是否有 WAF，以及是哪一家。
    2. 尝试绕过 WAF，看看是否能访问成功

    通过发送正常和恶意请求，收集相应中的这些指纹特征，与插件库中的 100+ 种 WAF 指纹进行匹配，从而判断 WAF 的类型。

隐藏 WAF 指纹的重要性：避免攻击者针对特定 WAF 的已知绕过技术进行攻击

tests文件夹结构
    1. conftest.py 测试配置和共享工具
    2. test_detection.py WAF 检测集成测试
        测试能否正确识别各种 WAF 产品
    3. test_evillib.py 测试底层HTTP请求功能模块
        测试内容:
            1.超时设置
            2.请求头管理
            3.请求计数器
            4.相应内容读取
    4. text_matching.py 指纹匹配逻辑测试
        测试各种 WAF 指纹匹配函数
            HTTP 头匹配
            Cookie 匹配
            状态码匹配
            状态原因短语匹配
    5. test_manager.py 插件管理器测试
        测试插件加载和管理系统
        测试内容：
            1. 插件能否正确加载
            2. 每个插件是否有必须的属性
            2. 每个插件是否又is_waf()函数
            3. 已知插件是否存在
    6. init.py 
        空文件，让python识别tests为一个包
        允许测试模块之间互相导入
    7.pycache 
        python自动生成的编译缓存
        提升重复运行测试的速度
        可以删除，会重新生成