<template>
  <div class="history-container">
    <div class="header-bar">
      <h1>📜 历史碳足迹记录</h1>
      <button @click="exportToCSV" class="export-btn">📥 导出为 CSV</button>
    </div>
    <div class="filter-bar">
      <select v-model="sortOrder" class="sort-select">
        <option value="desc">最新优先</option>
        <option value="asc">最早优先</option>
      </select>

      <button @click="filterLast7Days" class="filter-btn">📅 最近 7 天</button>
      <button @click="filterLast30Days" class="filter-btn">📅 最近 30 天</button>
      <button @click="resetFilter" class="filter-btn">🔄 全部显示</button>
    </div>
    <div v-if="history.length === 0" class="no-records">
      暂无历史记录，请先进行一次碳足迹计算。
    </div>

    <div v-else class="history-list">
      <div 
        v-for="record in displayedHistory" 
        :key="record.id"
        class="history-card"
        @click="viewDetails(record)"
      >
        <div class="card-header">
          <span class="updated_at">{{ formatTime(record.updated_at) }}</span>
        </div>
        <div class="card-body">
          <p><strong>总碳排放：</strong>{{ Math.round(record.carbon) }} kg CO₂</p>
          <p><strong>计算范围：</strong>{{ record.month }} 个月</p>
        </div>
        <div class="card-footer">
          <!-- <button @click.stop="deleteRecord(record.id)" class="delete-btn">🗑 删除</button> -->
        </div>
      </div>
    </div>

    <!-- 弹窗显示详情 -->
    <transition name="fade">
      <div v-if="selectedRecord" class="modal-overlay" @click.self="closeModal">
        <div class="modal-content zoom-in">
          <h2>📊 计算详情</h2>
          <p><strong>计算时间：</strong>{{ formatTime(selectedRecord.updated_at) }}</p>
          <p><strong>总碳排放：</strong>{{ Math.round(selectedRecord.carbon) }} kg CO₂</p>
          <p><strong>计算范围：</strong>{{ selectedRecord.month }} 个月</p>
          <div class="breakdown">
           <p>🚗 交通：{{ Math.round(selectedRecord.detail.transportation || 0) }} kg</p>
           <p>🏠 家庭：{{ Math.round(selectedRecord.detail.household || 0) }} kg</p>
           <p>🥦 饮食：{{ Math.round(selectedRecord.detail.diet || 0) }} kg</p>
           <p>🛍 购物：{{ Math.round(selectedRecord.detail.shopping || 0) }} kg</p>
          </div>
        </div>
      </div>
    </transition>
  </div>
</template>

<script setup>
import { ref, onMounted, computed  } from 'vue'
import api from '@/http'; // 引入全局配置的 axios 实例
const history = ref([])
const selectedRecord = ref(null)
const sortOrder = ref('desc') // 排序方式
const filteredHistory = ref([]) // 筛选后的历史记录

// 计算属性：排序 + 筛选后的历史记录
const displayedHistory = computed(() => {
  let list = filteredHistory.value.length > 0 ? filteredHistory.value : history.value

  // 排序逻辑
  return list.slice().sort((a, b) => {
    const timeA = new Date(a.updated_at).getTime();
    const timeB = new Date(b.updated_at).getTime();
    return sortOrder.value === 'desc' ? timeB - timeA : timeA - timeB;
  });
})

// 筛选最近 7 天
function filterLast7Days() {
  const now = Date.now()
  filteredHistory.value = history.value.filter(r => now - r.updated_at <= 7 * 86400000)
}

// 筛选最近 30 天
function filterLast30Days() {
  const now = Date.now()
  filteredHistory.value = history.value.filter(r => now - r.updated_at <= 30 * 86400000)
}

