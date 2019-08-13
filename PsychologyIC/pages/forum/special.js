const app = getApp(),
    allData = app.globalData,
    api = require('../../utils/api.js'),
    utils = require('../../utils/util.js'),
    ui = require('../../utils/Interface.js')
Page({

    data: {
    },

    onLoad: function (res) {
        app.setTitle1Width(this, res.title)
        Initialization(this, res.id)
    },

    onShareAppMessage: function () {

    },

    call(e) {
        wx.makePhoneCall({
            phoneNumber: e.currentTarget.dataset.ephone
        })
    },

    chat(e) {
        const id = e.currentTarget.dataset.id, pic = e.currentTarget.dataset.pic
        allData.chatName = id
        wx.navigateTo({ url: ui.page.chat_ui + "?pic=" + pic })
    },

})

function Initialization(that, id) {
    that.setData({ is_super: allData.is_super })
    wx.showLoading({
        title: '加载中',
        mask: true,
    })
    utils.GET('community', (e) => {
        // console.log(e)
        e.status == 0 ? that.setData({
            new_list: e.data
        }) : that.setData({ new_list: 'ErrorNetwork' }) &
            ui.showToast('错误:' + e.msg, 1)
    }, { sortby: 'time', order: 'desc', query: 'classifyId:' + id })
}