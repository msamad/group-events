import 'sdui_action.dart';
import 'sdui_component.dart';

class SduiDescriptor {
  const SduiDescriptor({
    required this.screenId,
    required this.title,
    this.subtitle,
    this.components = const <SduiComponent>[],
    this.actions = const <SduiAction>[],
    this.metadata = const <String, Object?>{},
  });

  factory SduiDescriptor.fromJson(Map<String, Object?> json) {
    return SduiDescriptor(
      screenId: json['screenId'] as String,
      title: json['title'] as String,
      subtitle: json['subtitle'] as String?,
      components: _readObjectList(
        json['components'],
      ).map(SduiComponent.fromJson).toList(growable: false),
      actions: _readObjectList(
        json['actions'],
      ).map(SduiAction.fromJson).toList(growable: false),
      metadata: _readObjectMap(json['metadata']),
    );
  }

  final String screenId;
  final String title;
  final String? subtitle;
  final List<SduiComponent> components;
  final List<SduiAction> actions;
  final Map<String, Object?> metadata;

  Map<String, Object?> toJson() {
    return <String, Object?>{
      'screenId': screenId,
      'title': title,
      if (subtitle != null) 'subtitle': subtitle,
      if (components.isNotEmpty)
        'components': components
            .map((component) => component.toJson())
            .toList(growable: false),
      if (actions.isNotEmpty)
        'actions': actions
            .map((action) => action.toJson())
            .toList(growable: false),
      if (metadata.isNotEmpty) 'metadata': metadata,
    };
  }
}

Map<String, Object?> _readObjectMap(Object? value) {
  if (value is Map<String, Object?>) {
    return value;
  }

  if (value is Map) {
    return value.map((key, mapValue) => MapEntry(key.toString(), mapValue));
  }

  return const <String, Object?>{};
}

List<Map<String, Object?>> _readObjectList(Object? value) {
  if (value is! List) {
    return const <Map<String, Object?>>[];
  }

  return value
      .whereType<Map>()
      .map(
        (entry) =>
            entry.map((key, mapValue) => MapEntry(key.toString(), mapValue)),
      )
      .toList(growable: false);
}
