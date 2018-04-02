angular.module('demo').controller('settings', ['$scope', '$http', settingsController]);


function settingsController($scope, $http) {
    var vm = this;

    vm.getSettings = function() {
        $http.get('/settings').then(function (value) {
            vm.settings = value.data;
        }, function (reason) {
            alert(reason);
        });
    }

    vm.getSettings();
}