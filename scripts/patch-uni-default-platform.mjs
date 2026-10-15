// 将 @dcloudio/vite-plugin-uni 的 CLI 默认平台从 h5 改为 app。
// 背景：HBuilderX 直接执行 `uni.js build`（不带 -p），CLI 内部把未指定平台兜底为 h5，
// 导致 App 打包拿到的却是 H5 产物。此脚本由 postinstall 触发，npm install 后自动重打补丁。
import fs from 'node:fs'
import path from 'node:path'

const target = path.resolve(
  'node_modules/@dcloudio/vite-plugin-uni/dist/cli/index.js'
)
const OLD = "default: 'h5',"
const NEW = "default: 'app',"

if (!fs.existsSync(target)) {
  console.error(`[patch-uni] 未找到 ${target}，请先执行 npm install 后再试。`)
  process.exit(1)
}

const source = fs.readFileSync(target, 'utf8')
if (source.includes(NEW)) {
  console.log('[patch-uni] 平台默认值已是 app，无需重复打补丁。')
  process.exit(0)
}
if (!source.includes(OLD)) {
  console.error(
    `[patch-uni] 未找到目标 "${OLD}"，@dcloudio/vite-plugin-uni 版本可能已变化，请人工核对。`
  )
  process.exit(1)
}

fs.writeFileSync(target, source.replace(OLD, NEW))
console.log('[patch-uni] 已修改 CLI 默认平台：h5 → app。')
