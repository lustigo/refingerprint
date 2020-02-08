import { UUID } from "./UUID";
import WidgetData from "./WidgetData";
/**
 * Describes the structure of the filled-in survey, which will be sent to the server.
 */
export default interface SurveyData {
    readonly id: UUID;
    readonly user: UUID;
    pages: WidgetData[][];
    // eslint-disable-next-line semi
}
