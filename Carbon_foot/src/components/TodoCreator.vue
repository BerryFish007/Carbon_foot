<template>
  <div class="input-wrap" :class="{ 'input-err ': todoState.invalid }">
      <input type="text" v-model="todoState.todo">
      <!-- 传给父组件TodoView -->
      <button @click="createTodo()">Create</button>
  </div>
  <p v-show="todoState.invalid" class="err-msg">{{ todoState.errMsg }}</p>
</template>

<script setup>
import { reactive } from "vue";
import api from '@/http'; // 引入全局配置的 axios 实例
const emit = defineEmits(["create-todo"]);
// 使用ref就是.value  使用reactive就是对象.属性
const todoState = reactive({
todo: "",
invalid: null,
errMsg: "",
});

const createTodo = async () => {
todoState.invalid = null;
if (todoState.todo !== "") {
  try {
    const response = await api.post('/plan/add', {
      content: todoState.todo
    });
    if (response.data.Code === 1) {
      emit("create-todo", {
        id: response.data.Data,
        content: todoState.todo,
        isCompleted: false
      });
      todoState.todo = "";
    } else {
      todoState.invalid = true;
      todoState.errMsg = response.data.Msg;
    }
  } catch (error) {
    todoState.invalid = true;
    todoState.errMsg = "请求出错，请稍后再试";
    console.error('请求出错:', error);
  }
} else {
  todoState.invalid = true;
  todoState.errMsg = "Todo value cannot be empty";
}
};
</script>

<style lang="scss" scoped>
button {
  cursor: pointer;
}
.input-wrap {
  display: flex;
  transition: 250ms ease;
  border: 2px solid #41b080;
  margin-left: -2vh;
  margin-right: 3vh;
  border-radius: 10px;
  &.input-err {
    border-color: red;
    box-shadow: 2px 2px 2px rgb(241, 108, 108);
    border-radius: 10px;
  }

  &:focus-within {
    box-shadow: 0 -4px 6px -1px rgb(0 0 0 / 0.1),
      0 -2px 4px -2px rgb(0 0 0 / 0.1);
  }
  input {
    width: 100%;
    padding: 8px 6px;
    border: none;
    border-top-left-radius: 10px;
    border-bottom-left-radius: 10px;
    &:focus {
      outline: none;
    }
  }
  button {
    padding: 8px 16px;
    border: none;
    border-top-right-radius: 10px;
    border-bottom-right-radius: 10px;
  }
}

.err-msg {
  margin-top: 6px;
  font-size: 12px;
  text-align: center;
  color: red;
}
</style>