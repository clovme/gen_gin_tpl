eventListener.addLoadEventListener(() => {
    eventListener.globalOnSwitchCaptcha('.verification-img', '#floatingCaptcha')
    eventListener.globalEmailValidateSend(".email-validate-code", ".float-form", "#floatingEmail")

    const request = (form) => {
        axios.post(form.action, new FormData(form)).then(result => {
            utils.storage.remove('time')
            utils.storage.remove('form')
            if (result.data.token !== '') {
                localStorage.setItem('TOKEN', result.data.token)
            } else {
                localStorage.removeItem('TOKEN')
            }
            meMsg.alert.success("登录提示", result.message, function () {
                window.location.reload()
            })
        }).catch(error => {
            for (let key in error.data) {
                const el = form.elements[key];
                if (el.getAttribute('type') === 'hidden') continue;
                if (key === 'email_code') {
                    document.querySelector('input[name="email_code"]').value = ''
                }
                utils.toast.createTippy(key, error.data[key]).show()
            }

            utils.storage.remove('time')
            utils.storage.remove('form')
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
})
