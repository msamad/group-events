import 'sdui_action.dart';

class SduiComponent {
  const SduiComponent({
    required this.id,
    required this.type,
    this.visible = true,
    this.title,
    this.body,
    this.data = const <String, Object?>{},
    this.actions = const <SduiAction>[],
    this.children = const <SduiComponent>[],
  });

  factory SduiComponent.fromJson(Map<String, Object?> json) {
    return SduiComponent(
      id: (json['id'] as String?) ?? '',
      type: json['type'] as String,
      visible: (json['visible'] as bool?) ?? true,
      title: json['title'] as String?,
      body: json['body'] as String?,
      data: _readObjectMap(json['data'] ?? json['props']),
      actions: _readObjectList(
        json['actions'],
      ).map(SduiAction.fromJson).toList(growable: false),
      children: _readObjectList(
        json['children'],
      ).map(SduiComponent.fromJson).toList(growable: false),
    );
  }

  final String id;
  final String type;
  final bool visible;
  final String? title;
  final String? body;
  final Map<String, Object?> data;
  final List<SduiAction> actions;
  final List<SduiComponent> children;

  Map<String, Object?> toJson() {
    return <String, Object?>{
      'id': id,
      'type': type,
      'visible': visible,
      if (title != null) 'title': title,
      if (body != null) 'body': body,
      if (data.isNotEmpty) 'data': data,
      if (actions.isNotEmpty)
        'actions': actions
            .map((action) => action.toJson())
            .toList(growable: false),
      if (children.isNotEmpty)
        'children': children
            .map((child) => child.toJson())
            .toList(growable: false),
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
