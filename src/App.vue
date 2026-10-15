<script setup lang="ts">
import { onLaunch } from '@dcloudio/uni-app'
import { useSessionStore } from './store/session'
import { restoreLocale } from './locales'
import { Capacitor } from '@capacitor/core'
import { App as NativeApp } from '@capacitor/app'
import { Browser } from '@capacitor/browser'
import { oauthCallbackRoute } from './native/oauth'
onLaunch(() => {
  restoreLocale(); useSessionStore().restore()
  if (Capacitor.isNativePlatform()) {
    const handleUrl = (url: string) => {
      const route = oauthCallbackRoute(url)
      if (route) { Browser.close(); uni.reLaunch({ url: route }) }
    }
    NativeApp.addListener('appUrlOpen', ({ url }) => handleUrl(url))
    NativeApp.getLaunchUrl().then(result => { if (result?.url) handleUrl(result.url) })
  }
})
</script>
<style>page{background:var(--cf-bg);color:var(--cf-text);font-family:Inter,"Noto Sans SC",sans-serif}view,text,button,input{box-sizing:border-box}button{border:0}</style>
