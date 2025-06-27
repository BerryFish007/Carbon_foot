<template>
  <div v-if="route.path !== '/'">
    <el-menu
      :default-active="activeIndex"
      class="el-menu-demo"
      mode="horizontal"
      @select="handleSelect"
    >
      <el-menu-item index="1">了解碳足迹</el-menu-item>
      <el-menu-item index="2">碳足迹计算</el-menu-item>
      <el-menu-item index="4">减少碳足迹</el-menu-item>
      <el-menu-item index="5">减碳计划表</el-menu-item>
      <el-menu-item index="3">历史计算结果</el-menu-item>

      <!-- 右侧用户信息 -->
      <div class="user-info">
        <el-dropdown @command="handleDropdownCommand">
          <span class="el-dropdown-link">
            {{ currentUserUsername }} <i class="el-icon-arrow-down el-icon--right"></i>
          </span>
          <template #dropdown>
            <el-dropdown-menu>
              <el-dropdown-item command="logout">退出登录</el-dropdown-item>
            </el-dropdown-menu>
          </template>
        </el-dropdown>
      </div>
    </el-menu>
  </div>

  <div class="main-content">
    <RouterView />
  </div>
</template>

<script setup>
import { ref, onMounted, onUnmounted } from 'vue';
import { useRouter, useRoute, } from 'vue-router';
import api from '@/http'; // 引入全局 axios 实例

const route = useRoute(); // 获取当前路由
const router = useRouter();
const currentUserUsername = ref('用户');

// 存储原始的 setItem 方法
const originalSetItem = localStorage.setItem;

// 重写 localStorage.setItem 以触发事件
localStorage.setItem = function(key, value) {
  const result = originalSetItem.apply(this, arguments);
  window.dispatchEvent(new Event(`storage-${key}`));
  return result;
};

// 更新用户名函数
const updateUsername = () => {
  const userStr = localStorage.getItem('user');
  if (userStr) {
    try {
      const user = JSON.parse(userStr);
      if (user && user.username) {
        currentUserUsername.value = user.username;
      }
    } catch (e) {
      console.error('解析 localStorage 中的 user 出错:', e);
    }
  }
};

onMounted(() => {
  // 初始加载时设置用户名
default_username:
  updateUsername();

  // 监听用户信息变化
  window.addEventListener('storage-user', updateUsername);
});

onUnmounted(() => {
  // 移除事件监听
  window.removeEventListener('storage-user', updateUsername);
});

// 处理下拉菜单点击
function handleDropdownCommand(command) {
  switch (command) {
    case 'logout':
    ElMessage.success("已退出登录");
      // 清除本地 token和用户信息
      localStorage.removeItem('token');
      localStorage.removeItem("user");
      // 删除 axios 请求头中的 authorization
      delete api.defaults.headers.common['Authorization'];
      router.push('/MyLogin'); // 跳转到首页或其他页面
      break;
    default:
      console.log('未知命令:', command);
  }
}

import { computed } from 'vue';
const activeIndex = computed(() => {
  switch (router.currentRoute.value.path) {
    case '/show':
      return '1';
    case '/person':
      return '2';
    case '/contactList':
      return '3';
    case '/adovocate':
      return '4';
    case '/todoList':
      return '5';
    default:
      return '1';
  }
});

const handleSelect = (index) => {
  switch (index) {
    case '1':
      router.push('/show');
      break;
    case '2':
      router.push('/person');
      break;
    case '3':
      router.push('/contactList');
      break;
    case '4':
      router.push('/adovocate');
      break;
    case '5':
      router.push('/todoList');
      break;
    default:
      router.push('/');
  }
}
</script>

<style lang="scss" scoped>
.el-menu-demo {
  position: fixed;
  top: 0;
  left: 0;
  width: 100%;
  z-index: 999;
  background-color: #fff;
  box-shadow: 0 2px 4px rgba(0, 0, 0, 0.1);
}

.main-content {
  margin-top: 60px;
}

.el-menu-demo .el-menu-item.is-active {
  background-color: transparent !important;
  color: #409EFF !important;
  font-weight: bold;
}

/* 右侧用户信息 */
.user-info {
  position: absolute;
  right: 20px;
  top: 0;
  height: 60px;
  display: flex;
  align-items: center;
}

.el-dropdown-link {
  cursor: pointer;
  font-size: 14px;
  color: #333;
}
</style>