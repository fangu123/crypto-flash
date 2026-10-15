# iOS 未签名 IPA 构建

工作流文件为 `.github/workflows/ios-unsigned.yml`，使用 GitHub 的 macOS runner 构建当前项目的 H5 资源，并通过 Capacitor 封装为 iOS App。向 `main` 或 `master` 分支推送时自动运行，也可以在 GitHub Actions 页面选择 `Build unsigned iOS IPA` 并点击 `Run workflow` 手动运行。

成功后从该次 workflow run 的 **Artifacts** 下载 `CryptoFlash-unsigned-ipa`。其中的 `CryptoFlash-unsigned.ipa` 包含 `Payload/App.app`，不含 Apple 签名证书或 provisioning profile，不能直接安装；后续需要使用与你的 Bundle ID 匹配的证书和描述文件签名。当前 iOS 工程的 Bundle ID 为 `com.yourname.cryptoflash`，正式签名前应改成你自己的 App ID，并让签名使用相同 Bundle ID。

生产 H5 构建从 `.env.real` 读取 `VITE_API_BASE_URL`。请确保该文件已提交且不包含密钥，或通过仓库配置提供此环境变量。签名相关材料不需要添加到 GitHub。
