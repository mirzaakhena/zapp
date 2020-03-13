export default function (Vue) {
  Vue.auth = {
    setToken (token) {
      localStorage.setItem('token', token)
    },
    getToken () {      
      var token = localStorage.getItem('token')
      if (!token) {      
        return null
      }
      if (Date.now() > parseInt(getExpired(token) + '000')) {
        this.removeToken()
        return null
      }
      return token
    },
    removeToken () {
      localStorage.removeItem('token')
    },
    isAuthenticated () {
      if (this.getToken()) {
        return true
      } else {
        return false
      }
    },
    getExtendData () {
      var token = localStorage.getItem('token')
      return getExtendDataFromJWT(token)
    },
  }

  Object.defineProperties(Vue.prototype, {
    $auth: {
      get () {
        return Vue.auth
      }
    }
  })

  function getExpired (token) {
    var tokens = token.split('.')
    var payloadDecoded = atob(tokens[1])
    return JSON.parse(payloadDecoded).exp
  }

  function getExtendDataFromJWT (token) {
    var tokens = token.split('.')
    var payloadDecoded = atob(tokens[1])
    return JSON.parse(payloadDecoded).extendData
  }  
}