<template>
    <div>
        <h3>node</h3>
        <el-table :data="tableData" style="width: 100%">
            <el-table-column prop="metadata.name" label="名称"/>
            <el-table-column prop="status.conditions[3].type" label="状态"/>
            <el-table-column prop="status.addresses[0].address" label="IP地址"/>
            <el-table-column prop="status.nodeInfo.kubeletVersion" label="kubeletVersion"/>
            <el-table-column prop="status.nodeInfo.osImage" label="osImage"/>
            <el-table-column label="创建时间">
                <template #default="scope">
                    {{getRelativeTime(scope.row.metadata.creationTimestamp)}}
                </template>
            </el-table-column>
        </el-table>
    </div>
</template>
<script setup>
import {getRelativeTime}  from "@/utils/time"
import {reactive, toRefs} from "vue";
import {NodeList}         from "@/api/k8s_base"
const data = reactive({
    tableData: []
})
const {tableData} = toRefs(data)
const getNodeList = async () => {
    const res = await NodeList()
    data.tableData = res.message.items
}
getNodeList()
</script>
