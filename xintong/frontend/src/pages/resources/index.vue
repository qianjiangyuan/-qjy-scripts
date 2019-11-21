<template>
    <d2-container>
    </d2-container>
</template>

<script>
    export default {
        name: 'index',
        data: function () {
            return {
            }
        },
        created: function () {
            window.open('http://myjupyter.com:31439/dashboard/db/kubernetes-cluster-monitoring-via-prometheus?refresh=10s&orgId=1', 'about:blank')
        }
    }
</script>
<style scoped>
</style>
