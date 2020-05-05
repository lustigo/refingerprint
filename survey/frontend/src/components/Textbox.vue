<template>
    <v-container>
        <v-col>
            {{structure.question}}
            <span v-if="req" style="color:red;">*</span>
        </v-col>
        <v-textarea 
            outlined 
            v-model="structure.answer" 
            v-on:change="onChanged"
            v-bind:rows="structure.rows"
            ></v-textarea>
    </v-container>
</template>

<script lang="ts">
import Vue from "vue";
import TextboxDescription from "../interfaces/TextboxDescription";
/**
 * Widget with a big Textbox as user-written input.
 */
export default Vue.extend({
    props: ["value","required"],
    data: () => ({
        req: false,
        structure: {} as TextboxDescription,
    }),
    methods: {
        onChanged: function(){
            this.$emit("input",this.structure);
            this.$emit("completed",!this.req || this.structure.answer != "");
        }
    },
    mounted() {
        this.req = this.required;
        this.structure = this.value;
        if(!this.structure.answer){
            this.structure.answer="";
        }
        this.$emit("completed",!this.req || this.structure.answer != "");
    },watch: {
        value: function(newValue) {
            this.structure = newValue;
            this.$emit("completed",!this.req || this.structure.answer != "");
        }}
});
</script>