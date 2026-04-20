import 'package:sdui/sdui.dart';

class Group {
  const Group({
    required this.id,
    required this.slug,
    required this.name,
    required this.description,
    required this.archived,
    required this.createdAt,
    this.ui,
  });

  factory Group.fromJson(Map<String, Object?> json) {
    final uiValue = json['ui'];
    final dataValue = json['data'];
    final source = dataValue is Map ? _normalizeMap(dataValue) : json;

    return Group(
      id: source['id'] as String,
      slug: source['slug'] as String,
      name: source['name'] as String,
      description: (source['description'] as String?) ?? '',
      archived: (source['archived'] as bool?) ?? false,
      createdAt: DateTime.parse(source['createdAt'] as String),
      ui: uiValue is Map<String, Object?>
          ? SduiDescriptor.fromJson(uiValue)
          : uiValue is Map
          ? SduiDescriptor.fromJson(_normalizeMap(uiValue))
          : null,
    );
  }

  final String id;
  final String slug;
  final String name;
  final String description;
  final bool archived;
  final DateTime createdAt;
  final SduiDescriptor? ui;

  Map<String, Object?> toJson() {
    return <String, Object?>{
      'id': id,
      'slug': slug,
      'name': name,
      'description': description,
      'archived': archived,
      'createdAt': createdAt.toIso8601String(),
      if (ui != null) 'ui': ui!.toJson(),
    };
  }
}

Map<String, Object?> _normalizeMap(Map<dynamic, dynamic> value) {
  return value.map((key, entry) => MapEntry(key.toString(), entry));
}
