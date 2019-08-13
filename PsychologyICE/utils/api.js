var root = require('../setinfo.js').RootUrl + "v1/";
function url(links) {
  var link, token = '?token=' + getApp().globalData.token,
    user = wx.getStorageSync("userInfo"),
    id = !getApp().globalData.userID ? wx.getStorageSync("user_id") : getApp().globalData.userID
  switch (links) {
    case 'upload':
      link = root + '/upload/upload'
      break
    case 'session_key':
      link = root + "users/session_key"
      break
      
    /* 用户 */
    case 'userInfo'://获取专家信息 or 更新用户信息
      link = root + 'expert/' + id
      break
    case 'userInfoOther'://获取其他用户信息
      link = root + 'users/'
      break
    case 'apply'://申请成为专家
      token = '?login_code=' + getApp().globalData.code
      link = root + 'expert_auth/' 
      break
    
    /* 资讯 */
    case 'getInformationClass'://获取文章分类列表
      link = root + 'info_class/'
      break
    case 'Information'://获取资讯列表 or 发布资讯
      link = root + 'info/'
      break
    case 'InformationInfo'://获取资讯详情 or 更新资讯详情 or 删除资讯
      link = root + 'info/' + getApp().globalData.idInformationInfoId
      break

    /* 问答 */
    case 'getCommunity_class':
      link = root + 'community_class/'
      break
    case 'Community'://获取帖子 or 发布帖子
      link = root + 'community/'
      break
    case 'CommunityInfo'://获取帖子详情
      link = root + 'community/' + getApp().globalData.CommunityInfo
      break
    case 'CommunityReply'://回复帖子
      link = root + 'community_reply/'
      break
    case 'CommunityUpdataDelete'://修改 or 删除回复帖子
      link = root + 'community_reply/' + getApp().globalData.CommunityReple
      break

    /* 意见反馈 */
    case 'opinion':
      link = root + 'opinion/'
      break;
  }
  return link + token
}
module.exports = {
  rootUrl: root,
  loadingImgUrl: require('../setinfo.js').RootUrl+'static/src/images/loading/',
  Anum: { //账号
    login: root + "expert/login",
    signup: root + "users/signup"
  },
  url,
}