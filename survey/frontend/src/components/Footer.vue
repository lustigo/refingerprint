<template>
    <v-footer  absolute height="auto" class="font-weight-medium" >   
        <v-layout>
            <v-flex class="text-left">
                <v-btn style="font-size:28px;" text v-bind:disabled="!navigationOptions.canBack" v-on:click="pageBackward">
                    <v-icon>fa-arrow-left</v-icon>
                </v-btn>
            </v-flex> 
            <v-flex class="text-right">
                <v-btn style="font-size:28px;" text v-on:click="$emit('send',{})" v-bind:disabled="!navigationOptions.canSend" v-if="navigationOptions.currentPage == navigationOptions.pageAmount-1">
                    <v-icon>fa-paper-plane</v-icon>
                </v-btn>
                <v-btn style="font-size:28px;" text v-bind:disabled="!navigationOptions.canForward" v-on:click="pageForward" v-else>
                    <v-icon>fa-arrow-right</v-icon>
                </v-btn>
            </v-flex>
        </v-layout>
    </v-footer>
</template>

<script lang="ts">
import Vue from "vue";
import NavigationOptions from "../interfaces/NavigationOptions";

export default Vue.extend({
    props: ["value"],
    data: () => ({
        navigationOptions: {} as NavigationOptions
    }),
    mounted(){
        this.navigationOptions = this.value;
    },
    methods: {
        pageForward: function(){
            this.navigationOptions.canForward = false;
            this.navigationOptions.canBack = true;
            this.navigationOptions.currentPage ++;
        },
        pageBackward: function(){
            this.navigationOptions.canForward = true;
            this.navigationOptions.currentPage --;
            this.navigationOptions.canBack = this.navigationOptions.currentPage > 0;
        },
    }
});
</script>