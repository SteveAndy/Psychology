const app = getApp(), api = require('../../utils/api.js'),
  utils = require('../../utils/util.js'), jim = require('../../utils/Jim.js'),
  UI = require('../../utils/Interface.js'), allData = app.globalData, LIMIT = "10"

import { $stopWuxRefresher } from '../../dist/index'
var loading = false, doubleClick = 0, item
/**
 * 计算scroll-view实际占界面的高度
 */
function setHeight(that) {
  var query = wx.createSelectorQuery();// 单位px；
  query.select('#nav').boundingClientRect(function (rect) {
    that.setData({
      contenHeight: rect.height
    })
  }).exec();
}
/**
 * 将分类和对应的资讯列表合并成一个数组
 * @param {Json} odata 获取的原有分类数据
 * @param {Int} ndata 获取到的对应分类的资讯列表
 * @param {Int} name 添加的字段名
 */
function dJSON(odata, ndata, name) {
  var JsonString
  JsonString = JSON.stringify(odata)
  JsonString = JsonString.substr(0, JsonString.length - 1);
  JsonString = "[" + JsonString + ",\"" + name + "\":" + ndata + ",\"moreStatus\":\"\"}]"
  JsonString = JSON.parse(JsonString)
  return JsonString[0]
}
Component({
  data: {
    TabCur: 0,
    ModalShow: false,
    scrollLeft: 0
  },
  pageLifetimes: {
    show() {
      if (typeof this.getTabBar === 'function' &&
        this.getTabBar()) {
        this.getTabBar().setData({
          selected: 0,
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
    onLoad: function () {
      app.setTitleWidth(this, '资讯')
    },
    /**
     * 获取分类列表
     */
    getInformationClass() {
      var t = this
      utils.GET('getInformationClass', function (res) {
        setHeight(t)
        res.status == 0 ? t.setData({
          informationClass: res.data
        }) & (allData.informationList = res.data) & t.Initialization(0) :
          UI.showToast('错误:' + res.msg, 1)
      })
    },
    /**
     * 初始化页面数据 tip:在utils.login()中执行
     * @param {Int} index 索引，用于查找要求请求资讯列表所属的分类Id
     * @param {Boolean} isStopRefresher 判断是否是下拉刷新触发的，如果是则关闭刷新动画
     */
    Initialization(index, isStopRefresher) {
      var t = this, data = allData.informationList[index], mdata
      //获取资讯列表
      utils.GET('Information', function (res) {
        // console.log(res)
        mdata = 'informationClass[' + index + ']'
        //判断是否是下拉刷新触发的，如果是则关闭刷新动画
        if (isStopRefresher) {
          loading = !loading
          $stopWuxRefresher("#refresher-" + index)
        }
        res.status == 0 ? t.setData({
          [mdata]: dJSON(data,
            JSON.stringify(res.data),
            "info_list")
        }) : t.setData({
          [mdata]: dJSON(data, "\"ErrorNetwork\"", "info_list")
        }) & UI.showToast('错误:' + res.msg, 1)
        wx.hideLoading()
      }, { query: 'ClassifyId:' + data.Id + ',Uid:' + allData.userID, sortby: 'time', order: 'desc' })
    },
    /**
     * 顶部分类Tab被单击事件
     * @param {*} e 
     */
    tabSelect(e) {
      if (loading)
        return
      var id = e.currentTarget.dataset.id, data = this.data.informationClass[id], t = this
      //判断是否初始化资讯列表，如果没有，则初始化  否则不执行减少请求次数
      if (!data.info_list) this.Initialization(id)
      t.setData({ TabCur: id })
    },
    /**
       * 下拉开始的回调函数
       */
    onPulling() { },
    /**
     * 正在刷新:下拉完成的回调函数
     */
    onRefresh() {
      var t = this
      loading = !loading
      /**
       * 下拉刷新时触发初始化获取资讯
       * tip:采用延时操作纯属为了UI好看
       */
      setTimeout(() => {
        t.Initialization(t.data.TabCur, true)
      }, 1000)
    },
    /**
     * 滚动到底部自动加载
     */
    onReachBottom() {
      var t = this, index = t.data.TabCur, data = t.data.informationClass[index],
        mdata = 'informationClass[' + index + '].info_list', moreStatus = 'informationClass[' + index + '].moreStatus',
        Status = t.data.informationClass[index].moreStatus
      
      //如果列表为空则不执行加载
      if (!data.info_list)return
      /**
       * 如果Status==over 则证明已经加载完了所有，就没必要再执行请求
       * 如果Status==loading 则证明正在向服务器请求，就不再执行，避免重复请求
       */
      if (Status == 'over' || Status == 'loading') return
      t.setData({ [moreStatus]: 'loading' })
      /**
       * 加载更多资讯
       * tip:采用延时操作纯属为了UI好看
       */
      setTimeout(() => {
        utils.GET('Information', function (res) {
          res.status == 0 ?
            (Array.isArray(res.data) ?
              t.setData({
                [mdata]: data.info_list.concat(res.data),
                [moreStatus]: res.data.length <= LIMIT ? 'over' : ''
              }) : t.setData({ [moreStatus]: 'over' })
            )
            : t.setData({ [moreStatus]: 'erro' })
        }, {
          query: 'ClassifyId:' + data.Id + ',Uid:' + allData.userID, sortby: 'time', order: 'desc', limit: LIMIT, offset: data.info_list.length })
      }, 1000)
    },
    /**
     * 在每次显示的时候判断是否获取userID，如果没有则执行登录
     */
    onShow: function () {
      if (!app.globalData.userID) {
        utils.login(this)
      }
    },
    /**
     * 顶部标题被双击返回顶部
     * @param {*} e 
     */
    doubleClick(e) {
      doubleClick++
      if (doubleClick == 2) {
        doubleClick = 0
        wx.pageScrollTo({
          scrollTop: 0,
          duration: 300
        })
      } else {
        setTimeout(() => {
          doubleClick = 0
        }, 450);
      }
    },
    /**
     * 跳转到 我的页面
     * @param {*} e 
     */
    toUser(e) {
      wx.navigateTo({ url: UI.page.user })
    },
    /**
     * 菜单显示/隐藏
     */
    Menu(e) {
      if (!this.data.ModalShow) {
        item = e.currentTarget.dataset.item
        allData.idInformationInfoId = item.Id
      }
      ModalShow(this)
    },
    /**
     * 修改
     */
    modify() {
      ModalShow(this)
      wx.navigateTo({ url: UI.page.release + '?pic=' + item.Icon + '&class_id=' + this.data.TabCur })
    },
    /**
     * 删除
     */
    delete() {
      const t = this
      ModalShow(t)
      UI.showLoading('删除中')
      utils.DELETE('InformationInfo', {}, (res) => {
        res.status == 0 ? UI.showToast('删除成功') & t.Initialization(t.data.TabCur) : UI.showToast('错误:' + res.msg, 1)
      }, 1)
    },
    /**
     * 显示动态面板
     * @param {*} e 
     */
    showModal(e) {
      this.setData({
        modalName: e.currentTarget.dataset.target
      })
    },
    /**
     * 隐藏动态面板
     * @param {*} e 
     */
    hideModal(e) {
      this.setData({
        modalName: null
      })
    },
    /**
     * 选择分类
     * @param {*} e 
     */
    tabSelect(e) {
      this.setData({
        TabCur: e.currentTarget.dataset.id,
        scrollLeft: (e.currentTarget.dataset.id - 1) * 60
      })
    }
  }
})

function ModalShow(that) {
  that.setData({ ModalShow: !that.data.ModalShow })
}