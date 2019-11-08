<template>
  <el-container class="task-form-container">
    <el-header>
      <h1>新規タスク追加</h1>
    </el-header>
    <el-main>
      <el-card class="todo-form-card">
        <el-form ref="taskForm" :model="todo" :rules="rules">
          <el-form-item label="タイトル" prop="title">
            <el-input v-model="todo.title" />
          </el-form-item>
          <el-form-item label="内容" prop="content">
            <el-input
              class="todo-content"
              type="textarea"
              rows="10"
              size="large"
              v-model="todo.content"
            />
          </el-form-item>
          <el-form-item label="期限" prop="limit">
            <el-date-picker type="date" v-model="todo.limit" />
          </el-form-item>
          <el-form-item>
            <el-button type="primary" @click="submit">登録</el-button>
          </el-form-item>
        </el-form>
      </el-card>
    </el-main>
  </el-container>
</template>

<script>
import Todo from "@/models/todo.js";
export default {
  name: "AddTodo",
  data() {
    return {
      rules: {
        title: [
          {
            required: true,
            message: "タイトルを入力してください",
            trigger: "change"
          }
        ],
        content: [
          {
            required: true,
            message: "内容を入力してください",
            trigger: "change"
          }
        ],
        limit: [
          {
            required: true,
            message: "期限を入力してください",
            trigger: "change"
          }
        ]
      },
      todo: new Todo()
    };
  },
  methods: {
    submit() {
      this.$refs.taskForm.validate(valid => {
        if (valid) {
          this.todo.store(response => {
            if (response.status === 200) {
              this.$notify({
                message: "タスクを登録しました",
                type: "success",
                duration: 1500
              });
            }
          });
        }
      });
    }
  }
};
</script>

<style scoped>
.task-form-container {
  margin-top: 1rem;
}

.todo-form-card {
  margin: 0 auto;
  width: 80%;
}
</style>