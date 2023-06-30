import { http } from "@/utils/http";

type Result = {
  success: boolean;
  data: Array<any>;
};

export const getAsyncRoutes = () => {
  return new Promise((resolve, reject)=>{
    resolve({
      success: true,
      data: []
    })
  });
};
