const app = getApp(), utils = require('../../../utils/util.js'), allData = getApp().globalData,
    UI = require('../../../utils/Interface.js')

Page({

    data: {
    },

    onLoad: function (options) {
        this.setData({
            status: options.status,
            msg: options.reject ? options.reject : '您的信息还在等待管理员审核哦'
        })
        !wx.getStorageSync('wave') ?
            UI.showLoading('加载中') & UI.seveImage(this, 'wave.gif', 'wave') :
            this.setData({
                wave: wx.getStorageSync('wave')
            })
    },
})