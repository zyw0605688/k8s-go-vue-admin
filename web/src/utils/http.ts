import axios, {AxiosInstance} from "axios"

const config = {
    baseURL: '/api',
    timeout: 1000,
}
class Http {
    service: AxiosInstance
    constructor(config) {
        this.service=axios.create(config)

        this.service.interceptors.request.use(function (config) {
            const token = JSON.parse(localStorage.getItem("token"))
            if (token) {
                config.headers.token = token
            }
            return config;
        }, function (error) {
            return Promise.reject(error);
        });


        this.service.interceptors.response.use(function (response) {
            if (response.data.code === 10401) {
                console.log("需要重新登录")
                localStorage.clear()
                location.href="/login"
            }
            return response.data;
        }, function (error) {
            return Promise.reject(error);
        });

    }


    get(url, params?: object, _object = {}) {
        return this.service.get(url, { params, ..._object });
    }
    post(url, params?: object | string, _object = {}) {
        return this.service.post(url, params, _object);
    }
    put(url, params?: object, _object = {}) {
        return this.service.put(url, params, _object);
    }
    delete(url, params?: any, _object = {}) {
        return this.service.delete(url, { params, ..._object });
    }
}

export const http =  new Http(config);
