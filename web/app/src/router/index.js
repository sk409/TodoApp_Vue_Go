import Vue from 'vue'
import VueRouter from 'vue-router'
import Todos from '../views/Todos.vue'
import AddTodo from "@/views/AddTodo.vue"

Vue.use(VueRouter)

const routes = [{
    path: '/',
    name: 'todos',
    component: Todos
  },
  {
    path: "/add_todo/",
    name: "add_todo",
    component: AddTodo
  }
]

const router = new VueRouter({
  mode: 'history',
  base: process.env.BASE_URL,
  routes
})

export default router