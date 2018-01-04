import axios from 'axios'
import apiConfig from './config/api'

let req = {};

function send(key, options) {
    return new Promise((resolve, reject) => {
        let config = apiConfig[key];
        options = options || {};
        let url = config.url;
        config.method = config.method.toLocaleLowerCase();

        // 替换 URL-PATH 调用参数
        if (options.params) {
            let params = options.params;
            for (let key in params) {
                if (params.hasOwnProperty(key)) {
                    url = url.replace(':' + key, params[key]);
                }
            }
        }

        if (config.method === 'post') {
            options.body = options.body || {};
        }

        // URL-QUERY 请求参数
        if (config.method === 'get') {
            if (options.query) {
                let query = options.query;
                let queryArray = [];
                for (let key in query) {
                    if (query.hasOwnProperty(key)) {
                        queryArray.push(key + '=' + query[key]);
                    }
                }
                if (queryArray.length > 0 ) {
                    url = url + '?' + queryArray.join('&');
                }
            }
        }

        let axiosConfig = {
            method: config.method,
            url: url,
            headers: {},
        };

        // 获取客户端
        let client = options.client;

        if (typeof window === 'undefined' && !client) {
            throw new Error(key + 'client 不能为空');
        }

        if (client && client.headers) {
            // 同步设置浏览器代理
            if (client.headers['user-agent']) {
                axiosConfig.headers['User-Agent'] = client.headers['user-agent'];
            }
            // 设置 cookie
            if (client.headers['cookie']) {
                axiosConfig.headers['Cookie'] = client.headers['cookie'];
            }
        }

        if (config.method === 'post') {
            axiosConfig.data = options.body;
        } else if (config.method === 'get') {
            axiosConfig.params = options.query;
        }

        let startTime = new Date().getTime();
        axios(axiosConfig)
            .then(function (response) {
                console.log({
                    url: url,
                    time: (new Date().getTime() - startTime) + 'ms',
                });
                return resolve(response.data)
            })
            .catch(function (error) {
                return reject(error)
            })
    })
}

for (let key in apiConfig) {
    // 封装成请求对象
    // 返回 key: function
    if (apiConfig.hasOwnProperty(key)) {
        req[key] = (options) => {
            return send(key, options)
        }
    }
}

export default req
