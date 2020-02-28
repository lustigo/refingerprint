import Module from "../interfaces/Module";

/**
 * Module which gets the Graphic Card Information and WebGL Extensions
 * 
 * created with help of https://gist.github.com/cvan/042b2448fcecefafbb6a91469484cdf8
 */
export default class WebGL implements Module {
    /**
     * Name of the module
     */
    public readonly name = "WebGL";

    /**
     * The Graphic Card Vendor
     */
    private vendor = "UNKNOWN";

    /**
     * The Graphic Card Model
     */
    private model = "UNKNOWN";


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
     * Gets the Information and stores it
     */
    private getWebGLInformation(): void {
        const context = this.canvas.getContext("webgl");
        if (!context) {
            this.model = "NOT AVAIL";
            this.vendor = "NOT AVAIL";
            this.extensions = [];
            return;
        }
        const debugInfo = context.getExtension("WEBGL_debug_renderer_info");
        if (!debugInfo) {
            this.vendor = "NOT AVAIL";
            this.model = "NOT AVAIL";
            return;
        }

        this.vendor = context.getParameter(debugInfo.UNMASKED_VENDOR_WEBGL);
        this.model = context.getParameter(debugInfo.UNMASKED_RENDERER_WEBGL);


        const ext = context.getSupportedExtensions();
        if (ext) {
            this.extensions = ext;
        }
    }

    /**
    * Will be called, when the Captcha is solved
    * Gets the Information and stores it.
    */
    public stop(): void {
        this.getWebGLInformation();
    }

    /**
     * returns the gathered Information
     * @returns Information
     */
    public getCollectedData(): object {
        return {
            gcVendor: this.vendor,
            gcModel: this.model,
            webglExtensions: this.extensions,
        };
    }


}
