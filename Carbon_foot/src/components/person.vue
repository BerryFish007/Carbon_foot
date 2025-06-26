<template>
  <el-backtop
			:right="35"
			:bottom="30" />
  <div class="bg">
    <div class="container">
      <header>
        <h1>碳足迹计算</h1>
        <p class="subtitle">
          碳足迹（Carbon Footprint）：一个人在一年内，因衣食住行等活动所消耗的能源，<br />
          最终转化为大气中二氧化碳(CO2)的量，用于衡量人类活动对于气候变化的影响。
        </p>
      </header>

      <main>
        <div class="calculator-container">
          <div class="tabs">
            <button
              class="tab-btn"
              :class="{ active: activeTab === 'transportation' }"
              @click="activeTab = 'transportation'"
              data-tab="transportation"
            >
              交通
            </button>
            <button
              class="tab-btn"
              :class="{ active: activeTab === 'home' }"
              @click="activeTab = 'home'"
              data-tab="home"
            >
              家庭
            </button>
            <button
              class="tab-btn"
              :class="{ active: activeTab === 'food' }"
              @click="activeTab = 'food'"
              data-tab="food"
            >
              饮食
            </button>
            <button
              class="tab-btn"
              :class="{ active: activeTab === 'shopping' }"
              @click="activeTab = 'shopping'"
              data-tab="shopping"
            >
              购物
            </button>
          </div>

          <div class="tab-content" :class="{ active: activeTab === 'transportation' }" id="transportation">
            <h2>交通碳排放计算</h2>
            <div class="input-group">
              <label for="car-mileage">汽车行驶里程 (公里/月):</label>
              <input type="number" id="car-mileage" min="0" v-model="formData.carMileage" />
              <select id="car-type" v-model="formData.carType">
                <option value="gasoline">汽油车</option>
                <option value="diesel">柴油车</option>
                <option value="hybrid">混合动力</option>
                <option value="electric">电动车</option>
              </select>
            </div>

            <div class="input-group">
              <label for="public-transport">公共交通里程 (公里/月):</label>
              <input type="number" id="public-transport" min="0" v-model="formData.publicTransport" />
              <select id="transport-type" v-model="formData.transportType">
                <option value="bus">公交车</option>
                <option value="subway">地铁</option>
                <option value="train">火车</option>
              </select>
            </div>

            <div class="input-group">
              <label for="flights">飞行里程 (公里):</label>
              <input type="number" id="flights" min="0" v-model="formData.flights" />
              <select id="flight-class" v-model="formData.flightClass">
                <option value="economy">经济舱</option>
                <option value="business">商务舱</option>
                <option value="first">头等舱</option>
              </select>
            </div>

          </div>

          <div class="tab-content" :class="{ active: activeTab === 'home' }" id="home">
            <h2>家庭能源碳排放计算</h2>
            <div class="input-group">
              <label for="electricity">每月用电量 (千瓦时):</label>
              <input type="number" id="electricity" min="0" v-model="formData.electricity" />
            </div>

            <div class="input-group">
              <label for="gas">每月天然气使用量 (立方米):</label>
              <input type="number" id="gas" min="0" v-model="formData.gas" />
            </div>

            <div class="input-group">
              <label for="water">每月用水量 (吨):</label>
              <input type="number" id="water" min="0" v-model="formData.water" />
            </div>

            <div class="checkbox-group">
              <label>
                <input type="checkbox" id="renewable-energy" v-model="formData.renewableEnergy" />
                我家使用可再生能源(太阳能/风能等)
              </label>
            </div>
          </div>

          <div class="tab-content" :class="{ active: activeTab === 'food' }" id="food">
            <h2>饮食碳排放计算</h2>
            <div class="input-group">
              <label for="meat-consumption">每周肉类消费 (克):</label>
              <input type="number" id="meat-consumption" min="0" v-model="formData.meatConsumption" />
            </div>

            <div class="input-group">
              <label for="dairy-consumption">每周乳制品消费 (克):</label>
              <input type="number" id="dairy-consumption" min="0" v-model="formData.dairyConsumption" />
            </div>

            <div class="input-group">
              <label for="organic-food">有机食品消费比例 (%):</label>
              <input
                type="number"
                id="organic-food"
                min="0"
                max="100"
                v-model="formData.organicFood"
              />
            </div>

            <div class="input-group">
              <label for="food-waste">每周食物浪费量 (克):</label>
              <input type="number" id="food-waste" min="0" v-model="formData.foodWaste" />
            </div>
          </div>

          <div class="tab-content" :class="{ active: activeTab === 'shopping' }" id="shopping">
            <h2>购物碳排放计算</h2>
            <div class="input-group">
              <label for="clothing">每月服装消费 (元):</label>
              <input type="number" id="clothing" min="0" v-model="formData.clothing" />
            </div>

            <div class="input-group">
              <label for="electronics">每月电子产品消费 (元):</label>
              <input type="number" id="electronics" min="0" v-model="formData.electronics" />
            </div>

            <div class="input-group">
              <label for="furniture">每月家具/家居用品消费 (元):</label>
              <input type="number" id="furniture" min="0" v-model="formData.furniture" />
            </div>

            <div class="checkbox-group">
              <label>
                <input type="checkbox" id="second-hand" v-model="formData.secondHand" />
                我经常购买二手商品
              </label>
            </div>
          </div>


          <div class="input-group">
              <label for="month-range">选择计算结果涵盖的时间段 (1-12个月):</label>
              <input 
                type="number" 
                id="month-range" 
                v-model.number="formData.monthRange"
                min="1" 
                max="12" 
                @input="validateMonthRange"
              />
              <small v-if="monthRangeError" style="color: red;">请输入1到12之间的整数。</small>
            </div>

          <div class="button-container">
            <button id="calculate-btn" class="calculate-btn" @click="calculateFootprint">
              计算我的碳足迹
          </button>
          </div>
        </div>

        <div id="result" class="result-container" :class="{ hidden: !showResults }">
          <h2>您的碳足迹结果</h2>
          <p id="time-range-info">（计算结果基于过去 {{ formData.monthRange }} 个月的数据）</p>
          <div class="carbon-circle">
            <div class="circle">
              <div id="carbon-value">{{ Math.round(footprints.total) }}</div>
              <div class="unit">千克 CO₂</div>
            </div>
          </div>

          <div class="breakdown">
            <h3>碳排放构成</h3>
            <div class="breakdown-item">
              <div class="category">交通</div>
              <div class="progress-container">
                <div
                  id="transport-progress"
                  class="progress-bar"
                  :style="{ width: transportPercent + '%' }"
                ></div>
              </div>
              <div id="transport-value" class="value">{{ Math.round(footprints.transport) }} kg</div>
            </div>

            <div class="breakdown-item">
              <div class="category">家庭</div>
              <div class="progress-container">
                <div
                  id="home-progress"
                  class="progress-bar"
                  :style="{ width: homePercent + '%' }"
                ></div>
              </div>
              <div id="home-value" class="value">{{ Math.round(footprints.home) }} kg</div>
            </div>

            <div class="breakdown-item">
              <div class="category">饮食</div>
              <div class="progress-container">
                <div
                  id="food-progress"
                  class="progress-bar"
                  :style="{ width: foodPercent + '%' }"
                ></div>
              </div>
              <div id="food-value" class="value">{{ Math.round(footprints.food) }} kg</div>
            </div>

            <div class="breakdown-item">
              <div class="category">购物</div>
              <div class="progress-container">
                <div
                  id="shopping-progress"
                  class="progress-bar"
                  :style="{ width: shoppingPercent + '%' }"
                ></div>
              </div>
              <div id="shopping-value" class="value">{{ Math.round(footprints.shopping) }} kg</div>
            </div>
          </div>

          <div class="comparison">
            <h3>对比分析</h3>
            <p id="comparison-text">{{ comparisonText }}</p>
            <div class="comparison-items">
              <div class="comparison-item">
                <div class="icon">🌳</div>
                <div id="trees" class="text">{{ Math.round(footprints.total / 20) }}棵树</div>
                <div class="description">需要种植来抵消您的排放</div>
              </div>
              <div class="comparison-item">
                <div class="icon">🚗</div>
                <div id="car-miles" class="text">{{ Math.round(footprints.total / 0.2) }}公里</div>
                <div class="description">相当于汽车行驶的里程</div>
              </div>
              <div class="comparison-item">
                <div class="icon">🏠</div>
                <div id="homes" class="text">{{ (footprints.total / avgFootprint).toFixed(1) }}个家庭</div>
                <div class="description">相当于普通家庭的年排放</div>
              </div>
            </div>
          </div>

          <div class="tips">
            <h3>减少碳足迹的建议</h3>
            <div class="tips-container">
              <div class="tip-card">
                <div class="tip-icon">🚆</div>
                <div class="tip-content">
                  <h4>绿色出行</h4>
                  <p>多使用公共交通、骑自行车或步行，减少私家车使用</p>
                </div>
              </div>
              <div class="tip-card">
                <div class="tip-icon">💡</div>
                <div class="tip-content">
                  <h4>节能家居</h4>
                  <p>使用节能电器，合理设置空调温度，随手关灯</p>
                </div>
              </div>
              <div class="tip-card">
                <div class="tip-icon">🥦</div>
                <div class="tip-content">
                  <h4>健康饮食</h4>
                  <p>减少肉类消费，增加植物性饮食，减少食物浪费</p>
                </div>
              </div>
              <div class="tip-card">
                <div class="tip-icon">🔄</div>
                <div class="tip-content">
                  <h4>可持续购物</h4>
                  <p>购买耐用品，选择二手商品，减少不必要消费</p>
                </div>
              </div>
            </div>
          </div>

          <div class="offset-options">
            <h3>碳补偿选项</h3>
            <p>考虑通过以下方式抵消您的碳排放:</p>
            <div class="offset-buttons">
              <router-link to="/adovocate">
                <button class="offset-btn">减少碳足迹</button>
              </router-link>
              <router-link to="/todoList">
                <button class="offset-btn">减碳计划表</button>
              </router-link>
            </div>
          </div>
        </div>
      </main>
    </div>
  </div>
