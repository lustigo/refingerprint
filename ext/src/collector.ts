import Module from './interfaces/Module';

/**
 * Data Collector
 */
export default class Collector {

    /**
     * Modules that should be initialized when a new Captcha will be rendered
     */
    private modules: Array<{ new(): Module }>;

    /**
     * Module Instances for the next or current Captcha
     */
    private collectors: Array<Module> = new Array();


    /**
     * Initializes the Constructor with the data collection modules
     * Add the Listener to the Document
     */
    public constructor(modules: Array<{ new(): Module }>) {
        console.log("Constructed");
        this.modules = modules;
        this.reset();
        document.getRootNode().addEventListener("captcha-created", this.start);
        document.getRootNode().addEventListener("captcha-solved", this.stop);
    }

    /**
     * Will be called when the Captcha is rendered.
     * Will start all Modules
     */
    public start() {
        this.collectors.forEach(col => col.start());
    }

    /**
     * Will be called when the Captcha was solved.
     * Will stop all Modules and collect the data
     */
    public stop() {
        this.collectors.forEach(col => col.stop());
        const data = this.collectData();
        this.saveData(data);
        this.reset();
    }

    /**
     * Collects the data gathered by the Modules and builds a single Object
     */
    private collectData(): Map<string, string | number | object> {
        const data = new Map<string, string | number | object>();
        this.collectors.forEach(col => data.set(col.name, col.getCollectedData()));
        return data;
    }

    /**
     * Resets the Collectors after a Captcha is solved 
     * or for the first time during Construction.
     */
    private reset() {
        this.collectors = this.modules.map(Mod => new Mod());
    }

    /**
     * Saves the collected data to 
     * @param data Data to save
     */
    private saveData(data: Map<string, string | number | object>) {
        console.log(data);
    }


}