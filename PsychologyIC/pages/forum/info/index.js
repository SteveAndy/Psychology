const app = getApp(),
  utils = require('../../../utils/util.js'),
  allData = getApp().globalData,
  api = require('../../../utils/api.js'),
  ui = require('../../../utils/Interface.js')
var doubleClick = 0
Page({

  data: {
    loadingImg: api.loadingImgUrl + 'loading-1.gif',
    ViewHeight: allData.windowHeight,
    InputBottom: 0,
    InputValue: '',
    user_id: wx.getStorageSync('user_id')
  },

  onLoad: function (options) {
    app.setTitle1Width(this, '详情')
    this.setData({ is_super: allData.is_super })
    allData.CommunityInfo = options.id
    this.getCommunityInfo()
  },

  getCommunityInfo() {
    var t = this, CommunityInfo
    utils.GET('CommunityInfo', function (res) {
      CommunityInfo = res.data
      res.status == 0 ? t.setData({ CommunityInfo })
        : t.setData({
          CommunityInfo: 'ErrorNetwork',
          indexTitle: '请求错误'
        }) & ui.showToast('错误:' + res.msg, 1)
    }, { offset: 0 })
  },

  //顶部标题被双击返回顶部
  doubleClick(e) {
    doubleClick++
    if (doubleClick == 2) {
      doubleClick = 0
      this.setData({ scrollTop: 0 })
    } else {
      setTimeout(() => {
        doubleClick = 0
      }, 450);
    }
  },

  call(e) {
    wx.makePhoneCall({
      phoneNumber: e.currentTarget.dataset.ephone
    })
  },

  chat(e) {
    const id = e.currentTarget.dataset.id, pic = e.currentTarget.dataset.pic
    allData.chatName = id
    wx.navigateTo({ url: ui.page.chat_ui + "?pic=" + pic })
  },
})