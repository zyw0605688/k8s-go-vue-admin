<template>
    <div>
        <Namespace @namespaceChanged="namespaceChanged" ref="namespaceRef"></Namespace>
        <el-button @click="showDialog" style="float: right">新增</el-button>
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
            <el-table-column label="操作">
                <template #default="scope">
                    <el-button @click="deleteOne(scope.row)">删除</el-button>
                </template>
            </el-table-column>
        </el-table>
        <el-dialog v-model="visible">
            <el-form :model="form" label-width="120px">
                <el-form-item label="名称">
                    <el-input v-model="form.ingress_name" />
                </el-form-item>
                <el-form-item label="命名空间">
                    <el-select v-model="form.ingress_namespace" @change="getList">
                        <el-option v-for="(item,index) in namespaceList" :index="index" :label="item.metadata.name" :value="item.metadata.name"></el-option>
                    </el-select>
                </el-form-item>
                <el-form-item label="域名">
                    <el-input v-model="form.ingress_host" />
                </el-form-item>
                <el-form-item label="服务名">
                    <el-input v-model.number="form.ingress_service_name" />
                </el-form-item>
                <el-form-item label="服务端口号">
                    <el-input-number v-model.number="form.ingress_service_port" />
                </el-form-item>
                <el-form-item>
                    <el-button @click="submit">提交</el-button>
                </el-form-item>
            </el-form>
        </el-dialog>
    </div>
</template>
<script setup>
import {getRelativeTime}                  from "@/utils/time"
import {onMounted, reactive, ref, toRefs} from "vue";
import Namespace                                from "@/components/namespace/index.vue";
import {AddIngress, DeleteIngress, IngressList} from "@/api/ingress";

const data = reactive({
    tableData: [],
    namespace: "",
    visible:false,
    namespaceList: [],
    form:{}
})
const {tableData, namespace,namespaceList,visible,form} = toRefs(data)

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
    await getNamespaceDataList()
})

const showDialog = ()=>{
    data.visible = true
}

const namespaceRef =ref(null)
const getNamespaceDataList = ()=>{
    data.namespaceList = namespaceRef.value.getNamespaceDataList()
}

const submit = async () => {
    await AddIngress(data.form)
    data.visible = false
    data.form={}
    await getList()
}

const deleteOne = async (row) => {
    const params = {
        namespace:row.metadata.namespace,
        name:row.metadata.name,
    }
    await DeleteIngress(params)
    await getList()
}
</script>

