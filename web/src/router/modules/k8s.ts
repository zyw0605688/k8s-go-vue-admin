const Layout = () => import("@/layout/index.vue");
export default {
  path: "/k8s",
  name: "Home",
  component: Layout,
  redirect: "/node",
  meta: {
    icon: "homeFilled",
    title: "集群",
    rank: 1
  },
  children: [
    {
      path: "/node",
      name: "node",
      component: () => import("@/views/k8s/node/index.vue"),
      meta: {
        title: "节点"
      }
    },
    {
      path: "/deployment",
      name: "deployment",
      component: () => import("@/views/k8s/deployment/index.vue"),
      meta: {
        title: "应用部署"
      }
    }
  ]
};
