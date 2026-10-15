import { defineConfig } from 'vite'
import uni from '@dcloudio/vite-plugin-uni'
export default defineConfig({ base: process.env.CAPACITOR_BUILD === '1' ? './' : '/', plugins: [uni()] })
