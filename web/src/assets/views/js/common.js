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
        newToast: () => {
            const messageBox = {
                Alert(options) {
                    const config = {
                        ...defaultMsgBoxConfig,
                        ...options,
                        buttons: {
                            ...defaultMsgBoxConfig.buttons,
                            ...(options.buttons || {})
                        }
                    };
                    renderMsgBox(config, "alert");
                },
                Confirm(options) {
                    const config = {
                        ...defaultMsgBoxConfig,
                        ...options,
                        buttons: {
                            ...defaultMsgBoxConfig.buttons,
                            ...(options.buttons || {})
                        }
                    };
                    renderMsgBox(config, "confirm");
                },
                Notify(options) {
                    const config = {
                        ...defaultNotificationConfig,
                        ...options
                    };

                    switch (config.location) {
                        case "top":
                            config.location = "locationT";
                            break;
                        case "right":
                            config.location = "locationR";
                            break;
                        default:
                            config.location = "locationR";
                            break;
                    }
                    renderNotification(config);
                },

                alert: {
                    normal: function (title, content, callback) {
                        messageBox.Alert({title: title, content: content, callback: callback})
                    },
                    success: function (title, content, callback) {
                        messageBox.Alert({title: title, content: content, type: "success", callback: callback})
                    },
                    warning: function (title, content, callback) {
                        messageBox.Alert({title: title, content: content, type: "warning", callback: callback})
                    },
                    error: function (title, content, callback) {
                        messageBox.Alert({title: title, content: content, type: "error", callback: callback})
                    }
                },
                confirm: {
                    question: function (title, content, callback) {
                        messageBox.Confirm({title: title, content: content, callback: callback})
                    },
                    warning: function (title, content, callback) {
                        messageBox.Confirm({title: title, content: content, type: "warning", callback: callback})
                    },
                    success: function (title, content, callback) {
                        messageBox.Confirm({title: title, content: content, type: "success", callback: callback})
                    },
                    error: function (title, content, callback) {
                        messageBox.Confirm({title: title, content: content, type: "error", callback: callback})
                    }
                },
                toast: {
                    normal: function (content) {
                        messageBox.Notify({content: content, showtime: 3000, location: "top"});
                    },
                    success: function (content) {
                        messageBox.Notify({content: content, type: "success", showtime: 3000, location: "top"});
                    },
                    warning: function (content) {
                        messageBox.Notify({content: content, type: "warning", showtime: 3000, location: "top"});
                    },
                    error: function (content) {
                        messageBox.Notify({content: content, type: "error", showtime: 3000, location: "top"});
                    },
                    info: function (content) {
                        messageBox.Notify({content: content, type: "info", showtime: 3000, location: "top"});
                    }
                },
                notify: {
                    normal: function (content, title="通知") {
                        messageBox.Notify({title: title, content: content, showtime: 3000});
                    },
                    success: function (content, title="通知") {
                        messageBox.Notify({title: title, content: content, type: "success", showtime: 3000});
                    },
                    warning: function (content, title="通知") {
                        messageBox.Notify({title: title, content: content, type: "warning", showtime: 3000});
                    },
                    error: function (content, title="通知") {
                        messageBox.Notify({title: title, content: content, type: "error", showtime: 3000});
                    },
                    info: function (content, title="通知") {
                        messageBox.Notify({title: title, content: content, type: "info", showtime: 3000});
                    }
                },
            };

            function renderMsgBox(config, mode) {
                const zIndex = calculateLayer("messagebox");

                // 背景层
                const backdrop = document.createElement("div");
                backdrop.className = "me-message-box-bg";
                backdrop.style.display = "block";

                // 主容器
                const box = document.createElement("div");
                box.className = "me-message-box-alert";

                if (zIndex >= 99999) {
                    backdrop.style.zIndex = zIndex - 1;
                    box.style.zIndex = zIndex;
                }

                // 顶部区域（用于拖动）
                const dragBar = document.createElement("div");
                dragBar.className = "distop";
                box.appendChild(dragBar);

                // 内容容器
                const contentBox = document.createElement("div");
                contentBox.className = "msgcontainer";
                box.appendChild(contentBox);

                // 图标
                if (mode === "confirm") {
                    const type = config.type && config.type !== "none" ? config.type : "question";
                    config.type = type;
                    const icon = document.createElement("div");
                    icon.className = `icon ${type}`;
                    contentBox.appendChild(icon);
                } else if (config.type && config.type !== "none") {
                    const icon = document.createElement("div");
                    icon.className = `icon ${config.type}`;
                    contentBox.appendChild(icon);
                }

                // 标题
                if (config.title) {
                    const titleEl = document.createElement("div");
                    titleEl.className = "msgtitle";
                    titleEl.textContent = config.title;
                    contentBox.appendChild(titleEl);
                }

                // 内容
                if (config.content) {
                    const text = document.createElement("div");
                    text.className = "msgcon";
                    text.textContent = config.content;
                    contentBox.appendChild(text);
                }

                // 底部按钮栏
                const footer = document.createElement("div");
                footer.className = "operatebar";
                box.appendChild(footer);

                const confirmBtn = document.createElement("button");
                confirmBtn.type = "button";
                confirmBtn.textContent = config.buttons.confirm.text;

                const cancelBtn = document.createElement("button");
                cancelBtn.type = "button";
                cancelBtn.textContent = config.buttons.cancel.text;
                cancelBtn.className = "cancel";

                // 按钮样式
                switch (config.type) {
                    case "success":
                        confirmBtn.classList.add("success");
                        break;
                    case "question":
                        confirmBtn.classList.add("normal");
                        break;
                    case "warning":
                        confirmBtn.classList.add("warning");
                        break;
                    case "error":
                        confirmBtn.classList.add("error");
                        break;
                    default:
                        confirmBtn.classList.add("normal");
                        contentBox.classList.add("typenone");
                        break;
                }

                // 不同模式的按钮组合
                if (mode === "alert") {
                    footer.appendChild(confirmBtn);
                } else if (mode === "confirm") {
                    footer.appendChild(confirmBtn);
                    footer.appendChild(cancelBtn);
                    confirmBtn.classList.add("beleft");
                    cancelBtn.classList.add("beright");
                } else {
                    footer.appendChild(confirmBtn);
                }

                // 是否已有弹窗
                if (document.querySelectorAll(".me-message-box-alert").length > 0) {
                    const timer = setInterval(() => {
                        if (document.querySelectorAll(".me-message-box-alert").length <= 0) {
                            clearInterval(timer);
                            show();
                        }
                    }, 500);
                } else {
                    show();
                }

                /** 展示逻辑 */
                function show() {
                    if (config.aero) {
                        document.body.style.overflowX = "hidden";
                        backdrop.classList.add("aero");
                    }

                    document.body.appendChild(backdrop);
                    document.body.appendChild(box);

                    // 居中计算
                    if (window.innerWidth < parseInt(getComputedStyle(box).maxWidth)) {
                        box.style.maxWidth = `${window.innerWidth - 10}px`;
                    }

                    adjustContentHeight();
                    window.addEventListener("resize", adjustContentHeight);

                    function adjustContentHeight() {
                        const availableHeight =
                            window.innerHeight -
                            dragBar.offsetHeight -
                            footer.offsetHeight -
                            120;
                        const msgCon = box.querySelector("div.msgcon");
                        if (msgCon) {
                            msgCon.style.maxHeight = `${availableHeight <= 100 ? 100 : availableHeight}px`;
                        }
                    }

                    backdrop.classList.add("me-message-box--motion", "me-message-box-bg--show");
                    box.style.left = `${(window.innerWidth - box.offsetWidth) / 2}px`;
                    box.style.top = `${(window.innerHeight - box.offsetHeight) / 2}px`;
                    box.style.display = "block";
                    box.classList.add("me-message-box--motion", "me-message-box-alert--open");

                    box.addEventListener("animationend", () => {
                        box.classList.remove("me-message-box-alert--open");
                        backdrop.addEventListener("mousedown", () => {
                            box.classList.add("me-message-box--leap");
                            box.addEventListener("animationend", () => {
                                box.classList.remove("me-message-box--leap");
                                if (cancelBtn) {
                                    cancelBtn.focus();
                                } else {
                                    confirmBtn.focus();
                                }
                            }, { once: true });
                        });
                    }, { once: true });

                    const icon = box.querySelector("div.icon");
                    if (icon) {
                        icon.classList.add("me-message-box--showicon", "me-message-box--motion");
                    }

                    initEvents();
                }

                /** 事件绑定 */
                function initEvents() {
                    if (mode === "alert") {
                        confirmBtn.focus();
                    } else {
                        cancelBtn.focus();
                    }

                    cancelBtn.addEventListener("click", () => close(false));
                    confirmBtn.addEventListener("click", () => close(true));
                }

                /** 关闭逻辑 */
                function close(result) {
                    box.classList.remove("me-message-box-alert--open");
                    box.classList.add("me-message-box-alert--close");

                    box.addEventListener("animationend", () => {
                        document.body.style.overflowX = "initial";
                        box.remove();
                        if (typeof config.callback === "function") {
                            config.callback(result);
                        }
                    }, { once: true });

                    backdrop.classList.remove("me-message-box-bg--show");
                    backdrop.classList.add("me-message-box-bg--hide");
                    backdrop.style.animationDelay = ".3s";
                    backdrop.addEventListener("animationend", () => backdrop.remove(), { once: true });
                }
            }

            function renderNotification(config) {
                // 容器
                let container = document.querySelector(
                    `.me-notification-container.${config.location}.${config.tipSort}`
                );

                if (!container) {
                    container = document.createElement("div");
                    container.className = `me-notification-container ${config.location} ${config.tipSort}`;
                    container.style.zIndex = calculateLayer("tips");
                    document.body.appendChild(container);
                }

                // 外壳
                const capsule = document.createElement("div");
                capsule.className = "notification-capsule";
                capsule.style.height = "0px";

                // 内容
                const notification = document.createElement("div");
                notification.className = "notification";

                // 类型图标
                if (config.type) {
                    notification.classList.add("carrystate", config.type);
                    const icon = document.createElement("i");
                    icon.className =
                        "notification-icon icon-state me-notification--motion me-message-box--showicon";
                    notification.appendChild(icon);
                    icon.addEventListener("animationend", e => e.stopPropagation());
                }

                // 标题
                if (config.title) {
                    const title = document.createElement("div");
                    title.className = "title me-notification--motion me-motion--inlinecon";
                    title.textContent = config.title;
                    notification.appendChild(title);
                    title.addEventListener("animationend", e => e.stopPropagation());
                }

                // 内容文本
                if (config.content) {
                    const content = document.createElement("div");
                    content.className = "con me-notification--motion me-motion--inlinecon";
                    if (notification.querySelector(".title")) {
                        content.style.marginTop = "5px";
                    }
                    content.textContent = config.content;
                    content.style.animationDelay = ".3s";
                    notification.appendChild(content);
                    content.addEventListener("animationend", e => e.stopPropagation());
                }

                // 插入容器（顶部 or 底部）
                if (config.tipSort === "top" && container.children.length > 0) {
                    container.insertBefore(capsule, container.firstChild);
                } else {
                    container.appendChild(capsule);
                }

                // 关闭按钮
                if (config.closable) {
                    const closeBox = document.createElement("div");
                    closeBox.className = "me-notification-close";

                    const closeBtn = document.createElement("button");
                    closeBtn.type = "button";
                    closeBtn.className = "close";
                    closeBox.appendChild(closeBtn);
                    closeBtn.addEventListener("click", () => close());
                    notification.appendChild(closeBox);
                }

                // 动画方向
                let showAnim = "me-notification-show--right";
                let hideAnim = "me-notification-hide--right";
                if (config.location === "locationT") {
                    showAnim = "me-notification-show--top";
                    hideAnim = "me-notification-hide--top";
                }

                // 插入内容
                capsule.appendChild(notification);
                capsule.style.height = notification.offsetHeight + 10 + "px";
                capsule.addEventListener("transitionend", () => {
                    document.body.style.overflowX = "hidden";
                });

                // 出现动画
                notification.classList.add("me-notification--motion", showAnim);
                notification.addEventListener("animationend", function onAnimEnd(e) {
                    capsule.style.height = "auto";
                    document.body.style.overflowX = "initial";

                    // 自动关闭逻辑
                    if (typeof config.showtime === "number") {
                        if (config.progressBar) {
                            const progress = document.createElement("div");
                            progress.className = "processbar";
                            const inner = document.createElement("div");
                            inner.className = "me-notification--motion me-notification--process";
                            inner.style.animationDuration = `${config.showtime / 1000}s`;

                            inner.addEventListener("animationend", ev => {
                                close();
                                ev.stopPropagation();
                            });

                            progress.appendChild(inner);
                            notification.appendChild(progress);

                            notification.addEventListener("mouseover", () => {
                                inner.style.animationPlayState = "paused";
                            });
                            notification.addEventListener("mouseout", () => {
                                inner.style.animationPlayState = "running";
                            });
                        } else {
                            setTimeout(() => close(), config.showtime);
                        }
                    }

                    e.stopPropagation();
                    notification.removeEventListener("animationend", onAnimEnd);
                });

                /** 关闭函数 */
                function close() {
                    capsule.style.height = capsule.offsetHeight + "px";
                    notification.classList.remove(showAnim);
                    notification.classList.add(hideAnim);

                    notification.addEventListener(
                        "animationend",
                        () => {
                            notification.remove();
                            capsule.style.height = "0";
                            capsule.addEventListener("transitionend", () => {
                                capsule.remove();
                                if (typeof config.callback === "function") {
                                    config.callback();
                                }
                                if (container.children.length === 0) {
                                    container.remove();
                                }
                            }, { once: true });
                        },
                        { once: true }
                    );
                }
            }

            function calculateLayer(type) {
                let layerIndex = 0;

                switch (type) {
                    case "messagebox":
                        break;

                    case "news": {
                        const alertBox = document.querySelector(".me-message-box-alert");
                        if (alertBox) {
                            const zIndex = parseFloat(getComputedStyle(alertBox).zIndex);
                            layerIndex = zIndex - 3;
                        }
                        break;
                    }

                    case "tips": {
                        const notification = document.querySelector(".me-notification-container");
                        if (notification) {
                            const zIndex = parseFloat(getComputedStyle(notification).zIndex);
                            layerIndex = zIndex + 1;
                        }
                        break;
                    }

                    default:
                        break;
                }

                if (layerIndex <= 0) {
                    // 遍历所有元素，取最大 z-index
                    const allElements = Array.from(document.body.querySelectorAll("*"));
                    const zIndices = allElements
                        .map(el => {
                            const style = getComputedStyle(el);
                            if (style.position !== "static") {
                                return parseInt(style.zIndex, 10) || -1;
                            }
                            return -1;
                        })
                        .filter(z => z >= 0);

                    layerIndex = zIndices.length > 0 ? Math.max(...zIndices) : -1;

                    if (layerIndex <= 0 || layerIndex === "auto") {
                        layerIndex = 9999;
                    }
                }

                return layerIndex;
            }

            // 默认消息框配置
            const defaultMsgBoxConfig = {
                title: "",
                content: "",
                type: "none",
                aero: true,
                buttons: {
                    confirm: { text: "确定" },
                    cancel: { text: "取消" }
                },
                callback: () => {}
            };

            // 默认提示条配置
            const defaultNotificationConfig = {
                title: "",
                content: "",
                location: "right",
                tipSort: "top",
                type: "",
                duration: null,
                closable: true,
                progressBar: true,
                callback: null
            };

            window.meMsg = messageBox || {}
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
                meMsg.toast.warning("请输入正确的邮箱地址！")
                return
            }

            const data = {email: emailElement.value}

            axios.post('/public/email/code', data).then(async result => {
                const rt = await utils.currentTime()
                meMsg.toast.success(result.message)
                utils.storage.set((rt + 60).toString(), 'time')
                // 启用倒计时
                await countdown(rt)
            })
        })
    },
    /**
     * 全局切换验证码事件
     * @param clickElement 点击事件元素，eg:#switch-captcha
     * @param inputElement 输入框元素，eg:#captcha-input
     */
    globalOnSwitchCaptcha: (clickElement, inputElement) => {
        document.querySelector(clickElement).addEventListener("click", function () {
            utils.onSwitchCaptcha()
        })
        document.querySelector(inputElement).addEventListener('focus', function () {
            utils.onSwitchCaptcha()
        })
    },
}

/*************************************************** 全局初始化方法 ************************************************************
 * @function {window.onload} - 全局页面初始化
 *************************************************** 全局初始化方法 ************************************************************/
eventListener.addLoadEventListener(() => {
    utils.toast.newToast()
    utils.applySystemTheme();
    matchMedia('(prefers-color-scheme: dark)').addEventListener('change', utils.applySystemTheme);

    // 绑定注销登录事件
    const logoutBtn = document.getElementById('logoutBtn')
    if (logoutBtn) {
        logoutBtn.addEventListener('click', () => {
            meMsg.confirm.question("注销提示", "是否确认注销当前登录用户！", function (isOk) {
                if (!isOk) {
                    return
                }
                axios.post('/logout').then(result => {
                    if (result.code !== 10000) {
                        meMsg.toast.warning(result.message)
                        return
                    }
                    utils.storage.remove('time')
                    utils.storage.remove('form')
                    localStorage.removeItem('TOKEN')
                    location.reload()
                })
            })
        })
    }
})
