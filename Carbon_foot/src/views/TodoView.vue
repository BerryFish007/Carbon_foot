<script setup>
import { ref, watch, computed, onMounted } from "vue";
import { Icon } from "@iconify/vue";
import TodoCreator from '../components/TodoCreator.vue';
import TodoItem from '../components/TodoItem.vue';
import { useRouter } from "vue-router";
import api from '@/http'; // 引入全局配置的 axios 实例

const todoList = ref([]);
const router = useRouter();

const back = () => {
  router.push('/');
};

watch(
  todoList,
  () => {
    setTodoListLocalStorage();
  },
  {
    deep: true,
  }
);

const todoCompleted = computed(() => {
  return todoList.value.every((todo) => todo.isCompleted);
});

const fetchTodoList = async () => {
  try {
    const response = await api.get('/plan/list');
    if (response.data.Code === 1) {
      todoList.value = response.data.Data.map(item => ({
        id: item.id,
        todo: item.content,
        isCompleted: item.finished === 1,
        isEditing: false
      }));
    } else {
      console.error('获取计划列表失败:', response.data.Msg);
    }
  } catch (error) {
    console.error('请求出错:', error);
  }
};

const setTodoListLocalStorage = () => {
  localStorage.setItem("todoList", JSON.stringify(todoList.value));
};

const createTodo = (newTodo) => {
  todoList.value.push({
    id: newTodo.id,
    todo: newTodo.content,
    isCompleted: newTodo.isCompleted,
    isEditing: false
  });
};

const toggleTodoComplete = (todoPos) => {
  todoList.value[todoPos].isCompleted = !todoList.value[todoPos].isCompleted;
};

const toggleEditTodo = (todoPos) => {
  todoList.value[todoPos].isEditing = !todoList.value[todoPos].isEditing;
};

const updateTodo = (todoVal, todoPos) => {
  todoList.value[todoPos].todo = todoVal;
};

const deleteTodo = (todoId) => {
  todoList.value = todoList.value.filter((todo) => todo.id !== todoId);
};

onMounted(() => {
  fetchTodoList();
});
</script>

<template>
  <div class="todo">
    <button @click="back"></button>
    <main>
      <h1>减碳计划表</h1>
      <TodoCreator @create-todo="createTodo"/>
      <ul class="todo-list" v-if="todoList.length > 0">
        <TodoItem 
          v-for="(todo, index) in todoList" 
          :todo="todo" 
          :index="index"
          @toggle-complete="toggleTodoComplete"
          @edit-todo="toggleEditTodo"
          @update-todo="updateTodo"
          @delete-todo="deleteTodo"
        />
      </ul>
      <p class="todos-msg" v-else>
        <Icon icon="noto-v1:sad-but-relieved-face" class="icon" width="22" />
        <span>你还没有减碳计划！添加一个！</span>
      </p>
      <p v-if="todoCompleted && todoList.length > 0" class="todos-msg">
        <Icon icon="noto-v1:party-popper" />
        <span>你已经完成了你的减碳计划!</span>
      </p>
    </main>
    <div class="place"></div>
  </div>

</template>

<style lang="scss" scoped>
.place {
  height: 80vh;
}
.todo {
  width: 100%;
  background: url('@/assets/bg1.jpg') no-repeat;
  background-size: cover;
  background-attachment: fixed;
  position: relative;
   button:hover {
      background-color: rgb(238, 232, 170, 0.7);
  }
}
main {
  display: flex;
  flex-direction: column;
  max-width: 500px;
  width: 100%;
  margin: 0 auto;
  padding: 40px 16px;

  h1 {
    margin-bottom: 16px;
    text-align: center;
    font-weight: 700;
    font-size: 50px;
    color: #020304;
  }

    .todo-list {
    display: flex;
    flex-direction: column;
    list-style: none;
    margin-top: 24px;
    gap: 20px;
    font-size: 25px;
    color: #020304;
  }
  .todos-msg {
    display: flex;
    align-items: center;
    justify-content: center;
    gap: 8px;
    margin-top: 24px;
    font-size: 23px;
    color: #020304;
  }
}
</style>
