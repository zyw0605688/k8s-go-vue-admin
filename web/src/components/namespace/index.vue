<template>
    <span>
        <span style="font-size: 14px">命名空间：</span>
        <el-select v-model="namespace" @change="namespaceChanged">
            <el-option v-for="(item,index) in namespaceList" :index="index" :label="item.metadata.name" :value="item.metadata.name"></el-option>
        </el-select>
    </span>
</template>
<script setup lang="ts" name="Namespace">
import {onMounted, reactive, toRefs} from "vue";
import {NamespaceList} from "@/api/namespace";

const props = defineProps({
    defaultNamespace: String
})

const data = reactive({
    namespaceList: [],
    namespace: props.defaultNamespace,
})
const {namespace, namespaceList} = toRefs(data)

const getNamespaceList = async () => {
    const res = await NamespaceList()
    data.namespaceList = res.message.items
}

onMounted(async () => {
    await getNamespaceList()
})
const emit = defineEmits(["namespaceChanged"]);
const namespaceChanged = ()=>{
    emit("namespaceChanged", data.namespace)
}
</script>
