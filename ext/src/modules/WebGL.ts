import Module from "../interfaces/Module";

/**
 * Module which gets the WebGL Renderer and Version  
 */
export default class WebGL implements Module {
    /**
     * Name of the module
     */
    public readonly name = "WebGL";

    /**
     * The WebGL Vendor
     */
    private vendor = "UNKNOWN";

    /**
     * The WebGL Model
     */
    private renderer = "UNKNOWN";

    /**
     * The WebGL Version
     */
    private version = "UNKNOWN";

    /**
     * Supported WebGL Extensions
     */
    private extensions = ["UNKNOWN"];

    /**
     * Reference to the Canvas
     */
    private canvas: HTMLCanvasElement = document.createElement("canvas");


    /**
     * Initializes the Canvas and renders it on the DOM 
     */
    private initCanvas(): void {
        this.canvas.id = "BrowserFingerprintSurvey";
        this.canvas.width = 1;
        this.canvas.height = 1;
        this.canvas.style.display = "none";
        document.body.appendChild(this.canvas);
    }

    /**
     * Will be called, when the Captcha is rendered
     * Renders the Canvas
     */
    public start(): void {
        this.initCanvas();
    }

    /**
    * Will be called, when the Captcha is solved
    * Gets the WebGL Information and stores it.
    */
    public stop(): void {
        const context = this.canvas.getContext("webgl");
        if (!context) {
            this.renderer = "NOT AVAIL";
            this.vendor = "NOT AVAIL";
            this.version = "NOT AVAIL";
            this.extensions = [];
            return;
        }
        this.vendor = context.getParameter(context.VENDOR);
        this.renderer = context.getParameter(context.RENDERER);
        this.version = context.getParameter(context.VERSION);

        const ext = context.getSupportedExtensions();
        if (ext) {
            this.extensions = ext;
        }
    }

    /**
     * returns the WebGL Information
     * @returns WebGL Information
     */
    public getCollectedData(): object {
        return {
            vendor: this.vendor,
            renderer: this.renderer,
            version: this.version,
            extensions: this.extensions,
        };
    }


}
