import cookies from './util.cookies'
import db from './util.db'
import log from './util.log'

const util = {
  cookies,
  db,
  log
}

/**
 * @description 更新标题
 * @param {String} title 标题
 */
util.title = function (titleText) {
  const processTitle = process.env.VUE_APP_TITLE || 'D2Admin'
  window.document.title = `${processTitle}${titleText ? ` | ${titleText}` : ''}`
}

/**
 * @description 打开新页面
 * @param {String} url 地址
 */
util.open = function (url) {
  var a = document.createElement('a')
  a.setAttribute('href', url)
  a.setAttribute('target', '_blank')
  a.setAttribute('id', 'd2admin-link-temp')
  document.body.appendChild(a)
  a.click()
  document.body.removeChild(document.getElementById('d2admin-link-temp'))
}

util.uuid = function () {
  var s = []
  var hexDigits = '0123456789abcdef'
  for (var i = 0; i < 36; i++) {
    s[i] = hexDigits.substr(Math.floor(Math.random() * 0x10), 1)
  }
  s[14] = '4' // bits 12-15 of the time_hi_and_version field to 0010
  s[19] = hexDigits.substr((s[19] & 0x3) | 0x8, 1)// bits 6-7 of the clock_seq_hi_and_reserved to 01
  s[8] = s[13] = s[18] = s[23] = '-'
  var uuid = s.join('')
  return uuid.toUpperCase()
}

util.platformDicts = [
  { label: 'OneFlow', value: 'OneFlow' },
  { label: 'tensorflow', value: 'TensorFlow' },
  { label: 'PyTorch', value: 'PyTorch' },
  { label: 'MxNet', value: 'MxNet' }]

util.labelDicts = [
  { label: '图像识别', value: '图像识别' },
  { label: '人脸识别', value: '人脸识别' },
  { label: '语音识别', value: '语音识别' },
  { label: '文本识别', value: '文本识别' },
  { label: '机器翻译', value: '机器翻译' },
  { label: '垃圾邮件', value: '垃圾邮件' } ]

export default util