</template>

<script setup>
import { ref, reactive, computed,onMounted } from 'vue'
import axios from "axios";  // 导入Axios库
import api from '@/http'; // 引入全局配置的 axios 实例
const activeTab = ref('transportation')
const showResults = ref(false)
const avgFootprint = 8000 // 中国平均每人年碳足迹约8吨
const monthRangeError = ref(false)


onMounted(() => {
  const savedToken = localStorage.getItem('token');
  if (savedToken) {
    api.defaults.headers.common['Authorization'] = `Bearer ${savedToken}`;
  } else {
    ElMessage.error("未检测到登录信息，请重新登录");
    router.push("/MyLogin");
  }
});


const formData = reactive({
  // 交通
  carMileage: 0,
  carType: 'gasoline',
  publicTransport: 0,
  transportType: 'bus',
  flights: 0,
  flightClass: 'economy',
  monthRange: 12, // 默认全年
  // 家庭
  electricity: 0,
  gas: 0,
  water: 0,
  renewableEnergy: false,
  
  // 饮食
  meatConsumption: 0,
  dairyConsumption: 0,
  organicFood: 0,
  foodWaste: 0,
  
  // 购物
  clothing: 0,
  electronics: 0,
  furniture: 0,
  secondHand: false
})

const footprints = reactive({
  transport: 0,
  home: 0,
  food: 0,
  shopping: 0,
  total: 0
})

