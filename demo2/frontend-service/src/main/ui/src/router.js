import Vue from 'vue'
import VueRouter from 'vue-router'

Vue.use(VueRouter);

// 1. 定义 (路由) 组件。
// 可以从其他文件 import 进来
import Help from './components/Help.vue'
import ListUser from './components/ListUser.vue'
import AddUser from './components/AddUser.vue'
import EditUser from './components/EditUser.vue'
import Index from "./components/Index";
import ShowUser from "./components/ShowUser";

// 2. 定义路由
// 每个路由应该映射一个组件。 其中"component" 可以是
// 通过 Vue.extend() 创建的组件构造器，
// 或者，只是一个组件配置对象。
// 我们晚点再讨论嵌套路由。
const routes = [
  { path: '/', component: Index, meta: {title: 'SpringMVC Demo 首页'} },
  { path: '/user/list', component: ListUser, meta: {title: 'SpringMVC 用户管理'} },
  { path: '/user/show/:id', component: ShowUser, meta: {title: 'SpringMVC 用户详情'}},
  { path: '/user/add', component: AddUser, meta: {title: "SpringMVC 添加用户"} },
  { path: '/user/edit/:id', component: EditUser, meta: {title: 'SpringMVC Demo 更新用户'} },
  { path: '/help', component: Help }
];

// 3. 创建 router 实例，然后传 `routes` 配置
// 你还可以传别的配置参数, 不过先这么简单着吧。
const router = new VueRouter({
  routes // (缩写) 相当于 routes: routes
});

const DEFAULT_TITLE = 'SpringMVC Demo';
router.afterEach((to, from) => {
  document.title = to.meta.title || DEFAULT_TITLE;
});

export default router