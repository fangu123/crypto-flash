import { spawnSync } from 'node:child_process'
import process from 'node:process'

const command = process.platform === 'win32' ? 'npm.cmd' : 'npm'
const result = spawnSync(command, ['run', 'build:prod:h5'], {
  stdio: 'inherit',
  shell: process.platform === 'win32',
  env: { ...process.env, CAPACITOR_BUILD: '1' }
})
process.exit(result.status ?? 1)