const factors = {
  // 交通
  car: {
    gasoline: 0.2,   // 汽油车 每公里排放
    diesel: 0.18,    // 柴油车
    hybrid: 0.12,    // 混合动力
    electric: 0.05   // 电动车 (假设使用电网平均电力)
  },
  publicTransport: {
    bus: 0.05,       // 公交车 每公里排放
    subway: 0.03,    // 地铁
    train: 0.04      // 火车
  },
  flight: {
    economy: 0.15,   // 经济舱 每公里排放
    business: 0.25,  // 商务舱
    first: 0.35      // 头等舱
  },
  // 家庭
  electricity: 0.5,    // 每千瓦时电力排放 (假设电网平均值)
  gas: 2.0,           // 每立方米天然气排放
  water: 0.3,          // 每吨水处理排放
  renewableEnergyMultiplier: 0.3, // 使用可再生能源的排放乘数
  // 饮食
  meat: 0.025,        // 每克肉类排放
  dairy: 0.01,        // 每克乳制品排放
  organicFoodMultiplier: 0.8, // 有机食品排放乘数
  foodWaste: 0.02,    // 每克食物浪费排放
  // 购物
  clothing: 0.02,      // 每元服装消费排放
  electronics: 0.015,  // 每元电子产品消费排放
  furniture: 0.03,     // 每元家具消费排放
  secondHandMultiplier: 0.3 // 二手商品排放乘数
}

