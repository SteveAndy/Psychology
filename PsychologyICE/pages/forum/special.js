const app = getApp(),allData = app.globalData,
api = require('../../utils/api.js'),utils = require('../../utils/util.js'),
UI = require('../../utils/Interface.js')
Page({

    data: {
    },

    onLoad: function (res) {
        app.setTitle1Width(this, res.title)
        Initialization(this, res.id)
    },

    onShareAppMessage: function () {

    }

})

function Initialization(that, id) {
    UI.showLoading('加载中')
    utils.GET('Community', (e) => {
        e.status == 0 ? that.setData({
            new_list: e.data
        }) : that.setData({ new_list: 'ErrorNetwork' }) &
            UI.showToast('错误:' + e.msg,1)
    }, { sortby: 'time', order: 'desc', query: 'classifyId:' + id })
}