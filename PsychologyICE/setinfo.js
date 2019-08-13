const md5 = require('utils/md5.min.js'),
    JgAppkey = '2e4504420cbf98563ef546a1',
    JgSecret = '02c0c4e49cd0ecbf4753876a', 
    RandomStr = 'E422A978DE37196588531CD0C42010B5',
    TimeStamp = (new Date()).getTime(),
    RootUrl = 'https://nxxlzx.tpengyun.com/',
    MiniProgramAppid = 'wx737c3211676c7ff5',//跳转小程序appid
    //IM用户配置s
    UserCode = 'e',
    UserPwd = '123456';
function signature() {
    return md5("appkey=" + JgAppkey
        + "&timestamp=" + TimeStamp
        + "&random_str=" + RandomStr + "&key=" + JgSecret)
}
//3995eddf8c8a8d0dac84030cdb831941
module.exports = { JgAppkey,JgSecret,RandomStr,TimeStamp,signature,RootUrl,UserCode,UserPwd }