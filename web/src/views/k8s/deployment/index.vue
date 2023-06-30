<template>
    <div>
        <h3>deployment</h3>
        请选择命名空间：
        <el-select v-model="namespace" @change="getList">
            <el-option v-for="(item,index) in namespaceList" :index="index" :key="item.metadata.name" :value="item.metadata.name"></el-option>
        </el-select>
        <el-button @click="showDialog">新增</el-button>
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
            <el-table-column label="操作">
                <template #default="scope">
                   <el-button @click="deleteOne(scope.row)">删除</el-button>
                </template>
            </el-table-column>
        </el-table>
        <el-dialog v-model="visible">
            <el-form :model="form" label-width="120px">
                <el-form-item label="名称">
                    <el-input v-model="form.pod_name" />
                </el-form-item>
                <el-form-item label="命名空间">
                    <el-input v-model="form.pod_namespace" />
                </el-form-item>
                <el-form-item label="镜像">
                    <el-input v-model="form.pod_image" />
                </el-form-item>
                <el-form-item label="端口号">
                    <el-input v-model.number="form.pod_port" />
                </el-form-item>
                <el-form-item label="副本数">
                    <el-input v-model.number="form.pod_replicas" />
                </el-form-item>
                <el-form-item>
                    <el-button @click="submit">提交</el-button>
                </el-form-item>
            </el-form>
        </el-dialog>
    </div>
</template>
<script setup>
import {getRelativeTime}             from "@/utils/time"
import {onMounted, reactive, toRefs}                     from "vue";
import {AddDeployment, DeleteDeployment, DeploymentList} from "@/api/deployment";
import {NamespaceList}                                   from "@/api/namespace";

const data = reactive({
    tableData: [],
    namespace: "",
    namespaceList: [],
    visible:false,
    form:{}
})
const {tableData, namespace, namespaceList,visible,form} = toRefs(data)
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

const showDialog = ()=>{
    data.visible = true
}

const submit = async () => {
    await AddDeployment(data.form)
    data.visible = false
    data.form={}
    await getList()
}

const deleteOne = async () => {
    const params = {}
    await DeleteDeployment(params)
    await getList()
}
</script>

