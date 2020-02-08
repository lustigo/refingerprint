<template>
  <v-container fluid fill-height>
    <v-container v-for="(widget, index) in widgets" v-bind:key="index">
      <!-- Check which type the Widget has and render it and assign corresponding
      structure-->
      <v-component
        v-bind:is="widget.type"
        v-model="widgets[index].structure"
        v-bind:required="widgets[index].required"
        v-on:completed="(data) => onCompleted(index,data)"
      />
    </v-container>
  </v-container>
</template>

<script lang="ts">
import Vue from "vue";
import WidgetDescription from "../interfaces/WidgetDescription";
import Captcha from "./Captcha.vue";
import Dropdown from "./Dropdown.vue";
import Matrix from "./Matrix.vue";
import Textfield from "./TextField.vue";
import Textwidget from "./Textwidget.vue";

/**
 * Single Page with multiple widgets
 */
export default Vue.extend({
    props: ["value"],
    components: {
        Captcha,
        Dropdown,
        Matrix,
        Textfield,
        Textwidget
    },
    data: () => ({
        widgets: [] as WidgetDescription[],
        completionMap: [] as Boolean[]
    }),
    methods: {
        /**
         * Checks if all widgets are filled-in
         */
        onCompleted: function(index:number, completed: boolean){
            this.completionMap[index] = completed;
            this.$emit("completed",
                this.completionMap.filter(x => x).length == this.completionMap.length);
        
        }
    },
    mounted() {
        this.widgets = this.value;
        this.completionMap = [];
        for(let i = 0; i < this.widgets.length; i++){
            this.completionMap.push(!this.widgets[i].required);
        }
    },
    watch: {
        value: function(newValue: WidgetDescription[]) {
            this.widgets = newValue;
            this.completionMap = [];
            for(let i = 0; i < this.widgets.length; i++){
                this.completionMap.push(!this.widgets[i].required);
            }
        }
    }
});
</script>
