<template>
  <div class="root">
    <top></top>

    <div class="mainheart">
      <el-container>
        <el-header>
          <!-- 标题 -->
          <div class="shppingcart_title">
            <h2>全部商品 {{ ShoppingLists.length }}</h2>
          </div>
          <!-- 表头 -->
        </el-header>

        <el-main>
          <!-- 显示商品清单 -->
          <el-table
            ref="multipleTable"
            :data="ShoppingLists"
            tooltip-effect="dark"
            style="width: 100%"
            @selection-change="handleSelectionChange"
          >
            <!-- 多选框 -->
            <el-table-column type="selection" width="55"> </el-table-column>
            <!-- 商品图片列 -->
            <el-table-column label="商品" width="120px">
              <template slot-scope="scope">
                <div class="content">
                  <img :src="scope.row.productImg" alt="" width="100px" />
                </div>
              </template>
            </el-table-column>

            <!-- 商品名 -->
            <el-table-column label="" width="280px">
              <template slot-scope="scope">
                <div class="content">
                  <!-- <p style="float: left"> -->
                  <router-link to="/web/v2/goods">{{
                    scope.row.productName
                  }}</router-link>
                  <!-- </p> -->
                </div>
              </template>
            </el-table-column>

            <!-- 选配情况 -->
            <el-table-column label="选配情况" width="200">
              <template slot-scope="scope">
                <p>{{ scope.row.color }}</p>
                <p>{{ scope.row.version }}</p>
              </template>
            </el-table-column>

            <!-- 单价 -->
            <el-table-column label="单价" width="120">
              <template slot-scope="scope">
                <p>￥{{ scope.row.productPrice }}</p>
              </template>
            </el-table-column>

            <!-- 商品数量 -->
            <el-table-column label="数量" show-overflow-tooltip>
              <template slot-scope="scope">
                <!-- 计数器 -->
                <el-input-number
                  size="mini"
                  v-model="scope.row.productNum"
                  :min="1"
                  :max="99"
                ></el-input-number>
              </template>
            </el-table-column>

            <el-table-column label="小计" show-overflow-tooltip>
              <!-- 计算属性 -->
              <template slot-scope="scope">
                <h3>￥{{ scope.row.productPrice * scope.row.productNum }}</h3>
              </template>
            </el-table-column>

            <el-table-column label="操作" show-overflow-tooltip>
              <!-- 删除 -->
              <template slot-scope="scope">
                <el-button
                  icon="el-icon-delete"
                  circle
                  size="small"
                  @click="
                    oneDelete(scope.$index, scope.row);
                    centerDialogVisible = true;
                  "
                ></el-button>
              </template>
            </el-table-column>
          </el-table>
        </el-main>

        <!-- 删除对话框 -->
        <el-dialog :visible.sync="centerDialogVisible" width="440px" center>
          <img
            slot="title"
            src="../../components/img/deleteicon.png"
            alt=""
            width="80px"
          />
          <h3 style="text-align: center">确定要删除该商品吗？</h3>
          <span slot="footer" class="dialog-footer">
            <el-button
              type="primary"
              @click="
                centerDialogVisible = false;
                needDeleteIds = [];
              "
              >我在想想</el-button
            >
            <el-button
              @click="
                centerDialogVisible = false;
                confirmDeletion();
              "
              >确 定</el-button
            >
          </span>
        </el-dialog>

        <el-footer height="50px">
          <!-- 结算 -->
          <div class="bottom">
            <div class="shoppingCart_bottom_left">
              <div class="shoppingCartSelectAll">
                <el-checkbox
                  v-model="selectAllChecked"
                  @change="selectAll(ShoppingLists)"
                  >全选</el-checkbox
                >
              </div>

              <div class="shopping_deleteSeleted">
                <p
                  @click="
                    centerDialogVisible = true;
                    moreDelete();
                  "
                >
                  删除选中的商品
                </p>
              </div>
            </div>

            <div class="shoppingCart_bottom_right">
              <div class="bottom_right_info">
                <p>已选<span>{{totalNum}}</span>件商品</p>
                <p>总价： <span>￥{{totalPrice}}</span> </p>
              </div>
             <el-button type="danger" @click="dialogFormVisible = true">立刻购买</el-button>
              
            </div>
            
          </div>
          <el-dialog title="支付" :visible.sync="dialogFormVisible">
            <el-form :model="form">
              <el-from-item label="价格" :label-width="formLabelWidth">
                <h2 style="color:red">￥{{totalPrice}}</h2>
              </el-from-item>
              <el-form-item label="扫码付款" :label-width="formLabelWidth">
                <img :src="imgs[img]" style="width: 180px;height: 180px;margin-right: 100px;" />
              </el-form-item>
              <el-form-item :label-width="formLabelWidth">
                <el-button type="primary" @click="onimg(0)">微信支付</el-button>
                <el-button type="primary" @click="onimg(1)" style="margin-right: 100px;">支付宝支付</el-button>
              </el-form-item>
            </el-form>
            <div slot="footer" class="dialog-footer">
              <el-button @click="dialogFormVisible = false">取 消</el-button>
              <el-button type="primary" @click="dialogFormVisible = false">确 定</el-button>
            </div>
          </el-dialog>
        </el-footer>
      </el-container>
    </div>
    <foot />
  </div>
