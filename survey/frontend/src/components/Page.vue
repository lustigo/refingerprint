<template>
  <div>
    <div v-for="(widget, index) in widgets" v-bind:key="index">
      /*
      * Check which type the Widget has and render it and assign
      * the corresponding structure
      */
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
import Widget from "./Widget.vue";

/**
 * Single Page with multiple widgets
 */
export default Vue.extend({
  props: ["value"],
  components: {
    Widget
  },
  data: () => ({
    widgets: [] as WidgetDescription[]
  }),
  methods: {},
  mounted() {
    this.widgets = this.value;
  }
});
</script>
