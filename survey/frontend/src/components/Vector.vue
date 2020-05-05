<template>
    <v-container class="font-regular font-weight-light text-center">
        <span cols="0"><strong>{{structure.question}}</strong> <span v-if="req" style="color:red;">*</span></span>
        <span style="margin-left:5%;">
            <span
                v-for="(value,key) in structure.selection" v-bind:key="value+key"
                style="margin-left:1%;">
                        <input  
                            class="v-input--selection-controls ripple"
                            type="radio"
                            v-bind:value="key"
                            v-on:change="changed"
                            v-model="structure.answer" 
                            v-bind:name="structure.selection.join('')"/>
                        {{ value }}

            </span>
        </span>
    </v-container>
</template>

<style scoped>
strong{
    font-weight: bold;
}
</style>

<script type="ts">
import Vue from "vue";

/**
 * Vector Widget:
 * Question with Radio Buttons
 */
export default Vue.extend({
    props: ["value","required"],
    data: () => ({
        req: false,
        structure: {}
    }),
    methods: {
        changed: function () {
            this.$emit("input",this.structure);
            this.$emit("completed",!this.req || this.structure.answer != -1);
        }
    },
    mounted() {
        this.structure = this.value;
        this.req = this.required;
        if(!this.structure.answer && this.structure.answer !== 0){
            this.structure.answer = -1;
        }
        this.$emit("completed",!this.req || this.structure.answer != -1);
    },watch: {
        value: function(newValue) {
            this.structure = newValue;
            this.$emit("completed",!this.req || this.structure.answer != -1);
        }}
});
</script>