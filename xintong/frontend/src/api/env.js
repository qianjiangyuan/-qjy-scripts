import 'process'

export default {
  development: {
    host: '10.5.24.113',
    port: '8080'
  },
  production: {
    host: '192.168.1.11',
    port: '8081'
  },
  url (ssl) {
    let protocol = ssl ? 'https://' : 'http://'
    let env = process.env.NODE_ENV === 'production' ? 'production' : 'development'

    return `${protocol}${this[env].host}:${this[env].port}`
  }
}
let Art = '                    __ _               \n' +
  '                   / _| |              \n' +
  '   ___  _ __   ___| |_| | _____      __\n' +
  '  / _ \\| \'_ \\ / _ \\  _| |/ _ \\ \\ /\\ / /\n' +
  ' | (_) | | | |  __/ | | | (_) \\ V  V / \n' +
  '  \\___/|_| |_|\\___|_| |_|\\___/ \\_/\\_/  \n' +
  '                                       \n' +
  '                                       '

export { Art }
