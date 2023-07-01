<template>
    <div>
        <Namespace @namespaceChanged="namespaceChanged"></Namespace>
        <el-table :data="tableData" style="width: 100%;margin-top: 12px">
            <el-table-column prop="metadata.name" label="名称" width="220"/>
            <el-table-column prop="status.podIP" label="IP地址" width="120"/>
            <el-table-column prop="status.phase" label="phase" width="120"/>
            <el-table-column label="所在节点" width="260">
                <template #default="scope">
                    {{scope.row.spec.nodeName}}/{{scope.row.status.hostIP}}
                </template>
            </el-table-column>
            <el-table-column label="容器">
                <template #default="scope">
                    <div v-for="(item,index) in scope.row.status.containerStatuses">
                        <div>容器名：{{item.name}}</div>
                        <div>容器:{{item.imageID}}</div>
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
import {getRelativeTime} from "@/utils/time"
import {onMounted, reactive, toRefs} from "vue";
import {PodList}      from "@/api/k8s_base";
import {NamespaceList} from "@/api/namespace";
import Namespace from "@/components/namespace/index.vue";

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
    const res = await PodList({namespace:data.namespace})
    data.tableData = res.message.items
}

onMounted(async () => {
    await getList()
})
</script>
