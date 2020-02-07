<template>
  <v-row>
    <v-col class="d-flex font-regular font-weight-light" cols="6">{{ structure.question }}</v-col>
    <v-col class="d-flex" cols="4">
      <v-select
        v-bind:items="structure.items"
        solo
        v-bind:label="structure.label"
        v-model="structure.selected"
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
    props: ["value"],
    data: () => ({
        structure: {} as DropdownDescription
    }),
    mounted() {
        this.structure = this.value;
        this.structure.selected = "";
    },
    methods: {
        selected() {
            this.$emit("input", this.structure);
        }
    }
});
</script>
