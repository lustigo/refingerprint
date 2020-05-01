<template>
    <v-icon 
     v-on:refextLoaded="display"
     v-bind:class="{block: noShow}" 
     id="refingerprint-ext">
            fa-check-circle
    </v-icon>
</template>

<style scoped>
    .block {
        display: none;
    }
    
    #refingerprint-ext{
        color: green;
    }
</style>

<script lang="ts">
import Vue from "vue";
/**
 * Displays a checkmark if the extension is loaded (a refextLoaded event is received)
 */
export default Vue.extend({
    data: () => ({
        noShow: true
    }),
    methods: {
        display: function() {
            this.noShow = false;
        }
    }
});
</script>
