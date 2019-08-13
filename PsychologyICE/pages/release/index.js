const app = getApp(),
  utils = require('../../utils/util.js'),
  allData = getApp().globalData,
  api = require('../../utils/api.js'),
  UI = require('../../utils/Interface.js')
/**
 * communityClass：所有资讯获取的分类
 * isUpdate：是否更新 isUpdate： 0:添加资料  1:更新资料  other:错误信息
 * lastClassID：更改资讯分类前的分类id
 * page：页面对象，用来获取上级页面
 * picUrl：资讯图片 仅用于更新资讯
 */
var communityClass, isUpdate = 0, lastClassID=-1, page, picUrl
Page({
  data: {
    dataArray: [],//资讯分类单选数组
    dataIndex: 0,//普通选择默认选中
    cursor: 0,//内容输入的字符串总长度
    title: '',//资讯标题初始化
    content: '',//资讯内容初始化
  },

  /**
   * 生命周期函数--监听页面加载
   */
  onLoad: function (e) {
    var t = this, dataArray = [], data
    app.setTitle1Width(t, '发布')
    utils.GET('getInformationClass', (res) => {//获取资讯所有分类
      if (res.status != 0) {
        isUpdate = 2
        UI.howToast('错误:' + res.msg,1)
        return null
      }
      communityClass = data = res.data
      for (let i = 0; i < data.length; i++) {
        dataArray[i] = data[i].ClassName
      }
      t.setData({ dataArray })
    })
    if (e.pic) {
      lastClassID = e.class_id
      picUrl = e.pic
      utils.GET('InformationInfo', function (res) {
        res.status == 0 ? (isUpdate = 1) & t.setData({
          'img[0]': e.pic,
          title: res.data.title,
          content: res.data.content,
          dataIndex: e.class_id,
          cursor: res.data.content.length
        }) : (isUpdate = 2) & t.setData({
          indexTitle: '请求错误'
        }) & UI.showToast('错误:' + res.msg,1)
      })
    } else isUpdate = 0
  },
  /**
   * 所属领域 picker value 改变时触发 change 事件
   * @param {*} e
   */
  Change(e) {
    this.setData({ dataIndex: e.detail.value })
  },
  /**
   * picker 设置value
   * @param {*} e 
   */
  ColumnChange(e) {
    var value = e.detail.value
    this.setData({ dataIndex: value });
  },
  /**
   * 提交按钮事件
   * @param {*} e 
   */
  sendCommunity(e) {
    page = getCurrentPages()
    //判断是否是更新资料 
    switch (isUpdate) {
      case 0:
        this.postCommunity(e)
        break;
      case 1:
        this.putCommunity(e)
        break;
      default: UI.showToast('获取资料错误')
        break;
    }
  },
  /**
   * 检测填写的内容是否符合要求
   * @param {Array} data 
   */
  isTrue(data) {
    if (data.title == '') {
      wx.showModal({
        content: '请输入资讯标题',
        showCancel: false,
      })
      return false
    }
    if (!this.data.img) {
      wx.showModal({
        content: '未选择资讯图片',
        showCancel: false,
      })
      return false
    }
    if (data.textarea == '') {
      wx.showModal({
        content: '请输入资讯内容',
        showCancel: false,
      })
      return false
    }
    if (data.textarea.length < 30) {
      wx.showModal({
        content: '您输入的文字小于30字,请补充!',
        showCancel: false,
      })
      return false
    }
    return true
  },
  /**
   * 提交更新信息
   * @param {*} e 
   */
  putCommunity(e) {
    const data = e.detail.value, t = this
    //填写检测，检测是否有填写不正常
    if (!t.isTrue(data)) return null
    //判断资讯图片是否被改变，如果改变过则先上传在更新资料
    if (t.data.img[0] != picUrl) {
      UI.showLoading('上传中')
      utils.UpLoadFile('upload', (e) => {
        e.status == 0 ? UI.showLoading('更新中')
          & (picUrl = e.data.fileUrl) & t.updateInfo(data)
          : UI.showToast(e.msg)
      }, t.data.img[0])
    } else {
      UI.showLoading('更新中')
      t.updateInfo(data)
    }
  },
  /**
   * 更新信息
   * @param {Array} data 
   */
  updateInfo(data) {
    const t = this
    utils.PUT('InformationInfo', {
      ClassifyId: communityClass[t.data.dataIndex].Id,
      Content: data.textarea,
      Icon: picUrl,
      Title: data.title,
      Uid: allData.userID
    }, (res) => {
      res.status == 0 ? isOK(t, '更新成功') : UI.showToast('错误:' + res.msg,1)
    }, 1)
  },
  /**
   * 添加资讯
   * @param {*} e 
   */
  postCommunity(e) {
    const data = e.detail.value, t = this
    if (!t.isTrue(data)) return null
    UI.showLoading('上传中')
    //上传资讯图片
    utils.UpLoadFile('upload', (e) => {
      e.status == 0 ? UI.showLoading('发布中')
        & utils.POST('Information', {//添加资讯详情
          Title: data.title,
          ClassifyId: communityClass[t.data.dataIndex].Id,//资讯分类
          Content: data.textarea,//资讯内容
          Icon: e.data.fileUrl,//资讯图片
          Uid: allData.userID,//用户ID
          AuthorType:1
        }, (res) => {
          res.status == 0 ? isOK(t, '发布成功') : UI.showToast('错误:' + res.msg,1)
        }, 1)
        : UI.showToast('错误:' + e.msg,1)
    }, t.data.img[0])
  },
  /**
   * 选择资讯图片
   */
  ChooseImage() {
    wx.chooseImage({
      count: 1, //默认9
      sizeType: ['original', 'compressed'], //可以指定是原图还是压缩图，默认二者都有
      sourceType: ['album'], //从相册选择
      success: (res) => {
        var img = []
        img[0] = picUrl = res.tempFilePaths[0]
        this.setData({ img })
      }
    });
  },
  /**
   * 查看资讯图片大图
   */
  ViewImage() {
    wx.previewImage({ current: this.data.img[0], urls: this.data.img })
  },
  /**
   * 删除资讯图片
   */
  DelImg() {
    this.setData({ img: null })
  },
  /**
   * 实时获取 内容编辑框 输入的字符串长度
   * @param {*} e 
   */
  textarea(e) {
    this.setData({ cursor: e.detail.cursor })
  }
})
/**
 * 添加或者更新成功的后续操作
 * @param {Object} that 
 * @param {String} title 
 */
function isOK(that, title) {
  UI.showToast(title)
  //判断分类ID是否被改变
  if (lastClassID != that.data.dataIndex && lastClassID!=-1) {
    page[page.length - 2].Initialization(lastClassID)
    page[page.length - 2].Initialization(that.data.dataIndex)
  } else {
    page[page.length - 2].Initialization(that.data.dataIndex)
  }
  setTimeout(() => {
    wx.navigateBack({
      delta: 1
    })
  }, 1500)
}