function dateTime () {
  const date = new Date()

  const padZero = (n: number) => (n < 10 ? '0' + n : n)

  const y = date.getFullYear()
  const m = padZero(date.getMonth() + 1)
  const d = padZero(date.getDate())
  const h = padZero(date.getHours())
  const min = padZero(date.getMinutes())
  const s = padZero(date.getSeconds())

  return `${y}-${m}-${d} ${h}:${min}:${s}`
}

function linkifyLog (logText: string) {
  const urlPattern = /(https?:\/\/\S+)/g
  return logText.replace(urlPattern, (url) => {
    return `&nbsp;<a target="_blank" href="${url}">${url}</a>&nbsp;`
  })
}

export function addLineLog (html: string, local?: boolean) {
  const logContainerContent = document.querySelector('.log-container-content')
  if (!logContainerContent) {
    return
  }
  const p = document.createElement('pre')
  if (local) {
    p.innerHTML = `[${dateTime()}] [INFO] > ${html}`
  } else {
    p.innerHTML = linkifyLog(html)
  }
  logContainerContent.appendChild(p)

  const logContainer = document.querySelector('.log-container')
  if (!logContainer) {
    return
  }
  logContainer.scrollTop = logContainer.scrollHeight
}
