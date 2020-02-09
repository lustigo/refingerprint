<template>
  <v-row>
    <v-col class="d-flex font-regular font-weight-light text-center" style="margin-left:1%;" cols="6">
      {{ structure.question }}
      <span v-if="req" style="color:red;">*</span>  
    </v-col>
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
        req: false,
        structure: {} as DropdownDescription,
    }),
    mounted() {
        this.structure = this.value;
        if(!this.structure.selected){
            this.structure.selected = "";
        }
        this.req = this.required;
        if(!this.req || this.structure.selected != ""){
            this.$emit("completed",true);
        }
    }
});
</script>
