/*************************************************** 全局工具方法 ************************************************************
 * @function {toast} - 全局 toast 方法
 * @function {currentTime} - 获取服务器时间戳
 * @function {storage} - localStorage存储工具
 * @function {validateForm} - 表单校验工具
 * @function {onSwitchCaptcha} - 验证码切换工具
 **************************************************** 全局工具方法 ***********************************************************/
const utils = {
    /**
     * 全局 toast 方法
     */
    toast: {
        /**
         * toast 模块
         * @returns {{msg: toast, error: (function(*): void), warning: (function(*): void), success: (function(*): void)}}
         */
        newToast: () => {
            const animationTime = 390;
            let options, defaults, container, icon, layout, popStyle, positions, close, headStyle;

            const _Toast = function (template, style) {
                this.defaults = {
                    template: null,
                    style: 'info',
                    autoclose: 3000,
                    position: 'top-center',
                    icon: true,
                    group: "toast-message-group",
                    onOpen: false,
                    onClose: false
                };

                headStyle = 'toast-pop--head-style'
                defaults = extend(this.defaults, toast.defaults);
                if (typeof template === 'string' || typeof style === 'string') {
                    options = {template: template, style: style || defaults.style};
                } else if (typeof template === 'object') {
                    options = template;
                } else {
                    console.error('Invalid arguments.');
                    return false;
                }
                this.opt = extend(defaults, options);
                if ($(`toast-pop--${this.opt.group}`)) {
                    this.remove($('toast-pop--' + this.opt.group));
                }
                this.open();
            };

            _Toast.prototype.create = function (template) {
                // 设置消息弹框样式
                const head = document.head || document.querySelector('head');
                if (!$(headStyle)) {
                    const style = document.createElement('style');
                    style.id = headStyle;
                    style.innerHTML = `@charset "UTF-8";.toast-pop-container{z-index:9999;position:fixed}.toast-pop-container,.toast-pop-container *,.toast-pop-container :after,.toast-pop-container :before{-webkit-box-sizing:border-box;-moz-box-sizing:border-box;box-sizing:border-box}.toast-pop--top-left{top:0;left:0}.toast-pop--top-left .toast-pop{-webkit-transform-origin:0 0;-ms-transform-origin:0 0;transform-origin:0 0}.toast-pop--top-center{top:0;left:50%;-webkit-transform:translateX(-50%);-ms-transform:translateX(-50%);transform:translateX(-50%)}.toast-pop--top-center .toast-pop{-webkit-transform-origin:50% 0;-ms-transform-origin:50% 0;transform-origin:50% 0}.toast-pop--top-right{top:0;right:0}.toast-pop--top-right .toast-pop{-webkit-transform-origin:100% 0;-ms-transform-origin:100% 0;transform-origin:100% 0}.toast-pop--center{top:50%;left:50%;-webkit-transform:translate3d(-50%,-50%,0);transform:translate3d(-50%,-50%,0)}.toast-pop--center .toast-pop{-webkit-transform-origin:50% 0;-ms-transform-origin:50% 0;transform-origin:50% 0}.toast-pop--bottom-left{bottom:0;left:0}.toast-pop--bottom-left .toast-pop{-webkit-transform-origin:0 100%;-ms-transform-origin:0 100%;transform-origin:0 100%}.toast-pop--bottom-center{bottom:0;left:50%;-webkit-transform:translateX(-50%);-ms-transform:translateX(-50%);transform:translateX(-50%)}.toast-pop--bottom-center .toast-pop{-webkit-transform-origin:50% 100%;-ms-transform-origin:50% 100%;transform-origin:50% 100%}.toast-pop--bottom-right{bottom:0;right:0}.toast-pop--bottom-right .toast-pop{-webkit-transform-origin:100% 100%;-ms-transform-origin:100% 100%;transform-origin:100% 100%}@media screen and (max-width:30em){.toast-pop--bottom-center,.toast-pop--bottom-left,.toast-pop--bottom-right,.toast-pop--top-center,.toast-pop--top-left,.toast-pop--top-right{top:auto;bottom:0;left:0;right:0;margin-left:0;-webkit-transform:translateX(0);-ms-transform:translateX(0);transform:translateX(0)}.toast-pop--bottom-center .toast-pop,.toast-pop--bottom-left .toast-pop,.toast-pop--bottom-right .toast-pop,.toast-pop--top-center .toast-pop,.toast-pop--top-left .toast-pop,.toast-pop--top-right .toast-pop{-webkit-transform-origin:50% 100%;-ms-transform-origin:50% 100%;transform-origin:50% 100%}.toast-pop{border-bottom:1px solid rgba(0,0,0,.15)}}.toast-pop{font-size:14px;-webkit-transform:translateZ(0);transform:translateZ(0);display:-webkit-box;display:-webkit-flex;display:-moz-box;display:-ms-flexbox;display:flex;-webkit-box-align:center;-webkit-align-items:center;-moz-box-align:center;-ms-flex-align:center;align-items:center}@media screen and (min-width:30em){.toast-pop{border-radius:2px;margin:.7em}}.toast-pop--error,.toast-pop--info,.toast-pop--success,.toast-pop--warning{color:#fff;background-color:#454a56}@-webkit-keyframes a{0%{-webkit-transform:scale(.2);transform:scale(.2)}95%{-webkit-transform:scale(1.1);transform:scale(1.1)}to{-webkit-transform:scale(1);transform:scale(1)}}@keyframes a{0%{-webkit-transform:scale(.2);transform:scale(.2)}95%{-webkit-transform:scale(1.1);transform:scale(1.1)}to{-webkit-transform:scale(1);transform:scale(1)}}@-webkit-keyframes b{0%{opacity:1;-webkit-transform:scale(1);transform:scale(1)}20%{-webkit-transform:scale(1.1);transform:scale(1.1)}to{opacity:0;-webkit-transform:scale(0);transform:scale(0)}}@keyframes b{0%{opacity:1;-webkit-transform:scale(1);transform:scale(1)}20%{-webkit-transform:scale(1.1);transform:scale(1.1)}to{opacity:0;-webkit-transform:scale(0);transform:scale(0)}}.toast-pop--out{-webkit-animation:b .4s ease-in-out;animation:b .4s ease-in-out}.toast--in{-webkit-animation:a .4s ease-in-out;animation:a .4s ease-in-out}.toast-pop-body{-webkit-box-flex:1;-webkit-flex:1;-moz-box-flex:1;-ms-flex:1;flex:1;padding-top:1px}.toast-pop-body p{margin:0}.toast-pop-body a{color:#fff;text-decoration:underline}.toast-pop-body a:hover{color:hsla(0,0%,100%,.8);text-decoration:none}.toast-pop-title{margin-top:0;margin-bottom:.25em;color:#fff}.toast-pop-close{height:32px;width:32px;padding-top:7px;padding-right:1px;font-size:22px;font-weight:700;text-align:center;line-height:.8;color:#fff;opacity:.5}.toast-pop-close:hover{opacity:.7;cursor:pointer}.toast-pop-icon{position:relative;margin:7px;width:30px;height:30px;border-radius:50%;-webkit-animation:a .4s .4s ease-in-out;animation:a .4s .4s ease-in-out}.toast-pop-icon:after,.toast-pop-icon:before{content:"";position:absolute;display:block}.toast-pop-icon--error,.toast-pop-icon--info{border:2px solid #3a95ed}.toast-pop-icon--error:before,.toast-pop-icon--info:before{top:5px;left:11px;width:4px;height:4px;background-color:#3a95ed}.toast-pop-icon--error:after,.toast-pop-icon--info:after{top:12px;left:11px;width:4px;height:9px;background-color:#3a95ed}.toast-pop-icon--error{border-color:#ff5656}.toast-pop-icon--error:before{top:16px;background-color:#ff5656}.toast-pop-icon--error:after{top:5px;background-color:#ff5656}.toast-pop-icon--success{border:2px solid #2ecc54}.toast-pop-icon--success:before{top:7px;left:7px;width:13px;height:8px;border-bottom:3px solid #2ecc54;border-left:3px solid #2ecc54;-webkit-transform:rotate(-45deg);-ms-transform:rotate(-45deg);transform:rotate(-45deg)}.toast-pop-icon--warning{border:2px solid #fcd000}.toast-pop-icon--warning:before{top:7px;left:7px;width:0;height:0;border-style:solid;border-color:transparent transparent #fcd000;border-width:0 6px 10px}`
                    head.appendChild(style);
                }

                container = $(this.getPosition('toast-pop--', this.opt.position));
                icon = (!this.opt.icon) ? '' : `<i class="toast-pop-icon ${this.getStyle('toast-pop-icon--', this.opt.style)}"></i>`;
                layout = `${icon}<div class="toast-pop-body">${template}</div><div class="toast-pop-close" data-pop-toast="close" aria-label="Close">&times;</div>`;

                if (!container) {
                    this.popContainer = document.createElement('div');
                    this.popContainer.setAttribute('class', `toast-pop-container ${this.getPosition('toast-pop--', this.opt.position)}`);
                    this.popContainer.setAttribute('id', this.getPosition('toast-pop--', this.opt.position));
                    document.body.appendChild(this.popContainer);
                    container = $(this.getPosition('toast-pop--', this.opt.position));
                }
                this.pop = document.createElement('div');
                this.pop.setAttribute('class', `toast-pop toast-pop--out toast--in ${this.getStyle('toast-pop--', this.opt.style)}`);
                if (this.opt.group && typeof this.opt.group === 'string') {
                    this.pop.setAttribute('id', 'toast-pop--' + this.opt.group);
                }
                this.pop.setAttribute('role', 'alert');
                this.pop.innerHTML = layout;
                container.appendChild(this.pop);
            };

            _Toast.prototype.getStyle = function (sufix, arg) {
                popStyle = {
                    'success': 'success',
                    'error': 'error',
                    'warning': 'warning'
                };
                return sufix + (popStyle[arg] || 'info');
            };

            _Toast.prototype.getPosition = function (sufix, position) {
                positions = {
                    'top-left': 'top-left',
                    'top-center': 'top-center',
                    'top-right': 'top-right',
                    'bottom-left': 'bottom-left',
                    'bottom-center': 'bottom-center',
                    'bottom-right': 'bottom-right'
                };
                return sufix + (positions[position] || 'top-right');
            };

            _Toast.prototype.open = function () {
                this.create(this.opt.template);
                if (this.opt.onOpen) {
                    this.opt.onOpen();
                }
                this.close();
            };

            _Toast.prototype.close = function () {
                if (this.opt.autoclose && typeof this.opt.autoclose === 'number') {
                    this.autocloseTimer = setTimeout(this.remove.bind(this, this.pop), this.opt.autoclose);
                }
                this.pop.addEventListener('click', this.addListeners.bind(this), false);
            };

            _Toast.prototype.addListeners = function (event) {
                close = event.target.getAttribute('data-pop-toast');
                if (close === 'close') {
                    if (this.autocloseTimer) {
                        clearTimeout(this.autocloseTimer);
                    }
                    this.remove(this.pop);
                }
            };

            _Toast.prototype.remove = function (elm) {
                const _this = this;
                if (this.opt.onClose) {
                    this.opt.onClose();
                }
                removeClass(elm, 'toast--in');
                setTimeout(function () {
                    if (document.body.contains(elm)) {
                        elm.parentNode.removeChild(elm);
                    }
                    const toastHeadStyle = $(headStyle)
                    const toastContainer = $(_this.getPosition('toast-pop--', _this.opt.position));
                    if (toastContainer.childNodes.length <= 0) {
                        toastContainer.remove()
                        if (toastHeadStyle) {
                            toastHeadStyle.remove();
                        }
                    }
                }, animationTime);
            };


            // Helpers
            function $(el, con) {
                return typeof el === 'string' ? (con || document).getElementById(el) : el || null;
            }

            function removeClass(el, className) {
                if (el.classList) {
                    el.classList.remove(className);
                } else {
                    el.className = el.className.replace(new RegExp('(^|\\b)' + className.split(' ').join('|') + '(\\b|$)', 'gi'), ' ');
                }
            }

            function extend(obj, src) {
                for (const key in src) {
                    if (src.hasOwnProperty(key)) obj[key] = src[key];
                }
                return obj;
            }

            const toast = function (template, style) {
                if (!template || !window.addEventListener) {
                    return false;
                }
                return new _Toast(template, style);
            };

            window.toast = {
                msg: toast,
                error: function (msg) {
                    toast(msg, 'error');
                },
                warning: function (msg) {
                    toast(msg, 'warning');
                },
                success: function (msg) {
                    toast(msg, 'success');
                }
            }
        },
        /**
         * 创建 tippy 实例
         * @param name input name
         * @param content tippy show 内容
         * @returns {*}
         */
        createTippy: (name, content='') => {
            const el = document.querySelector(`[data-validate="${name}"]`)
            if (!el._tippy) {
                // 只初始化一次
                tippy(el, {
                    content: content,
                    trigger: 'manual',
                    placement: 'right',
                    animation: 'scale',
                    hideOnClick: false,   // 👈 关键，防止点其他地方自动 hide
                    interactive: true     // 👈 防止鼠标移上去又消失
                })
            }
            el._tippy.setContent(content)
            return el._tippy
        }
    },
    /**
     * 获取当前时间戳
     * @param {string} key - 时间戳的 key，默认 second
     * @returns {Promise<number>}
     */
    currentTime: async (key = 'second')=> {
        const result = await axios.get('/public/time')
        return result.data[key]
    },
    /**
     * 存储工具
     */
    storage: {
        /**
         * 设置 set 存储
         * @param value
         * @param suffix
         */
        set: (value, suffix=null) => {
            if (suffix) {
                localStorage.setItem(`${clientId}-${suffix}`, value)
                return
            }
            localStorage.setItem(clientId, value)
        },
        /**
         * 获取 get 存储
         * @param suffix
         * @returns {string|null}
         */
        get: (suffix=null) =>  {
            if (suffix) {
                return localStorage.getItem(`${clientId}-${suffix}`)
            }
            return localStorage.getItem(clientId)
        },
        remove: (suffix=null) => {
            if (suffix) {
                localStorage.removeItem(`${clientId}-${suffix}`)
            }
            localStorage.removeItem(clientId)
        }
    },
    /**
     * 表单验证工具
     * @param {string} formSelectors 表单选择器
     * @param {Function} callback 验证通过后的回调
     * @param {Object} rules 验证规则对象，key 为 name，value 为 function(value, form) => errorMsg | null
     */
    validateForm: (formSelectors, callback, rules) =>  {
        const form = document.querySelector(formSelectors)

        // 校验单个字段
        const validateField = el => {
            if (el.getAttribute('type') ===  'hidden') return true;
            const name = el.getAttribute('name')
            if (!name || !rules[name]) return true // 没有规则的跳过

            const error = rules[name](el.value, form)

            const tip = utils.toast.createTippy(name, '')

            if (error) {
                tip.setContent(error)
                tip.show()
                el.classList.add('is-invalid')
                el.classList.remove('is-valid')
                return false
            } else if (el.checkValidity()) {
                tip.hide()
                el.classList.remove('is-invalid')
                el.classList.add('is-valid')
                return true
            } else {
                tip.hide()
                el.classList.add('is-invalid')
                el.classList.remove('is-valid')
                return true
            }
        }

        // 实时监听输入事件
        form.querySelectorAll('[name]').forEach(el => {
            el.addEventListener('input', () => {
                validateField(el)
            })
        })

        form.querySelector('button[data-type="submit"]').addEventListener('click', () => {
            let hasError = false
            // 遍历所有需要验证的字段
            form.querySelectorAll('[name]').forEach(el => {
                if (!validateField(el)) {
                    hasError = true
                }
            })
            if (!hasError) {
                // 所有通过，执行回调
                callback(form)
                // 添加 was-validated 样式，方便 Bootstrap 样式反馈
                form.classList.add('was-validated')
            }
        })
    },
    /**
     * 切换验证码
     */
    onSwitchCaptcha: () => {
        const elBox = document.querySelector('div[data-validate="captcha"]')
        const imgEl = document.querySelector('img')
        const inputEl = elBox.querySelector('input')
        inputEl.value = ''
        imgEl.src = "__LOCALHOST__/public/captcha.png?t=" + new Date().getTime();
    },
    /**
     * 应用系统主题
     */
    applySystemTheme: () =>  {
        const theme = window.matchMedia('(prefers-color-scheme: dark)').matches ? 'dark' : 'light';
        document.documentElement.setAttribute('data-bs-theme', theme);
    }
}