const transportPercent = computed(() => {
  const total = footprints.transport + footprints.home + footprints.food + footprints.shopping
  return total > 0 ? (footprints.transport / total) * 100 : 0
})

const homePercent = computed(() => {
  const total = footprints.transport + footprints.home + footprints.food + footprints.shopping
  return total > 0 ? (footprints.home / total) * 100 : 0
})

const foodPercent = computed(() => {
  const total = footprints.transport + footprints.home + footprints.food + footprints.shopping
  return total > 0 ? (footprints.food / total) * 100 : 0
})

const shoppingPercent = computed(() => {
  const total = footprints.transport + footprints.home + footprints.food + footprints.shopping
  return total > 0 ? (footprints.shopping / total) * 100 : 0
})

const comparisonText = computed(() => {
  const difference = footprints.total - avgFootprint
  if (difference < 0) {
    return `您的碳足迹比中国平均水平低${Math.abs(difference).toFixed(0)}千克。做得很好！`
  } else if (difference > 0) {
    return `您的碳足迹比中国平均水平高${difference.toFixed(0)}千克。考虑采取一些减排措施吧。`
  } else {
    return '您的碳足迹与中国平均水平相当。'
  }
})
//验证函数防止非法输入
function validateMonthRange(event) {
  const value = parseInt(event.target.value)
  if (value >= 1 && value <= 12) {
    formData.monthRange = value
    monthRangeError.value = false
  } else {
    monthRangeError.value = true
  }
}

