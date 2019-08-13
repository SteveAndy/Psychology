const app = getApp(),
  utils = require('../../../utils/util.js'),
  allData = getApp().globalData,
  api = require('../../../utils/api.js'),
  UI = require('../../../utils/Interface.js')
var doubleClick = 0, isReply = true
Page({

  data: {
    loadingImg: api.loadingImgUrl + 'loading-1.gif',
    ViewHeight: wx.getSystemInfoSync().screenHeight,
    InputBottom: 0,
    InputValue: '',
    user_id: wx.getStorageSync('user_id')
  },

  onLoad: function (options) {
    app.setTitle1Width(this, '详情')
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
        }) & UI.showToast('错误:' + res.msg, 1)
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

  InputFocus(e) {
    this.setData({
      InputBottom: e.detail.height
    })
  },

  InputBlur(e) {
    this.setData({
      InputBottom: 0
    })
  },
  /**
   * 提交回复信息
   * @param {*} data from提交的数据
   */
  Reply(data) {
    const content = data.detail.value.content, t = this
    if (content == '') {
      UI.showToast('回复内容不能为空')
      return null
    } else if (content.length < 10) {
      UI.showToast('回复内容不能太少啦')
      return null
    }
    UI.showLoading('回复中...')
    isReply ? Reply(t, content) : Updata(t, content)
  },
  /**
   * 菜单操作
   * @param {*} e 
   */
  menu(e) {
    const index = e.currentTarget.dataset.index, type = e.currentTarget.dataset.type, message_list = this.data.CommunityInfo.message_list[index]
    allData.CommunityReple = message_list.id //将要被操作的回复id存起来
    switch (type) {
      case 'updata':
        isReply = false
        this.setData({ InputValue: message_list.content })
        break;
      case 'delete':
        isReply = true
        Delete(this)
        break;
    }
  },
})

/**
 * 删除回复内容
 * @param {*} that 
 */
function Delete(that) {
  utils.DELETE('CommunityUpdataDelete', {}, (res) => {
    res.status == 0 ? UI.showToast('删除成功')
      & that.getCommunityInfo()//重新获取帖子详情，刷新数据
      : UI.showToast('错误:' + res.msg, 1)
  })
}

/**
 * 修改回复内容
 * @param {*} that 
 * @param {*} content 
 */
function Updata(that, content) {
  utils.PUT('CommunityUpdataDelete', {
    Cid: allData.CommunityInfo, //帖子ID
    Content: content,//回复内容
    UserId: allData.userID
  }, (res) => {
    res.status == 0 ? that.setData({ InputValue: '' })
      & UI.showToast('修改成功')
      & that.getCommunityInfo()//重新获取帖子详情，刷新数据
      : UI.showToast('错误:' + res.msg, 1)
  })
}

/**
 * 回复帖子
 * @param {*} that 
 * @param {*} content 
 */
function Reply(that, content) {
  utils.POST('CommunityReply', {
    Cid: allData.CommunityInfo, //帖子ID
    Content: content,//回复内容
    UserId: allData.userID
  }, (res) => {
    res.status == 0 ? that.setData({ InputValue: '' })
      & UI.showToast('回复成功')
      & that.getCommunityInfo()//重新获取帖子详情，刷新数据
      : UI.showToast('错误:' + res.msg, 1)
  })
}