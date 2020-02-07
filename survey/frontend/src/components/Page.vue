<template>
  <div>
    <div v-for="(widget, index) in widgets" v-bind:key="index">
      <!-- Check which type the Widget has and render it and assign corresponding
      structure-->
      <v-component
        :is="widget.type"
        v-model="widgets[index].structure"
        v-bind:required="widgets[index].required"
      />
    </div>
  </div>
</template>

<script lang="ts">
import Vue from "vue";
import WidgetDescription from "../interfaces/WidgetDescription";
import Matrix from "./Matrix.vue";
import Dropdown from "./Dropdown.vue";

/**
 * Single Page with multiple widgets
 */
export default Vue.extend({
    props: ["value"],
    components: {
        Matrix,
        Dropdown
    },
    data: () => ({
        widgets: [] as WidgetDescription[]
    }),
    methods: {},
    mounted() {
        this.widgets = this.value;
    },
    watch: {
        value: function(newValue: WidgetDescription[]) {
            this.widgets = newValue;
        }
    }
});
</script>
