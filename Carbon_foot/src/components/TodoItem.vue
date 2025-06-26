<script setup>
import { ref } from "vue";
import { Icon } from "@iconify/vue";
import api from '@/http'; // 引入全局配置的 axios 实例

const props = defineProps({
  todo: {
    type: Object,
    required: true,
  },
  index: {
    type: Number,
    required: true,
  },
});

const emit = defineEmits(["toggle-complete", "edit-todo", "update-todo", "delete-todo"]);
const editedContent = ref(props.todo.todo);

const handleInputChange = (value) => {
  editedContent.value = value;
};

const saveEditedTodo = async () => {
  try {
    const response = await api.post('/plan/edit', {
      id: props.todo.id,
      content: editedContent.value
    });

    if (response.data.Code === 1) {
      emit('update-todo', editedContent.value, props.index);
      emit('edit-todo', props.index); // 关闭编辑状态
    } else {
      console.error('编辑计划失败:', response.data.Msg);
    }
  } catch (error) {
    console.error('请求出错:', error);
  }
};

const toggleComplete = async () => {
  try {
    const response = await api.post('/plan/finished', {
      id: props.todo.id
    });

    if (response.data.Code === 1) {
      emit('toggle-complete', props.index);
    } else {
      console.error('完成计划失败:', response.data.Msg);
    }
  } catch (error) {
    console.error('请求出错:', error);
  }
};

const deleteTodo = async () => {
  try {
    const response = await api.post('/plan/delete', {
      id: props.todo.id
    });

    if (response.data.Code === 1) {
      emit('delete-todo', props.todo.id);
    } else {
      console.error('删除计划失败:', response.data.Msg);
    }
  } catch (error) {
    console.error('请求出错:', error);
  }
};
</script>

<template>
  <li>
    <input
      type="checkbox"
      :checked="todo.isCompleted"
      @input="toggleComplete"
    />
    <div class="todo">
      <input
        v-if="todo.isEditing"
        type="text"
        :value="todo.todo"
        @input="handleInputChange($event.target.value)"
      />
      <span v-else :class="{ 'completed-todo': todo.isCompleted }">{{
        todo.todo
      }}</span>
    </div>
    <div class="todo-actions">
      <Icon
        v-if="todo.isEditing"
        icon="ph:check-circle"
        class="icon"
        color="#41b080"
        width="30"
        @click="saveEditedTodo"
      />
      <Icon
        v-else
        icon="ph:pencil-fill"
        class="icon"
        color="#41b080"
        width="30"
        @click="$emit('edit-todo', index)"
      />
      <Icon
        icon="ph:trash"
        class="icon"
        color="#f95e5e"
        width="30"
        @click="deleteTodo"
      />
    </div>
  </li>
</template>

<style lang="scss" scoped>
li {
  display: flex;
  align-items: center;
  gap: 10px;
  padding: 14px 10px;
  background-color: #f9f8e48f;
  border-radius: 12px;
  box-shadow: 0 20px 25px -5px rgb(0 0 0 / 0.1),
    0 8px 10px -6px rgb(0 0 0 / 0.1);
  margin-left: -5vh;
  &:hover {
    .todo-actions {
      opacity: 1;
    }
  }
  input[type="checkbox"] {
    appearance: none;
    width: 20px;
    height: 20px;
    cursor: pointer;
    background-color: #fff;
    border-radius: 50%;
    box-shadow: 0 4px 6px -1px rgb(0 0 0 / 0.1), 0 2px 4px -2px rgb(0 0 0 / 0.1);
    &:checked {
      background-color: #41b080;
    }
  }
  .todo {
    flex: 1;
    .completed-todo {
      text-decoration: line-through;
    }
    input[type="text"] {
      width: 100%;
      padding: 7px 6px;
      border: 2px solid #41b080;
      outline: none;
      font-size: 17px;
    }
  }
  .todo-actions {
    display: flex;
    gap: 6px;
    opacity: 0;
    transition: 150ms ease-in-out;
    .icon {
      cursor: pointer;
    }
  }
}
</style>