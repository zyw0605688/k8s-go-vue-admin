<template>
    <div>
        <h2>命名空间</h2>
        <el-input v-model="namespace" style="width: 200px" clearable></el-input>
        <el-button style="margin-left: 12px" @click="add">新增</el-button>
        <el-table :data="tableData" width="100%" style="margin-top: 12px">
            <el-table-column prop="metadata.name" label="名称"/>
            <el-table-column prop="status.phase" label="状态"/>
            <el-table-column label="创建时间">
                <template #default="scope">
                    {{getRelativeTime(scope.row.metadata.creationTimestamp)}}
                </template>
            </el-table-column>
            <el-table-column label="操作">
                <template #default="scope">
                    <el-button @click="deleteOne(scope.row)">删除</el-button>
                </template>
            </el-table-column>
        </el-table>
    </div>
</template>
<script setup lang="ts" name="namespace">
import {getRelativeTime} from "@/utils/time"
import {onMounted, reactive, toRefs} from "vue";
import {AddNamespace, DeleteNamespace, NamespaceList} from "@/api/namespace";

const data = reactive({
    tableData: [],
    namespace: "",
})
const {tableData, namespace} = toRefs(data)
const getList = async () => {
    const res = await NamespaceList()
    data.tableData = res.message.items
}

const add = async () => {
    await AddNamespace({namespace: data.namespace})
    data.namespace = ""
    await getList()
}

const deleteOne = async (row) => {
    await DeleteNamespace({namespace: row.metadata.name})
    await getList()
}

onMounted(async () => {
    await getList()
})
</script>
