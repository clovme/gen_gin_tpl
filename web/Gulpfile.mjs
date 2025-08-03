import gulp from 'gulp';
import rename from 'gulp-rename';
import browserSync from 'browser-sync';
import terser from 'gulp-terser';
import babel from 'gulp-babel';
import sourcemaps from 'gulp-sourcemaps';
import dartSass from 'sass';
import gulpSass from 'gulp-sass';
import postcss from 'gulp-postcss';
import autoprefixer from 'autoprefixer';
import cssnano from 'gulp-cssnano';
import {deleteAsync} from 'del';
import {createRequire} from 'module';
import clone from 'gulp-clone';

const require = createRequire(import.meta.url);
const revCollector = require('gulp-rev-collector');
const replace = require('gulp-replace');
import gulpIf from 'gulp-if';

const {src, dest, watch, series, parallel} = gulp;
const sass = gulpSass(dartSass);
const bs = browserSync.create();

const conf = {
    dest: '../api/public/web',
    src: 'src',
    watch: {events: ['add', 'change', 'unlink'], queue: true},
};

const isDev = process.env.MODE === 'development';
if (isDev) {
    conf.dest = 'dist'
}

// 正则
const rmspaceOpen = /%}\s+\{\{/g;
const rmspaceClose = /}}\s+\{%/g;
const host = 'http://192.168.1.2:9527'

const hostReplace =()  => {
    if (isDev) {
        return host;
    } else {
        return '';
    }
}

// 删除文件
const cleanCss = () => deleteAsync([`${conf.dest}/assets/css`], {force: true});
const cleanJs = () => deleteAsync([`${conf.dest}/assets/js`], {force: true});
const cleanImg = () => deleteAsync([`${conf.dest}/assets/images`], {force: true});
const cleanHtml = () => deleteAsync([`${conf.dest}/*.html`], {force: true});
const cleanPlugins = () => deleteAsync([`${conf.dest}/assets/plugins`], {force: true});

// 处理 HTML
const html = series(cleanHtml, () =>
    src(`${conf.src}/*.html`)
        .pipe(revCollector({
            replaceReved: true,
            dirReplacements: {
                js: conf.dest,
                css: conf.dest,
            },
        }))
        .pipe(replace('__LOCALHOST__', hostReplace()))
        .pipe(replace('.scss', '.css'))
        .pipe(replace(rmspaceOpen, '%}{{'))
        .pipe(replace(rmspaceClose, '}}{%'))
        .pipe(dest(conf.dest))
        .on('end', () => {
            console.log('HTML写入完成，磁盘已 flush');
        })
        .pipe(bs.stream())
);

// 编译SCSS → CSS
const css = series(cleanCss, () =>
    src(`${conf.src}/assets/css/**/*.scss`)
        .pipe(gulpIf(isDev, sourcemaps.init()))
        .pipe(sass().on('error', sass.logError))
        .pipe(replace(/url\(([^)]+)\.scss\)/g, "url($1.css)"))
        .pipe(gulpIf(isDev, dest(`${conf.dest}/assets/css/`)))
        .pipe(postcss([autoprefixer()]))
        .pipe(gulpIf(!isDev, cssnano()))
        // .pipe(gulpIf(!isDev, rename({suffix: '.min'})))
        .pipe(gulpIf(isDev, sourcemaps.write('./maps', {addComment: false})))
        .pipe(gulpIf(!isDev, dest(`${conf.dest}/assets/css/`)))
        .on('end', () => {
            console.log('CSS写入完成，磁盘已 flush');
        })
        .pipe(bs.stream())
);

// 编译JS
const js = series(cleanJs, () => {
    const source = src(`${conf.src}/assets/js/**/*.js`)
        .pipe(replace('__LOCALHOST__', hostReplace()))
        .pipe(gulpIf(isDev, sourcemaps.init()));
    const sourceClone = source.pipe(clone());

    // 原文件
    source.pipe(gulpIf(isDev, dest(`${conf.dest}/assets/js/`)));

    // 压缩文件
    sourceClone
        .pipe(gulpIf(!isDev, babel()))
        .pipe(gulpIf(!isDev, terser()))
        // .pipe(gulpIf(!isDev, rename({suffix: '.min'})))
        .pipe(gulpIf(isDev, sourcemaps.write('./maps', {addComment: false})))
        .pipe(gulpIf(!isDev, dest(`${conf.dest}/assets/js/`)))
        .on('end', () => {
            console.log('JS写入完成，磁盘已 flush');
        })
        .pipe(bs.stream());

    return sourceClone;
});

// 复制图片
const images = series(cleanImg, () =>
    src(`${conf.src}/assets/images/**`, {encoding: false})
        .pipe(dest(`${conf.dest}/assets/images/`))
        .on('end', () => {
            console.log('图片写入完成，磁盘已 flush');
        })
        .pipe(bs.stream())
);

// 复制插件文件
const plugins = series(cleanPlugins, () =>
    src(`${conf.src}/assets/plugins/**`, {encoding: false})
        .pipe(dest(`${conf.dest}/assets/plugins/`))
        .on('end', () => {
            console.log('插件写入完成，磁盘已 flush');
        })
        .pipe(bs.stream())
);

// 监听文件变化
const watchFiles = (cb) => {
    if (!isDev) {
        deleteAsync(['dist'], {force: true})
        console.log('Build 完成...')
        console.log(`Dist 目录：${conf.dest}`)
        cb()
        return
    }
    bs.init({
        server: {
            baseDir: conf.dest,
            livereload: true,
        },
        ui: false,
        open: false,
        notify: false,
    });

    watch(`${conf.src}/assets/js/**/*.js`, conf.watch, js);
    watch(`${conf.src}/assets/css/**/*.scss`, conf.watch, css);
    watch(`${conf.src}/*.html`, conf.watch, html);
    watch(`${conf.src}/assets/plugins/**/*.*`, conf.watch, plugins);
    watch(
        [
            `${conf.src}/assets/images/**/*.png`,
            `${conf.src}/assets/images/**/*.jpg`,
            `${conf.src}/assets/images/**/*.jpeg`,
            `${conf.src}/assets/images/**/*.gif`,
        ],
        conf.watch,
        images
    );

    cb();
};

export default series(parallel(js, css, images, plugins, html), watchFiles);
