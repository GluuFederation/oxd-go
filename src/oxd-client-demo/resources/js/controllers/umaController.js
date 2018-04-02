angular.module('demo').controller('uma', ['$scope', '$http', umaController]);


function umaController($scope, $http) {
    var vm = this;
    vm.uma = [];
    vm.settings =[];
    vm.introspect =[];

    vm.getSettings = function() {
        $http.get('/settings').then(function (value) {
            vm.settings = value.data;
            vm.uma = vm.settings.Uma;
            vm.introspect = JSON.stringify(vm.uma.IntrospectRpt, null, 2);
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

    vm.umaRsProtect= function () {
        $http.get('/umaRsProtect').then(function (value) {
            vm.getSettings();
        }, function (reason) {
            alert(reason);
        });
    }

    vm.umaCheckAccess= function () {
        $http.get('/umaCheckAccess').then(function (value) {
            alert(value.data);
            vm.getSettings();
        }, function (reason) {
            alert(reason);
        });
    }

    vm.umaGetRpt= function () {
        $http.get('/umaRpt').then(function (value) {
            vm.getSettings();
        }, function (reason) {
            alert(reason);
        });
    }

    vm.umaGetClaimsUrl= function () {
        $http.get('/umaClaimsUrl').then(function (value) {
            vm.getSettings();
        }, function (reason) {
            alert(reason);
        });
    }

    vm.umaIntrospectRpt= function () {
        $http.get('/umaIntrospectRpt').then(function (value) {
            vm.getSettings();
        }, function (reason) {
            alert(reason);
        });
    }

    vm.getSettings();
}