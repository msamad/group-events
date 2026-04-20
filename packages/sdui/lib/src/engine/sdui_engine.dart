import '../models/sdui_action.dart';
import '../models/sdui_component.dart';
import '../models/sdui_descriptor.dart';

class SduiEngine {
  const SduiEngine();

  SduiDescriptor parseDescriptor(Map<String, Object?> json) {
    return SduiDescriptor.fromJson(json);
  }

  List<SduiComponent> resolveComponents(SduiDescriptor descriptor) {
    return List<SduiComponent>.unmodifiable(descriptor.components);
  }

  List<SduiAction> resolveActions(SduiDescriptor descriptor) {
    return List<SduiAction>.unmodifiable(descriptor.actions);
  }

  bool supportsComponent(SduiComponent component) {
    switch (component.type) {
      case 'badge':
      case 'bullet_list':
      case 'button':
      case 'card':
      case 'info_panel':
      case 'list':
      case 'section':
      case 'stack':
      case 'text':
        return true;
      default:
        return false;
    }
  }
}
