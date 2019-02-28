
const routes = [
  {
    path: '/',
    component: () => import('layouts/home.vue'),
    children: [
      { path: '', component: () => import('pages/home.vue') }
    ]
  }
]

if (process.env.MODE !== 'ssr') {
  routes.push({
    path: '*',
    component: () => import('pages/404.vue')
  })
}

export default routes