function calculateFootprint() {
  // 计算交通碳排放
  let carFootprint = formData.carMileage * factors.car[formData.carType] * formData.monthRange
  let transportFootprint = formData.publicTransport * factors.publicTransport[formData.transportType] * formData.monthRange
  let flightFootprint = formData.flights * factors.flight[formData.flightClass]
  footprints.transport = carFootprint + transportFootprint + flightFootprint
  
  // 计算家庭碳排放
  let electricityFootprint = formData.electricity * factors.electricity * formData.monthRange
  if (formData.renewableEnergy) {
    electricityFootprint *= factors.renewableEnergyMultiplier
  }

  let gasFootprint = formData.gas * factors.gas * formData.monthRange
  let waterFootprint = formData.water * factors.water * formData.monthRange
  footprints.home = electricityFootprint + gasFootprint + waterFootprint
  
  // 计算饮食碳排放
  // 每周数据 × 周数（月份数 / 12 * 52）
  let totalWeeks = (formData.monthRange / 12) * 52
  let meatFootprint = formData.meatConsumption * factors.meat * totalWeeks
  let dairyFootprint = formData.dairyConsumption * factors.dairy * totalWeeks
  let foodWasteFootprint = formData.foodWaste * factors.foodWaste * totalWeeks
  
  // 有机食品调整
  let organicAdjustment = 1 - (formData.organicFood / 100) * (1 - factors.organicFoodMultiplier)
  footprints.food = (meatFootprint + dairyFootprint + foodWasteFootprint) * organicAdjustment
  
  // 计算购物碳排放
  let clothingFootprint = formData.clothing * factors.clothing * formData.monthRange
  let electronicsFootprint = formData.electronics * factors.electronics * formData.monthRange
  let furnitureFootprint = formData.furniture * factors.furniture * formData.monthRange
  
  // 二手商品调整
  let shoppingTotal = clothingFootprint + electronicsFootprint + furnitureFootprint
  if (formData.secondHand) {
    shoppingTotal *= factors.secondHandMultiplier
  }
  footprints.shopping = shoppingTotal
  
  // 计算总量
  footprints.total = footprints.transport + footprints.home + footprints.food + footprints.shopping
  
  if (footprints.total <= 0) {
    ElMessage.error("请填写有效的数据！");
    return;
  }

  // 显示结果
  showResults.value = true

  // 滚动到结果
  setTimeout(() => {
    const resultSection = document.getElementById('result')
    if (resultSection) {
      window.scrollTo({
        top: resultSection.offsetTop - 70, // 向上偏移 70px，即结果向下显示一点
        behavior: 'smooth'
      })
    }
  }, 100)

  // 检查是否已登录
const userStr = localStorage.getItem("user");
if (!userStr) {
  ElMessage.error("请先登录");
  router.push("/MyLogin");
  return;
}

// 构建碳排放数据
const carbonData = {
  Carbon: Math.round(footprints.total), // 总排放量
  month: formData.monthRange, // 时间范围（月）
  detail: {
    transportation: Math.round(footprints.transport),
    household: Math.round(footprints.home),
    diet: Math.round(footprints.food),
    shopping: Math.round(footprints.shopping)
  }
};

const url = "/user/calculate";


// 发送请求并处理异常
try {
  const response =  api.post(url, carbonData);
  ElMessage.success("碳足迹数据已成功保存");
} catch (error) {
  console.error("请求失败:", error);

  if (error.response && error.response.status === 401) {
    ElMessage.error("登录超时，请重新登录");
    localStorage.removeItem("token");
    delete api.defaults.headers.common["Authorization"];
    router.push("/MyLogin");
  } else {
    ElMessage.error("无法保存计算结果，请稍后再试");
  }
}
}
</script>

<style lang="scss" scoped>
/* 全局样式 */
* {
  margin: 0;
  padding: 0;
  box-sizing: border-box;
}

.bg {
  width: 100%;
  background: url(@/assets/person1.png);
  min-height: 100vh;
  padding: 20px;
}

.container {
  max-width: 1000px;
  margin: 0 auto;
  background-color: #fff;
  border-radius: 10px;
  box-shadow: 0 20px 25px -5px rgb(0 0 0 / 0.5),
    0 8px 10px -6px rgb(0 0 0 / 0.5);
  overflow: hidden;
}

header {
  background-color: #2ecc71;
  color: white;
  padding: 30px 20px;
  text-align: center;
}

header h1 {
  font-size: 2.5rem;
  margin-bottom: 10px;
  font-weight: 700;
}

.subtitle {
  font-size: 1.2rem;
  opacity: 0.9;
}

/* 计算器容器样式 */
.calculator-container {
  padding: 30px;
  background-color: #f0f8ff;
  border: 2px solid #add8e6;
  border-radius: 15px;
  margin: 20px;
}

.tabs {
  display: flex;
  margin-bottom: 20px;
  margin-left: 65px;
  border-bottom: none;
  justify-self: center;
}

.tab-btn {
  padding: 12px 20px;
  background-color: #e0f7fa;
  border: 1px solid #b2ebf2;
  cursor: pointer;
  font-size: 1rem;
  font-weight: 500;
  color: #00838f;
  margin-right: 50px;
  border-radius: 5px;
  transition: all 0.3s;
}

.tab-btn:hover {
  background-color: #b2ebf2;
  color: #006064;
}

.tab-btn.active {
  background-color: #00838f;
  color: white;
  border: 1px solid #00838f;
}

.tab-content {
  display: none;
  padding: 20px 0;
}

.tab-content.active {
  display: block;
}

.tab-content h2 {
  margin-bottom: 20px;
  color: #2ecc71;
  font-size: 1.5rem;
  text-align: center; // 水平居中
}

