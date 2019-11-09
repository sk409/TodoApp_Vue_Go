<template>
  <div id="todos">
    <el-container>
      <el-header>
        <el-menu mode="horizontal">
          <el-menu-item>全{{todoCount}}件</el-menu-item>
          <el-menu-item>
            <i class="el-icon-search"></i>
            <el-input v-model="searchKeyword" />
          </el-menu-item>
        </el-menu>
      </el-header>
      <el-main>
        <el-card v-for="todo in todos" :key="todo.id" class="todo-card">
          <div
            slot="header"
            contenteditable="true"
            @input="updateTitle($event, todo.id)"
          >{{todo.title}}</div>
          <div>
            <div
              class="todo-content"
              contenteditable="true"
              @input="updateContent($event, todo.id)"
            >{{todo.content}}</div>
            <el-divider></el-divider>
            <div class="todo-button">
              <el-button type="primary">DONE</el-button>
            </div>
          </div>
        </el-card>
      </el-main>
    </el-container>
  </div>
</template>

<script>
import Todo from "@/models/todo.js";
export default {
  name: "todos",
  data() {
    return {
      searchKeyword: "",
      todos: []
    };
  },
  computed: {
    todoCount() {
      return this.todos.length;
    }
  },
  created() {
    const that = this;
    Todo.fetch(null, response => {
      response.data.forEach(data => {
        that.todos.push(
          new Todo(data.id, data.title, data.content, data.limit)
        );
      });
    });
  },
  methods: {
    updateTitle(e, id) {
      const todo = this.todos.find(todo => todo.id === id);
      if (!todo) {
        return;
      }
      todo.title = e.target.value;
      todo.update();
    },
    updateContent(e, id) {
      const todo = this.todos.find(todo => todo.id === id);
      if (!todo) {
        return;
      }
      todo.content = e.target.value;
      todo.update();
    }
  }
};
</script>

<style scoped>
#todos {
  width: 100%;
  height: 100%;
}

.el-container {
  height: 100%;
}

.todo-card {
  width: 70%;
  margin: 0 auto 2rem auto;
}

.todo-content {
  margin-bottom: 1.5rem;
}

.todo-button {
  text-align: center;
}
</style>