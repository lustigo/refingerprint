import { UUID } from "./UUID";
import WidgetDescription from "./WidgetDescription";

/**
 * Describes the Structure of the Survey
 */
export default interface SurveyDescription {
    readonly survey: UUID;
    readonly pages: WidgetDescription[][];
    // eslint-disable-next-line semi
}
