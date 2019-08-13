const app = getApp(),IMG_SRC = '/images/tabBar/'
Component({
  options: {
    addGlobalClass: true,
  },
  data: {
    selected: 0,
    color: "#47505b",
    selectedColor: "#0081ff",
    backgroundColor: "#FFF",
    borderStyle: "white",
    msgNum: app.globalData.msgNum, // 简化的定义方式
    list: [{
        iconPath: IMG_SRC+"zixun.png",
        selectedIconPath: IMG_SRC+"zixun_i.png",
        pagePath: "/pages/index/index",
        text: "资讯"
      },
      {
        iconPath: IMG_SRC+"wenda.png",
        selectedIconPath: IMG_SRC+"wenda_i.png",
        pagePath: "/pages/forum/forum",
        text: "问答"
      },
      {
        iconPath: IMG_SRC+"message.png",
        selectedIconPath: IMG_SRC+"message_i.png",
        pagePath: "/pages/chat/chat",
        text: "消息"
      }
    ],
  },
  methods: {
    switchTab(e) {
      const data = e.currentTarget.dataset
      const url = data.path
      wx.switchTab({
        url
      })
      this.setData({
        selected: data.index,
        msgNum: app.globalData.msgNum,
      })
      //setTimeout(this._animation(e), 500)
    },
    _animation(e) {
      var anmiaton = e.currentTarget.dataset.class;
      var that = this;
      that.setData({
        animation: anmiaton
      })
      setTimeout(function() {
        that.setData({
          animation: ''
        })
      }, 1000)
    },
  }
})