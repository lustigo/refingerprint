import { UUID } from "./UUID";
import WidgetDescription from "./WidgetDescription";

/**
 * Describes the Structure of the Survey
 */
export default interface SurveyDescription {
  readonly id: UUID;
  readonly pages: WidgetDescription[][];
}
