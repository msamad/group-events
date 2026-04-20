class SduiAction {
  const SduiAction({
    required this.id,
    required this.label,
    required this.type,
    this.endpoint,
    this.confirm = false,
    this.visible = true,
    this.payload = const <String, Object?>{},
  });

  factory SduiAction.fromJson(Map<String, Object?> json) {
    return SduiAction(
      id: json['id'] as String,
      label: json['label'] as String,
      type: json['type'] as String,
      endpoint: (json['endpoint'] ?? json['route']) as String?,
      confirm: (json['confirm'] as bool?) ?? false,
      visible: (json['visible'] as bool?) ?? true,
      payload: _readObjectMap(json['payload']),
    );
  }

  final String id;
  final String label;
  final String type;
  final String? endpoint;
  final bool confirm;
  final bool visible;
  final Map<String, Object?> payload;

  Map<String, Object?> toJson() {
    return <String, Object?>{
      'id': id,
      'label': label,
      'type': type,
      if (endpoint != null) 'endpoint': endpoint,
      'confirm': confirm,
      'visible': visible,
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
