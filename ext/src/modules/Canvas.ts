import Module from "../interfaces/Module";
import jsSHA from "jssha";

/**
 * Module which computes a Canvas Fingerprint
 * 
 * Freely adapted from Acar, Gunes & Eubank, Christian & Englehardt, Steven & Juarez, Marc & Narayanan, Arvind & Diaz, Claudia. (2014). The Web Never Forgets: Persistent Tracking Mechanisms in the Wild. Proceedings of the ACM Conference on Computer and Communications Security. 674-689. 10.1145/2660267.2660347.
 */
export default class Canvas implements Module {
    /**
     * Name of the module
     */
    public readonly name = "Canvas";

    /**
     * The Canvas Fingerprint
     */
    private fingerprint = "UNKNOWN";

    /**
     * Reference to the Canvas
     */
    private canvas: HTMLCanvasElement = document.createElement("canvas");


    /**
     * Initializes the Canvas and renders it on the DOM 
     */
    private initCanvas(): void {
        this.canvas.id = "PixelPerfect";
        this.canvas.width = 100;
        this.canvas.height = 100;
        this.canvas.style.display = "none";
        document.body.appendChild(this.canvas);
    }

    /**
     * Fills the Canvas with Text
     */
    private fillCanvas(): void {
        const context = this.canvas.getContext("2d");
        if (context) {
            context.font = "invalid font specification. sic!";
            context.fillStyle = "#12E3DB";
            context.textBaseline = "top";
            context.fillText("Gotcha 🖱️ 帆布", 2, 2);

            context.font = "20pt no-real-font-never-forget";
            context.fillStyle = "#C0B733";
            context.fillText("Gotcha 🖱️ 帆布", 2, 2);
        }
    }

    /**
     * Will be called, when the Captcha is rendered
     * Renders the Canvas with its content
     */
    public start(): void {
        this.initCanvas();
        this.fillCanvas();
    }

    /**
     * Gets the Content of the Canvas (and throws the Data-Png-Information away)
     * @returns ImageData of the Canvas
     */
    private getImageData(): string {
        return this.canvas ? this.canvas.toDataURL().substring(22) : "UNKNOWN";
    }

    /**
     * Computes the Fingerprint of the given string
     * @param imageData Data to computer Fingerprint from
     * @returns Fingerprint
     */
    private computeHash(imageData: string): string {
        const hasher = new jsSHA("SHA-512", "B64");
        hasher.update(imageData);
        return hasher.getHash("B64");
    }

    /**
    * Will be called, when the Captcha is solved
    * Gets the Image Data of the Canvas and stores the Fingerprint
    */
    public stop(): void {
        this.fingerprint = this.computeHash(this.getImageData());
    }

    /**
     * returns the Fingerprint
     * @returns Canvas Fingerprint
     */
    public getCollectedData(): string {
        return this.fingerprint;
    }


}