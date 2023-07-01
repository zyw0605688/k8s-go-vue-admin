import {http} from "@/utils/http";

export const NodeList = () => {
  return http.get("/node")
}

export const PodList = (params) => {
  return http.get("/pod", params)
}