.input-group {
  margin-bottom: 15px;
  // margin-left: -55px;
}

.input-group label {
  display: block;
  margin-bottom: 8px;
  font-weight: 500;
  color: #444;
}

.input-group input,
.input-group select {
  width: 100%;
  padding: 15px;
  border: 1px solid #90caf9;
  border-radius: 8px;
  font-size: 1rem;
  margin-bottom: 15px;
  background-color: #fafafa;
  transition: all 0.3s;
  -webkit-appearance: none; /* 移除默认样式 */
  appearance: none;
  background-image: url('data:image/svg+xml;utf8,<svg fill="%231976d2" height="24" viewBox="0 0 24 24" width="24" xmlns="http://www.w3.org/2000/svg"><path d="M7 10l5 5 5-5z"/></svg>');
  background-repeat: no-repeat;
  background-position: right 10px center;
  background-size: 20px;
  padding-right: 35px;
  border-radius: 12px; /* 新增圆角 */
  box-shadow: 0 2px 4px rgba(0, 0, 0, 0.05); /* 新增阴影效果 */
  transition: all 0.3s ease; /* 平滑过渡效果 */
}

.input-group select {
  background-color: #fafafa;
  border: 1px solid #90caf9;
  color: #333;
  font-size: 1rem;
}

.input-group select:focus {
  outline: none;
  border-color: #1565c0;
  background-color: white;
}

.input-group select:hover {
  border-color: #1565c0;
  background-color: #f5fbff;
}

.input-group .option {
  padding: 10px;
  background-color: #fff;
  transition: background-color 0.3s;
}

.input-group .option:hover {
  background-color: #e3f2fd;
}

.checkbox-group {
  margin: 20px 0;
}

.checkbox-group label {
  display: flex;
  align-items: center;
  cursor: pointer;
  color: #444;
}

.checkbox-group input {
  margin-right: 10px;
}

.calculate-btn {
  width: auto;
  padding: 15px;
  background-color: #2ecc71;
  color: white;
  border: none;
  border-radius: 5px;
  font-size: 1.1rem;
  font-weight: 500;
  cursor: pointer;
  margin-top: 20px;
  transition: background-color 0.3s;
  
}

.calculate-btn:hover {
  background-color: #27ae60;
}

/* 结果容器样式 */
.result-container {
  padding: 30px;
  background-color: #f9f9f9;
  border-radius: 10px;
  margin-top: 20px;
}

.result-container h2 {
  color: #2ecc71;
  text-align: center;
  margin-bottom: 10px;
  font-size: 1.8rem;
}

.result-container p {
  color: #2ecc71;
  text-align: center;
  margin-bottom: 30px;
  // font-size: 1.8rem;
}

.carbon-circle {
  display: flex;
  justify-content: center;
  margin-bottom: 30px;
}

.circle {
  width: 200px;
  height: 200px;
  border-radius: 50%;
  background-color: #2ecc71;
  display: flex;
  flex-direction: column;
  justify-content: center;
  align-items: center;
  color: white;
  box-shadow: 0 0 20px rgba(46, 204, 113, 0.3);
}

.circle #carbon-value {
  font-size: 2.5rem;
  font-weight: 700;
}

.circle .unit {
  font-size: 1rem;
  opacity: 0.9;
}

.breakdown {
  margin-bottom: 30px;
}

.breakdown h3 {
  margin-bottom: 15px;
  color: #444;
}

.breakdown-item {
  display: flex;
  align-items: center;
  margin-bottom: 15px;
}

.category {
  width: 80px;
  font-weight: 500;
  color: #555;
}

.progress-container {
  flex: 1;
  height: 20px;
  background-color: #eee;
  border-radius: 10px;
  margin: 0 15px;
  overflow: hidden;
}

.progress-bar {
  height: 100%;
  background-color: #2ecc71;
  border-radius: 10px;
  transition: width 0.5s;
}

.value {
  width: 80px;
  text-align: right;
  font-weight: 500;
  color: #555;
}

.comparison {
  margin-bottom: 30px;
}

