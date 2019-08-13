const app = getApp(), allData = app.globalData, api = require('../../utils/api.js'),
  utils = require('../../utils/util.js'), UI = require('../../utils/Interface.js');
var qw, fb
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
  data: {},
  /**
   * 生命周期函数--监听页面加载
   */
  onLoad: function (options) {
    app.setTitleWidth(this)
    setTopHeight(this)
  },
  Input(e) {
    const type = e.currentTarget.dataset.type, value = e.detail.value
    switch (type) {
      case '0':
        qw = value
        break;
      case '1':
        fb = value
        break;
    }
  },

  postFB() {
    if (!qw) {
      UI.showToast('请输入QQ/微信')
      return
    }
    if (!fb) {
      UI.showToast('请输入反馈内容')
      return
    }
    if (fb.length < 10) {
      UI.showToast('请输入10个汉字以上')
      return
    }
    utils.POST('opinion', {
      Content: fb,
      Telephone: qw,
      UserId: allData.UserId
    }, (res) => {
      res.status == 0 ? UI.showToast('提交成功')
        & setTimeout(() => { wx.navigateBack({ delta: 1 }) }, 1500)
        : UI.showToast('提交错误:' + res.msg, 1)
    }, 1)
  }

})