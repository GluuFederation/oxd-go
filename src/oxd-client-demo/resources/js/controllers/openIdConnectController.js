angular.module('demo').controller('openIdConnect', ['$scope', '$http', openIdConnectController]);


function openIdConnectController($scope, $http) {
    var vm = this;
    vm.openIdConnect = [];
    vm.claims = ""



    vm.getSettings = function() {
        $http.get('/settings').then(function (value) {
            vm.openIdConnect = value.data.OpenIdConnect;
            vm.claims = JSON.stringify(vm.openIdConnect.Claims, null, 2);
        }, function (reason) {
            alert(reason);
        });
    }

    vm.getAuthorizationUrl = function() {
        $http.get('/authorizationUrl').then(function (value) {
            vm.getSettings();
        }, function (reason) {
            alert(reason);
        });
    }

    vm.getTokens = function() {
        $http.get('/token').then(function (value) {
            vm.getSettings();
        }, function (reason) {
            alert(reason);
        });
    }

    vm.refreshTokens = function() {
        $http.get('/refresh').then(function (value) {
            vm.getSettings();
        }, function (reason) {
            alert(reason);
        });
    }

    vm.getUserInfo = function() {
        $http.get('/userInfo').then(function (value) {
            vm.getSettings();

        }, function (reason) {
            alert(reason);
        });
    }

    vm.getLogoutUrl = function() {
        $http.get('/logoutUrl').then(function (value) {
            vm.getSettings();
        }, function (reason) {
            alert(reason);
        });
    }

    vm.getSettings();
}