// 恢复全部
function resetFilter() {
  filteredHistory.value = []
}
function exportToCSV() {

  if (history.value.length === 0) {
    ElMessage.warning("没有可导出的数据");
    return;
  }

  const csvRows = [];

  // 表头
  csvRows.push([
    "计算时间",
    "总碳排放(kg CO₂)",
    "交通(kg)",
    "家庭(kg)",
    "饮食(kg)",
    "购物(kg)",
    "计算范围(月)"
  ].join(","));

  // 数据行
  for (const record of history.value) {
    csvRows.push([
      formatTime(record.updated_at),
      Math.round(record.carbon),
      Math.round(record.detail.transportation),
      Math.round(record.detail.household),
      Math.round(record.detail.diet),
      Math.round(record.detail.shopping),
      record.month
    ].join(","));
  }

  // 创建下载链接
  const csvString = csvRows.join("\n");
  const blob = new Blob([`\uFEFF${csvString}`], { type: 'text/csv;charset=utf-8;' });
  const url = URL.createObjectURL(blob);
  const link = document.createElement("a");


  link.setAttribute("href", url);
  const userStr = localStorage.getItem('user');
  let username = '用户'; // 默认值

  if (userStr) {
    try {
      const user = JSON.parse(userStr);
      if (user && user.username) {
        username = user.username;
      }
    } catch (e) {
      console.error('解析 localStorage 中的 user 出错:', e);
    }
  }

  const date = new Date().toISOString().slice(0, 10);
  const filename = `${username}的历史碳足迹记录_${date}.csv`;
  link.setAttribute("download", filename);
  link.style.visibility = "hidden";
  document.body.appendChild(link);
  link.click();
  document.body.removeChild(link);
}

// 格式化时间
function formatTime(updated_at) {
  const date = new Date(updated_at)
  const YYYY = date.getFullYear()
  const MM = ('0' + (date.getMonth() + 1)).slice(-2)
  const DD = ('0' + date.getDate()).slice(-2)
  const hh = ('0' + date.getHours()).slice(-2)
  const mm = ('0' + date.getMinutes()).slice(-2)
  const ss = ('0' + date.getSeconds()).slice(-2)
  
  // 加上单引号，强制识别为文本
  return `'${YYYY}-${MM}-${DD} ${hh}:${mm}:${ss}`
}

// 获取历史记录
async function loadHistory() {
    // 检查是否已登录
    const userStr = localStorage.getItem("user");
    if (!userStr) {
      ElMessage.error("请先登录");
      router.push("/MyLogin");
      return;
    }

    const url = "/user/history";
    // 发送请求并处理异常
  try {
    const response = await api.get(url);
    console.log(response.data.Data);
      history.value = response.data.Data.map((record, index) => ({
        id: index + 1,
        carbon: record.carbon,
        month: record.month,
        detail: {
          transportation: record.detail.transportation,
          household: record.detail.household,
          diet: record.detail.diet,
          shopping: record.detail.shopping
        },
        updated_at: record.updated_at
      }));
  } catch (error) {
    console.error("请求失败:", error);
    if (error.response && error.response.status === 401) {
      ElMessage.error("登录超时，请重新登录");
      localStorage.removeItem("token");
      delete api.defaults.headers.common["Authorization"];
      router.push("/MyLogin");
    }
  }
  
}

// // 删除记录
// function deleteRecord(id) {
//   history.value = history.value.filter(r => r.id !== id)
//   localStorage.setItem('carbonFootprintHistory', JSON.stringify(history.value))
// }

// 查看详情
function viewDetails(record) {
  selectedRecord.value = record
}

// 关闭弹窗
function closeModal() {
  selectedRecord.value = null
}

onMounted(() => {
  loadHistory()
})
</script>

<style scoped>
.history-container {
  padding: 40px 20px;
  height: 100vh;
  margin: auto;
  font-family: 'Segoe UI', sans-serif;
  background-color: #caffe0;
  box-sizing: border-box;
}

h1 {
  text-align: center;
  color: #2c3e50;
  margin-bottom: 30px;
}

.no-records {
  text-align: center;
  color: #7f8c8d;
  font-size: 1.1rem;
  padding: 40px;
  background-color: #ecf0f1;
  border-radius: 10px;
  box-shadow: 0 2px 6px rgba(0, 0, 0, 0.1);
}

.history-list {
  display: grid;
  grid-template-columns: repeat(auto-fill, minmax(280px, 1fr));
  gap: 20px;
}

.history-card {
  background-color: white;
  border-radius: 10px;
  box-shadow: 0 4px 10px rgba(0, 0, 0, 0.1);
  cursor: pointer;
  transition: transform 0.3s ease, box-shadow 0.3s ease;
  overflow: hidden;
  display: flex;
  flex-direction: column;
  height: 100%;
}