/*************************************************** 全局监听事件 ************************************************************
 * @function {emailValidateSend} - 绑定邮箱验证码发送事件
 * @function {userRegeditForm} - 绑定用户注册事件
 * @function {userRegeditForm} - 绑定用户注册事件
 **************************************************** 全局监听事件 ************************************************************/
const eventListener = {
    addLoadEventListener: (callback) => {
        window.addEventListener("load", callback)
    },
    /**
     * 绑定邮箱验证码发送事件
     * @param btnElementClick 发送验证码按钮元素，点击事件，eg:#send-regedit-email-code
     * @param formElementName 表单元素，eg:#form-regedit
     * @param emailInputElement 输入框元素，也就是邮箱地址，eg:#email-regedit-input
     */
    globalEmailValidateSend: (btnElementClick, formElementName, emailInputElement) => {
        const btnElement = document.querySelector(btnElementClick);
        const formElement = document.querySelector(formElementName);
        const emailElement = document.querySelector(emailInputElement);

        /**
         * 计算倒计时
         * @param oldTime 旧时间戳
         * @param newTime 新时间戳，默认不传，会自动获取当前时间戳
         * @returns {number}
         */
        function now(oldTime, newTime = 0) {
            return oldTime - newTime;
        }

        /**
         * 倒计时
         * @param rt 当前时间戳，eg:1690707070000，默认不传，会自动获取当前时间戳
         */
        async function countdown(rt) {
            const numStr = utils.storage.get('time')
            if (!numStr) {
                utils.storage.remove('form')
                return;
            }

            if (!rt) {
                rt = await utils.currentTime()
            }

            let time = parseInt(numStr, 10);

            if (isNaN(time) || time < 0) {
                utils.storage.remove('time')
                utils.storage.remove('form')
                return
            }

            // form 倒计时持久化操作
            const formLocal = utils.storage.get('form')
            if (formLocal) {
                Object.entries(JSON.parse(formLocal)).forEach(([key, value]) => {
                    formElement.elements[key].value = value;
                });
            } else {
                let formData = new FormData(formElement);
                let formDataObj = {};
                formData.forEach((value, key) => {
                    formDataObj[key] = value;
                });
                utils.storage.set(JSON.stringify(formDataObj), 'form');
            }

            btnElement.setAttribute('disabled', 'disabled');
            btnElement.innerHTML = `发送验证码(<i>${now(time - rt)}</i>s)`;

            let timer = setInterval(() => {
                if (now(time - rt) <= 0) {
                    clearInterval(timer);
                    btnElement.innerHTML = `重新发送验证码`;
                    btnElement.removeAttribute('disabled');
                    utils.storage.remove('time')
                    utils.storage.remove('form')
                    return;
                }
                btnElement.setAttribute('disabled', 'disabled');
                btnElement.innerHTML = `发送验证码(<i>${now(time - rt)}</i>s)`;
                rt++
            }, 1000);
        }

        // 初始化倒计时
        countdown();
        btnElement.addEventListener("click", function () {
            const emailRegex = /^[a-zA-Z0-9._%+\-]+@[a-zA-Z0-9.\-]+\.[a-zA-Z]{2,}$/
            if (!emailRegex.test(emailElement.value)) {
                toast.error("请输入正确的邮箱地址！")
                return
            }

            const data = {email: emailElement.value}

            axios.post('/public/email/code', data).then(async result => {
                const rt = await utils.currentTime()
                toast.success(result.message)
                utils.storage.set((rt + 60).toString(), 'time')
                // 启用倒计时
                await countdown(rt)
            })
        })
    },
    globalOnSwitchCaptcha: (clickElement) => {
        document.querySelector(clickElement).addEventListener("click", function () {
            utils.onSwitchCaptcha()
        })
    },
    /**
     * 绑定用户注册事件
     */
    userRegeditForm: () => {
        const request = (form) => {
            axios.post(form.action, new FormData(form)).then(result => {
                utils.storage.remove('time')
                utils.storage.remove('form')
                toast.success(result.message)
                const trimer = setTimeout(() => {
                    location.reload()
                    clearTimeout(trimer)
                }, 2000)
            }).catch(error => {
                for (let key in error.data) {
                    const el = form.elements[key];
                    if (el.getAttribute('type') === 'hidden') continue;
                    utils.toast.createTippy(key, error.data[key]).show()
                }
                document.querySelector('input[name="email_code"]').value = ''
                utils.onSwitchCaptcha()
            })
        }

        utils.validateForm('.float-form', request, {
            username: value => {
                if (!value) return "用户名不能为空"
                const normalRegex = /^[a-zA-Z0-9]{5,20}$/
                if (!(normalRegex.test(value))) {
                    return "请输入5-20位的字母、数字作为用户名"
                }
                return null
            },
            email: value => {
                if (!value) return "邮箱不能为空"
                const emailRegex = /^[a-zA-Z0-9._%+\-]+@[a-zA-Z0-9.\-]+\.[a-zA-Z]{2,}$/
                if (!(emailRegex.test(value))) {
                    return "请输入有效的邮箱地址"
                }
                return null
            },
            email_code: value => {
                if (!value) return "邮箱验证码不能为空"
                return null
            },
            password: value => {
                if (!value) return "密码不能为空"
                if (value.length < 6 || value.length > 20) return "密码必须包含字母、数字和特殊字符，长度6-20"
                const hasLetter = /[A-Za-z]/.test(value)
                const hasNumber = /\d/.test(value)
                const hasSpecial = /[^A-Za-z\d]/.test(value)
                if (!(hasLetter && hasNumber && hasSpecial)) {
                    return "密码必须包含字母、数字和特殊字符，长度6-20"
                }
                return null
            },
            confirm_password: (value, form) => {
                const passwordValue = form.querySelector('[name="password"]').value
                if (!value) return "确认密码不能为空"
                if (value !== passwordValue) return "两次输入的密码不一致"
                return null
            },
            captcha: value => {
                if (!value) return "验证码不能为空"
                return null
            }
        })
    },
    /**
     * 登录表单校验
     */
    userLoginForm: () => {
        const request = (form) => {
            axios.post(form.action, new FormData(form)).then(result => {
                toast.success(result.message)
                const trimer = setTimeout(() => {
                    location.reload()
                    clearTimeout(trimer)
                }, 2000)
            }).catch(error => {
                for (let key in error.data) {
                    const el = form.elements[key];
                    if (el.getAttribute('type') === 'hidden') continue;
                    utils.toast.createTippy(key, error.data[key]).show()
                }
                utils.onSwitchCaptcha()
            })
        }

        utils.validateForm('.float-form', request, {
            username: value => {
                if (!value) return "用户名不能为空"
                return null
            },
            password: value => {
                if (!value) return "密码不能为空"
                return null
            },
            captcha: value => {
                if (!value) return "验证码不能为空"
                return null
            }
        })
    }
}

/*************************************************** 全局初始化方法 ************************************************************
 * @function {window.onload} - 全局页面初始化
 *************************************************** 全局初始化方法 ************************************************************/
eventListener.addLoadEventListener(() => {
    utils.toast.newToast()
    utils.applySystemTheme();
    matchMedia('(prefers-color-scheme: dark)').addEventListener('change', utils.applySystemTheme);
})
