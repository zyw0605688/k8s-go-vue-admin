import {http} from "@/utils/http";

export const NodeList = () => {
  return http.get("/node")
}

export const NamespaceList = () => {
  return http.get("/namespace")
}

export const PodList = (params) => {
  return http.get("/pod", params)
}

export const IngressList = (params) => {
  return http.get("/ingress", params)
}

export const ServiceList = (params) => {
  return http.get("/service", params)
}
