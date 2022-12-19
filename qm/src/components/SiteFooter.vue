<template>
  <div class="site-footer">
    <ul class="site-service">
      <li class="service-item"
      v-for="(item, index) in serviceInfo"
      :key="index">
        <a :href="item.url" target="_blank">
          <i :class="item.font"></i>
          <span>{{item.value}}</span>
        </a>
      </li>
    </ul>
    <div class="site-link">
      <div class="links">
        <div class="link-container"
        v-for="(list, key) in serviceLink"
        :key="key">
          <span class="link-title">{{list.value}}</span>
          <ul class="link-list">
            <li class="link-item"
            v-for="(item, index) in list.subLinks"
            :key="index">
              <a :href="item.url" @click="showWindow" v-if="item.value === '官方微信'">
                {{item.value}}
              </a>
              <a :href="item.url" target="_blank" v-else>
                {{item.value}}
              </a>
            </li>
          </ul>
        </div>
      </div>
      <div class="link-contact">
        <p class="phone">400-100-5678</p>
        <p class="desc">周一至周日 8:00-18:00<br>（仅收市话费）</p>
        <a href="https://support.kefu.mi.com/" target="_blank">
          <i class="iconfont icon-icon-test1"></i>
          <span>人工客服</span>
        </a>
      </div>
    </div>
  </div>
</template>

<script>
import bus from '@/bus.js'

export default {
  data () {
    return {
      serviceInfo: [
        {value: '预约维修服务', font:'iconfont icon-buoumaotubiao46', url: '/'},
        {value: '7天无理由退货', font:'iconfont icon-7tianwuliyoutuihuo', url: '/'},
        {value: '15天免费换货', font:'iconfont icon-15tianwuliyoutuihuo', url: '/'},
        {value: '满150元包邮', font:'iconfont icon-liwu', url: '/'},
        {value: '520余家售后网点', font:'iconfont icon-icon-test', url: '/'}
      ],
      serviceLink: [
        {value: '帮助中心', subLinks: [
          {value: '账户管理', url: '/'},
          {value: '购物指南', url: '/'},
          {value: '订单操作', url: '/'}
        ]},
        {value: '服务支持', subLinks: [
          {value: '售后政策', url: '/'},
          {value: '自助服务', url: '/'},
          {value: '相关下载', url: '/'}
        ]},
        {value: '线下门店', subLinks: [
          {value: '狗东之家', url: '/'},
          {value: '服务网点', url: '/'},
          {value: '授权体验店', url: '/'}
        ]},
        {value: '关于狗东', subLinks: [
          {value: '了解狗东', url: '/'},
          {value: '加入狗东', url: '/'},
          {value: '投资者关系', url: '/'}
        ]},
        {value: '关注我们', subLinks: [
          {value: '新浪微博', url: '/'},
          {value: '官方微信', url: 'javaScript:void(0);'},
          {value: '联系我们', url: '/'}
        ]},
        {value: '特色服务', subLinks: [
          {value: 'F 码通道', url: '/'},
          {value: '礼物码', url: '/'},
          {value: '防伪查询', url: '/'}
        ]}
      ]
    }
  },
  methods: {
    showWindow () {
      bus.$emit('weChatWindowEvent', true);
    }
  }
}
</script>

<style lang="less">
.site-footer {
    background-color: #2D3A4B;;
    position: relative;
    width: 1226px;
    height: auto;
    margin: 0 auto;
    .site-service {
      display: flex;
      padding: 27px 0;
      border-bottom: 1px solid #e0e0e0;
    }

    .site-link {
      display: flex;
      justify-content: space-between;
      padding: 40px 0;
      .links {
        display: flex;
      }
    }
  }

  .service-item {
    width: 19.8%;
    height: 25px;
    border-left: 1px solid #e0e0e0;
    font-size: 16px;
    line-height: 25px;
    text-align: center;
    &:first-child {
      border-left: none;
    }

    a {
      color: #616161;
      transition: color .2s;
      &:hover {
        color: #ff6709;
      }
      i {
        margin-right: 6px;
        font-size: 24px;
        line-height: 24px;
        vertical-align: -4px;
      }

      span {
        font-size: 16px;
        line-height: 25px;
      }
    }
  }

  .link-container {
    width: 160px;
    height: 112px;
    margin: 0;
    text-align: left;
    .link-title {
      display: block;
      margin: -1px 0 26px;
      font-size: 14px;
      line-height: 1.25;
      color: #424242;
    }

    .link-list, .link-item {
      margin: 10px 0 0;
      font-size: 12px;
      a {
        color: #757575;
        transition: color .2s;
        &:hover {
          color: #ff6709;
        }
      }
    }
  }

  .link-contact {
    width: 256px;
    height: 112px;
    border-left: 1px solid #e0e0e0;
    text-align: center;
    color: #616161;
    .phone {
      margin: 0 0 5px;
      font-size: 22px;
      line-height: 1;
      color: #ff6700;
    }

    .desc {
      margin: 0 0 16px;
      font-size: 12px;
    }

    a {
      display: inline-flex;
      justify-content: center;
      width: 118px;
      height: 28px;
      font-size: 12px;
      line-height: 28px;
      border: 1px solid #ff6700;
      background: #fff;
      color: #ff6700;
      transition: all .4s;
      &:hover {
        color: #fff;
        background-color: #f25807;
        border-color: #f25807;
      }
    }
  }
</style>