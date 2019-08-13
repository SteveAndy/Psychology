const app = getApp(), api = require('api.js'),
    USER_PAGE = {
        index: '/pages/index/',
        user: '/pages/user/',
        forum: '/pages/forum/',
        release: '/pages/release/',
        chat: '/pages/chat/',
        info: '/pages/info/',
        webview: '/pages/webview/',
        feedback:'/pages/feedback/'
    }
/**
 * 保存图片文件
 * @param {*} that 
 * @param {*} url 链接
 * @param {*} name 保存名称
 */
function seveImage(that, url, name) {
    wx.downloadFile({
        url: api.loadingImgUrl + url,
        success: function (res) {
            if (res.statusCode === 200) {
                // 第二步: 使用小程序的文件系统，通过小程序的api获取到全局唯一的文件管理器
                const fs = wx.getFileSystemManager()
                //  fs为全局唯一的文件管理器。那么文件管理器的作用是什么，我们可以用来做什么呢？
                //   文件管理器的作用之一就是可以根据临时文件路径，通过saveFile把文件保存到本地缓存.
                fs.saveFile({
                    tempFilePath: res.tempFilePath, // 传入一个临时文件路径
                    success(res) {
                        console.log('图片缓存成功', res.savedFilePath) // res.savedFilePath 为一个本地缓存文件路径
                        wx.hideLoading()
                        wx.setStorageSync(name, res.savedFilePath)
                        that.onReady()
                    }
                })
            } else {
                console.log('响应失败', res.statusCode)
            }
        }
    })
}

/**
 * 弹出提示
 * @param {String} title 标题
 * @param {Int} type 输出类型 1:错误类型
 */
function showToast(title, type) {
    wx.showToast({
        title,
        icon: 'none'
    })
    switch (type) {
        case 1:
            console.log("Error:" + title)
            break;

        default: console.log("Toast:" + title)
            break;
    }
}

/**
 * 显示加载框
 * @param {*} title 标题
 */
function showLoading(title) {
    wx.showLoading({
        title: title,
        mask: true
    })
    console.log("showLoading:" + title)
}


module.exports = {
    seveImage,
    showLoading,
    showToast,
    page: {
        index: USER_PAGE.index + 'index',
        user: USER_PAGE.user + 'user',
        user_auth: USER_PAGE.user + 'auth',
        user_editInfo: USER_PAGE.user + 'editInfo',
        user_apply_index: USER_PAGE.user + 'apply_expert/index',
        user_apply_details: USER_PAGE.user + 'apply_expert/details',
        forum: USER_PAGE.forum + 'forum',
        forum_special: USER_PAGE.forum + 'special',
        forum_info: USER_PAGE.forum + 'info/index',
        release: USER_PAGE.release + 'index',
        chat: USER_PAGE.chat + 'chat',
        chat_ui: USER_PAGE.chat + 'chatUI/index',
        info: USER_PAGE.info + 'index',
        webview: USER_PAGE.webview + 'webview',
        feedback:USER_PAGE.feedback+'index'
    }
}