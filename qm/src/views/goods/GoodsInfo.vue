<!-- 商品分页 -->
<template>
  <div>
    <top></top>
    <div class="mainheart">
      <!-- 购买部分 -->
      <div class="buy">
        <el-container>
          <el-aside width="450px">
            <!-- 商品轮播图 -->
            <el-carousel
              indicator-position="outside"
              :interval="5000"
              arrow="always"
              height="450px"
            >
              <el-carousel-item
                v-for="(item, index) in ProductList"
                :key="index"
              >
                <img :src="item.imgURL" alt="" />
              </el-carousel-item>
            </el-carousel>
          </el-aside>

          <el-container class="infoContainer">
            <el-header height="150px">
              <!-- 名称 -->
              <h1>{{ ProductList[0].name }}</h1>
              <!-- 描述 -->
              <h3>{{ ProductList[0].desc }}</h3>
              <!-- 价格 -->
              <h2>{{ ProductList[0].price }}</h2>
            </el-header>

            <el-divider></el-divider>

            <el-main>
              <!-- 选版本 -->
              <h2>选择版本</h2>
              <div class="versionchoose">
                <el-radio-group v-model="versionradio">
                  <el-radio-button
                    :label="item"
                    v-for="(item, index) in versions"
                    :key="index"
                  ></el-radio-button>
                </el-radio-group>
              </div>

              <!-- 选颜色 -->
              <h2>选择颜色</h2>
              <div class="colorchoose">
                <el-radio-group v-model="colorradio">
                  <el-radio-button
                    :label="item"
                    v-for="(item, index) in colors"
                    :key="index"
                  ></el-radio-button>
                </el-radio-group>
              </div>
            </el-main>

            <el-divider></el-divider>

            <el-footer>
              <!-- 购买数量->计数器 -->
              <el-input-number
                v-model="countnum"
                :min="1"
                :max="99"
                label="描述文字"
              ></el-input-number>
              <!-- 按钮->加购物车 -->
              <!-- <el-button type="danger" round>加入购物车</el-button> -->
              <el-button type="danger" plain>加入购物车</el-button>
              <!-- 直接购买 -->
              <el-button type="danger" @click="dialogFormVisible = true">立刻购买</el-button>
              <el-dialog title="支付" :visible.sync="dialogFormVisible">
                <el-form :model="form">
                  <el-from-item label="价格" :label-width="formLabelWidth">
                    <h2 style="margin-left: 80px;color:red">￥{{ ProductList[0].price }}</h2>
                  </el-from-item>
                  <el-form-item label="扫码付款" :label-width="formLabelWidth">
                   <img :src="imgs[img]" style="width: 180px;height: 180px;margin-left: 80px;"/>
                  </el-form-item>
                  <el-form-item  :label-width="formLabelWidth">
                    <el-button type="primary" @click="onimg(0)">微信支付</el-button>
                    <el-button type="primary" @click="onimg(1)">支付宝支付</el-button>
                  </el-form-item>
                </el-form>
                <div slot="footer" class="dialog-footer">
                  <el-button @click="dialogFormVisible = false">取 消</el-button>
                  <el-button type="primary" @click="dialogFormVisible = false">确 定</el-button>
                </div>
              </el-dialog>
            </el-footer>
          </el-container>
        </el-container>
      </div>

      <!-- 商品介绍和评论 -->
      <div class="other">
        <el-tabs v-model="activeName" type="border-card">
          <!-- 商品介绍 -->
          <el-tab-pane label="商品介绍" name="first">
            <introduce></introduce>
          </el-tab-pane>

          <!-- 评论 -->
          <el-tab-pane label="商品评价" name="second">
            <comments :idFamily="ProductList[0].idFamily"></comments>
          </el-tab-pane>
        </el-tabs>
      </div>
    </div>
    <!-- footer -->
    <foot/>
  </div>
</template>

<script>
import top from "@/components/TopBar.vue";
import introduce from "./ProductIntroduction.vue";
import comments from "./ProductComment.vue";
import foot from "@/components/SiteInfo.vue";

