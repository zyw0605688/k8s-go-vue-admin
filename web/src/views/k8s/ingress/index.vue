<template>
    <div>
        <Namespace @namespaceChanged="namespaceChanged"></Namespace>
        <el-table :data="tableData" style="width: 100%;margin-top: 12px">
            <el-table-column prop="metadata.name" label="名称" width="220"/>
            <el-table-column label="域名">
                <template #default="scope">
                    <div v-for="(item,index) in scope.row.spec.rules">
                        <div>{{item.host}}</div>
                    </div>
                </template>
            </el-table-column>
            <el-table-column label="路径">
                <template #default="scope">
                    <div v-for="(item,index) in scope.row.spec.rules">
                        <div v-for="j in item.http.paths">{{j.path}}</div>
                    </div>
                </template>
            </el-table-column>
            <el-table-column label="类型">
                <template #default="scope">
                    <div v-for="(item,index) in scope.row.spec.rules">
                        <div v-for="j in item.http.paths">{{j.pathType}}</div>
                    </div>
                </template>
            </el-table-column>
            <el-table-column label="创建时间" width="120">
                <template #default="scope">
                    {{getRelativeTime(scope.row.metadata.creationTimestamp)}}
                </template>
            </el-table-column>
        </el-table>
    </div>
</template>
<script setup>
import {getRelativeTime}             from "@/utils/time"
import {onMounted, reactive, toRefs} from "vue";
import {IngressList}  from "@/api/k8s_base";
import Namespace                     from "@/components/namespace/index.vue";

const data = reactive({
    tableData: [],
    namespace: "",
})
const {tableData, namespace} = toRefs(data)

const namespaceChanged = (val) => {
    data.namespace = val;
    getList()
}
const getList = async () => {
    const res = await IngressList({namespace: data.namespace})
    data.tableData = res.message.items
}

onMounted(async () => {
    await getList()
})
</script>

