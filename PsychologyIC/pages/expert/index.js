const app = getApp(), allData = app.globalData, api = require('../../utils/api.js'),
  utils = require('../../utils/util.js'), jim = require('../../utils/Jim.js'), ui = require('../../utils/Interface.js')
Component({
  data: {
    width: wx.getSystemInfoSync().windowWidth,
    info_class: { style: '', index: '' },
  },
  pageLifetimes: {
    show() {
      if (typeof this.getTabBar === 'function' &&
        this.getTabBar()) {
        this.getTabBar().setData({
          selected: 1,
          msgNum: allData.msgNum
        })
        jim.setThat(this)
      }
    }
  },
  methods: {
    /**
     * 生命周期函数--监听页面加载
     */
    onLoad: function (options) {
      app.setTitleWidth(this, true)
      Initialization(this)
    },

    /**
     * 拨打电话
     * @param {*} e 
     */
    call: function (e) {
      var num = this.data.expert_list[e.currentTarget.dataset.index].PhoneNum
      if (num && num != '') {
        wx.makePhoneCall({
          phoneNumber: num
        })
      } else {
        ui.showToast('该专家没有预留联系方式')
      }

    },

    /**
     * 查看详细介绍
     * @param {*} e 
     */
    seeInfo(e) {
      var index = e.currentTarget.dataset.id
      var style = this.data.info_class.style != '' &&
        this.data.info_class.index == index ?
        '' : "height: 100%;background-color:#FFF;" +
        "word-break:break-all;white-space:pre-wrap;" +
        "line-height:35rpx;"
      this.setData({
        info_class: {
          style: style,
          index: index
        }
      })
    },

    Online(e) {
      allData.chatName = "e" + e.currentTarget.dataset.id
      allData.chatId = 0
      wx.navigateTo({ url: ui.page.chat_ui + '?pic=' + e.currentTarget.dataset.pic })
    },
  }
})

function Initialization(that) {
  var expert_list
  that.setData({
    is_super: allData.is_super
  })
  wx.showLoading({
    title: '加载中',
    mask: true,
  })

  //获取专家分类
  utils.GET('getExpertClass', (res) => {
    //Do....
    //获取专家列表
    utils.GET('getExpert', (res) => {
      expert_list = res.data
      res.status == 0 ? that.setData({ expert_list })
        : that.setData({ expert_list: 'ErrorNetwork' }) & ui.showToast('错误:' + res.msg, 1)
      wx.hideLoading()
    })
  })

}