import {http} from "@/utils/http";

export const NamespaceList = () => {
  return http.get("/namespace")
}

export const AddNamespace = (params) => {
  return http.post("/namespace",params)
}

export const DeleteNamespace = (params) => {
  return http.delete("/namespace",params)
}
