<template>
  <div id="app">
    <el-container>
      <el-aside id="side-bar" class="el-aside">
        <div id="logo">
          <h1>TodoApp</h1>
        </div>
        <el-menu background-color="#545c64" text-color="#fff" active-text-color="#ffd04b">
          <el-menu-item index="1">
            <router-link tag="div" :to="{name: 'todos'}">一覧</router-link>
          </el-menu-item>
          <el-menu-item index="2">
            <router-link tag="div" :to="{name: 'add_todo'}">追加</router-link>
          </el-menu-item>
        </el-menu>
      </el-aside>
      <router-view />
    </el-container>
  </div>
</template>

<style>
* {
  margin: 0;
  padding: 0;
  box-sizing: border-box;
}

#app {
  font-family: "Avenir", Helvetica, Arial, sans-serif;
  -webkit-font-smoothing: antialiased;
  -moz-osx-font-smoothing: grayscale;
  color: #2c3e50;
  height: 100vh;
}

#side-bar {
  text-align: center;
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
</style>
