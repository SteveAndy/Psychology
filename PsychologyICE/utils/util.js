const api = require('api'), app = getApp(), globalData = app.globalData,
  HEADER_X = "application/x-www-form-urlencoded", Header_J = 'application/json', jim = require('Jim.js'),UI = require('Interface.js')
var mthat;

/**
 * 返回当前时间
 * @param {Int} type 1：返回当前时间;2：返回当前年月日;other：返回当前年月日及时间
 */
function formatTime(type) {
  var date = new Date()
  var year = date.getFullYear()
  var month = date.getMonth() + 1
  var day = date.getDate()
  var hour = date.getHours()
  var minute = date.getMinutes()
  var second = date.getSeconds()
  var time = [hour, minute, second].map(formatNumber).join(':'),
    ydate = [year, month, day].map(formatNumber).join('/')
  switch (type) {
    case 1:
      return time
      break;
    case 2:
      return ydate
      break;
    default:
      return ydate + ' ' + time
      break;
  }
}

const formatNumber = n => {
  n = n.toString()
  return n[1] ? n : '0' + n
}

//登录
function login(that) {
  mthat = that
  UI.showLoading('加载中')
  if (!wx.getStorageSync("userData")) {
    wx.navigateTo({
      url: UI.page.user_auth,
    })
    return null;
  }
  wx.login({
    success: function (res) {
      wx.request({
        url: api.Anum.login,
        data: { expert_code: res.code },
        method: 'POST',
        header: { "content-type": HEADER_X },
        success: function (e) {
          const status = e.data.status, data = e.data.data, msg = e.data.msg
          switch (status) {
            case 0://登录成功
              loginOK(data)
              break;
            case 1://申请中
              wx.redirectTo({ url: UI.page.user_apply_details + '?status=' + status})
              break;
            case 2://未申请
              wx.redirectTo({ url: UI.page.user_apply_index+'?openid='+data.openid })
              break;
            case 3://被驳回
              wx.redirectTo({ url: UI.page.user_apply_details + '?status=' + status + '&reject=' + data })
              break;
            default:
              wx.hideLoading()
              UI.showToast('登录遇到未知错误')
              console.log('error:', e.data)
              break;
          }
        },
      })
    },
  })
}

/**
 * 登录成功
 * @param {*} data 
 */
function loginOK(data) {
  var user = wx.getStorageSync("userInfo")
  globalData.token = data.token
  wx.setStorageSync("user_id", data.id)
  globalData.userID = data.id
  globalData.name = data.name
  globalData.iconUrl = !data.photo ? user.avatarUrl : data.photo
  console.log('用户已登录:', data)
  jim.init(this)
  mthat.getInformationClass()//初始化页面数据
}

/**
 * UPDATE_FILE 上传修改的资料
 * @param {String} urlkey URL标识符
 * @param {function} cb 回调函数
 * @param {String} temp 选择的图片路径
 */
function uploadFile(urlkey, cb, temp, formData) {
  var url = api.url(urlkey)
  // console.log(url)
  wx.uploadFile({
    url: url,
    filePath: temp,
    name: 'file',
    formData: formData,
    success: function (res) {
      // console.log(res)
      var obj = JSON.parse(res.data);
      obj.status != 0 ? wx.hideLoading() : null
      wx.hideLoading();
      return typeof cb == "function" && cb(obj)
    },
    fail: function (res) {
      wx.hideLoading();
      wx.showModal({
        title: '网络错误',
        content: '网络出错，请刷新重试',
        showCancel: false
      })
      return typeof cb == "function" && cb(false)
    },
  })
}

/**
 * GET
 * @param {String} urlkey URL标识符
 * @param {function} cb 回调函数
 * @param {String} token 令牌
 */
function getReq(urlkey, cb, data) {
  var url = api.url(urlkey)
  //console.log('请求链接:', url)
  wx.request({
    url: url,
    method: 'GET',
    data: data,
    header: {
      "content-type": HEADER_X
    },
    success: function (res) {
      wx.hideLoading();
      return typeof cb == "function" && cb(res.data)
    },
    fail: function () {
      wx.hideLoading();
      wx.showModal({
        title: '网络错误',
        content: '网络出错，请刷新重试',
        showCancel: false
      })
      return typeof cb == "function" && cb(false)
    }
  })
}
/**
 * POST
 * @param {String} urlkey URL标识符
 * @param {String[]} data 提交的数据
 * @param {function} cb 回调函数
 * @param {String} Header 自定义请求头 null或者不填: application/x-www-form-urlencoded; 1：application/json; Header：自定义请求头
 */
function postReq(urlkey, data, cb, Header) {
  var url = api.url(urlkey)
  // console.log('请求链接:', url)
  // console.log('请求数据:', data)
  wx.request({
    url: url,
    data: data,
    method: 'POST',
    header: {
      "content-type": !Header ? HEADER_X : (Header != 1 ? Header : Header_J)
    },
    success: function (res) {
      wx.hideLoading();
      return typeof cb == "function" && cb(res.data)
    },
    fail: function () {
      wx.hideLoading();
      wx.showModal({
        title: '网络错误',
        content: '网络出错，请刷新重试',
        showCancel: false
      })
      return typeof cb == "function" && cb(false)
    }
  })

}

/**
 * PUT
 * @param {String} urlkey URL标识符
 * @param {String[]} data 提交的数据
 * @param {function} cb 回调函数
 * @param {String} Header 自定义请求头 null或者不填: application/x-www-form-urlencoded; 1：application/json; Header：自定义请求头
 */
function putReq(urlkey, data, cb, Header) {
  var url = api.url(urlkey)
  // console.log('请求链接:', url)
  // console.log('请求数据:', data)
  wx.request({
    url: url,
    data: data,
    method: 'PUT',
    header: {
      "content-type": !Header ? HEADER_X : (Header != 1 ? Header : Header_J)
    },
    success: function (res) {
      wx.hideLoading();
      return typeof cb == "function" && cb(res.data)
    },
    fail: function () {
      wx.hideLoading();
      wx.showModal({
        title: '网络错误',
        content: '网络出错，请刷新重试',
        showCancel: false
      })
      return typeof cb == "function" && cb(false)
    }
  })
}

/**
 * DELETE
 * @param {String} urlkey URL标识符
 * @param {String[]} data 提交的数据
 * @param {function} cb 回调函数
 * @param {String} Header 自定义请求头 null或者不填: application/x-www-form-urlencoded; 1：application/json; Header：自定义请求头
 */
function DeleteReq(urlkey, data, cb, Header) {
  var url = api.url(urlkey)
  wx.request({
    url: url,
    data: data,
    method: 'DELETE',
    header: {
      "content-type": !Header ? HEADER_X : (Header != 1 ? Header : Header_J)
    },
    success: function (res) {
      wx.hideLoading();
      return typeof cb == "function" && cb(res.data)
    },
    fail: function () {
      wx.hideLoading();
      wx.showModal({
        title: '网络错误',
        content: '网络出错，请刷新重试',
        showCancel: false
      })
      return typeof cb == "function" && cb(false)
    }
  })
}

module.exports = {
  formatTime,
  login,
  GET: getReq,
  POST: postReq,
  PUT: putReq,
  DELETE: DeleteReq,
  UpLoadFile: uploadFile
}