import { http } from "@/utils/http";

export type UserResult = {
  success: boolean;
  data: {
    /** 用户名 */
    username: string;
    /** 当前登陆用户的角色 */
    roles: Array<string>;
    /** `token` */
    accessToken: string;
    /** 用于调用刷新`accessToken`的接口时所需的`token` */
    refreshToken: string;
    /** `accessToken`的过期时间（格式'xxxx/xx/xx xx:xx:xx'） */
    expires: Date;
  };
};

export type RefreshTokenResult = {
  success: boolean;
  data: {
    /** `token` */
    accessToken: string;
    /** 用于调用刷新`accessToken`的接口时所需的`token` */
    refreshToken: string;
    /** `accessToken`的过期时间（格式'xxxx/xx/xx xx:xx:xx'） */
    expires: Date;
  };
};

/** 登录 */
export const getLogin = (data?: object) => {
  return new Promise((resolve, reject)=>{
    resolve({
      success: true,
      data: {
        username: "admin",
        // 一个用户可能有多个角色
        roles: ["admin"],
        accessToken: "eyJhbGciOiJIUzUxMiJ9.admin",
        refreshToken: "eyJhbGciOiJIUzUxMiJ9.adminRefresh",
        expires: "2023/10/30 00:00:00"
      }
    })
  })
};

/** 刷新token */
export const refreshTokenApi = (data?: object) => {
  return new Promise((resolve)=>{
    resolve({
      success: true,
      data: {
        accessToken: "eyJhbGciOiJIUzUxMiJ9.newAdmin",
        refreshToken: "eyJhbGciOiJIUzUxMiJ9.newAdminRefresh",
        expires: "2023/10/30 23:59:59"
      }
    })
  })
};
