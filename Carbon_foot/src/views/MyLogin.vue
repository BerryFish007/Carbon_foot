<template>
    <div class="dowebok" :class="{ 's--signup': !isLogin }">
      <!-- 登录表单 -->
      <div class="form sign-in">
        <h2>欢迎回来</h2>
        <label>
          <span>用户名</span>
          <input type="text" v-model="form.username" placeholder="请输入用户名"/>
        </label>
        <label>
          <span>密码</span>
          <input type="password" v-model="form.password" placeholder="请输入密码"/>
        </label>
        <button type="button" class="submit" @click="doSubmit" :disabled="isLoading">
          {{ isLoading ? "处理中..." : "登 录" }}
        </button>
      </div>
  
      <div class="sub-cont">
        <div class="img">
          <div class="img__text m--up">
            <h2>没有账号？</h2>
          </div>
          <div class="img__text m--in">
            <h2>已有账号？</h2>
          </div>
          <div class="img__btn" @click="switchMode">
            <span class="m--up">去 注 册</span>
            <span class="m--in">去 登 录</span>
          </div>
        </div>
        
        <!-- 注册表单 -->
        <div class="form sign-up">
          <h2>立即注册</h2>
          <label>
            <span>用户名</span>
            <input type="text" v-model="form.username" placeholder="请输入用户名"/>
          </label>
          <label>
            <span>密码</span>
            <input type="password" v-model="form.password" placeholder="请输入密码"/>
          </label>
          <button type="button" class="submit" @click="doSubmit" :disabled="isLoading">
            {{ isLoading ? "处理中..." : "注 册" }}
          </button>
        </div>
      </div>
    </div>
  </template>
  
  <script>
  import { ref, reactive } from "vue";  // 引入Vue的ref和reactive，用于创建响应式数据
  import { ElMessage } from "element-plus";  // 导入Element Plus的消息提示组件
  import axios from "axios";  // 导入Axios库
  import { useStore } from 'vuex';  // 导入 Vuex store
  import { useRouter } from "vue-router";  // 导入 Vue Router，用于页面跳转
  import api from '@/http'; // 引入全局配置的 axios 实例
  import {jwtDecode} from 'jwt-decode'; // 引入 jwt 解码库
  
  // const api = axios.create({
  //   baseURL: "/api", 
  //   timeout: 5000, 
  // });
  
  export default {
    name: "LoginRegisterPage",
    setup() {
      const router = useRouter();
      const store = useStore(); // Vuex Store，管理用户状态
      const isLogin = ref(true);  // 控制当前是否是登录状态
      const isLoading = ref(false);  // 控制是否正在处理请求
      const form = reactive({
        username: "",
        password: "",
      });
  
      
      // 登录/注册提交方法
      const doSubmit = async () => {
        if (!form.username || !form.password) {
          ElMessage.error("用户名和密码不能为空");
          return;
        }
  
        try {
          isLoading.value = true;
          const url = isLogin.value ? "/login" : "/register";
          const params = {
            username: form.username,
            password: form.password
          };
          
          // 发送请求
          const response = await api.post(url, params);
          console.log(response.data);
          if (isLogin.value) {
            // 登录成功，返回用户ID
            if (response.data.Msg == "成功") {
              const token = response.data.Data;
              // 保存 token 到 localStorage
              localStorage.setItem("token", token);

              // 设置 axios 请求头
              api.defaults.headers.common["Authorization"] = `Bearer ${token}`;

              // 解码 JWT 获取用户信息
              const decodedToken = jwtDecode(token);
              const user = {
                userId: decodedToken.Uid,
                username: decodedToken.Username
              };

              // // 将用户信息存入 localStorage
              localStorage.setItem("user", JSON.stringify(user));
              console.log(user);
              setTimeout(() => {
                ElMessage.success("登录成功");  // 弹出成功提示
                // 重定向到首页
                router.push("/show");
                isLoading.value = true;
              }, 500);
              
            } else {
              ElMessage.error("用户名或密码错误");
              form.password = "";
            }
          } else {
            // 注册成功，返回"OK"
            if (response.data.Msg === "成功") {
              ElMessage.success("注册成功，请登录");
              isLogin.value = true;  // 切换回登录模式
              form.password = "";
            } else {
              // console.log(response.data);
              ElMessage.error("注册失败，用户名已被占用");
              form.password = "";
            }
          }
        } catch (err) {
          console.error("请求错误:", err);
          ElMessage.error("网络或服务器错误");
        } finally {
          setTimeout(() => {
            isLoading.value = false;
              }, 500);
        }
      };

      
  
      // 切换登录/注册状态
      const switchMode = () => {
        isLogin.value = !isLogin.value;
        form.username = "";
        form.password = "";
      };
  
      return {
        form,
        isLogin,
        isLoading,
        doSubmit,
        switchMode,
      };
    },
  };
  </script>
  
  
  
  <style scoped>
  
  *, *:before, *:after {
      box-sizing: border-box;
      margin: 0;
      padding: 0;
  }
  
  body {
      font-family: 'Open Sans', Helvetica, Arial, sans-serif;
      background: #ededed;
  }
  
  input, button {
      border: none;
      outline: none;
      background: none;
      font-family: 'Open Sans', Helvetica, Arial, sans-serif;
  }
  
  .tip {
      font-size: 20px;
      margin: 40px auto 50px;
      text-align: center;
  }
  
  .dowebok {
      overflow: hidden;
      position: absolute;
      left: 50%;
      top: 56%;
      width: 900px;
      height: 550px;
      margin: -300px 0 0 -450px;
      background: #fff;
      box-shadow: 0 4px 8px rgba(0, 0, 0, 0.2);
  }
  
  
  .form {
      position: relative;
      width: 640px;
      height: 100%;
      transition: -webkit-transform 0.6s ease-in-out;
      transition: transform 0.6s ease-in-out;
      transition: transform 0.6s ease-in-out, -webkit-transform 0.6s ease-in-out;
      padding: 50px 30px 0;
      padding: 130px 30px 0;  /* 增加上方的 padding 来使内容居中 */
  }
  
  .sub-cont {
      overflow: hidden;
      position: absolute;
      left: 640px;
      top: 0;
      width: 900px;
      height: 100%;
      padding-left: 260px;
      background: #fff;
      transition: -webkit-transform 0.6s ease-in-out;
      transition: transform 0.6s ease-in-out;
      transition: transform 0.6s ease-in-out, -webkit-transform 0.6s ease-in-out;
  }
  
  .dowebok.s--signup .sub-cont {
      -webkit-transform: translate3d(-640px, 0, 0);
      transform: translate3d(-640px, 0, 0);
  }
  
  button {
      display: block;
      margin: 0 auto;
      width: 260px;
      height: 36px;
      border-radius: 30px;
      background: #409eff;
      color: #fff;
      font-size: 15px;
      cursor: pointer;
  }
  
  .img {
      overflow: hidden;
      z-index: 2;
      position: absolute;
      left: 0;
      top: 0;
      width: 260px;
      height: 100%;
      padding-top: 360px;
  }
  
  .img:before {
      content: '';
      position: absolute;
      right: 0;
      top: 0;
      width: 900px;
      height: 100%;
      background-image: url(../assets/bg2.jpg);
      background-size: cover;
      transition: -webkit-transform 0.6s ease-in-out;
      transition: transform 0.6s ease-in-out;
      transition: transform 0.6s ease-in-out, -webkit-transform 0.6s ease-in-out;
  }
  
  .img:after {
      content: '';
      position: absolute;
      left: 0;
      top: 0;
      width: 100%;
      height: 100%;
      background: rgba(0, 0, 0, 0.6);
  }
  
  .dowebok.s--signup .img:before {
      -webkit-transform: translate3d(640px, 0, 0);
      transform: translate3d(640px, 0, 0);
  }
  
  .img__text {
      z-index: 2;
      position: absolute;
      left: 0;
      top: 110px;
      width: 100%;
      padding: 0 20px;
      text-align: center;
      color: #fff;
      transition: -webkit-transform 0.6s ease-in-out;
      transition: transform 0.6s ease-in-out;
      transition: transform 0.6s ease-in-out, -webkit-transform 0.6s ease-in-out;
  }
  
  .img__text h2 {
      margin-bottom: 10px;
      font-weight: normal;
  }
  
  .img__text p {
      font-size: 14px;
      line-height: 1.5;
  }
  
  .dowebok.s--signup .img__text.m--up {
      -webkit-transform: translateX(520px);
      transform: translateX(520px);   
  }
  .img__text.m--in {
      -webkit-transform: translateX(-520px);
      transform: translateX(-520px);
  }
  
  .dowebok.s--signup .img__text.m--in {
      -webkit-transform: translateX(0);
      transform: translateX(0);
  }
  
  .img__btn {
      overflow: hidden;
      z-index: 2;
      position: relative;
      width: 100px;
      height: 36px;
      margin: 0 auto;
      background: transparent;
      color: #fff;
      text-transform: uppercase;
      font-size: 15px;
      cursor: pointer;
  }
  .img__btn:after {
      content: '';
      z-index: 2;
      position: absolute;
      left: 0;
      top: 0;
      width: 100%;
      height: 100%;
      border: 2px solid #fff;
      border-radius: 30px;
  }
  
  .img__btn span {
      position: absolute;
      left: 0;
      top: 0;
      display: flex;
      justify-content: center;
      align-items: center;
      width: 100%;
      height: 100%;
      transition: -webkit-transform 0.6s;
      transition: transform 0.6s;
      transition: transform 0.6s, -webkit-transform 0.6s;
  }
  
  .img__btn span.m--in {
      -webkit-transform: translateY(-72px);
      transform: translateY(-72px);
  }
  
  .dowebok.s--signup .img__btn span.m--in {
      -webkit-transform: translateY(0);
      transform: translateY(0);
  }
  
  .dowebok.s--signup .img__btn span.m--up {
      -webkit-transform: translateY(72px);
      transform: translateY(72px);
  }
  
  h2 {
      width: 100%;
      font-size: 26px;
      text-align: center;
  }
  
  label {
      display: block;
      width: 260px;
      margin: 25px auto 0;
      text-align: center;
  }
  
  label span {
      font-size: 12px;
      color: #909399;
      text-transform: uppercase;
  }
  
  input {
    width: 100%;
    padding: 12px;
    font-size: 14px;
    border-bottom: 1px solid rgba(0, 0, 0, 0.4);
    border-radius: 5px;
    box-sizing: border-box;
  }
  
  .submit {
      margin-top: 40px;
      margin-bottom: 20px;
      background: #409eff;
      text-transform: uppercase;
  }
  
  .sign-in {
      transition-timing-function: ease-out;
  }
  .dowebok.s--signup .sign-in {
      transition-timing-function: ease-in-out;
      transition-duration: 0.6s;
      -webkit-transform: translate3d(640px, 0, 0);
      transform: translate3d(640px, 0, 0);
  }
  
  .sign-up {
      -webkit-transform: translate3d(-900px, 0, 0);
      transform: translate3d(-900px, 0, 0);
  }
  .dowebok.s--signup .sign-up {
      -webkit-transform: translate3d(0, 0, 0);
      transform: translate3d(0, 0, 0);
  }
  
  </style>