<template>
    <div>
        <Namespace @namespaceChanged="namespaceChanged" ref="namespaceRef"></Namespace>
        <el-button @click="showDialog" style="float: right">新增</el-button>
        <el-table :data="tableData" style="width: 100%;margin-top: 12px">
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
                    <el-select v-model="form.pod_namespace" @change="getList">
                        <el-option v-for="(item,index) in namespaceList" :index="index" :label="item.metadata.name" :value="item.metadata.name"></el-option>
                    </el-select>
                </el-form-item>
                <el-form-item label="镜像">
                    <el-input v-model="form.pod_image" />
                </el-form-item>
                <el-form-item label="端口号">
                    <el-input v-model.number="form.pod_port" />
                </el-form-item>
                <el-form-item label="副本数">
                    <el-input-number v-model.number="form.pod_replicas" />
                </el-form-item>
                <el-form-item>
                    <el-button @click="submit">提交</el-button>
                </el-form-item>
            </el-form>
        </el-dialog>
    </div>
</template>
<script setup>
import {getRelativeTime}                                 from "@/utils/time"
import {onMounted, reactive, ref, toRefs}                from "vue";
import {AddDeployment, DeleteDeployment, DeploymentList} from "@/api/deployment";
import Namespace from "@/components/namespace/index.vue"

const data = reactive({
    tableData: [],
    namespace: "",
    namespaceList: [],
    visible:false,
    form:{
        pod_replicas:1,
    }
})
const {tableData, namespace, namespaceList,visible,form} = toRefs(data)

const namespaceChanged = (val) => {
    data.namespace = val;
    getList()
}
const getList = async () => {
    data.tableData = []
    const res = await DeploymentList({namespace:data.namespace})
    data.tableData = res.message.items
}

onMounted(async () => {
    await getList()
    await getNamespaceDataList()
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

const deleteOne = async (row) => {
    const params = {
        namespace:row.metadata.namespace,
        name:row.metadata.name,
    }
    await DeleteDeployment(params)
    await getList()
}

const namespaceRef =ref(null)
const getNamespaceDataList = ()=>{
    data.namespaceList = namespaceRef.value.getNamespaceDataList()
}
</script>

