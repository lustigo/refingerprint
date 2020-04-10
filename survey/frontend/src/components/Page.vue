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
        v-on:input="$emit('input', widgets)"
      />
    </v-container>
    <v-container class="text-left" absolute height="auto" style="margin-bottom:4%; bottom: 0; position:fixed" v-if="isSomethingRequired">
        <span style="color:red;">*</span> Erforderlich
    </v-container>
  </v-container>
</template>

<script lang="ts">
import Vue from "vue";
import WidgetDescription from "../interfaces/WidgetDescription";
import Captcha from "./Captcha.vue";
import Checkbox from "./Checkbox.vue";
import Dropdown from "./Dropdown.vue";
import Matrix from "./Matrix.vue";
import Textbox from "./Textbox.vue";
import Textfield from "./TextField.vue";
import Textwidget from "./Textwidget.vue";
import Vector from "./Vector.vue";

/**
 * Single Page with multiple widgets
 */
export default Vue.extend({
    props: ["value"],
    components: {
        Captcha,
        Checkbox,
        Dropdown,
        Matrix,
        Textbox,
        Textfield,
        Textwidget,
        Vector
    },
    data: () => ({
        isSomethingRequired: false,
        completionMap: [] as Boolean[],
        widgets: [] as WidgetDescription[],
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
            if(this.widgets[i].required){
                this.isSomethingRequired = true;
            }
        }
    },
    watch: {
        value: function(newValue: WidgetDescription[]) {
            this.isSomethingRequired = false;
            this.widgets = newValue;
            this.completionMap = [];
            for(let i = 0; i < this.widgets.length; i++){
                this.completionMap.push(!this.widgets[i].required);
                if(this.widgets[i].required){
                    this.isSomethingRequired = true;
                }
            }
        }
    }
});
</script>
