/**
 * Describes the structure of a Widget.
 */
export default interface WidgetDescription {
  readonly type: String;
  readonly structure: Object;
  readonly required: boolean;
}
