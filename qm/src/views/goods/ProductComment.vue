<template>
  <div id="root">
    <div class="header1">
      <h3>商品评价</h3>
    </div>
    <div class="filter">
      <el-tabs v-model="tabName" @tab-click="handleClick()">
        <el-tab-pane label="全部" name="all">
          <div class="commentbox" v-for="item in CommentsList" :key="item.id">
            <el-container>
              <el-header height="40px">
                <!-- id, 星级，时间 -->
                <div class="userinfo">
                  <h3>{{ item.username }}</h3>
                  <p>{{ item.time }}</p>
                </div>
                <div class="love">
                  <el-rate :value="item.rates" :colors="lovecolors" disabled>
                  </el-rate>
                </div>
              </el-header>
              <el-main>
                <!-- 评论 -->
                <p>{{ item.comment }}</p>
                <el-divider></el-divider>
              </el-main>
            </el-container>
          </div>
        </el-tab-pane>
        <el-tab-pane label="好评" name="better">
          <div
            class="commentbox"
            v-for="item in TempCommentsList"
            :key="item.id"
          >
            <el-container>
              <el-header height="40px">
                <!-- id, 星级，时间 -->
                <div class="userinfo">
                  <h3>{{ item.username }}</h3>
                  <p>{{ item.time }}</p>
                </div>
                <div class="love">
                  <el-rate :value="item.rates" :colors="lovecolors" disabled>
                  </el-rate>
                </div>
              </el-header>
              <el-main>
                <!-- 评论 -->
                <p>{{ item.comment }}</p>
                <el-divider></el-divider>
              </el-main>
            </el-container></div
        ></el-tab-pane>
        <el-tab-pane label="中评" name="common"
          ><div
            class="commentbox"
            v-for="item in TempCommentsList"
            :key="item.id"
          >
            <el-container>
              <el-header height="40px">
                <!-- id, 星级，时间 -->
                <div class="userinfo">
                  <h3>{{ item.username }}</h3>
                  <p>{{ item.time }}</p>
                </div>
                <div class="love">
                  <el-rate :value="item.rates" :colors="lovecolors" disabled>
                  </el-rate>
                </div>
              </el-header>
              <el-main>
                <!-- 评论 -->
                <p>{{ item.comment }}</p>
                <el-divider></el-divider>
              </el-main>
            </el-container></div
        ></el-tab-pane>
        <el-tab-pane label="差评" name="bad"
          ><div
            class="commentbox"
            v-for="item in TempCommentsList"
            :key="item.id"
          >
            <el-container>
              <el-header height="40px">
                <!-- id, 星级，时间 -->
                <div class="userinfo">
                  <h3>{{ item.username }}</h3>
                  <p>{{ item.time }}</p>
                </div>
                <div class="love">
                  <el-rate :value="item.rates" :colors="lovecolors" disabled>
                  </el-rate>
                </div>
              </el-header>
              <el-main>
                <!-- 评论 -->
                <p>{{ item.comment }}</p>
                <el-divider></el-divider>
              </el-main>
            </el-container></div
        ></el-tab-pane>
      </el-tabs>
    </div>
    <!-- 分页 -->
    <div class="page">
      <el-pagination
        background
        layout="prev, pager, next"
        :total="CommentsList.length"
        page-size="2"
      >
      </el-pagination>
    </div>
  </div>
</template>

<script>
export default {
  name: "productComment",
  props: ["idFamily"],
  data() {
    return {
      tabName: "all",
      lovevalue: null,
      lovecolors: ["#99A9BF", "#F7BA2A", "#FF9900"],
      TempCommentsList: [],
      CommentsList: [
        {
          id: 0,
          productId: 1,
          productIdFamily: 1,
          username: " YZ",
          comment:
            "性价比：性价比很高，别无二致 运行速度：第一代骁龙8，无与伦比 拍照效果：日常生活够用了 电池续航：5000毫安大电池，没毛病 外观设计：时尚大气 MIUI系统：更新了米优13，很好用了",
          rates: 4,
          time: "2022/12/11",
        },
        {
          id: 1,
          productId: 2,
          productIdFamily: 1,
          username: " YZ",
          comment:
            "无论材质还是做工都超出了预期，超值，穿起来也很好看哦。卖家发货也比预计提早了很多，还会再来。买吧，不会后悔的哦！全五分。大品牌，用起来很放心，质量非常好，赶上活动购买非常的划算,客服真的很不错哦～一如既往的好品质，非常喜欢的一款产品，非常棒！o(*￣︶￣*)o",
          rates: 5,
          time: "2022/12/11",
        },
        {
          id: 2,
          productId: 3,
          productIdFamily: 1,
          username: " YZ",
          comment:
            "整体装配工艺差，换了两次次手机摄像头还是歪的，实在懒得换了，其它的还没有体验，疫情期间送货慢可以理解",
          rates: 3,
          time: "2022/12/11",
        },
        {
          id: 3,
          productId: 4,
          productIdFamily: 1,
          username: " YZ",
          comment:
            "不到三十天，手机就出现问题，卡二的手机卡会识别不了，直接掉了显示没插卡，寄修了还说没有问题给我退回来了，过两天估计又要折腾。真的垃圾品控",
          rates: 1,
          time: "2022/12/11",
        },
      ],
    };
  },
  methods: {
    handleClick() {
      var target = [];
      if (this.tabName == "all") return;
      if (this.tabName == "better") {
        target = [4, 5];
      } else if (this.tabName == "common") {
        target = [2, 3];
      } else if (this.tabName == "bad") {
        target = [1];
      }
      this.TempCommentsList = [];
      this.CommentsList.forEach((element) => {
        if (target.includes(element.rates)) {
          this.TempCommentsList.push(element);
        }
      });
      console.log(this.TempCommentsList);
    },
  },
};
</script>

<style>
#root {
  /* height: 1000px; */
  margin-top: 10px;
  /* background: burlywood; */
}

.header1 {
  text-align: left;
  color: #666;
  background-color: #f7f7f7;
  border: 1px solid #eee;
  margin-bottom: 10px;
}

.userinfo {
  float: left;
}

.userinfo p {
  color: #999;
}

.love {
  float: right;
}

.el-main p {
  font-size: 14px;
}

.commentbox {
  margin: 10px 0;
}
</style>