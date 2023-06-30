import { http } from "@/utils/http";

export const AddDeployment = (params) => {
  return http.post("/deployment",params)
}

export const DeleteDeployment = (params) => {
  return http.delete("/deployment",params)
}
export const DeploymentList = (params) => {
  return http.get("/deployment",params)
}
