import {http} from "@/utils/http";

export const AddIngress = (params) => {
  return http.post("/ingress", params)
}

export const DeleteIngress = (params) => {
  return http.delete("/ingress", params)
}
export const IngressList = (params) => {
  return http.get("/ingress", params)
}
