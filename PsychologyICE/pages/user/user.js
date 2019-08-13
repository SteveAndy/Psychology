const app = getApp(),
  allData = app.globalData,
  utils = require('../../utils/util.js'),
  UI = require('../../utils/Interface.js')

Page({
  /**
   * 页面的初始数据
   */
  data: {
    user_info: wx.getStorageSync("userInfo"),
  },
    /**
     * 生命周期函数--监听页面加载
     */
    onLoad: function (options) {
      if (!allData.isGetUser) UI.showLoading('加载中')
      !wx.getStorageSync('wave') ?
        UI.seveImage(this, 'wave.gif', 'wave') :
        this.setData({
          wave: wx.getStorageSync('wave')
        }) & this.getUserInfo()
    },

    getUserInfo: function () {
      var t = this
        app.setTitleWidth(this);
        utils.GET('userInfo', function (res) {
          console.log('获取用户数据:', res) //回调数据调试
          var iconUrl = !res.data.Icon ? allData.iconUrl : res.data.Icon
          allData.name = res.data.Name
          allData.address = res.data.Address
          allData.phoneNum = res.data.PhoneNum
          allData.iconUrl = iconUrl
          allData.gender = res.data.Gender
          allData.age = res.data.WorkAge
          allData.info = res.data.Info

          // //防止点击过快出现数据为缓存后，无法自动再次加载的尴尬情况
          allData.isGetUser = !res.data.address && !res.data.user_name && !res.data.gender ? false : true
          t.setData({
            address: res.data.Address,
            userName: res.data.Name,
            gender: res.data.Gender,
            iconUrl: iconUrl,
          })
        }, true)
    },
    /**
     * 跳转页面
     * @param {*} e 
     */
    toPage(e) {
      const page = UI.page[[e.currentTarget.dataset.page]]
      allData.that = this
      wx.navigateTo({ url: page })
    }
})