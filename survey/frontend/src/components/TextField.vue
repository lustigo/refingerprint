<template>
    <v-row>
        <v-col class="d-flex font-regular font-weight-light text-center" style="margin-left:1%;" cols="6">{{ structure.question }}</v-col>
        <v-col class="d-flex" cols="4">
        <v-text-field
            v-model="structure.content"
            outlined
            v-on:change="$emit('completed', !req || structure.content != '')"
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
        structure: {} as TextFieldDescription,
        req: false,
    }),
    mounted() {
        this.structure = this.value;
        if(!this.structure.content){
            this.structure.content = "";
        }
        this.req = this.required;
        this.$emit("completed",!this.req || this.structure.content.trim() != "");
    }
});
</script>