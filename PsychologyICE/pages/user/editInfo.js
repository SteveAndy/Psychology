const app = getApp(), allData = app.globalData, api = require('../../utils/api.js'),
  utils = require('../../utils/util.js'), UI = require('../../utils/Interface.js')
var isChoice, name, age, phone, info;
/**
 * 计算顶部实际占界面的高度
 */
function setTopHeight(that) {
  var query = wx.createSelectorQuery()//单位px；
  query.select('#top').boundingClientRect(function (rect) {
    that.setData({
      TopHeight: rect.height
    })
  }).exec();
}
Page({

  data: { },

  /**
   * 生命周期函数--监听页面加载
   */
  onLoad: function (options) {
    isChoice = false
    setTopHeight(this)
    app.setTitleWidth(this)
  },

  RegionChange: function (e) {
    this.setData({
      region: e.detail.value
    })
  },

  onShow: function (res) {
    var t = this,
      gender = (allData.gender == 1 ? true : false)
    t.setData({
      iconUrl: allData.iconUrl,
      userName: allData.name,
      age: allData.age,
      gender: gender,
      phoneNum: allData.phoneNum,
      region: allData.address,
      info: allData.info
    })
  },

  switch2Change(e) {
    this.setData({
      gender: e.detail.value
    })
  },
  /**
   * 选择图片
   */
  choice_img() {
    var t = this
    wx.chooseImage({
      count: 1,
      success: function (res) {
        isChoice = true
        t.setData({
          iconUrl: res.tempFilePaths[0]
        })
      }
    })
  },

  /**
   * input的value改变
   * @param {*} e 
   */
  changeInput(e) {
    var key = e.target.dataset.name, value = e.detail.value
    // console.log(value)
    switch (key) {
      case 'name':
        name = value
        break;
      case 'age':
        age = value
        break;
      case 'phone':
        phone = value
        break;
      case 'info':
        info = value
        break;
    }
  },

  /**
   * 更新头像
   */
  updataUserInfo() {
    var t = this
    UI.showLoading('提交中...')
    //如果头像没有被更改，则不执行上传返回，否则先上传头像后再执行更新资料方法
    !isChoice ? t.updateData() : utils.UpLoadFile('upload', function (res) {
      res.status == 0 ?
        t.updateData(res.data.fileUrl) :
        UI.showToast('上传头像失败，请稍候重试！')
    }, t.data.iconUrl)
  },

  /**
   * 更新资料
   */
  updateData(e) {
    var t = this.data, region = t.region
    utils.PUT('userInfo', {
      "Address": t.region.length == 3 ? region[0] + ',' + region[1] + ',' + region[2] : region,
      "Gender": t.gender ? '1' : '2',
      "Icon": isChoice ? e : t.iconUrl,
      "Info": info ? info : t.info,
      "Name": name ? name : t.userName,
      "PhoneNum": phone ? phone : t.phoneNum,
      "WorkAge": age ? age : t.age
    }, (res) => {
      // console.log(res)
      res.status == 0 ?
        allData.isGetUser = false & allData.that.getUserInfo() & UI.showToast('保存成功')
        : UI.showToast('保存失败')
      },1)
  }
})