.comparison h3 {
  margin-bottom: 15px;
  color: #444;
}

.comparison-items {
  display: flex;
  justify-content: space-between;
  margin-top: 20px;
}

.comparison-item {
  flex: 1;
  text-align: center;
  padding: 15px;
  background-color: white;
  border-radius: 8px;
  margin: 0 10px;
  box-shadow: 0 0 10px rgba(0, 0, 0, 0.05);
}

.comparison-item .icon {
  font-size: 2rem;
  margin-bottom: 10px;
}

.comparison-item .text {
  font-size: 1.2rem;
  font-weight: 700;
  color: #2ecc71;
  margin-bottom: 5px;
}

.comparison-item .description {
  font-size: 0.9rem;
  color: #777;
}

.tips {
  margin-bottom: 30px;
}

.tips h3 {
  margin-bottom: 15px;
  color: #444;
}

.tips-container {
  display: grid;
  grid-template-columns: repeat(auto-fit, minmax(200px, 1fr));
  gap: 15px;
  margin-top: 20px;
}

.tip-card {
  background-color: white;
  border-radius: 8px;
  padding: 15px;
  box-shadow: 0 0 10px rgba(0, 0, 0, 0.05);
  display: flex;
  align-items: flex-start;
}

.tip-icon {
  font-size: 1.5rem;
  margin-right: 15px;
  color: #2ecc71;
}

.tip-content h4 {
  margin-bottom: 5px;
  color: #444;
}

.tip-content p {
  font-size: 0.9rem;
  color: #666;
}

.offset-options {
  margin-bottom: 20px;
}

.offset-options h3 {
  margin-bottom: 15px;
  color: #444;
}

.offset-options p {
  margin-bottom: 15px;
  color: #666;
}

.offset-buttons {
  display: flex;
  gap: 15px;
  justify-content: center;
}

.offset-btn {
  flex: 1;
  padding: 12px;
  background-color: white;
  border: 1px solid #2ecc71;
  color: #2ecc71;
  border-radius: 5px;
  font-weight: 500;
  cursor: pointer;
  transition: all 0.3s;
  margin-right: 20px;
}

.offset-btn:hover {
  background-color: #2ecc71;
  color: white;
}

.hidden {
  display: none;
}

.input-group input,
  .input-group select {
    width: 100%;
    appearance: none;
    background-image: url('data:image/svg+xml;utf8,<svg fill="%231976d2" height="24" viewBox="0 0 24 24" width="24" xmlns="http://www.w3.org/2000/svg"><path d="M7 10l5 5 5-5z"/></svg>');
    background-repeat: no-repeat;
    background-position: right 10px center;
    background-size: 20px;
    padding-right: 35px;
    border-radius: 12px; /* 新增圆角 */
    box-shadow: 0 2px 4px rgba(0, 0, 0, 0.05); /* 新增阴影效果 */
    transition: all 0.3s ease; /* 平滑过渡效果 */
  }

  .input-group input:hover,
  .input-group select:hover {
    border-color: #1565c0; /* 更深的蓝色 */
    box-shadow: 0 4px 8px rgba(21, 101, 192, 0.2); /* 增强悬停阴影 */
  }

  .input-group input:focus,
  .input-group select:focus {
    outline: none;
    border-color: #1565c0; /* 统一焦点边框颜色 */
    box-shadow: 0 0 0 3px rgba(21, 101, 192, 0.2); /* 聚焦时的发光效果 */
  }

  .button-container {
    display: flex;
    justify-content: center;
    align-items: center;
    // background-color: #fa2424;
  }

/* 响应式设计 */
@media (max-width: 768px) {
  .comparison-items {
    flex-direction: column;
  }

  .comparison-item {
    margin: 10px 0;
  }

  .tips-container {
    grid-template-columns: 1fr;
  }

  .offset-buttons {
    flex-direction: column;
  }

  .tabs {
    margin-left: 0;
    flex-wrap: wrap;
    justify-content: center;
  }

  .tab-btn {
    margin: 5px;
  }

  .input-group {
    margin-left: 0;
  }


}
</style>