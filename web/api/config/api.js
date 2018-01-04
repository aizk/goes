let config = {
    apiURL: '/api',
    backApiURL: 'http://127.0.0.1:8023/api'
};

let url = config.apiURL;

const api = {
    getCategories: {
        url: url + '/categories',
        method: 'GET',
        desc: '获取分类列表',
    }
};

export default api;
