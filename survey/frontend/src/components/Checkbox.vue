<template>
    <v-container class="font-regular font-weight-light text-center">
        <span>{{structure.question}} <span v-if="req" style="color:red;">*</span></span>
        <span style="margin-left:5%">
            <span v-for="(value, index) in structure.selection" v-bind:key="value+index" style="margin-left:1%">
                <input type="checkbox" v-bind:checked="structure.checked[index]" v-on:change="onChecked(index)"/>
                {{value}}
            </span>
        </span>
    </v-container>
</template>

<script lang="ts">
import Vue from "vue";
import CheckboxDescription from "../interfaces/CheckboxDescription";
/**
 * Checkbox Widget
 */
export default Vue.extend({
    props: ["value","required"],
    data: () => ({
        req: false,
        structure: {} as CheckboxDescription
    }),
    methods:{
        onChecked: function (index: number){
            this.structure.checked[index] = !this.structure.checked[index];
            this.$emit("input",this.structure);
            this.$emit("completed",!this.req || this.structure.checked.find(v => v) != undefined);
        }
    },
    mounted() {
        this.req = this.required;
        this.structure = this.value;
        if(!this.structure.checked){
            this.structure.checked = this.structure.selection.map(_ => false);
        }
        // Widget is completed, if it is not required or at least one checkbox is filled
        this.$emit("completed",!this.req || this.structure.checked.find(v => v) != undefined);
    }
});
</script>