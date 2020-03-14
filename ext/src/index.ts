import Collector from "./collector";
import Module from "./interfaces/Module";

import Audio from "./modules/Audio";
import BrowserInfo from "./modules/BowserInfo";
import Canvas from "./modules/Canvas";
import FramePosition from "./modules/FramePosition";
import Id from "./modules/Id";
import MouseMovementCheckbox from "./modules/MouseMovementCheckbox";
import MouseMovementRest from "./modules/MouseMovementRest";
import ScreenResolution from "./modules/ScreenResolution";
import Time from "./modules/Time";
import WebGL from "./modules/WebGL";

/**
 * All module classes that should be used for the Collection
 * Time-critical modules should be placed at the beginning.
 */
const modules: Array<{ new(): Module }> = [Time, Canvas, Audio, MouseMovementCheckbox, MouseMovementRest, FramePosition, BrowserInfo, ScreenResolution, Id, WebGL];

/**
 * Start Data Collector
 */
new Collector(modules);
