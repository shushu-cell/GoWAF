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

    