import WxValidate from '../../../utils/WxValidate.js'
const app = getApp(),
  utils = require('../../../utils/util.js'),
  allData = getApp().globalData,
  UI = require('../../../utils/Interface.js')
var subArr = [],
  communityClass, Validate, Openid

/**
 * 计算view实际占界面的高度
 */
function setViewHeight(that) {
  const query = wx.createSelectorQuery(); //单位px；
  query.select('#top').boundingClientRect(function(rect) {
    that.setData({
      ViewHeight: rect.height
    })
  }).exec();
}

Page({

  /**
   * 页面的初始数据
   */
  data: {
    StatusBar: allData.StatusBar,
    CustomBar: allData.CustomBar,
    imgList: [],
    photo: [],
    dataIndex: [0, 0],
    region: [],
    dataArray: [],
    show: false,
  },
  onLoad: function(e) {
    setViewHeight(this)
    initValidate()
    Openid = e.openid
    var t = this,
      dataArray = [],
      mainArray = [],
      subArray, data, sub
    utils.GET('getCommunity_class', (res) => { //获取专家所有分类
      communityClass = data = res.data
      console.log(res)
      for (let i = 0; i < data.length; i++) {
        subArray = []
        sub = !data[i].sub ? 0 : data[i].sub.length
        for (let a = 0; a < sub; a++) {
          subArray[a] = data[i].sub[a].Title
        }
        mainArray[i] = data[i].name
        subArr[i] = subArray
      }
      dataArray = [mainArray, subArr[0]]
      t.setData({
        dataArray
      })
    })
  },
  /**
   * 选择图片
   * @param {*} e 
   */
  ChooseImage(e) {
    const t = this,
      type = e.currentTarget.dataset.type
    wx.chooseImage({
      count: type == 'photo' ? 1 : 2, //默认9
      sizeType: ['compressed'],
      success: (res) => {
        updataImg(t, res.tempFilePaths, type)
      }
    });
  },
  /**
   * 查看图片
   * @param {*} e 
   */
  ViewImage(e) {
    const url = e.currentTarget.dataset.url,
      type = e.currentTarget.dataset.type,
      data = this.data
    wx.previewImage({
      urls: type == 'photo' ? data.photo : data.imgList,
      current: url
    });
  },
  /**
   * 删除图片
   * @param {*} e 
   */
  DelImg(e) {
    const imgList = this.data.imgList,
      index = e.currentTarget.dataset.index,
      type = e.currentTarget.dataset.type
    imgList.splice(index, 1);
    if (type == 'photo') {
      this.setData({
        photo: []
      })
    } else {
      this.setData({
        imgList
      })
    }
  },
  /**
   * 所属领域 picker value 改变时触发 change 事件
   * @param {*} e 
   */
  Change(e) {
    this.setData({
      dataIndex: e.detail.value
    })
  },
  /**
   * picker 设置value
   * @param {*} e 
   */
  ColumnChange(e) {
    var value = e.detail.value,
      column = e.detail.column,
      index;
    let data = {
      dataArray: this.data.dataArray,
      dataIndex: this.data.dataIndex
    };
    data.dataIndex[column] = value
    switch (column) {
      case 0:
        data.dataArray[1] = subArr[value]
        data.dataIndex[1] = 0;
        break;
    }
    this.setData(data);
  },
  /**
   * 地址选择
   * @param {*} e 
   */
  RegionChange(e) {
    this.setData({
      region: e.detail.value
    })
  },
  /**
   * 提交信息
   */
  formSubmit(e) {
    const data = e.detail.value,
      t = this.data,
      Address = t.region.length > 0 ? (t.region[0] + ',' + t.region[1] + ',' + t.region[2]) : ''
    //校验表单
    if (!Validate.checkForm(data)) {
      UI.showToast(Validate.errorList[0].msg)
      return false
    }
    if (Address == '') {
      UI.showToast('请选择地址')
      return null
    }
    if (t.photo.length == 0) {
      UI.showToast('请上传本人头像')
      return null
    }
    if (t.imgList.length < 1) {
      UI.showToast('请上传身份证正反面')
      return null
    }
    UI.showLoading('提交中...')
    allData.code = data.code
    utils.POST('apply', {
      Address: Address,
      Age: data.workAge,
      ClassifyId: communityClass[t.dataIndex[0]].sub[t.dataIndex[1]].Id,
      IdcardB: t.imgList[0],
      IdcardF: t.imgList[1],
      Openid: Openid,
      Icon: allData.iconUrl,
      Info: data.introduction,
      Name: data.name,
      PhoneNum: data.phoneNum,
      Photo: t.photo[0],
    }, (res) => {
      res.status == 0 ? UI.showToast('提交审核成功') & setTimeout(() => {
        wx.redirectTo({
          url: UI.page.user_apply_details + '?status=1'
        })
      }, 1500) : UI.showToast('错误' + res.msg, 1)
    }, 1)
  },
  /**
   * 控制 获取登录码 dialog显示状态
   */
  getCode() {
    this.setData({
      show: !this.data.show
    })
  },
  /**
   * 跳转小程序
   */
  ToMinProgram() {
    const t = this
    wx.navigateToMiniProgram({
      appId: app.setinfo.MiniProgramAppid
    })
    t.getCode()
  }
})

//验证函数
function initValidate() {
  const rules = {
    code: {
      required: true,
      minlength: 4,
    },
    name: {
      required: true,
      minlength: 2,
    },
    idCard: {
      required: true,
      idcard: true
    },
    workAge: {
      required: true,
      minlength: 1,
      maxlength: 2,
    },
    phoneNum: {
      required: true,
      tel: true
    },
    introduction: {
      required: true,
      minlength: 10,
    },
  }
  const messages = {
    code: {
      required: '请填输入登录码',
      minlength: '请输入正确的登录码',
    },
    name: {
      required: '请填写姓名',
      minlength: '请输入正确的姓名'
    },
    idCard: {
      required: '请输入身份证号码',
      idcard: '请输入正确的身份证号码',
    },
    workAge: {
      required: '请输入从业年龄',
      minlength: '请输入正确的从业年龄',
      maxlength: '请输入正确的从业年龄',
    },
    phoneNum: {
      required: '请输入手机号',
      tel: '请输入正确的手机号'
    },
    introduction: {
      required: '请输入您的个人简介',
      minlength: '个人简介不得少于10个字'
    },
  }
  Validate = new WxValidate(rules, messages)
}
/**
 * 上传图片
 * @param {this} that 当前对象
 * @param {String[]} pic 图片路径集合
 * @param {String} type 类型
 */
function updataImg(that, pic, type) {
  const data = pic
  for (let i = 0; i < data.length; i++) {
    UI.showLoading('上传中')
    utils.UpLoadFile('upload', (res) => {
      const imgUrl = res.data.fileUrl
      if (type == 'photo') {
        that.setData({
          "photo[0]": imgUrl
        })
      } else {
        var imgList = 'imgList[' + (that.data.imgList.length != 0 ? 1 : 0) + ']'
        that.setData({
          [imgList]: imgUrl
        })
      }
    }, data[i])
  }

}