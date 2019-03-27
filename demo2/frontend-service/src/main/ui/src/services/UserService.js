import Vue from 'vue'
import VueResource from 'vue-resource'

Vue.use(VueResource)

const getUserList = function(args, successCallback, errorCallback) {
    var memthodName = 'getUserList'
    Vue.http.get('http://127.0.0.1:8084/api/users').then(response => {
        if (response.status !== 200) {
            Vue.console.log("status: " + response.status + ", body: " + response.body)
            errorCallback(memthodName + ' request failed')
        } else {
            successCallback(response.body)
        }
    }, response => {
        Vue.console.log("status: " + response.status + ", body: " + response.body)
        errorCallback(memthodName + ' request failed')
    })
}

const getUserById = function (args, successCallback, errorCallback) {
    var memthodName = 'getUserById'
    Vue.http.get('http://127.0.0.1:8084/api/users/' + args.userId).then(response => {
        if (response.status !== 200) {
            Vue.console.log("status: " + response.status + ", body: " + response.body)
            errorCallback(memthodName + ' request failed')
        } else {
            successCallback(response.body)
        }
    }, response => {
        Vue.console.log("status: " + response.status + ", body: " + response.body)
        errorCallback(memthodName + ' request failed')
    })
}

const editUser = function (args, successCallback, errorCallback) {
    var memthodName = 'getUserById'
    Vue.http.put('http://127.0.0.1:8084/api/users/' + args.userId, args.user).then(response => {
        if (response.status !== 200) {
            Vue.console.log("status: " + response.status + ", body: " + response.body)
            errorCallback(memthodName + ' request failed')
        } else {
            successCallback(response.body)
        }
    }, response => {
        Vue.console.log("status: " + response.status + ", body: " + response.body)
        errorCallback(memthodName + ' request failed')
    })
}

const addUser = function (args, successCallback, errorCallback) {
    var memthodName = 'getUserById'
    Vue.http.post('http://127.0.0.1:8084/api/users', args.user).then(response => {
        if (response.status !== 200) {
            Vue.console.log("status: " + response.status + ", body: " + response.body)
            errorCallback(memthodName + ' request failed')
        } else {
            successCallback(response.body)
        }
    }, response => {
        Vue.console.log("status: " + response.status + ", body: " + response.body)
        errorCallback(memthodName + ' request failed')
    })
}

const deleteUser = function (args, successCallback, errorCallback) {
    var memthodName = 'deleteUser'
    Vue.http.delete('http://127.0.0.1:8084/api/users/' + args.userId).then(response => {
        if (response.status !== 200) {
            Vue.console.log("status: " + response.status + ", body: " + response.body)
            errorCallback(memthodName + ' request failed')
        } else {
            successCallback(response.body)
        }
    }, response => {
        Vue.console.log("status: " + response.status + ", body: " + response.body)
        errorCallback(memthodName + ' request failed')
    })
}

export default {
    getUserList,
    getUserById,
    editUser,
    addUser,
    deleteUser
}