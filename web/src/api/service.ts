import { http } from "@/utils/http";

export const AddService = (params) => {
  return http.post("/service", params)
}

export const DeleteService = (params) => {
  return http.delete("/service", params)
}

export const ServiceList = (params) => {
  return http.get("/service", params)
}
