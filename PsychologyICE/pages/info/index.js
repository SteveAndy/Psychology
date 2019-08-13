const app = getApp(),
  utils = require('../../utils/util.js'),
  allData = getApp().globalData,
  api = require('../../utils/api.js'),
  UI = require('../../utils/Interface.js')
var doubleClick = 0
Page({

  data: {
    loadingImg: api.loadingImgUrl + 'loading-1.gif',
    ViewHeight: wx.getSystemInfoSync().screenHeight
  },

  onLoad: function (options) {
    var t = this, idInformationInfo
    app.setTitle1Width(this)
    UI.showLoading('加载中')
    allData.idInformationInfoId = options.id
    utils.GET('InformationInfo', function (res) {
      idInformationInfo = res.data
      res.status == 0 ? t.setData({
        idInformationInfo,
        imgurl: options.img,
        indexTitle: '资讯详情'
      }) : t.setData({
        idInformationInfo: 'ErrorNetwork',
        indexTitle: '请求错误'
      }) & UI.showToast('错误:' + res.msg,1)
    })
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

  /*menu(e){
    console.log(e)
    var page = getCurrentPages(), type = e.currentTarget.dataset.type
    page = page[page.length - 2]
    type == 'delete' ? page.delete() : page.modify()
  }*/
})