export default {
  name: "goodsInfo",
  components: {
    top,
    introduce,
    comments,
    foot,
  },
  data() {
    return {
      imgs:[
        require("../../assets/image/weixin.jpg"),
        require("../../assets/image/zhifubao.jpg")
      ],
      img: "0",
      dialogTableVisible: false,
      dialogFormVisible: false,
      form: {
        name: '',
        region: '',
        date1: '',
        date2: '',
        delivery: false,
        type: [],
        resource: '',
        desc: ''
      },
      formLabelWidth: '120px',
      activeName: "first", //默认显示第一个tab
      countnum: 1, //计数器数值
      versionradio: "", //版本选择值
      colorradio: "", //颜色选择值
      ProductList: [
        {
          id: 1,
          name: "Redmi K50 至尊版",
          idFamily: 1,
          price: 2699,
          imgURL: require("../../assets/goods/01/info/0.jpg"),
          desc: "「享6期分期免息；12+512GB券后到手价3369元；12+256GB券后到手价3069元；整点购机限量赢定制品」骁龙8+「狂暴调校」｜ 定制 1.5K 旗舰直屏 ｜ 120W神仙秒充丨1 亿像素光学防抖相机｜ 电竞级 VC 散热 | 屏下指纹",
          color: "雅黑",
          version: "8GB+128GB",
        },
        {
          id: 2,
          name: "Redmi K50 至尊版",
          idFamily: 1,
          price: 2799,
          imgURL: require("../../assets/goods/01/info/1.jpg"),
          desc: "「享6期分期免息；12+512GB券后到手价3369元；12+256GB券后到手价3069元；整点购机限量赢定制品」骁龙8+「狂暴调校」｜ 定制 1.5K 旗舰直屏 ｜ 120W神仙秒充丨1 亿像素光学防抖相机｜ 电竞级 VC 散热 | 屏下指纹",
          color: "银迹",
          version: "8GB+256GB",
        },
        {
          id: 3,
          name: "Redmi K50 至尊版",
          idFamily: 1,
          price: 3099,
          imgURL: require("../../assets/goods/01/info/2.jpg"),
          desc: "「享6期分期免息；12+512GB券后到手价3369元；12+256GB券后到手价3069元；整点购机限量赢定制品」骁龙8+「狂暴调校」｜ 定制 1.5K 旗舰直屏 ｜ 120W神仙秒充丨1 亿像素光学防抖相机｜ 电竞级 VC 散热 | 屏下指纹",
          color: "冰蓝",
          version: "12GB+256GB",
        },
        {
          id: 4,
          name: "Redmi K50 至尊版",
          idFamily: 1,
          price: 3399,
          imgURL: require("../../assets/goods/01/info/3.jpg"),
          desc: "「享6期分期免息；12+512GB券后到手价3369元；12+256GB券后到手价3069元；整点购机限量赢定制品」骁龙8+「狂暴调校」｜ 定制 1.5K 旗舰直屏 ｜ 120W神仙秒充丨1 亿像素光学防抖相机｜ 电竞级 VC 散热 | 屏下指纹",
          color: "冰蓝",
          version: "12GB+512GB",
        },
      ],
    };
  },
  computed: {
    fare() {
      var pri = "暂无定价";
      this.ProductList.forEach((element) => {
        if (
          element.version == this.versionradio &&
          element.color == this.colorradio
        ) {
          console.log(element.price);
          pri = element.price + "元";
        }
      });
      return pri;
    },
    versions() {
      let versions = [];
      this.ProductList.forEach((element) => {
        versions.push(element.version);
      });
      return new Set(versions);
    },
    colors() {
      let colors = [];
      this.ProductList.forEach((element) => {
        colors.push(element.color);
      });
      return new Set(colors);
    },
  },
  methods: {
    onimg(index){
        if (index == 0){
          this.img = "0";
        }else{
          this.img = "1";
        }
    },
    toTop() {
      document.documentElement.scrollTop = 0;
    },
  },
};
</script>

<style>
/* 版心 */
.mainheart {
  width: 1240px;
  margin: 0 auto;
  margin-bottom: 100px;
}

.buy {
  margin-top: 20px;
  margin-bottom: 100px;
}

/* container */
.el-header,
.el-footer {
  color: #333;
  text-align: left;
}

/* header */
.infoContainer .el-header h1,
h2,
h3 {
  padding: 5px;
}

.infoContainer .el-header h2 {
  color: #f56c6c;
}

.infoContainer .el-header h3 {
  color: #909399;
}

.el-aside {
  text-align: center;
}

.el-main {
  text-align: left;
}

/* 版本，颜色 */
.versionchoose {
  margin-bottom: 30px;
}

.el-main h2 {
  color: #606266;
}

/* 两组单选框 */
.el-radio-group :nth-child(n + 2) {
  padding-left: 20px;
  border-left: 1px solid #dcdfe6;
}

.versionchoose,
.colorchoose {
  margin-top: 5px;
}

/* footer */
.el-input-number {
  width: 140px;
}

.el-footer .el-button {
  margin-left: 40px;
}

body > .el-container {
  margin-bottom: 40px;
}
</style>