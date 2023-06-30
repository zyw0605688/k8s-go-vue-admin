import { http } from "@/utils/http";

export const DeploymentList = (params) => {
  return http.get("/deployment",params)
}
