<template>
    <div>
        <top></top>
        <h1>基本信息</h1>
        <hr/>
        <el-row :gutter="15" style="margin-top:50px;">
            <el-form ref="elForm" :model="formData" :rules="rules" size="medium" label-width="100px">
                <el-col :span="13">
                    <el-form-item label="用户名" prop="field101">
                        <el-input v-model="formData.field101" placeholder="请输入用户名" clearable prefix-icon='el-icon-user'
                            :style="{ width: '100%' }"></el-input>
                    </el-form-item>
                </el-col>
                <el-col :span="13">
                    <el-form-item label="邮箱" prop="field102">
                        <el-input v-model="formData.field102" placeholder="请输入邮箱" clearable :style="{ width: '100%' }">
                        </el-input>
                    </el-form-item>
                </el-col>
                <el-col :span="13">
                    <el-form-item label="收货地址" prop="field103">
                        <el-cascader v-model="formData.field103" :options="field103Options" :props="field103Props"
                            :style="{ width: '100%' }" placeholder="请选择收货地址" clearable></el-cascader>
                    </el-form-item>
                </el-col>
                <el-col :span="24">
                    <el-form-item label="详细地址" prop="field106">
                        <el-input v-model="formData.field106" placeholder="详细到小区" clearable :style="{ width: '100%' }">
                        </el-input>
                    </el-form-item>
                </el-col>
                <el-col :span="24">
                    <el-form-item label="个人简介" prop="field107">
                        <el-input v-model="formData.field107" type="textarea" placeholder="请输入个人简介"
                            :autosize="{ minRows: 4, maxRows: 4 }" :style="{ width: '100%' }"></el-input>
                    </el-form-item>
                </el-col>
                
                <el-col :span="24">
                    <el-form-item size="large">
                        <el-button type="primary" @click="submitForm">提交</el-button>
                        <el-button @click="resetForm">重置</el-button>
                    </el-form-item>
                </el-col>
                <el-col :span="24">
                    <el-form-item label="" prop="field104">
                        <el-button type="primary" icon="el-icon-thumb" size="medium" @click="tuichu()"> 退出登录 </el-button>
                    </el-form-item>
                </el-col>
            </el-form>
        </el-row>
    </div>
</template>
<script>
import top from "@/components/TopBar.vue"
import user from "./user.json"
export default {
    components: {
        top
    },
    props: [],
    data() {
        return {
            formData: {
                field101: "itcyy",
                field102: "itcyy@student.gdufe.edu.cn",
                field103: ["北京市", "海淀区", "市区"],
                field106: undefined,
                field107: undefined,
                field104: undefined,
            },
            rules: {
                field101: [{
                    required: true,
                    message: '请输入用户名',
                    trigger: 'blur'
                }],
                field102: [{
                    required: true,
                    message: '请输入邮箱',
                    trigger: 'blur'
                }],
                field103: [{
                    required: false,
                    type: 'array',
                    message: '请至少选择一个收货地址',
                    trigger: 'change'
                }],
                field106: [{
                    required: false ,
                    message: '详细到小区',
                    trigger: 'blur'
                }],
                field107: [{
                    required: false,
                    message: '请输入个人简介',
                    trigger: 'blur'
                }],
            },
            field103Options: user,
            field103Props: {
                "multiple": false,
                "label": "value",
                "value": "value",
                "children": "children"
            },
        }
    },
    computed: {},
    watch: {},
    created() {
        this.getField103Options()
    },
    mounted() { },
    methods: {
        submitForm() {
            this.$refs['elForm'].validate(valid => {
                if (!valid) return
                // TODO 提交表单
            })
        },
        resetForm() {
            this.$refs['elForm'].resetFields()
        },
        tuichu(){
            console.log(localStorage.getItem("id"), 55555555888)
            localStorage.setItem("id", "0");
            console.log(localStorage.getItem("id"),55555555888)
            this.$router.push({path:"/"})
        },
        getField103Options() {
            // 注意：this.$axios是通过Vue.prototype.$axios = axios挂载产生的
            this.$axios({
                method: 'get',
                url: ''
            }).then(resp => {
                var {
                    data
                } = resp
                this.field103Options = data.list
            })
        },
    }
}

</script>
<style>

</style>
