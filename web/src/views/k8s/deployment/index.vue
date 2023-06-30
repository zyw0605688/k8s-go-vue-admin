<template>
    <div>
        <h3>deployment</h3>
        请选择命名空间：
        <el-select v-model="namespace" @change="getList">
            <el-option v-for="(item,index) in namespaceList" :index="index" :key="item.metadata.name" :value="item.metadata.name"></el-option>
        </el-select>
        <el-table :data="tableData" style="width: 100%">
            <el-table-column prop="metadata.name" label="名称" width="220"/>
            <el-table-column label="所在节点" prop="spec.template.spec.nodeName"></el-table-column>
            <el-table-column label="">
                <template #default="scope">
                    <div v-for="(item,index) in scope.row.spec.template.spec.containers">
                        <div>镜像：{{item.image}}</div>
                        <div>端口:{{item.ports}}</div>
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
import {DeploymentList}              from "@/api/deployment";
import {NamespaceList}               from "@/api/namespace";

const data = reactive({
    tableData: [],
    namespace: "",
    namespaceList: []
})
const {tableData, namespace, namespaceList} = toRefs(data)
const getList = async () => {
    data.tableData = []
    const res = await DeploymentList({namespace:data.namespace})
    data.tableData = res.message.items
}

const getNamespaceList = async () => {
    const res = await NamespaceList()
    data.namespaceList = res.message.items
    data.namespace = data.namespaceList[0].metadata.name
}

onMounted(async () => {
    await getNamespaceList()
    await getList()
})
</script>

