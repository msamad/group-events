class SduiAction {
  const SduiAction({
    required this.id,
    required this.label,
    required this.type,
    this.route,
    this.payload = const <String, Object?>{},
  });

  factory SduiAction.fromJson(Map<String, Object?> json) {
    return SduiAction(
      id: json['id'] as String,
      label: json['label'] as String,
      type: json['type'] as String,
      route: json['route'] as String?,
      payload: _readObjectMap(json['payload']),
    );
  }

  final String id;
  final String label;
  final String type;
  final String? route;
  final Map<String, Object?> payload;

  Map<String, Object?> toJson() {
    return <String, Object?>{
      'id': id,
      'label': label,
      'type': type,
      if (route != null) 'route': route,
      if (payload.isNotEmpty) 'payload': payload,
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
