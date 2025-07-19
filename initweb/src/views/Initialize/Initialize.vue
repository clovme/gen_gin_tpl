<script lang="ts" setup>
import PageView from '@/views/layout/PageView.vue'
import { ref, watch } from 'vue'
import { VxeUI, VxeFormProps, VxeFormListeners } from 'vxe-pc-ui'
import { getFormConfig, postFormConfig } from '@/api/initiate.ts'
import { FormDataVO } from '@/utils/interfaces/form.ts'
import { useAppStore } from '@/store/app.ts'

const appStore = useAppStore()
let configData = {}
const watchType = {}
const formOptions = ref<VxeFormProps<FormDataVO>>()

function filterVisibleItems (items, data) {
  return items.filter(group => {
    if (group.showWhen) {
      const {
        field,
        value
      } = group.showWhen
      watchType[field] = value
      return data[field] === value
    }
    return true
  })
}

getFormConfig().then((config) => {
  VxeUI.loading.open({
    text: '正则加载数据，请稍后...'
  })
  configData = JSON.parse(JSON.stringify(config.data.items))
  formOptions.value = config.data as VxeFormProps<FormDataVO>

  // 初始过滤
  formOptions.value.items = filterVisibleItems(formOptions.value.items, formOptions.value.data)

  for (const watchTypeKey in watchType) {
    // @ts-ignore
    watch(() => formOptions.value?.data[watchTypeKey], () => {
      if (formOptions.value) {
        formOptions.value.items = filterVisibleItems(configData, formOptions.value.data)
      }
    })
  }
  VxeUI.loading.close()
})

const formEvents: VxeFormListeners<FormDataVO> = {
  submit () {
    appStore.setIsMaximize(true)
    VxeUI.loading.open({
      text: '正在检测配置项，请稍后...'
    })
    postFormConfig(formOptions.value?.data as FormDataVO).then(() => {
      VxeUI.loading.open({ text: '初始化成功，请稍后...' })
    }).catch(() => {
      VxeUI.loading.close()
    })
  }
}
</script>

<template>
  <PageView class="initialize-page" :padding="true" :background="true">
    <div>
      <vxe-form class-name="initialize-page-form" v-bind="formOptions" v-on="formEvents"/>
    </div>
  </PageView>
</template>

<style lang="scss" scoped>
.initialize-page {
  display: flex;
  justify-content: center;

  .initialize-page-form {
    width: var(--min-width);
  }

  :deep(.grid-lines-5-box) {
    display: grid;
    grid-template-columns: 1fr 1fr 1fr 1fr 1fr;
    grid-gap: 5px;
  }

  :deep(.grid-lines-4-box) {
    display: grid;
    grid-template-columns: 1fr 1fr 1fr 1fr;
    grid-gap: 5px;
  }

  :deep(.grid-lines-3-box) {
    display: grid;
    grid-template-columns: 1fr 1fr 1fr;
    grid-gap: 5px;
  }
}
</style>
