const app = getApp(),
  jim = require('../../utils/Jim.js'),
  allData = app.globalData,
  api = require('../../utils/api.js'),
  utils = require('../../utils/util.js'),
  UI = require('../../utils/Interface.js')
/**
 * 计算顶部高度
 */
function getNavHeight(that) {
  var query = wx.createSelectorQuery(); //单位px；
  query.select('#nav').boundingClientRect(function (rect) {
    if (rect) {
      that.setData({
        NavHeight: rect.height + app.globalData.CustomBar
      })
    }
  }).exec();
}
Component({
  data: {
    TabCur: 0,
    scrollLeft: 0,
    news: {
      classNames: 'wux-animate--fadeInLeft',
      enter: true,
      exit: false,
      in: true,
    },
    special: {
      classNames: 'wux-animate--fadeInRight',
      enter: true,
      exit: false,
      in: false,
    },
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
    onLoad: function (options) {
      getNavHeight(this)
      app.setTitleWidth(this, true);
      !allData.CommunityClass ? Initialization(this) : null
    },

    /**
     * 顶部TAB被点击事件
     * @param {*} e 
     */
    tabSelect(e) {
      var id = e.currentTarget.dataset.id,
        news = 'news.in',
        special = 'special.in',
        show = id == '0' ? true : false,
        noshow = id == '1' ? true : false
      this.setData({
        TabCur: id,
        scrollLeft: (e.currentTarget.dataset.id - 1) * 60,
        [news]: show,
        [special]: noshow
      })
    },

    /**
     * 跳转专题页面
     * @param {*} e 
     */
    toSpecial(e) {
      var title = e.currentTarget.dataset.title,
        id = e.currentTarget.dataset.id
      wx.navigateTo({
        url: UI.page.forum_special + "?title=" + title + "&id=" + id,
      })
    },


    /**
     * 获取帖子列表
     * @param {*} that 
     */
    getCommunity() {
      var that = this
      utils.GET('Community', (e) => {
        e.status == 0 ? that.setData({
          new_list: e.data
        }) : that.setData({
          new_list: 'ErrorNetwork'
        }) &
          UI.showToast('错误:' + e.msg)
      }, {
          sortby: 'time',
          order: 'desc'
        })
    },

    send() {
      wx.navigateTo({
        url: UI.page.release
      })
    },
  }
})

/**
 * 初始化
 * @param {*} that 
 */
function Initialization(that) {
  UI.showLoading('加载中')
  utils.GET('getCommunity_class', (res) => {
    res.status == 0 ? that.setData({
      sp_list: res.data
    }) : that.setData({
      sp_list: 'ErrorNetwork'
    }) &
      UI.showToast('错误:' + res.msg)
    that.getCommunity()
  })
}