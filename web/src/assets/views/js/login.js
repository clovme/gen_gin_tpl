eventListener.addLoadEventListener(() => {
    eventListener.globalOnSwitchCaptcha('.verification-img', '#floatingCaptcha')
    const request = (form) => {
        axios.post(form.action, new FormData(form)).then(result => {
            toast.success(result.message)
            const trimer = setTimeout(() => {
                if (result.data.token !== '') {
                    localStorage.setItem('TOKEN', result.data.token)
                } else {
                    localStorage.removeItem('TOKEN')
                }
                location.reload()
                clearTimeout(trimer)
            }, 2000)
        }).catch(error => {
            for (let key in error.data) {
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
})
