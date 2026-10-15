import { Capacitor } from '@capacitor/core'
import { Browser } from '@capacitor/browser'

const callbackBase = 'cryptoflash://oauth'

export function oauthRedirectUrl(provider: 'google' | 'microsoft', action = '') {
  if (Capacitor.isNativePlatform()) {
    const query = new URLSearchParams({ oauthProvider: provider })
    if (action) query.set('oauthAction', action)
    return `${callbackBase}?${query}`
  }
  const query = new URLSearchParams({ oauthProvider: provider })
  if (action) query.set('oauthAction', action)
  return `${location.origin}${location.pathname}?${query}`
}

export async function openOAuthUrl(url: string) {
  if (Capacitor.isNativePlatform()) await Browser.open({ url })
  else window.location.href = url
}

export function oauthCallbackRoute(url: string) {
  const parsed = new URL(url)
  if (parsed.protocol !== 'cryptoflash:' || parsed.hostname !== 'oauth') return ''
  const provider = parsed.searchParams.get('oauthProvider')
  if (provider !== 'google' && provider !== 'microsoft') return ''
  const query = new URLSearchParams(parsed.search)
  return `/pages/index/index?${query}`
}
