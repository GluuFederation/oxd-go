angular.module('demo').controller('settings', ['$scope', '$http', settingsController]);


function settingsController($scope, $http) {
    var vm = this;
    vm.settings = [];



    vm.getSettings = function() {
        $http.get('/settings').then(function (value) {
            vm.settings = value.data;
        }, function (reason) {
            alert(reason);
        });
    }

    vm.setSettings = function () {
        $http.post('/settings',vm.settings).then(function (value) {
            vm.getSettings();
        }, function (reason) {
            alert(reason);
        });
    }

    vm.getClientToken = function () {
        $http.get('/getAccessToken',vm.settings).then(function (value) {
            vm.getSettings();
        }, function (reason) {
            alert(reason);
        });
    }

    vm.getSettings();
}