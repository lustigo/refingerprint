<template>
    <vue-markdown v-bind:source="structure.content"/>
</template>

<script lang="ts">
import Vue from "vue";
import VueMarkdown from "vue-markdown";
import TextWidgetDescription from "../interfaces/TextWidgetDescription";

/**
 * Text-Widget, works with markdown
 */
export default Vue.extend({
    props: ["value"],
    components: {
        "vue-markdown": VueMarkdown
    },
    data: () => ({
        structure: {} as TextWidgetDescription,
    }),
    mounted() {
        this.structure = this.value;
        this.$emit("completed",true);
    }, watch: {
        value: function(newValue: TextWidgetDescription) {
            this.structure = newValue;
            this.$emit("completed",true);
        }}
});
</script>