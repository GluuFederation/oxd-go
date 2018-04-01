angular.module('demo').controller('uma', ['$scope', '$http', umaController]);


function umaController($scope, $http) {
    var vm = this;
    vm.settings = [];

    vm.getSettings = function() {
        $http.get('/settings').then(function (value) {
            vm.settings = value.data;
        }, function (reason) {
            alert(reason);
        });
    }

    vm.getSettings();
}