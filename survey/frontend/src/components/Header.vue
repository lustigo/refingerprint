<template>
    <v-toolbar>
        <v-toolbar-items>
            <v-img src="../assets/logo.png" aspect-ratio="1" width="5rem"/>
            <extension></extension>
        </v-toolbar-items>
        <v-toolbar-title style="margin-left:42%">Umfrage</v-toolbar-title>
    </v-toolbar>
</template>

<script lang="ts">
import Vue from "vue";
import Extension from "./Extension.vue";
/**
 * Navigationbar
 */
export default Vue.extend({
    components: {
        Extension
    }
});
</script>