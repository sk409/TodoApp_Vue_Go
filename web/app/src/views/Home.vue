<template>
  <div id="home">
    <el-container>
      <el-aside class="el-aside">
        <div id="logo">
          <h1>TodoApp</h1>
        </div>
        <el-menu background-color="#545c64" text-color="#fff" active-text-color="#ffd04b">
          <el-menu-item index="1">一覧</el-menu-item>
          <el-menu-item index="2">
            <router-link :to="{name: 'add_todo'}">追加</router-link>
          </el-menu-item>
        </el-menu>
      </el-aside>
      <el-container>
        <el-header>
          <el-menu mode="horizontal">
            <el-menu-item>全{{taskCount}}件</el-menu-item>
            <el-menu-item>
              <i class="el-icon-search"></i>
              <el-input v-model="searchKeyword" />
            </el-menu-item>
          </el-menu>
        </el-header>
        <el-main></el-main>
      </el-container>
    </el-container>
  </div>
</template>

<script>
export default {
  name: "home",
  data() {
    return {
      taskCount: 1,
      searchKeyword: ""
    };
  }
};
</script>

<style scoped>
#home {
  height: 100%;
}

#logo {
  height: 50px;
  font-size: 1.5rem;
  background: rgb(69, 75, 81);
  color: white;
}

.el-container {
  height: 100%;
}

.el-aside {
  height: 100%;
  background: #545c64;
}
/* 69, 75, 81 */
</style>