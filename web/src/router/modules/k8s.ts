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
      path: "/namespace",
      name: "namespace",
      component: () => import("@/views/k8s/namespace/index.vue"),
      meta: {
        title: "命名空间"
      }
    },
    {
      path: "/deployment",
      name: "deployment",
      component: () => import("@/views/k8s/deployment/index.vue"),
      meta: {
        title: "应用部署"
      }
    },
    {
      path: "/pod",
      name: "pod",
      component: () => import("@/views/k8s/pod/index.vue"),
      meta: {
        title: "容器组"
      }
    },
    {
      path: "/service.ts",
      name: "service",
      component: () => import("@/views/k8s/service/index.vue"),
      meta: {
        title: "服务"
      }
    },
    {
      path: "/ingress",
      name: "ingress",
      component: () => import("@/views/k8s/ingress/index.vue"),
      meta: {
        title: "应用路由"
      }
    },

  ]
};
