<script lang="ts" setup>
import { ref } from 'vue'
import XEUtils from 'xe-utils'
import { getCopyrightConfig } from '@/api/initiate.ts'

const data = ref({
  name: '[{ .Title }]',
  startTime: XEUtils.toDateString(new Date(), 'yyyy'),
  endTime: XEUtils.toDateString(new Date(), 'yyyy')
})

getCopyrightConfig().then(result => {
  data.value = result.data
})
</script>

<template>
  <div class="page-footer">
    <span v-if="data.startTime != data.endTime">版权所有 ©{{ data.name }} {{ data.startTime }}-{{ data.endTime }}</span>
    <span v-else>版权所有 ©{{ data.name }} {{ data.endTime }}</span>
  </div>
</template>

<style lang="scss" scoped>
.page-footer {
  text-align: center;
  padding: 8px;
  background-color: var(--page-layout-background-color);
}
</style>