</template>

<script>
import top from "@/components/TopBar.vue";
import foot from "@/components/SiteInfo.vue";


export default {
  name: "shoppingCart",
  data() {
    return {
      imgs: [
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
      needDeleteIds: [], //要删除的产品id
      centerDialogVisible: false, //设置对话框是否可见
      selectAllChecked: false,
      ShoppingLists: [
        {
          productId: 1,
          productName:
            "Redmi K50 至尊版",
          productNum: 1,
          // 通过id查的
          productPrice: 2699,
          productImg: require("../../assets/goods/01/info/0.jpg"),
          productFamily: 1,
          color: "雅黑",
          version: "8GB+128GB",
        },
        {
          productId: 2,
          productName: "Redmi K50",
          productNum: 2,
          // 通过id查的
          productPrice: 2799,
          productImg: require("../../assets/goods/01/info/1.jpg"),
          productFamily: 1,
          color: "银迹",
          version: "8GB+256GB",
        },
        {
          productId: 3,
          productName: "晓龙 gen 8",
          productNum: 3,
          // 通过id查的
          productPrice: 3099,
          productImg: require("../../assets/goods/01/info/2.jpg"),
          productFamily: 1,
          color: "冰蓝",
          version: "12GB+256GB",
        },
        {
          productId: 4,
          productName: "王者note8",
          productNum: 4,
          // 通过id查的
          productPrice: 3399,
          productImg: require("../../assets/goods/01/info/3.jpg"),
          productFamily: 1,
          color: "冰蓝",
          version: "12GB+512GB",
        },
      ],
      // 被选中的框的集合
      multipleSelection: [],
    };
  },
  computed: {
    //总金额
    totalPrice() {
      var price = 0;
      this.multipleSelection.forEach(element => {
        price += element.productNum * element.productPrice;
      });
      return price;
    },
    //总件数
    totalNum(){
      var nums = 0;
      this.multipleSelection.forEach(element => {
        nums += element.productNum;
      });
      return nums;
    }
  },
  components: {
    top,
    foot,
  },
  methods: {
    onimg(index) {
      if (index == 0) {
        this.img = "0";
      } else {
        this.img = "1";
      }
    },
    confirmDeletion() {
      //确定删除订单
      this.ShoppingLists = this.ShoppingLists.filter((item) => {
          return !this.needDeleteIds.includes(item.productId);
      });

      //删完清空数组
      this.needDeleteIds = [];
    },
    //下面的全选框用
    selectAll(rows) {
      if (this.selectAllChecked == true) {
        rows.forEach((row) => {
          this.$refs.multipleTable.toggleRowSelection(row);
        });
      } else {
        this.$refs.multipleTable.clearSelection();
      }
    },
    handleSelectionChange(val) {
      // console.log(val);
      this.multipleSelection = val;
    },
    //单个删除
    oneDelete(index, row) {
      // 返回索引
      this.needDeleteIds.push(row.productId);
    },
    //批量删除
    moreDelete() {
      this.multipleSelection.forEach(element => {
        this.needDeleteIds.push(element.productId);
      });
    },
  },
};
</script>

<style>
/* 版心 */
.mainheart {
  width: 1240px;
  margin: 0 auto;
  margin-top: 50px;
  margin-bottom: 100px;
}

.shppingcart_title {
  float: left;
}


.content a {
  color: black;
}

.content a:hover {
  color: #f56c6c;
}

/* 计数器缩小 */
.el-input-number {
  width: 100px;
}

/* 删除键里的图标 */
.deleteicon {
  width: 50px;
}

/* footer */
.shoppingCart_bottom_left {
  float: left;
  width: 150px;
  height: 50px;
  line-height: 50px;
}

.shoppingCartSelectAll {
  float: left;
}

.shoppingCart_bottom_left p {
  float: right;
}

.shoppingCart_bottom_left p:hover {
  color: #f56c6c;
  cursor: pointer;
}

.shoppingCart_bottom_right {
  float: right;
  width: 350px;
  height: 50px;
  line-height: 50px;
}

.shoppingCart_bottom_right .bottom_right_info {
  float: left;
}

.shoppingCart_bottom_right .el-button {
  float: right;
}

.bottom_right_info p{
  float: left;
  margin-right: 20px;
}

.bottom_right_info p span{
  color: #f56c6c;
  font-size: 15px;
  margin: 0 3px;
}
</style>