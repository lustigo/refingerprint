<template>
    <v-row>
        <v-col class="d-flex font-regular font-weight-light text-center" style="margin-left:1%;" cols="6">
            {{ structure.question }}
            <span v-if="req" style="color:red;">*</span>
        </v-col>
        <v-col class="d-flex" cols="4">
        <v-text-field
            v-model="structure.content"
            outlined
            v-on:change="onChange"
        ></v-text-field>
        </v-col>
    </v-row>
</template>

<script lang="ts">
import Vue from "vue";
import TextFieldDescription from "../interfaces/TextFieldDescription";

/**
 * TextField-Widget
 */
export default Vue.extend({
    props: ["value","required"],
    data: () => ({
        req: false,
        structure: {} as TextFieldDescription,
    }),
    methods: {
        onChange: function (){
            this.$emit("completed", !this.req || this.structure.content != "");
            this.$emit("input", this.structure);
        },
    },
    mounted() {
        this.structure = this.value;
        if(!this.structure.content){
            this.structure.content = "";
        }
        this.req = this.required;
        this.$emit("completed",!this.req || this.structure.content.trim() != "");
    },watch: {
        value: function(newValue) {
            this.structure = newValue;
            this.$emit("completed",!this.req || this.structure.content != "");
        }}
});
</script>