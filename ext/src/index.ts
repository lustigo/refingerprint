import Collector from "./collector";
import Module from "./interfaces/Module";

import Canvas from "./modules/Canvas";
import Id from "./modules/Id";

/**
 * All module classes that should be used for the Collection
 * Time-critical modules should be placed at the beginning.
 */
const modules: Array<{ new(): Module }> = [Canvas, Id];

/**
 * Start Data Collector
 */
new Collector(modules);