.history-card:hover {
  transform: translateY(-5px);
  box-shadow: 0 6px 16px rgba(0, 0, 0, 0.15);
}

.card-header {
  background-color: skyblue;
  padding: 10px 15px;
}

.card-header .updated_at {
  color: white;
  font-weight: bold;
  font-size: 0.9rem;
}

.card-body {
  padding: 15px;
  flex-grow: 1;
  font-size: 1rem;
  color: #2c3e50;
}

.card-footer {
  display: flex;
  justify-content: flex-end;
  padding: 10px 15px;
  background-color: #f8f9fa;
  border-top: 1px solid #ddd;
}

.delete-btn {
  background-color: #e74c3c;
  color: white;
  border: none;
  padding: 6px 12px;
  border-radius: 5px;
  font-size: 0.9rem;
  cursor: pointer;
  transition: background-color 0.3s ease;
}

.delete-btn:hover {
  background-color: #c0392b;
}

/* 弹窗样式 */
.modal-overlay {
  position: fixed;
  top: 0; left: 0; right: 0; bottom: 0;
  background-color: rgba(0, 0, 0, 0.7);
  display: flex;
  justify-content: center;
  align-items: center;
  z-index: 1000;
}

.modal-content {
  background-color: white;
  padding: 30px;
  border-radius: 12px;
  width: 90%;
  max-width: 500px;
  box-shadow: 0 10px 30px rgba(0, 0, 0, 0.2);
  animation: zoom-in 0.3s ease forwards;
}

.modal-content h2 {
  margin-top: 0;
  margin-bottom: 15px;
  color: #2c3e50;
}

.breakdown {
  margin-top: 10px;
}

.breakdown p {
  margin: 8px 0;
  font-size: 1rem;
  color: #34495e;
}

/* 过渡动画 */
.fade-enter-active,
.fade-leave-active {
  transition: opacity 0.3s ease;
}
.fade-enter-from,
.fade-leave-to {
  opacity: 0;
}

@keyframes zoom-in {
  from {
    transform: scale(0.9);
    opacity: 0;
  }
  to {
    transform: scale(1);
    opacity: 1;
  }
}

.range-tag {
  background-color: #3498db;
  color: white;
  padding: 2px 6px;
  border-radius: 4px;
  font-size: 0.9rem;
  margin-left: 5px;
}

.export-btn {
  display: block;
  margin: 20px auto;
  background-color: #3498db;
  color: white;
  border: none;
  padding: 10px 20px;
  font-size: 1rem;
  border-radius: 6px;
  cursor: pointer;
  transition: background-color 0.3s ease;
}

.export-btn:hover {
  background-color: #2980b9;
}

.filter-bar {
  display: flex;
  justify-content: center;
  gap: 15px;
  margin-bottom: 20px;
  flex-wrap: wrap;
}

.sort-select,
.filter-btn {
  padding: 8px 12px;
  font-size: 0.95rem;
  border-radius: 4px;
  border: 1px solid #ccc;
  cursor: pointer;
}

.sort-select {
  background-color: white;
}

.filter-btn {
  background-color: #0fc4f1e4;
  color: #fff;
  border: none;
  transition: background-color 0.3s;
}

.filter-btn:hover {
  background-color: #2980b9;
}

.header-bar {
  position: relative;
  text-align: center;
  padding: 10px 0;
  margin-bottom: 10px;
}

.header-bar h1 {
  margin: 0;
  font-size: 1.8rem;
  color: #2c3e50;
}

.export-btn {
  position: absolute;
  right: 0;
  top: -25px; /* 调整垂直位置 */
  background-color: #3498db;
  color: white;
  border: none;
  padding: 10px 20px;
  font-size: 1rem;
  border-radius: 6px;
  cursor: pointer;
  transition: background-color 0.3s ease;
  z-index: 10;
}

.export-btn:hover {
  background-color: #2980b9;
}

/* 响应式设计 */
@media (max-width: 600px) {
  .modal-content {
    width: 95%;
    padding: 20px;
  }

  .card-header .updated_at {
    font-size: 0.8rem;
  }

  .card-body {
    font-size: 0.95rem;
  }
}
</style>