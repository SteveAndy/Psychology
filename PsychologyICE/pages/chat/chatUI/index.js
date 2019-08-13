const app = getApp(), utils = require('../../../utils/util.js'),
  Jim = require('../../../utils/Jim.js'), allData = app.globalData
var chatContentList = [], msg = '';

/**
 * 初始化数据
 */
function initData(that, pic) {
  setScrollViewHeight(that)
  const messageList = allData.Off_lineMsg
  chatContentList = !messageList ? null : messageList[allData.chatName]
  chatContentList = {
    name: allData.chatName,
    msgs: !chatContentList ? [] : chatContentList.msgs,
    unread_msg_count: !chatContentList ? [] : chatContentList.unread_msg_count
  }
  that.setData({
    StatusBar: allData.StatusBar,
    CustomBar: allData.CustomBar,
    chatContentList,
    selfAvatar: allData.iconUrl,
    otherAvatar: !allData.iconArr ? pic : allData.iconArr[allData.chatName],
  });
}

/**
 * 计算聊天scroll-view实际占界面的高度
 */
function setScrollViewHeight(that) {
  var TopHeight = allData.CustomBar, //获取顶部高度
    ScreenHeight = wx.getSystemInfoSync().windowHeight, //获取屏幕高度
    query = wx.createSelectorQuery();
  //获取底部输入框所占屏幕高度  单位px；
  query.select('#inputView').boundingClientRect((res) => {
    that.setData({
      ScrollViewHeight: ScreenHeight - TopHeight - res.height
    })
  }).exec();
}
Page({

  data: { },

  onLoad: function (e) {
    initData(this, !e ? null : e.pic)
    this.setToMsg()
  },

  setToMsg(){
    if (chatContentList.msgs.length>0) {
      this.setData({
        toMsg: chatContentList.msgs.length - 1
      })
    }
  },

  onReady(){
    this.setToMsg()
  },

  setChatContentList(content) {
    var length = chatContentList?chatContentList.msgs.length:0
    if (!chatContentList){
      chatContentList = { name: allData.chatName,msgs: [] }
    }
    chatContentList.msgs[length] = { content: { msg_body: { text: content }, create_time: utils.formatTime(1)} }
    this.setData({
      chatContentList, toMsg: chatContentList.msgs.length - 1, MessageInputValue: ''
    })
    Jim.sendSingleMsg({
      target_username: chatContentList.name,
      content: msg
    })
  },

  //将输入框输入的内容保存到msg
  inputSendMessage(e) {
    var value = e.detail.value;
    msg = value
  },

  /**
   * 发送点击监听
   */
  sendClick: function (e) {
    this.setChatContentList(e.detail.value)
  },
});
