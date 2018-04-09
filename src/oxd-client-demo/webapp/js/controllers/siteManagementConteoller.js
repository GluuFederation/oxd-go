angular.module('demo').controller('siteManagement', ['$scope', '$http', siteManagementController]);


function siteManagementController($scope, $http) {
    var vm = this;
    vm.registerSettings = [];

    vm.registerSite = function() {
        $http.get('/registerSite').then(function (value) {
            vm.getSettings()
        }, function (reason) {
            alert(reason);
        });
    }

    vm.updateSite = function() {
        $http.get('/updateSite',vm.registerSettings).then(function (value) {
            vm.getSettings()
        }, function (reason) {
            alert(reason);
        });
    }

    vm.removeSite = function() {
        $http.get('/removeSite',vm.registerSettings).then(function (value) {
            vm.getSettings()
        }, function (reason) {
            alert(reason);
        });
    }

    vm.getSettings = function() {
        $http.get('/settings').then(function (value) {
            vm.registerSettings = value.data.RegisterSiteParams;
        }, function (reason) {
            alert(reason);
        });
    }

    vm.getSettings();
}