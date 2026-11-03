import { defineConfig } from 'vite'
import uni from '@dcloudio/vite-plugin-uni'

const removeDcloudShadowPreload = {
  name: 'remove-dcloud-shadow-preload',
  apply: 'build' as const,
  generateBundle(_options: unknown, bundle: Record<string, any>) {
    const preloadRule = /body:after\{position:fixed;content:"";left:-1000px;top:-1000px;animation:shadow-preload \.1s;animation-delay:3s\}@keyframes shadow-preload\{0%\{background-image:url\(https:\/\/cdn\.dcloud\.net\.cn\/img\/shadow-grey\.png\)\}to\{background-image:url\(https:\/\/cdn\.dcloud\.net\.cn\/img\/shadow-grey\.png\)\}\}/g
    for (const asset of Object.values(bundle)) {
      if (asset.type === 'asset' && asset.fileName.endsWith('.css') && typeof asset.source === 'string') {
        asset.source = asset.source.replace(preloadRule, '')
      }
    }
  }
}

export default defineConfig({ base: process.env.CAPACITOR_BUILD === '1' ? './' : '/', plugins: [uni(), removeDcloudShadowPreload] })
