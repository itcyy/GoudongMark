<template>
  <div>
    <Particles></Particles>
    <el-card class="box-card" v-if="Bool">
      <div slot="header" class="clearfix">
        <span class="sys-name">狗东商城</span>
      </div>
      <el-container>
        <el-aside width="200px"
          ><img
            src="../components/img/1.png"
            style="width: 200px; height: 300px"
        /></el-aside>
        <el-main>
          <div class="card-body">
            <el-form ref="form" :model="form" :rules="ruleForm">
              <el-form-item prop="name">
                <el-input
                  type="text"
                  v-model="form.loginName"
                  auto-complete="off"
                  placeholder="请输入用户名"
                >
                  <template slot="prepend"
                    ><i style="font-size: 20px" class="el-icon-user-solid"></i
                  ></template>
                </el-input>
              </el-form-item>
              <el-form-item prop="password">
                <el-input
                  type="password"
                  v-model="form.loginPass"
                  auto-complete="off"
                  placeholder="请输入用户密码"
                  show-password
                >
                  <template slot="prepend"
                    ><i style="font-size: 20px" class="el-icon-s-goods"></i
                  ></template>
                </el-input>
              </el-form-item>
              <el-form-item>
                <el-radio-group v-model="radio">
                  <el-checkbox v-model="checked" style="font-size: 2px;color: aliceblue;"
                    >同意并登陆狗东商城<a 
                      href="https://static.account.xiaomi.com/html/agreement/user/zh_CN.html"
                      >《用户协议》</a
                    >和<a href="https://privacy.mi.com/miaccount/zh_CN/"
                      >《隐私协议》</a
                    ></el-checkbox
                  >
                </el-radio-group>
              </el-form-item>

              <el-form-item>
                <el-button
                  style="width: 100%"
                  type="primary"
                  @click="onHome"
                  :loading="logining"
                  >登录
                </el-button>
                <div @click="onClink()">
                  <a class="text_in_login" href="#">注册账号>></a>
                </div>
              </el-form-item>
            </el-form>
          </div>
        </el-main>
      </el-container>
    </el-card>
    <el-card class="box-card" style="width: 480px" v-if="Bool1">
      <div slot="header" class="clearfix">
        <span class="sys-name">狗东商城</span>
      </div>
      <div class="card-body">
        <el-form ref="form" :model="login" :rules="ruleForm">
          <el-form-item prop="name">
            <el-input type="text" v-model="login.loginName" auto-complete="off" placeholder="请输入用户名">
              <template slot="prepend"><i style="font-size: 20px" class="el-icon-user-solid"></i></template>
            </el-input>
          </el-form-item>
          <el-form-item prop="name">
            <el-input
              type="text"
              v-model="login.loginMail"
              auto-complete="off"
              placeholder="请输入邮箱"
            >
              <template slot="prepend"
                ><i style="font-size: 20px" class="el-icon-user-solid"></i
              ></template>
            </el-input>
          </el-form-item>

          <el-form-item prop="password">
            <el-input
              type="password"
              v-model="login.loginPass"
              auto-complete="off"
              placeholder="请输入密码"
            >
              <template slot="prepend"
                ><i style="font-size: 20px" class="el-icon-s-goods"></i
              ></template>
            </el-input>
          </el-form-item>
          <el-form-item prop="password">
            <el-input
              type="password"
              v-model="login.loginPass1"
              auto-complete="off"
              placeholder="请再次输入密码"
            >
              <template slot="prepend"
                ><i style="font-size: 20px" class="el-icon-s-goods"></i
              ></template>
            </el-input>
          </el-form-item>
          <el-form-item>
            <el-radio-group v-model="radio">
              <el-checkbox v-model="checked" style="font-size: 2px">同意并登陆狗东商城<a
                  href="https://static.account.xiaomi.com/html/agreement/user/zh_CN.html">《用户协议》</a>和<a
                  href="https://privacy.mi.com/miaccount/zh_CN/">《隐私协议》</a></el-checkbox>
            </el-radio-group>
          </el-form-item>

          <el-form-item>
            <el-button
              style="width: 100%; margin-top: 20px"
              type="primary"
              @click="onRES"
              :loading="logining"
              >注册
            </el-button>
          </el-form-item>
          <el-form-item
            ><div class="el-icon-arrow-left" @click="onLogin()">
              <a style="text-decoration: none; color: whitesmoke" href="#"><i class="el-icon-arrow-left"></i>返回登录 </a>
            </div></el-form-item
          >
        </el-form>
      </div>
    </el-card>
    <el-card class="box-card" style="width: 480px" v-if="Bool2">
      <div slot="header" class="clearfix">
        <span class="sys-name">狗东商城</span>
      </div>
      <div class="card-body">
        <slide-verify
          :l="42"
          :r="10"
          :w="380"
          :h="180"
          
          slider-text="向右滑动"
          @success="onSuccess"
          @fail="onFail"
          :imgs="img" 
          @refresh="onRefresh"
          show="true"
        ></slide-verify>
        <div>{{ msg }}</div>
      </div>
    </el-card>

  </div>
