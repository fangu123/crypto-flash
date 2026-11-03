<script setup lang="ts">
import { onLaunch } from '@dcloudio/uni-app'
import { useSessionStore } from './store/session'
import { restoreLocale } from './locales'
import { Capacitor } from '@capacitor/core'
import { App as NativeApp } from '@capacitor/app'
import { Browser } from '@capacitor/browser'
import { oauthCallbackRoute } from './native/oauth'
import { recordAppLog } from './api/diagnostics'
onLaunch(() => {
  restoreLocale(); useSessionStore().restore()
  if (Capacitor.isNativePlatform()) {
    const handleUrl = (url: string) => {
      const route = oauthCallbackRoute(url)
      if (route) { Browser.close(); uni.reLaunch({ url: route }) }
    }
    NativeApp.addListener('appUrlOpen', ({ url }) => handleUrl(url))
    NativeApp.getLaunchUrl().then(result => { if (result?.url) handleUrl(result.url) }).catch((error: any) => {
      recordAppLog('native.getLaunchUrl', 'failure', error?.message || '读取启动链接失败')
    })
  }
})
</script>
<style lang="scss">
@import './styles/ui-scale.scss';

page{background:var(--cf-bg);color:var(--cf-text);font-family:Inter,"Noto Sans SC",sans-serif}
view,text,button,input{box-sizing:border-box}
button{display:flex!important;align-items:center!important;justify-content:center!important;border:0;line-height:1.2!important}
.search-symbol{display:inline-flex;align-items:center;justify-content:center;flex:0 0 38rpx;width:38rpx;height:38rpx}
.search-symbol svg{display:block;width:38rpx;height:38rpx;fill:none;stroke:currentColor;stroke-width:1.8;stroke-linecap:round;stroke-linejoin:round}
.home-content .search-symbol{flex-basis:34rpx;width:34rpx;height:34rpx}
.home-content .search-symbol svg{width:34rpx;height:34rpx}
.home-content svg.search-symbol{fill:none;stroke:currentColor;stroke-width:1.8;stroke-linecap:round;stroke-linejoin:round}
.search-bar .search-symbol{flex-basis:52rpx;width:52rpx;height:64rpx}
.search-bar .search-symbol svg{width:40rpx;height:40rpx}
.inner-nav .back svg,.detail-nav .nav-back svg,.point-nav .nav-back svg,.notifications-nav .back svg{display:block;width:56rpx;height:56rpx;fill:none;stroke:currentColor;stroke-width:2.2;stroke-linecap:round;stroke-linejoin:round}
.detail-nav .nav-share svg,.detail-nav .nav-bookmark svg,.point-nav .nav-share svg,.point-nav .nav-bookmark svg{display:block;width:44rpx;height:44rpx;fill:none;stroke:currentColor;stroke-width:1.8;stroke-linecap:round;stroke-linejoin:round}
.detail-actions button:nth-child(2) svg{fill:none;stroke:currentColor}
.unlock-icon{width:38rpx;height:38rpx;fill:none;stroke:currentColor;stroke-width:1.8;stroke-linecap:round;stroke-linejoin:round}
</style>
