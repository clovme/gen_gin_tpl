import { createApp } from 'vue'
import App from './App.vue'
import pinia from './store'

import VxeTable from 'vxe-table'
import 'vxe-table/lib/style.css'

import VxeUI from 'vxe-pc-ui'
import 'vxe-pc-ui/lib/style.css'

import './style/style.scss'

import VxeUIPluginValidator from '@vxe-ui/plugin-validator'

VxeUI.use(VxeUIPluginValidator)

const app = createApp(App)

app.use(pinia).use(VxeUI).use(VxeTable).mount('#app')
