import Collector from "./collector";
import Module from "./interfaces/Module";

import BrowserInfo from "./modules/BowserInfo";
import Canvas from "./modules/Canvas";
import Id from "./modules/Id";
import ScreenResolution from "./modules/ScreenResolution";
import WebGL from "./modules/WebGL";

/**
 * All module classes that should be used for the Collection
 * Time-critical modules should be placed at the beginning.
 */
const modules: Array<{ new(): Module }> = [Canvas, BrowserInfo, ScreenResolution, Id, WebGL];

/**
 * Start Data Collector
 */
new Collector(modules);
