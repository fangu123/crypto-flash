import { spawnSync } from 'node:child_process'
import { existsSync } from 'node:fs'
import { join, resolve } from 'node:path'
import process from 'node:process'

const variant = process.argv[2]
if (variant !== 'debug' && variant !== 'release') throw new Error('Use debug or release')
const wrapper = resolve('android', process.platform === 'win32' ? 'gradlew.bat' : 'gradlew')
const sdk = process.env.ANDROID_HOME || process.env.ANDROID_SDK_ROOT || (process.platform === 'win32' ? join(process.env.LOCALAPPDATA || '', 'Android', 'Sdk') : '')
if (!sdk || !existsSync(sdk)) throw new Error('Android SDK not found. Set ANDROID_HOME.')
const result = spawnSync(wrapper, [`assemble${variant[0].toUpperCase()}${variant.slice(1)}`], {
  cwd: resolve('android'),
  stdio: 'inherit',
  shell: process.platform === 'win32',
  env: { ...process.env, ANDROID_HOME: sdk }
})
if (result.status !== 0) process.exit(result.status ?? 1)
const apkName = variant === 'release' ? 'app-release-unsigned.apk' : 'app-debug.apk'
const apk = resolve('android', 'app', 'build', 'outputs', 'apk', variant, apkName)
if (!existsSync(apk)) throw new Error(`APK not found: ${apk}`)
console.log(`APK: ${apk}`)
