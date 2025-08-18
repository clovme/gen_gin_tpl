var __assign = (this && this.__assign) || function () {
    __assign = Object.assign || function(t) {
        for (var s, i = 1, n = arguments.length; i < n; i++) {
            s = arguments[i];
            for (var p in s) if (Object.prototype.hasOwnProperty.call(s, p))
                t[p] = s[p];
        }
        return t;
    };
    return __assign.apply(this, arguments);
};
import { loadEnv } from 'vite';
import path from 'path';
import XEUtils from 'xe-utils';
import vue from '@vitejs/plugin-vue';
import vueJsx from '@vitejs/plugin-vue-jsx';
import { createHtmlPlugin } from 'vite-plugin-html';
import externalGlobals from 'rollup-plugin-external-globals';
// https://vitejs.dev/config/
export default (function (_a) {
    var mode = _a.mode;
    var env = loadEnv(mode, process.cwd(), 'VITE_');
    return {
        base: '/me',
        plugins: [
            vue(),
            vueJsx(),
            createHtmlPlugin({
                inject: {
                    data: __assign({ VITE_APP_BUILD_TIME: XEUtils.toDateString(new Date(), 'yyyy-MM-dd HH:mm:ss') }, env)
                }
            })
        ],
        resolve: {
            alias: {
                '@': path.join(__dirname, './src')
            },
            extensions: ['.js', '.vue', '.json', '.ts', '.tsx']
        },
        server: {
            port: 8084
        },
        build: {
            assetsDir: '.',
            outDir: env.VITE_OUT_DIR || 'dist',
            rollupOptions: {
                // 不打包依赖
                external: ['echarts'],
                plugins: [
                    // 不打包依赖映射的对象
                    externalGlobals({
                        echarts: 'echarts'
                    })
                ]
            }
        }
    };
});
