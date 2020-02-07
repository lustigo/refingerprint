<template>
  <v-row>
    <v-col class="d-flex font-regular font-weight-light text-center" style="margin-left:1%;" cols="6">{{ structure.question }}</v-col>
    <v-col class="d-flex" cols="4">
      <v-select
        v-bind:items="structure.items"
        solo
        v-bind:label="structure.label"
        v-model="structure.selected"
        v-on:change="$emit('completed', !req || structure.selected != '')"
      ></v-select>
    </v-col>
  </v-row>
</template>

<script lang="ts">
import Vue from "vue";
import DropdownDescription from "../interfaces/DropdownDescription";

/**
 * Dropdown-Widget
 */
export default Vue.extend({
    props: ["value","required"],
    data: () => ({
        structure: {} as DropdownDescription,
        req: false
    }),
    mounted() {
        this.structure = this.value;
        if(!this.structure.selected){
            this.structure.selected = "";
        }
        this.req = this.required;
        if(!this.req){
            this.$emit("completed",true);
        }
    }
});
</script>
