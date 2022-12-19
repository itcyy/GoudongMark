<template>
  <div style="border-bottom: 1px solid #dfdfdf; font-size: 14px; width: 100%;">
    <div
      style="
        display: flex;
        height: 40px;
        padding: 0 200px;
        background-color: #00bfff;
      "
    >
      <div
        class="top-title"
        style="
          padding-right: 20px;
          border-right: 1px solid rgb(230 230 230 / 40%);
          line-height: 40px;
          color: #ffffff;
        "
      >
        手机版
      </div>
      <div
        class="top-title"
        style="
          padding-right: 20px;
          border-right: 1px solid rgb(230 230 230 / 40%);
          line-height: 40px;
          color: #ffffff;
          margin-left: 20px;
        "
      >
        留言反馈
      </div>
      <div
        class="top-title"
        style="
          padding-right: 20px;
          border-right: 1px solid rgb(230 230 230 / 40%);
          line-height: 40px;
          color: #ffffff;
          margin-left: 20px;
        "
      >
        会员中心
      </div>
      <div
        class="top-title"
        style="line-height: 40px; color: #ffffff; margin-left: 20px"
      >
        收藏本站
      </div>
      <div style="flex: 1; display: flex">
        <div style="flex: 1"></div>
        <div style="width: 170px; display: flex">
          <el-button
            type="text"
            style="color: #ffffff"
            class="top-title"
            @click="JumpeLogin(0)"
            >登录</el-button
          >
          <el-button
            type="text"
            style="color: #ffffff; margin-right: 10px"
            class="top-title"
            @click="JumpeLogin(1)"
            >注册</el-button
          >
          <div
            style="
              background-color: rgb(223 221 221 / 40%);
              display: flex;
              width: 100px;
              padding: 0 10px;
            "
            class="top-title"
          >
            <img
              src="../../components/img/购物车.png"
              style="
                width: 20px;
                height: 20px;
                margin-top: 10px;
                margin-right: 5px;
              "
            />

            <el-button
              type="text"
              style="color: #ffffff"
              class="top-title"
              @click="JumpShoppingCart()"
              >购物车</el-button
            >
          </div>
        </div>
      </div>
    </div>
    <div style="margin: 0 200px; height: 80px; line-height: 80px">
      <div style="display: flex">
        <img
          src="../../components/img/狗.png"
          style="cursor: pointer; height: 60px; margin-top: 10px"
        />
        <div
          class="title"
          style="
            margin-right: 20px;
            cursor: pointer;
            color: gray;
            margin-left: 30px;
          "
        >
          <router-link to="/">首页</router-link>
        </div>
        <div
          class="title"
          style="margin-right: 20px; cursor: pointer; color: gray"
        >
          电子数码
        </div>
        <div
          class="title"
          style="margin-right: 20px; cursor: pointer; color: gray"
        >
          家居百货
        </div>
        <div
          class="title"
          style="margin-right: 20px; cursor: pointer; color: gray"
        >
          电脑硬件
        </div>
        <div
          class="title"
          style="margin-right: 20px; cursor: pointer; color: gray"
        >
          关于我们
        </div>
        <div
          class="title"
          style="margin-right: 20px; cursor: pointer; color: gray"
        >
          联系我们
        </div>
        <div style="flex: 1; text-align: right">
          <el-input
            placeholder="请输入内容"
            v-model="input"
            style="width: 300px"
          >
            <el-button slot="append" icon="el-icon-search"></el-button>
          </el-input>
        </div>
      </div>
    </div>
  </div>
</template>
 
<script>
export default {
  data() {
    return {
      input: "",
    };
  },
  mounted() {},
  methods: {
    homePage() {
      this.$router.push({ path: "/" });
    },
    JumpeLogin(index) {
      this.$router.push({ path: "/Login", query: { id: index } });
    },
    JumpShoppingCart() {
      this.$router.push({ path: "/shoppingcart" });
    },
  },
};
</script>
 
<style>
/* 其他组件宽度被flex */
.top{
  width: 100%;
}

.top-title:hover {
  cursor: pointer;
  color: #e5e5e5 !important;
}

.title {
  font-weight: bold;
  font-size: 18px;
}

.title:hover {
  color: #c250c6 !important;
}
</style>