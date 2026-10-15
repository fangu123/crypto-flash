import fs from 'node:fs'

const envFile = '.env.real'
const fileText = fs.existsSync(envFile) ? fs.readFileSync(envFile, 'utf8') : ''
const fileMatch = fileText.match(/^\s*VITE_API_BASE_URL\s*=\s*(.*)\s*$/m)
const configured = process.env.VITE_API_BASE_URL || fileMatch?.[1] || ''
const value = configured.replace(/^['"]|['"]$/g, '').replace(/\/$/, '')

if (!value || !/^https?:\/\//i.test(value)) {
  console.error('正式版构建已停止：请在 .env.real 或环境变量 VITE_API_BASE_URL 中配置 http(s) 后端地址。')
  process.exit(1)
}

console.log(`正式版后端地址：${value}`)
