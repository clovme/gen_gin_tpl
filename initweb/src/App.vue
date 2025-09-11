<script lang="ts" setup>
import UserLayout from '@/views/layout/UserLayout.vue'
import { onMounted, onUnmounted } from 'vue'
import { useAppStore } from '@/store/app.ts'
import { addLineLog } from '@/utils/log.ts'

const appStore = useAppStore()

let ws: WebSocket

onMounted(() => {
  if (import.meta.env.VITE_APP_BASE_API.includes('http')) {
    const appBase = import.meta.env.VITE_APP_BASE_API.split(':')

    const temp = window.location.host.split(':')
    temp[1] = appBase[appBase.length - 1]

    ws = new WebSocket(`ws://${temp.join(':')}/logs`)
  } else {
    ws = new WebSocket(`ws://${window.location.host}/logs`)
  }

  appStore.setIsMaximize(true)
  ws.onopen = () => {
    addLineLog('WebSocket 已正常链接日志后台', true)
    appStore.setIsMaximize(false)
  }
  ws.onmessage = (event) => { addLineLog(event.data) }
  ws.onerror = (err) => { addLineLog(`WebSocket 链接错误，请检查网络是否连通或者其他异常，错误信息：${err}`, true) }
  ws.onclose = () => { addLineLog('WebSocket 已关闭，程序已退出，请重新启动程序进入正式模式！', true) }
})

onUnmounted(() => {
  ws && ws.close()
})

window.addEventListener('resize', () => {
  if (appStore.isMaximize) {
    const height = window.innerHeight - 50
    const logContainer = document.querySelector('.log-container') as HTMLElement
    logContainer.style.height = `${height}px`
  }
})
</script>

<template>
  <div class="app-container">
    <UserLayout />
    <div class="log-container" style="padding: 0; height: 0;">
      <div class="log-container-content"></div>
    </div>
  </div>
</template>

<style scoped lang="scss">
.app-container {
  .log-container {
    position: fixed;
    bottom: 0;
    left: 0;
    width: 100%;
    z-index: -1;
    background-color: #1e1f22;
    transition: all 0.3s ease;
    overflow: auto;
    padding-bottom: 10px !important;

    .log-container-content {
      width: 100%;
      color: #67c23a;
      white-space: nowrap;
    }

    :deep(pre) {
      font-family: var(--vxe-ui-font-family), serif;

      a {
        color: #409eff;
        text-decoration: none;

        &:hover {
          color: #409eff;
          text-decoration: underline;
        }
      }
    }
  }
}
</style>
