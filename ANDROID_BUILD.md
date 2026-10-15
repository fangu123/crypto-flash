# Android 终端构建

使用 JDK 17、Android SDK 35 和 Node.js 18 以上。首次执行 `npm ci` 安装依赖，然后在项目根目录运行：

```powershell
npm run apk:debug
```

命令会构建正式后端的 H5 资源、同步到 Capacitor 工程，并用 Gradle 生成 `android/app/build/outputs/apk/debug/app-debug.apk`。后端地址来自 `.env.real` 的 `VITE_API_BASE_URL`。可用 `adb install -r android/app/build/outputs/apk/debug/app-debug.apk` 安装到已连接设备。

`npm run apk:release` 生成 `android/app/build/outputs/apk/release/app-release-unsigned.apk`。此文件尚未签名，不能直接发布；正式发布需使用自己的 Android 发布密钥签名。请妥善保存密钥，后续更新必须使用同一密钥。

Google 和 Microsoft 授权使用系统浏览器，回调地址为 `cryptoflash://oauth?oauthProvider=google` 或 `cryptoflash://oauth?oauthProvider=microsoft`。服务端 OAuth 配置必须允许授权完成后跳回该地址。当前后端授权接口返回 `code: 500`，需服务端修复后在真机上完成第三方登录验证。