</template>
<script>
import Particles from '@/components/particles/index';
import axios from "axios";
import { getUser, setUser, SetUser, insertUser } from "@/api";
import { encryptDes, decryptDes } from "@/api"


// @ is an alias to /src
export default {
  name: "Login",
  components:{
    Particles
  },
  data() {
    return {
      params: {
        username: "root",
        password: "123456",
      },
      img: [require("../assets/image/1.jpg"),
        require("../assets/image/2.jpg"),
        require("../assets/image/3.jpg"),
        require("../assets/image/4.jpg"),],
        
      tk:"",
      checked: false,
      radio: "",
      msg: "",
      loginfrom: [

  ],
      Bool: true,
      timer: "",
      Bool1: false,
      Bool2: false,
      Bool3: false,
      logining: false,
      form: {
        loginName: "",
        loginPass: "",
        loginPass1: "",
      },
      login: {
        loginMail:"",
        loginName: "",
        loginPass: "",
        loginPass1: "",
      },
     
      set: [],
      ruleForm: {
        loginName: [
          {
            required: true,
            message: "请输入账号",
            trigger: "blur",
          },
        ],
        loginPass: [
          {
            required: true,
            message: "请输入密码",
            trigger: "blur",
          },
        ],
      },
    };
  },
  created() {
    this.setusers()
  },
  beforeMount() {
    console.log(this.$route.query.id);
    console.log("000");
    let tf = this.$route.query.seesion;
    console.log(tf);
    if (tf == "1") {
      console.log("000");
      this.Bool = false;
      this.Bool1 = true;
    } else {
      console.log("111");
      this.Bool = true, 
      this.Bool1 = false;
    }
  },
  methods: {
    omfocus(event){
      console.log(8456456456456454, event)
    },
    Emailpan(value){
      console.log(8456456456456454, value)

    },
    async setusers(uaername,password) {
      //异步  async与await
      const res = await setUser(uaername, password);
      
      console.log(res.data.message,10086);

      console.log("55455");
    },
    async onHome(){
      if (!this.checked){
        this.$alert("请勾选隐私协议按钮", "温馨提示", {
          confirmButtonText: "确定",
          callback: (action) => {
            action:"请勾选隐私协议按钮"
          },
        });
        return 0
      }else{

      
      //异步  async与await
      var flag = await setUser(this.form.loginName, this.form.loginPass)
      
      if(flag.data.message == "OK"){
        this.Bool2 = true;
        this.Bool = false;
      } else {
        this.$alert("密码或账号错误，请重新输入！", "温馨提示", {
          confirmButtonText: "确定",
          callback: (action) => {
            this.form.loginName = "";
            this.form.loginPass = "";
          },
        });

        console.log("error submit!");

        return false;
      }
      
      }
    },
    async getShops() {
      //异步  async与await
      const row = await getUser();
      console.log(row.data);

      console.log("56454545");
    },

    onSuccess() {
      this.msg = "验证成功";
      this.Bool2 = false;
      const h = this.$createElement;
      this.$msgbox({
        title: "消息",
        message: h("p", null, [
          h("span", null, "验证成功 "),
          h("i", { style: "color: teal" }, "是否登陆？"),
        ]),
        showCancelButton: true,
        confirmButtonText: "确定",
        cancelButtonText: "取消",
        beforeClose: (action, instance, done) => {
          if (action === "confirm") {
            instance.confirmButtonLoading = true;
            instance.confirmButtonText = "登录中...";
            setTimeout(() => {
              done();
              setTimeout(() => {
                instance.confirmButtonLoading = true;
              }, 3000);
            }, 3000)
            localStorage.setItem("id","01");
            this.$router.push({
              name: "home",
              query: { id: "01" },
            });
          } else {
            done();
          }
        },
      }).then((action) => {
        this.$message({
          type: "info",
          message: "提示:登录成功 ",
        });
      });
    },
    onFail() {
      this.msg = "验证失败，请重试";
    },
    onRefresh() {
      this.msg = "";
    },
/*
    submit() {
      var arr = Object.keys(this.login.user);
      var len = arr.length;
      console.log(this.login.user[0].password);
      console.log(len, 200);

      this.$refs.form.validate(async (valid) => {
        console.log(this.login.user[0].password);
        console.log(this.login.user.length, 100);
        for (var i = 0; i < len; i++) {
          if (
            this.form.loginName == this.login.user[i].username &&
            this.form.loginPass == this.login.user[i].password
          ) {
            if (this.login.user[i].username == "root") {
              this.$router.push({
                name: "home",
                query: { id: "01" },
              });
            }
            this.flag = 1;
            console.log(this.login.user[0].password, 6);
            break;
          }
        }
        console.log(this.login.user[0].password, 1);
        console.log(this.login.user[0].username, 1);
        if (this.flag == 1) {
          console.log(this.login.user[0].password, 2);
          console.log(this.login.user[0].username, 2);
          this.Bool2 = true;
          this.Bool = false;
        } else {
          this.$alert("密码或账号错误，请重新输入！", "温馨提示", {
            confirmButtonText: "确定",
            callback: (action) => {
              this.form.loginName = "";
              this.form.loginPass = "";
            },
          });

          console.log("error submit!");

          return false;
        }

        //this.logining = true;
        /*if (this.form.name === 'admin' &&
                        this.form.password === '123456') {
                        this.logining = false;
                        sessionStorage.setItem('user', this.form.name);
                        this.$router.push({
                            name: 'home'
                        });
                    } else {
                        this.logining = false;
                        this.$alert('name or password wrong!', 'info', {
                            confirmButtonText: 'ok'
                        })
                    }
                   
      });
    }, */
    onClink: function () {
      console.log("000");
      this.Bool = false;
      this.Bool1 = true;
      console.log(this.Bool, this.Bool1);
    },
    
    onLogin() {
        this.Bool = true;
        this.Bool1 = false;
    },
    async onRES() {
      if (!this.checked) {
        this.$alert("请勾选隐私协议按钮", "温馨提示", {
          confirmButtonText: "确定",
          callback: (action) => {
            action: "请勾选隐私协议按钮"
          },
        });
        return 0
      }else{
      if (this.login.loginPass.length < 6 | this.login.loginPass.length > 16) {
        console.log(this.login.loginPass.length, this.login.loginPass.length, 100)
        this.$alert("密码长度为6-16", "温馨提示", {
          confirmButtonText: "确定",

        });
      }
      var res = await SetUser(this.login.loginName)
      console.log(res.data,55)
      if(res.data.message == "OK"){
        this.$alert("用户名已存在！", "温馨提示", {
          confirmButtonText: "确定",
          callback: (action) => {
          },
        });
        return false;
      }
      if (this.login.loginMail == "") {
        this.$alert("Email地址不能为空", "温馨提示", {
          confirmButtonText: "确定",
          callback: (action) => {
      
          },
        });
        return false;
      }
  if (this.login.loginMail.indexOf("@") == -1) {
        this.$alert("Email地址格式不正确，必须包含@", "温馨提示", {
          confirmButtonText: "确定",
       
        });
        return false;
      }
  if (this.login.loginMail.indexOf(".") == -1) {
        this.$alert("Email地址格式不正确，必须包含.", "温馨提示", {
          confirmButtonText: "确定",
         
        });
  
        return false;
      }
 
    
      if (this.login.loginPass != this.login.loginPass1) {
        this.$alert("两次密码输入不一致！", "温馨提示", {
          confirmButtonText: "确定",
      
        });
      }
      if(this.login.loginPass==""){
        this.$alert("请输入密码", "温馨提示", {
          confirmButtonText: "确定",
          
        });
      }
      console.log(this.login.loginPass.length, this.login.loginPass.length, 10011)
      }
      var res = await insertUser(this.login.loginName,this.login.loginPass,this.login.loginMail)
      if(res.data.message=="OK"){
        this.$alert("注册成功成功,前往登录", "温馨提示", {
          confirmButtonText: "确定",
          callback: (action)=>{
                this.Bool = true,
                this.Bool1 = false;
            }
        });
        
      }
    },
    onradio: function () {
      radio: "";
      console.log(10086);
    },
  },
};
</script>
 
<style>
.center_in_col {
  border: 1px solid green;
}
.sys-name {
  font-size: 18px;
  text-align: center;
  color: white;
}

.box-card {
  width: 640px;
  margin: 0 auto;
  margin-top: 160px;
  background-color:#1824446C ;
}

.card-body {
  padding: 0 30px;
}
.text_in_login {
  margin-left: 220px;
  margin-bottom: 20px;
  color: rgb(25, 187, 33);
  text-decoration: underline;
}
</style>